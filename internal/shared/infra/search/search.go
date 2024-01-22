package search

import (
	"log"
	"net"

	"gitlab.com/colmeia/desafio-google-search/internal/shared/infra/environment"
	searchserver "gitlab.com/colmeia/desafio-google-search/internal/shared/infra/grpc/search_server"
	googlesearch "gitlab.com/colmeia/desafio-google-search/internal/shared/services/google_search"
	"gitlab.com/colmeia/desafio-google-search/pkg/grpc/searchproto/searchproto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type SearchConfig struct {
	Environment environment.IEnvironment
	Logger      *zap.Logger
}

type ISearch interface {
	Setup()
}

type Search struct {
	Environment environment.IEnvironment
	Logger      *zap.Logger
	Listener    net.Listener
	Server      *grpc.Server
}

func New(config *SearchConfig) ISearch {
	return Search{
		Environment: config.Environment,
		Logger:      config.Logger,
	}
}

func (s Search) Setup() {
	s.SetupPort()
	s.SetupServer()
	s.Serve()
}

func (s *Search) SetupPort() {
	lis, err := net.Listen("tcp", ":"+environment.GetVar(environment.VAR_SERVER_PORT))
	if err != nil {
		s.Logger.Error(err.Error())
		log.Fatalf("failed to listen: %v", err)
	}

	s.Listener = lis
}

func (s *Search) SetupServer() {
	server := grpc.NewServer()
	search_sv := &searchserver.SearchServer{
		SearchService: googlesearch.New(s.Environment),
	}

	searchproto.RegisterSearchServer(server, search_sv)

	s.Server = server
}

func (s *Search) Serve() {
	log.Printf("Listening at %v", s.Listener.Addr())
	if err := s.Server.Serve(s.Listener); err != nil {
		s.Logger.Error(err.Error())
		log.Fatalf("failed to serve: %v", err)
	}
}
