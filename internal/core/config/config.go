package config

import (
	"path/filepath"

	"github.com/spf13/viper"
)

const (
	// database
	dalMocked = "PORAHAT_DB_MOCKED"
	dbUrl     = "PORAHAT_DB_URL"

	// rest
	restPort        = "PORAHAT_REST_PORT"
	restSize        = "PORAHAT_REST_SIZE_LIMIT"
	restHeader      = "PORAHAT_REST_HEADER_SIZE"
	restIdleTimeout = "PORAHAT_REST_IDLE_TIMEOUT"
)

type Config struct {
	DalConfig        *DalConfig
	RestServerConfig *RestConfig
	Mocks            *Mocks
}

type RestConfig struct {
	Port        int
	SizeLimit   int
	HeaderSize  int
	IdleTimeout int
}

type DalConfig struct {
	Url string
}

type Mocks struct {
	DalMocked bool
}

func LoadConfig(envFilePath string) (*Config, error) {
	v := viper.New()
	v.SetDefault(dalMocked, false)
	v.SetDefault(restPort, 8080)
	v.SetDefault(restSize, 4*1024*1024)
	v.SetDefault(restHeader, 4*1024)
	v.SetDefault(restIdleTimeout, 120)
	v.SetDefault(dbUrl, "postgres://postgres:1231312Xx@localhost:5432/flower_management")

	if envFilePath != "" {
		dir, file := filepath.Split(envFilePath)

		v.AddConfigPath(dir)
		v.SetConfigName(file)
		v.SetConfigType("env")

		if err := v.ReadInConfig(); err != nil {
			return nil, err
		}
	} else {
		v.AutomaticEnv()
	}

	return &Config{
		DalConfig: &DalConfig{
			Url: v.GetString(dbUrl),
		},
		RestServerConfig: &RestConfig{
			Port:        v.GetInt(restPort),
			SizeLimit:   v.GetInt(restSize),
			HeaderSize:  v.GetInt(restHeader),
			IdleTimeout: v.GetInt(restIdleTimeout),
		},
		Mocks: &Mocks{
			DalMocked: v.GetBool(dalMocked),
		},
	}, nil
}
