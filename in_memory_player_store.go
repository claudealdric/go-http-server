package main

type InMemoryPlayerStore struct {
	store map[string]int
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		map[string]int{},
	}
}

func (i *InMemoryPlayerStore) GetLeague() []Player {
	return nil
}

func (i *InMemoryPlayerStore) GetPlayerScore(player string) int {
	return i.store[player]
}

func (i *InMemoryPlayerStore) RecordWin(player string) {
	i.store[player]++
}
