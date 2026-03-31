package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type ServerEx struct {
	Log    *log.Logger
	Server http.Server
}

func CreateNewServer(l *log.Logger) *ServerEx {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HandlerForm)
	mux.HandleFunc("/upload", handlers.HandlerUpload)
	return &ServerEx{Log: l, Server: http.Server{Addr: ":8080", Handler: mux, ErrorLog: l, ReadTimeout: 5 * time.Second, WriteTimeout: 10 * time.Second, IdleTimeout: 15 * time.Second}}
}
