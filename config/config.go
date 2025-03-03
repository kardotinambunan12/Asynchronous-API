package config

import (
	"os"
	errorhandler "spe_test/error_handler"

	"github.com/joho/godotenv"
)

type Config interface {
	Get(key string) string
}

type configImpl struct {
}

func (config *configImpl) Get(key string) string {
	return os.Getenv(key)
}

func New(filenames ...string) Config {
	err := godotenv.Load(filenames...)
	errorhandler.PanicIfNeeded(err)
	return &configImpl{}
}
