package config

type (
	Config struct {
		App      App      `yaml:"app"`
		Postgres Postgres `yaml:"postgresql"`
		Mongo    Mongo    `yaml:"mongodb"`
		Redis    Redis    `yaml:"redis"`
		Account  Account  `yaml:"account"`
		Auth     Auth     `yaml:"auth"`
	}

	App struct {
		Rest Rest `yaml:"rest"`
		Grpc Grpc `yaml:"grpc"`
	}

	Auth struct {
		SecretKey            string `yaml:"secret_key"`
		Audience             string `yaml:"audience"`
		Issuer               string `yaml:"issuer"`
		AccessTokenDuration  uint64 `yaml:"access_token_duration"`
		RefreshTokenDuration uint64 `yaml:"refresh_token_duration"`
	}

	Rest struct {
		Name    string `yaml:"name"`
		Address string `yaml:"address"`
	}

	Grpc struct {
		Name    string `yaml:"name"`
		Address string `yaml:"address"`
	}

	Postgres struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		DBName   string `yaml:"db_name"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
	}

	Mongo struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		DBName   string `yaml:"db_name"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
	}

	Redis struct {
		Address  string `yaml:"address"`
		Password string `yaml:"password"`
		DB       int    `yaml:"db"`
	}

	Account struct {
		MinUsernameLength int `yaml:"min_username_length"`
	}
)
