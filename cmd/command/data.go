package command


type FrefData struct {
	FreffileNotice string `yaml:"freffile_notice,omitempty"`
	FreffileVersion string `yaml:"freffile_version,omitempty"`
	Description string `yaml:"description,omitempty"`
	Source string `yaml:"source,omitempty"`
	Identifiers []string `yaml:"identifiers"`
}

