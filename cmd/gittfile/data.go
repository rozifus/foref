package gittfile


type GittfileData struct {
	GittfileNotice string `yaml:"gittfile_notice,omitempty"`
	GittfileVersion string `yaml:"gittfile_version,omitempty"`
	Description string `yaml:"description,omitempty"`
	Source string `yaml:"source,omitempty"`
	Identifiers []string `yaml:"identifiers"`
}

