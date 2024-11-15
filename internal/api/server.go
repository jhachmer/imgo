package api

import (
	"log"
	"net/http"

	"github.com/jhachmer/imgo/internal/config"
	"github.com/jhachmer/imgo/internal/store"
	"github.com/jhachmer/imgo/internal/utils"
)

type Server struct {
	Addr   string
	Prefix string
	Store  *store.Store
}

func NewServer(addr, prefix string, store *store.Store) *Server {
	return &Server{
		Addr:   addr,
		Prefix: prefix,
		Store:  store,
	}
}

func (apiServer *Server) setupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET "+apiServer.Prefix+"/{$}", Chain(handleIndex, Logging()))
	mux.HandleFunc("POST "+apiServer.Prefix+"/upload", Chain(handleUpload, Logging()))
	mux.HandleFunc("GET "+apiServer.Prefix+"/fourier/{filename}", Chain(handleFourier, Logging()))
	mux.Handle(config.FILESERVER, http.StripPrefix(config.FILESERVER, fileHandler(config.IMAGELOCATION)))
}

func (apiServer *Server) setup() error {
	err := utils.SetupDir(config.IMAGELOCATION)
	if err != nil {
		return err
	}
	return nil
}

func (apiServer *Server) Serve() {
	mux := http.NewServeMux()
	if err := apiServer.setup(); err != nil {
		log.Fatal(err)
	}
	apiServer.setupRoutes(mux)
	log.Println("Starting server on " + apiServer.Addr)
	log.Fatal(http.ListenAndServe(apiServer.Addr, clearTrailingSlash(logFileServer(mux))))
}
