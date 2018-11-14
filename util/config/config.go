package config

import (
	"bigproject/entity"
	"github.com/spf13/viper"
)

func GetConfig()  entity.Config {
	// get config from config.json
	conf := entity.Config{
		DBHost: viper.GetString(`database.host`),
		DBPort : viper.GetString(`database.port`),
		DBUser : viper.GetString(`database.user`),
		DBPass : viper.GetString(`database.pass`),
		DBName : viper.GetString(`database.name`),
		RedisMaxActive: viper.GetInt(`redis.maxActive`),
		RedisMaxIdle: viper.GetInt(`redis.maxIdle`),
		RedisHost: viper.GetString(`redis.host`),
		RedisPort: viper.GetString(`redis.port`),
		MQDefaultConsumerMaxAttempts: viper.GetInt(`mq.defaultConsumerMaxAttempts`),
		MQDefaultConsumerMaxInFlight: viper.GetInt(`mq.defaultConsumerMaxInFlights`),

	}
	return conf
}