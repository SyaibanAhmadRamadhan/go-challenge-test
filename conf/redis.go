package conf

import (
	"os"
	"strconv"

	"challenge-test-synapsis/helper"
)

type RedisConf struct {
	Host string
	Port int
	DB   int
	Pass string
}

func EnvRedisConf() RedisConf {
	host := os.Getenv("REDIS_HOST")
	if host == "" {
		panic("REDIS_HOST is not exist")
	}
	pass := os.Getenv("REDIS_PASS")
	if pass == "" {
		panic("REDIS_PASS is not exist")
	}
	port := os.Getenv("REDIS_PORT")
	if port == "" {
		panic("REDIS_PORT is not exist")
	}
	portInt, err := strconv.Atoi(port)
	if err != nil {
		helper.PanicIf(err)
	}
	db := os.Getenv("REDIS_DB")
	if db == "" {
		panic("REDIS_DB is not exist")
	}
	dbInt, err := strconv.Atoi(db)
	if err != nil {
		helper.PanicIf(err)
	}
	return RedisConf{
		Host: host,
		Port: portInt,
		DB:   dbInt,
		Pass: pass,
	}
}
