package config

type (
	Config struct {
		State      string
		RestfulAPI struct {
			Host string `yaml:"host"`
			Port string `yaml:"port"`
		} `yaml:"restful_api"`
		DB struct {
			Host     string `yaml:"host"`
			Port     string `yaml:"port"`
			Database string `yaml:"database"`
			Username string `yaml:"username"`
			Password string `yaml:"password"`
		} `yaml:"db"`
		Seller struct {
			Notification struct {
				Type     []string `yaml:"type"`
				Template struct {
					Sms   string `yaml:"sms"`
					Email struct {
						Subject string `yaml:"subject"`
						Body    string `yaml:"body"`
						Sender  struct {
							Add      string `yaml:"add"`
							Host     string `yaml:"host"`
							From     string `yaml:"from"`
							Password string `yaml:"password"`
						} `yaml:"sender"`
					} `yaml:"email"`
				} `yaml:"template"`
			} `yaml:"notification"`
		} `yaml:"seller"`
	}
)
