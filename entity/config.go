package entity


type Config struct {
	DBHost 	string
	DBPort	string
	DBUser	string
	DBPass	string
	DBName	string
	RedisMaxActive int
	RedisMaxIdle int
	RedisHost string
	RedisPort string
	MQDefaultConsumerMaxAttempts int
	MQDefaultConsumerMaxInFlight int
}