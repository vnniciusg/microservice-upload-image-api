package config

import "github.com/spf13/viper"

type conf struct {
	CloudName   string `mapstructure:"CLOUDNAME"`
	CloudAPIKey string `mapstructure:"CLOUDAPIKEY"`
	CloudSecret string `mapstructure:"CLOUDSECRET"`
}

func LoadConfig() (*conf, error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	cfg := &conf{
		CloudName:   viper.GetString("CLOUDNAME"),
		CloudAPIKey: viper.GetString("CLOUDAPIKEY"),
		CloudSecret: viper.GetString("CLOUDSECRET"),
	}

	return cfg, err
}
