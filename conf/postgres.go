package conf

import (
	"fmt"
	"os"
	"strconv"

	"challenge-test-synapsis/helper"
)

type PostgresConf struct {
	User     string
	Password string
	Host     string
	Port     int
	DB       string
	SSL      string
}

func EnvPostgresConf() PostgresConf {
	UserMaster := os.Getenv("POSTGRES_USER")
	if UserMaster == "" {
		panic("POSTGRES_USER is not exist")
	}
	PasswordMaster := os.Getenv("POSTGRES_PASSWORD")
	if PasswordMaster == "" {
		panic("POSTGRES_PASSWORD is not exist")
	}
	HostMaster := os.Getenv("POSTGRES_HOST")
	if HostMaster == "" {
		panic("POSTGRES_HOST is not exist")
	}
	PortMaster := os.Getenv("POSTGRES_PORT")
	if PortMaster == "" {
		panic("POSTGRES_PORT is not exist")
	}
	portMasterInt, err := strconv.Atoi(PortMaster)
	if err != nil {
		helper.PanicIf(err)
	}
	dbMaster := os.Getenv("POSTGRES_DB")
	if dbMaster == "" {
		panic("POSTGRES_DB is not exist")
	}
	sslMaster := os.Getenv("POSTGRES_SSL")
	if sslMaster == "" {
		panic("POSTGRES_SSL is not exist")
	}

	conf := PostgresConf{
		User:     UserMaster,
		Password: PasswordMaster,
		Host:     HostMaster,
		Port:     portMasterInt,
		DB:       dbMaster,
		SSL:      sslMaster,
	}
	return conf
}

func (p PostgresConf) DBUrl() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		p.Host, p.Port, p.User, p.Password, p.DB, p.SSL)
}
