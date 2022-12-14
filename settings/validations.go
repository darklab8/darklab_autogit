package settings

type ValidationScheme struct {
	Sections struct {
		Hook struct {
			CommitMsg struct {
				Enabled bool `yaml:"enabled"`
			} `yaml:"commitMsg"`
		} `yaml:"hook"`
		Changelog struct {
			Enabled bool `yaml:"enabled"`
		} `yaml:"changelog"`
	} `yaml:"sections"`
	Rules struct {
		Issue struct {
			Present bool `yaml:"present"`
		} `yaml:"issue"`
		Header struct {
			MaxLength int `yaml:"maxLength"`
			Type      struct {
				Lowercase bool `yaml:"lowercase"`
			} `yaml:"type"`
			Scope struct {
				Present   bool     `yaml:"present"`
				Lowercase bool     `yaml:"lowercase"`
				Whitelist []string `yaml:"whitelist"`
			} `yaml:"scope"`
			Subject struct {
				MinWords int `yaml:"minWords"`
			} `yaml:"subject"`
		} `yaml:"header"`
	} `yaml:"rules"`
}
