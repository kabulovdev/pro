package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	Environment        string
	PostgresHost       string
	PostgresPort       int
	PostgresDatabase   string
	PostgresUser       string
	PostgresPassword   string
	LogLevel           string
	RPCPort            string
	PostServiceHost    string
	PostServicePort    int
	ReatingServiceHost string
	ReatingServicePort int
	ReviewServiceHost  string
	ReviewServicePort  string
	Partitions         int
	KafkaHost          string
	KafkaPort          string
}

func Load() Config {
	c := Config{}
	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))
	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "database-1.c9lxq3r1itbt.us-east-1.rds.amazonaws.com"))
	c.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	c.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "custumer_db"))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "abduazim"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "1234"))
	c.Partitions = cast.ToInt(getOrReturnDefault("PARITIIONS", 0))

	c.ReatingServiceHost = cast.ToString(getOrReturnDefault("REATING_HOST", "reating_service"))
	c.ReatingServicePort = cast.ToInt(getOrReturnDefault("REATING_PORT", 9084))

	c.PostServiceHost = cast.ToString(getOrReturnDefault("POST_HOST", "post_service"))
	c.PostServicePort = cast.ToInt(getOrReturnDefault("POST_PORT", 9083))
	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))

	//c.KafkaHost = cast.ToString(getOrReturnDefault("KAFKA_HOST", "kafka"))
	//c.KafkaPort = cast.ToString(getOrReturnDefault("KAFKA_PORT", "9092"))

	c.RPCPort = cast.ToString(getOrReturnDefault("RPC_PORT", ":9088"))
	return c
}

func getOrReturnDefault(key string, defaulValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaulValue
}
