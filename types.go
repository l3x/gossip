package main

import (
	"fmt"
)

type Host struct {
	Port          int					// 7001
	MovieInfo
	Peers         []string				// http://localhost:7001
	PeersMovieInfo map[string]MovieInfo	//
}

type MovieInfo struct {
	FavMovie	string `json:"fav_movie"`	// "Leaving Las Vegas"
	FavCntr	int	`json:"fav_cntr"` 			// Sequential count of favorite movie choices
}
func (bi MovieInfo) String() string {
	return fmt.Sprintf("FavMovie: %s, FavCntr: %d", bi.FavMovie, bi.FavCntr)
}

func (h Host) String() string {
	return fmt.Sprintf(
		"Port:  %d\nFavMovie:  %s\nFavCntr:  %d\nPeers: %+v",
		h.Port, h.FavMovie, h.FavCntr, h.Peers)
}

func hostAddressFromPort(port int) string {
	return fmt.Sprintf("http://localhost:%d", port)
}


func newHost(port int) (Host, error) {
	movieInfo := &MovieInfo{FavMovie:randomMovie(), FavCntr:0}
	host := &Host{Port: port, MovieInfo: *movieInfo, PeersMovieInfo: make(map[string]MovieInfo)}

	fmt.Println("Host", Self)
	return *host, nil
}

type PeersMsg struct {
	Peers []string `json:"peers"`  // [ "http://localhost:7001", "http://localhost:7002" ]
}

type FavMovieMsg struct {
	PeerAddress string  `json:"peer_address"`
	MovieInfo MovieInfo `json:"movie_info"`  // [ "http://localhost:7001", "http://localhost:7002" ]
}

