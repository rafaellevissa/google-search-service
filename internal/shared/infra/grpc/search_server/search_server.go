package searchserver

import (
	"context"

	googlesearch "gitlab.com/colmeia/desafio-google-search/internal/shared/services/google_search"
	"gitlab.com/colmeia/desafio-google-search/pkg/grpc/searchproto/searchproto"
)

type SearchServer struct {
	searchproto.UnimplementedSearchServer
	SearchService googlesearch.IGoogleSearch
}

func (s *SearchServer) ExecuteSearch(ctx context.Context, req *searchproto.SearchRequest) (*searchproto.SearchResponse, error) {
	result, err := s.SearchService.Search(req.Query, req.Country, req.Lang)
	if err != nil {
		return &searchproto.SearchResponse{
			Error: err.Error(),
			Items: nil,
		}, nil
	}

	var items []*searchproto.SearchItem
	for _, item := range result.Items {
		items = append(items, &searchproto.SearchItem{
			Title:   item.Title,
			Snippet: item.Snippet,
			Url:     item.Url,
		})
	}

	return &searchproto.SearchResponse{
		Error: "",
		Items: items,
	}, nil
}
