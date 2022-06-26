package config

/*
import "github.com/spf13/viper"

type Config struct {
	Port               string `mapstructure:"PORT"`
	CustomerSvcBaseUrl string `mapstructure:"CUSTOMER_SVC_BASE_URL"`
	DbConnectionUrl    string `mapstructure:"DB_CONNECTION_URL"`
	DbCollectionName   string `mapstructure:"DB_COLLECTION_NAME"`
	DbName             string `mapstructure:"DB_NAME"`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./src/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)

	return
}
*/
