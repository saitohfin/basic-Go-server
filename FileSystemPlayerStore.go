package main

import (
	"encoding/json"
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
}

func (f *FileSystemPlayerStore) GetLeague() League {
	f.database.Seek(0, 0)
	var results, _ = NewLeague(f.database)
	return results
}

func (f *FileSystemPlayerStore) GetPlayerScore(player_name string) int {
	var score int
	var league = f.GetLeague()
	player := league.Find(player_name)

	if player != nil {
		score = player.Wins
	}

	return score
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	var league = f.GetLeague()
	player := league.Find(name)
	if player != nil {
		player.Wins++
	}

	f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(&league)
}
