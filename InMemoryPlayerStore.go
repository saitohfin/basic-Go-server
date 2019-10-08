package main

type InMemoryPlayerStore struct {
	scores map[string]int
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.scores[name]
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.scores[name] = i.scores[name] + 1
}

func EmptyMemoryPlayerStoreFactory() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}
