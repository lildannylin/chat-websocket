package main

import (
	"github.com/bmizerany/pat"
	"github.com/lildannylin/chat-websocket/internal/handlers"
	"net/http"
)

func routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Home))
	mux.Get("/ws", http.HandlerFunc(handlers.WsEndPoint))

	return mux
}
