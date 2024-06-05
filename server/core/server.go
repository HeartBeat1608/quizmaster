package core

import (
	"net/http"
	"path"
	"quizmaster-server/core/handlers"
	"quizmaster-server/core/middlewares"
	"quizmaster-server/pkg/database"
	"strings"
)

type AppServerMux struct {
	*http.ServeMux
	server *AppServer
	prefix string
}

func (g *AppServerMux) HandleFunc(pattern string, handler http.HandlerFunc) {
  parts := strings.Split(pattern, " ")

  var method string
  var pathname string

  if len(parts) == 2 {
    method = parts[0]
    pathname = parts[1]
  } else {
    method = "GET"
    pathname = parts[0]
  }

  patt := path.Join(g.prefix, pathname)
	g.ServeMux.HandleFunc(method + " " + patt, handler)
}

type AppServer struct {
	listenAddr  string
	mux         *http.ServeMux
	middlewares []func(http.Handler) http.Handler
}

func NewAppServer(listenAddr string) *AppServer {
	mux := http.NewServeMux()
	return &AppServer{listenAddr, mux, make([]func(http.Handler) http.Handler, 0)}
}

func (s *AppServer) ListenAddr() string {
	return s.listenAddr
}

func (s *AppServer) Serve() error {
	return http.ListenAndServe(s.listenAddr, middlewares.Logger(s.mux))
}

func (s *AppServer) Group(prefix string, fn func(mux *AppServerMux)) {
	groupMux := &AppServerMux{s.mux, s, prefix}
	fn(groupMux)
}

func (s *AppServer) Use(middleware func(http.Handler) http.Handler) {
	s.middlewares = append(s.middlewares, middleware)
}

func (s *AppServer) SetupMiddlewares() {
	s.Use(middlewares.Logger)
}

func (s *AppServer) SetupRoutes(db *database.AppDatabase) {
	authHandler := handlers.NewAuthHandler(db)
	s.Group("/auth", func(mux *AppServerMux) {
		mux.HandleFunc("POST /login", authHandler.HandleLogin)
		mux.HandleFunc("POST /register", authHandler.HandleRegister)
	})
}
