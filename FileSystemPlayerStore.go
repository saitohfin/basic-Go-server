package main

import (
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadSeeker
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	f.database.Seek(0, 0)
	var results, _ = NewLeague(f.database)
	return results
}

func (f *FileSystemPlayerStore) GetPlayerScore(player_name string) int {
	var score int
	var league = f.GetLeague()
	for _, player := range league {
		if player_name == player.Name {
			score = player.Wins
			break
		}
	}

	return score
}
