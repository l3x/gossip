package main

// Ex: addPeer("http://localhost:7001")
func addPeer(peer string) {
	// If not already in peers list (do include my address into peers list)
	if !stringInSlice(peer, Self.Peers) {
		mutex.Lock()
		Self.Peers = append(Self.Peers, peer)
		gotNewPeers = true
		mutex.Unlock()
	}
}

func randomMovie() string {
	return movies[random(0, numberOfMovies)]
}

func setFavMovie() {
	Self.FavMovie = randomMovie()
	Self.FavCntr += 1
	broadcastMyFavMovie()
}

func updatePeerFavMovie(peerAddress string, bi MovieInfo) {
	mutex.Lock()
	Self.PeersMovieInfo[peerAddress] = bi
	mutex.Unlock()
}