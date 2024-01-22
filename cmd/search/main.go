package main

import (
	"gitlab.com/colmeia/desafio-google-search/internal/shared/infra/environment"
	"gitlab.com/colmeia/desafio-google-search/internal/shared/infra/search"
)

func main() {
	env := environment.New()
	err := env.Load()
	if err != nil {
		panic(err)
	}

	service := search.New(&search.SearchConfig{
		Environment: env,
	})

	service.Setup()
}
