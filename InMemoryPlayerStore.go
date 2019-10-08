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

func (i *InMemoryPlayerStore) GetLeague() []Player {
	league := []Player{}
	for name, wins := range i.scores {
		league = append(league, Player{name, wins})
	}
	return league
}

func EmptyMemoryPlayerStoreFactory() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}
