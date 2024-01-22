package googlesearch

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"gitlab.com/colmeia/desafio-google-search/internal/shared/infra/environment"
)

type IGoogleSearch interface {
	Search(query, country, lang string) (Search, error)
}

type GoogleSearch struct {
	environment environment.IEnvironment
}

type Search struct {
	Error string
	Items []SearchItem
}

type SearchItem struct {
	Title   string
	Snippet string
	Url     string
}

func New(env environment.IEnvironment) IGoogleSearch {
	return GoogleSearch{
		environment: env,
	}
}

func (g GoogleSearch) Search(query, country, lang string) (Search, error) {
	apiKey, cx := g.environment.Var(environment.VAR_G_API_KEY), g.environment.Var(environment.VAR_G_CX)

	result, err := ConsumeSearchAPI(apiKey, cx, query, country, lang)
	if err != nil {
		return Search{
			Error: err.Error(),
		}, err
	}

	var search Search
	if items, ok := result["items"].([]interface{}); ok {
		for _, item := range items {
			itemData := item.(map[string]interface{})
			searchItem := SearchItem{
				Title:   itemData["title"].(string),
				Snippet: itemData["snippet"].(string),
				Url:     itemData["link"].(string),
			}
			search.Items = append(search.Items, searchItem)
		}
	}

	return search, nil
}

func ConsumeSearchAPI(apiKey, cx, query, country, lang string) (map[string]interface{}, error) {
	url := fmt.Sprintf("https://www.googleapis.com/customsearch/v1?key=%s&cx=%s&q=%s&cr=%s&lr=%s", apiKey, cx, query, country, lang)

	response, err := http.Get(url)
	if err != nil {
		return map[string]interface{}{}, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return map[string]interface{}{}, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return map[string]interface{}{}, err
	}

	if errorMsg, ok := result["error"]; ok {
		return map[string]interface{}{}, fmt.Errorf("erro na resposta da api: %v", errorMsg)
	}

	return result, nil
}
