package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type FileSystemPlayerStore struct {
	database *json.Encoder
	league League
}

func NewFileSystemPlayerStore(database *os.File) (*FileSystemPlayerStore, error){
	err := initialisePlayerDBFile(database)
	if err != nil {
		return nil, fmt.Errorf("problem initialising player db file, %v", err)
	}

	league, err := NewLeague(database)
	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", database.Name(), err)
	}
	return &FileSystemPlayerStore{database: json.NewEncoder(&tape{database}), league: league}, nil
}

func initialisePlayerDBFile(database *os.File) error {
	database.Seek(0, 0)
	info, err := database.Stat()
	if err != nil {
		return fmt.Errorf("problem getting file info from file %s, %v", database.Name(), err)
	}
	if info.Size() == 0 {
		database.Write([]byte("[]"))
		database.Seek(0, 0)
	}
	return nil
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

	f.database.Encode(f.league)
}
