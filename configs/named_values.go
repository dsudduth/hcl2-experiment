package configs

import (
	"fmt"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/hashicorp/hcl2/hcl/hclsyntax"
	"github.com/zclconf/go-cty/cty"
)

// Variable represents a "variable" block in a module or file.
type Variable struct {
	Name           string
	Description    string
	Default        cty.Value
	Type           cty.Type
	ParsingMode    VariableParsingMode
	DescriptionSet bool
	DeclRange      hcl.Range
}

// VariableParsingMode defines how values of a particular variable given by
// text-only mechanisms (command line arguments and environment variables)
// should be parsed to produce the final value.
type VariableParsingMode rune

// VariableParseLiteral is a variable parsing mode that just takes the given
// string directly as a cty.String value.
const VariableParseLiteral VariableParsingMode = 'L'

// VariableParseHCL is a variable parsing mode that attempts to parse the given
// string as an HCL expression and returns the result.
const VariableParseHCL VariableParsingMode = 'H'

// Parse uses the receiving parsing mode to process the given variable value
// string, returning the result along with any diagnostics.
//
// A VariableParsingMode does not know the expected type of the corresponding
// variable, so it's the caller's responsibility to attempt to convert the
// result to the appropriate type and return to the user any diagnostics that
// conversion may produce.
//
// The given name is used to create a synthetic filename in case any diagnostics
// must be generated about the given string value. This should be the name
// of the root module variable whose value will be populated from the given
// string.
//
// If the returned diagnostics has errors, the returned value may not be
// valid.
func (m VariableParsingMode) Parse(name, value string) (cty.Value, hcl.Diagnostics) {
	switch m {
	case VariableParseLiteral:
		return cty.StringVal(value), nil
	case VariableParseHCL:
		fakeFilename := fmt.Sprintf("<value for var.%s>", name)
		expr, diags := hclsyntax.ParseExpression([]byte(value), fakeFilename, hcl.Pos{Line: 1, Column: 1})
		if diags.HasErrors() {
			return cty.DynamicVal, diags
		}
		val, valDiags := expr.Value(nil)
		diags = append(diags, valDiags...)
		return val, diags
	default:
		// Should never happen
		panic(fmt.Errorf("Parse called on invalid VariableParsingMode %#v", m))
	}
}
