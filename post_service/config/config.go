package config

import (
	"os"
	"github.com/spf13/cast"
)

type Config struct{
	Environment string
	PostgresHost string
	PostgresPort int
	PostgresDatabase string
	PostgresUser string
	PostgresPassword string
	LogLevel string
	RPCPort string
	KafkaHost          string
	KafkaPort          string
	ReviewServiceHost string
	ReviewServicePort string
}
func Load() Config{
	c:=Config{}
	c.Environment=cast.ToString(getOrReturnDefault("ENVIRONMENT","develop"))
	c.PostgresHost=cast.ToString(getOrReturnDefault("POSTGRES_HOST","localhost"))
	c.PostgresPort=cast.ToInt(getOrReturnDefault("POSTGRES_PORT",5432))
	c.PostgresDatabase=cast.ToString(getOrReturnDefault("POSTGRES_DATABASE","post_db"))
	c.PostgresUser=cast.ToString(getOrReturnDefault("POSTGRES_USER","postgres"))
	c.PostgresPassword=cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD","sdy12197"))
	//c.KafkaHost = cast.ToString(getOrReturnDefault("KAFKA_HOST", "kafka"))
	//c.KafkaPort = cast.ToString(getOrReturnDefault("KAFKA_PORT", "9092"))
	c.LogLevel=cast.ToString(getOrReturnDefault("LOG_LEVEL","debug"))

	c.RPCPort=cast.ToString(getOrReturnDefault("RPC_PORT",":9097"))
	return c
}

func getOrReturnDefault(key string,defaulValue interface{}) interface{}{
	_,exists:=os.LookupEnv(key)
	if exists{
		return os.Getenv(key)
	}
	return defaulValue
}

