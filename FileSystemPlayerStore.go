package main

import (
	"encoding/json"
	"io"
	"os"
)

type FileSystemPlayerStore struct {
	database io.Writer
	league League
}

func NewFileSystemPlayerStore(database *os.File) *FileSystemPlayerStore{
	database.Seek(0,0)
	league, _ := NewLeague(database)
	return &FileSystemPlayerStore{database: &tape{database}, league: league}
}

func (f *FileSystemPlayerStore) GetLeague() League {
	return f.league
}

func (f *FileSystemPlayerStore) GetPlayerScore(playerName string) int {
	var score int
	var league = f.GetLeague()
	player := league.Find(playerName)

	if player != nil {
		score = player.Wins
	}

	return score
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	player := f.league.Find(name)
	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}

	json.NewEncoder(f.database).Encode(f.league)
}
