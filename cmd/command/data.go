package command


type ForefData struct {
	ForeffileNotice string `yaml:"foreffile_notice,omitempty"`
	ForeffileVersion string `yaml:"foreffile_version,omitempty"`
	Description string `yaml:"description,omitempty"`
	Source string `yaml:"source,omitempty"`
	Identifiers []string `yaml:"identifiers"`
}

