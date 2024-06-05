package core

import "net/http"

type AppServer struct {
  listenAddr string
  mux *http.ServeMux
}

func NewAppServer(listenAddr string) *AppServer {
  mux := http.NewServeMux()
  return &AppServer{listenAddr, mux}
}



