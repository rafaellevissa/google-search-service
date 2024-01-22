package environment_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/colmeia/desafio-google-search/internal/shared/infra/environment"
)

type Test struct {
	Name   string
	Before func()
	Assert func()
}

func TestEnvironment(t *testing.T) {
	DEFAULT_ENV_INSTANCE := environment.New()

	tests := []Test{
		{
			Name: "IsLoadedEnv - Should return false when not loaded env",
			Assert: func() {
				assert.False(t, DEFAULT_ENV_INSTANCE.IsLoadedEnv())
			},
		},
		{
			Name: "IsLoadedEnv - Should return true when loaded env",
			Before: func() {
				os.Setenv(environment.VAR_APP_ENV, environment.APP_ENV_HOMOLOG)
			},
			Assert: func() {
				assert.True(t, DEFAULT_ENV_INSTANCE.IsLoadedEnv())
			},
		},
		{
			Name: "Load - Should return nil when already loaded env",
			Before: func() {
				os.Setenv(environment.VAR_APP_ENV, environment.APP_ENV_HOMOLOG)
			},
			Assert: func() {
				assert.Nil(t, DEFAULT_ENV_INSTANCE.Load())
			},
		},
		{
			Name: "Load - Should return err when load env",
			Before: func() {
				os.Setenv(environment.VAR_APP_ENV, "")
			},
			Assert: func() {
				assert.Error(t, DEFAULT_ENV_INSTANCE.Load())
			},
		},
		{
			Name: "Var - Should return env var value",
			Before: func() {
				os.Setenv("test", "123")
			},
			Assert: func() {
				testVar := DEFAULT_ENV_INSTANCE.Var("test")
				assert.Equal(t, "123", testVar)
			},
		},
	}

	for _, test := range tests {
		if test.Before != nil {
			test.Before()
		}
		test.Assert()
	}
}
