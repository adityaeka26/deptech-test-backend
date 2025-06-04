package config

import (
	"github.com/spf13/viper"
)

type EnvConfig struct {
	AppPort        string `mapstructure:"APP_PORT"`
	MySqlHost      string `mapstructure:"MYSQL_HOST"`
	MySqlPort      string `mapstructure:"MYSQL_PORT"`
	MySqlUsername  string `mapstructure:"MYSQL_USERNAME"`
	MySqlPassword  string `mapstructure:"MYSQL_PASSWORD"`
	MySqlDbName    string `mapstructure:"MYSQL_DB_NAME"`
	MySqlSslMode   string `mapstructure:"MYSQL_SSL_MODE"`
	JwtPrivateKey  string `mapstructure:"JWT_PRIVATE_KEY"`
	JwtPublicKey   string `mapstructure:"JWT_PUBLIC_KEY"`
	MinioEndpoint  string `mapstructure:"MINIO_ENDPOINT"`
	MinioAccessKey string `mapstructure:"MINIO_ACCESS_KEY"`
	MinioSecretKey string `mapstructure:"MINIO_SECRET_KEY"`
	MinioUseSSL    bool   `mapstructure:"MINIO_USE_SSL"`
}

func Load(filename string) (*EnvConfig, error) {
	var envCfg EnvConfig

	viper.AddConfigPath(".")
	viper.SetConfigName(filename)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&envCfg); err != nil {
		return nil, err
	}

	return &envCfg, nil
}
