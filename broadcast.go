package main

import (
	"fmt"
	"encoding/json"
	"bytes"
	"net/http"
	"log"
)

func broadcastPeers() {
	for _, peer := range Self.Peers {
		if peer != myAddress {
			// Don't broadcast to self!
			fmt.Printf("Broadcasting my peers to %s/peers\n", peer)
			postMyPeersTo(peer)
		}
	}
}


func broadcastMyFavMovie() {
	for _, peer := range Self.Peers {
		if peer != myAddress {
			// Don't broadcast to self!
			fmt.Printf("Broadcasting my peers to %s/peers\n", peer)
			postMyFavMovieTo(peer)
		}
	}
}


func postMyFavMovieTo(peer string) {

	fmt.Printf(">> postMyFavMovieTo (%v) < myPeers: %v\n", fmt.Sprintf("%s/newFavMovie", peer),  Self.Peers)

	msg := &FavMovieMsg{myAddress, Self.MovieInfo}
	bytesRepresentation, err := json.Marshal(msg)
	if err != nil {
		log.Println(err)
		return
	}

	resp, err := http.Post(
		fmt.Sprintf("%s/newFavMovie", peer),
		"application/json",
		bytes.NewBuffer(bytesRepresentation))

	if err != nil {
		log.Println(err)
		return
	}
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	log.Println("result", result)
}

func postMyPeersTo(peer string) {

	fmt.Printf(">> postMyPeersTo (%v) < myPeers: %v\n", fmt.Sprintf("%s/newPeers", peer),  Self.Peers)

	msg := &PeersMsg{Self.Peers}
	bytesRepresentation, err := json.Marshal(msg)
	if err != nil {
		log.Println(err)
		return
	}

	resp, err := http.Post(
		fmt.Sprintf("%s/newPeers", peer),
		"application/json",
		bytes.NewBuffer(bytesRepresentation))

	if err != nil {
		log.Println(err)
		return
	}
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	log.Println("result", result)
	gotNewPeers = false
}

