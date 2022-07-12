package config

type (
	Config struct {
		App      App      `yaml:"app"`
		Postgres Postgres `yaml:"postgresql"`
		Redis    Redis    `yaml:"redis"`
		Account  Account  `yaml:"account"`
		Auth     Auth     `yaml:"auth"`
	}

	App struct {
		Rest Rest `yaml:"rest" envconfig:"BASE_MICROSERVICE_APP_NAME"`
		Grpc Grpc `yaml:"grpc" envconfig:"BASE_MICROSERVICE_APP_ADDRESS"`
	}

	Auth struct {
		SecretKey            string `yaml:"secret_key"`
		Audience             string `yaml:"audience"`
		Issuer               string `yaml:"issuer"`
		AccessTokenDuration  uint64 `yaml:"access_token_duration"`
		RefreshTokenDuration uint64 `yaml:"refresh_token_duration"`
	}

	Rest struct {
		Name    string `yaml:"name" envconfig:"BASE_MICROSERVICE_APP_NAME"`
		Address string `yaml:"address" envconfig:"BASE_MICROSERVICE_APP_ADDRESS"`
	}

	Grpc struct {
		Name    string `yaml:"name" envconfig:"BASE_MICROSERVICE_APP_NAME"`
		Address string `yaml:"address" envconfig:"BASE_MICROSERVICE_APP_ADDRESS"`
	}

	Postgres struct {
		Username string `yaml:"username" envconfig:"BASE_MICROSERVICE_MYSQL_USERNAME"`
		Password string `yaml:"password" envconfig:"BASE_MICROSERVICE_MYSQL_PASSWORD"`
		DBName   string `yaml:"db_name" envconfig:"BASE_MICROSERVICE_MYSQL_DBNAME"`
		Host     string `yaml:"host" envconfig:"BASE_MICROSERVICE_MYSQL_HOST"`
		Port     string `yaml:"port" envconfig:"BASE_MICROSERVICE_MYSQL_PORT"`
	}

	Redis struct {
		Address  string `yaml:"address" envconfig:"REDIS_MICROSERVICE_URL"`
		Password string `yaml:"password" envconfig:"REDIS_MICROSERVICE_PASSWORD"`
		DB       int    `yaml:"db" envconfig:"REDIS_MICROSERVICE_DB"`
	}

	Account struct {
		MinUsernameLength int `yaml:"min_username_length"`
	}
)
