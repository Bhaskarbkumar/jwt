package api

import "database/sql"

type APIServer struct {
	addr string
	db   *sql.DB
}

func newAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {

}
