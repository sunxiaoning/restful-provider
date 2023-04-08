package config

import (
	"errors"
	"github.com/sunxiaoning/restful-provider/pkg/strutils"
)

type Config struct {
	Logger   *LoggerConfig   `yaml:"logger" mapstructure:"logger"`
	Server   *ServerConfig   `yaml:"server" mapstructure:"server"`
	Database *DatabaseConfig `yaml:"database" mapstructure:"database"`
}

type validation interface {
	Validate() error
}

func (c *Config) Validate() error {
	if c.Logger == nil {
		return errors.New("logger config not specified")
	}
	if strutils.IsBlank(c.Logger.Level) {
		return errors.New("logger level not specified")
	}
	if c.Server == nil {
		return errors.New("server config not specified")
	}
	if c.Server.Port <= 0 {
		return errors.New("invalid server port")
	}

	return nil
}

func NewConfig() *Config {
	return &Config{
		Logger: &LoggerConfig{
			Level: "debug",
		},
		Server: &ServerConfig{Port: 8080},
		Database: &DatabaseConfig{
			User:            "ldata_admin",
			Host:            "rm-wz91gshzzg1b25144.mysql.rds.aliyuncs.com",
			Port:            3306,
			DBName:          "ldata",
			LogLevel:        "info",
			LogColor:        false,
			BatchSize:       100,
			MaxIdleConns:    20,
			MaxOpenConns:    200,
			ConnMaxLifetime: 60,
		},
	}
}

type RedisConfig struct {
	Cluster []string `yaml:"user" mapstructure:"cluster"`
}

const (
	DBModRW = "rw"
	DBModR  = "r"
)

type DatabaseConfig struct {
	User            string `yaml:"user" mapstructure:"user"`
	Password        string `yaml:"password" mapstructure:"password"`
	Host            string `yaml:"host" mapstructure:"host"`
	Port            int    `yaml:"port" mapstructure:"port"`
	DBName          string `yaml:"dbname" mapstructure:"dbname"`
	DbMod           string `yaml:"db-mod" mapstructure:"db-mod"`
	LogLevel        string `yaml:"log-level" mapstructure:"log-level"`
	LogColor        bool   `yaml:"log-color" mapstructure:"log-color"`
	BatchSize       int    `yaml:"batch-size" mapstructure:"batch-size"`
	MaxIdleConns    int    `yaml:"max-idel-conns" mapstructure:"max-idel-conns"`
	MaxOpenConns    int    `yaml:"max-open-conns" mapstructure:"max-open-conns"`
	ConnMaxLifetime int    `yaml:"conn-max-lifetime" mapstructure:"conn-max-lifetime"`
}

func (d *DatabaseConfig) Validate() error {
	if d == nil {
		return errors.New("database config not specified")
	}
	if strutils.IsBlank(d.User) {
		return errors.New("database user not specified")
	}
	if strutils.IsBlank(d.Password) {
		return errors.New("database password not specified")
	}
	if strutils.IsBlank(d.Host) {
		return errors.New("database host not specified")
	}
	if d.Port < 0 || d.Port > 65535 {
		return errors.New("invalid database port")
	}
	if strutils.IsBlank(d.DBName) {
		return errors.New("database dbName not specified")
	}
	if d.MaxIdleConns == 0 {
		return errors.New("database max-idle-conns not specified")
	}
	if d.MaxOpenConns == 0 {
		return errors.New("database max-open-conns not specified")
	}
	if d.ConnMaxLifetime == 0 {
		return errors.New("database conn-max-lifetime not specified")
	}

	if d.MaxOpenConns < d.MaxIdleConns {
		d.MaxOpenConns = d.MaxIdleConns
	}

	if d.DbMod == DBModRW {
		if d.BatchSize < 1 {
			return errors.New("invalid database create batch_size")
		}
	}
	return nil
}

type ServerConfig struct {
	Port        int    `yaml:"port" mapstructure:"port"`
	Env         string `yaml:"env" mapstructure:"env"`
	SecurityKey string `yaml:"security_key" mapstructure:"security_key"`
}

type LoggerConfig struct {
	Level string
}
