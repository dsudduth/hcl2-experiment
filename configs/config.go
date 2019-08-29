package configs

// Config is the top-level config object
type Config struct {
	Project       Project       `hcl:"project,block" json:"project"`
	AutomationSet AutomationSet `hcl:"automation_set,block" json:"automationSet"`
}

// AutomationSet is a container for all automation-specific settings.
type AutomationSet struct {
	Name        string   `hcl:"name,label" json:"name"`
	Description string   `hcl:"description,optional" json:"description"`
	Machines    []string `hcl:"machines" json:"machines"`
}
