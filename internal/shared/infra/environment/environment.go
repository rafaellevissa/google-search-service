package environment

import (
	"os"

	"github.com/joho/godotenv"
)

const (
	APP_ENV_PROD    = "prod"
	APP_ENV_HOMOLOG = "staging"
	APP_ENV_DEV     = "dev"
)

const (
	VAR_APP_ENV     = "APP_ENV"
	VAR_SERVER_PORT = "SERVER_PORT"
	VAR_G_API_KEY   = "G_API_KEY"
	VAR_G_CX        = "G_CX"
)

func GetVar(name string) string {
	return os.Getenv(name)
}

type IEnvironment interface {
	Load() error
	Var(name string) string
	IsLoadedEnv() bool
}

type Environment struct{}

func New() IEnvironment {
	return Environment{}
}

func (e Environment) Load() error {
	if e.IsLoadedEnv() {
		return nil
	}

	return godotenv.Load()
}

func (e Environment) Var(name string) string {
	return os.Getenv(name)
}

func (e Environment) IsLoadedEnv() bool {
	env := e.Var(VAR_APP_ENV)
	return env == APP_ENV_HOMOLOG || env == APP_ENV_PROD
}
