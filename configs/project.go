package configs

// Project is a container for project-related configuration settings.
type Project struct {
	Name        string `hcl:"name,label" json:"name"`
	Description string `hcl:"description,optional" json:"description"`
	Version     string `hcl:"version" json:"version"`
}
