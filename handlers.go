package main

import (
	"fmt"
	"encoding/json"
	"io"
	"net/http"
)

func handleGetMyFav(w http.ResponseWriter, r *http.Request) {
	fmt.Println(">> GET  handleGetMyFav")
	bytes, err := json.MarshalIndent(Self.MovieInfo, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

func handleGetPeers(w http.ResponseWriter, r *http.Request) {
	fmt.Println(">> GET  handleGetPeers")
	bytes, err := json.MarshalIndent(Self.Peers, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}


func handleGetPeersFavMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println(">> GET  handleGetPeersFavMovie")
	bytes, err := json.MarshalIndent(Self.PeersMovieInfo, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}


func handleNewPeers(w http.ResponseWriter, r *http.Request) {
	fmt.Println(">> POST handleNewPeers ")
	w.Header().Set("Content-Type", "application/json")
	var m PeersMsg

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&m); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()

	for _, peer := range m.Peers {
		if !stringInSlice(peer, Self.Peers) {
			addPeer(peer)
		}
	}

	respondWithJSON(w, r, http.StatusCreated, Self.Peers)
}


func handleFavMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println(">> POST handleFavMovie ")
	w.Header().Set("Content-Type", "application/json")
	var m FavMovieMsg

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&m); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()

	for _, peer := range Self.Peers {
		if peer != myAddress {
			updatePeerFavMovie(m.PeerAddress, m.MovieInfo)
		}
	}

	respondWithJSON(w, r, http.StatusCreated, Self.Peers)
}



func handleSocketIo(w http.ResponseWriter, r *http.Request) {
	fmt.Println(">> POST handleFavMovie ")
	w.Header().Set("Content-Type", "application/json")
	var m FavMovieMsg

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&m); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()

	for _, peer := range Self.Peers {
		if peer != myAddress {
			updatePeerFavMovie(m.PeerAddress, m.MovieInfo)
		}
	}

	respondWithJSON(w, r, http.StatusCreated, Self.Peers)
}
