package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

// create handlers
func makeMuxRouter(h *hub) http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/favmovie", handleGetMyFav).Methods("GET")
	muxRouter.HandleFunc("/peers", handleGetPeers).Methods("GET")
	muxRouter.HandleFunc("/peersFavMovie", handleGetPeersFavMovie).Methods("GET")
	muxRouter.HandleFunc("/newPeers", handleNewPeers).Methods("POST")
	muxRouter.HandleFunc("/newFavMovie", handleFavMovie).Methods("POST")
	muxRouter.Handle("/ws", wsHandler{h: h})
	return muxRouter
}


