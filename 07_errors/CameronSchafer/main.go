package main

import (
	"fmt"
	"io"
	"os"
	"sync" //https://golang.org/pkg/sync/#Map
)

var out io.Writer = os.Stdout

const (
	errInvalidID   = 400
	errPupNotFound = 404
)

type Error struct {
	Message string
	Code    uint32
}

func (e *Error) Error() string {
	return fmt.Sprintf("%d: %v", e.Code, e.Message)
}

func (e *Error) setError(c uint32, m string) {
	e.Code, e.Message = c, m
}

func (e *Error) errInvalidValue() string {
	return fmt.Sprintf("%d: %v", e.Code, e.Message)
}

type Puppy struct {
	pid    int
	breed  string
	colour string
	value  string
}

type MapStore struct {
	store map[int]Puppy
}

type SyncStore struct {
	store sync.Map
}

type Storer interface {
	CreatePuppy(*Puppy) (int, error)
	ReadPuppy(int) (*Puppy, error)
	UpdatePuppy(*Puppy)
	DeletePuppy(int)
}

func (m *MapStore) CreatePuppy(p *Puppy) (int, error) {
	_, ok := m.store[p.pid]
	if !ok {
		m.store[p.pid] = *p
		return p.pid, nil
	}
	return -1, fmt.Errorf("Could not Create new puppy")
}

func (m *MapStore) ReadPuppy(pid int) (*Puppy, error) {
	pup, ok := m.store[pid]
	if ok {
		return &pup, nil
	}
	return nil, fmt.Errorf("Could not retrieve puppy with id: %d", pid)
}
func (m *MapStore) UpdatePuppy(p *Puppy) {
	m.store[p.pid] = *p
}
func (m *MapStore) DeletePuppy(pid int) {
	delete(m.store, pid)
}

func (s *SyncStore) CreatePuppy(p *Puppy) (int, error) {
	_, ok := s.store.Load(p.pid)
	if !ok {
		s.store.Store(p.pid, p)
		return p.pid, nil
	}
	return -1, fmt.Errorf("Could not Create new puppy")
}
func (s *SyncStore) ReadPuppy(pid int) (*Puppy, error) {
	pup, ok := s.store.Load(pid)
	if ok {
		return pup.(*Puppy), nil
	}
	return nil, fmt.Errorf("Could not retrieve puppy with id: %d", pid)
}
func (s *SyncStore) UpdatePuppy(p *Puppy) {
	s.store.Store(p.pid, p)
}
func (s *SyncStore) DeletePuppy(pid int) {
	s.store.Delete(pid)
}

// using the map store + methods
func usingNormMap(pups []Puppy) {
	// initialise the map store
	var puppyMS Storer = &MapStore{store: make(map[int]Puppy)}
	for _, n := range pups {
		pup := n
		_, e := puppyMS.CreatePuppy(&pup)
		if e != nil {
			fmt.Fprintln(out, e)
		}
	}
	if p, e := puppyMS.ReadPuppy(pups[1].pid); p != nil {
		fmt.Fprintln(out, *p)
	} else {
		fmt.Fprintln(out, e)
	}

	upDog := pups[0]
	upDog.colour = "red"
	puppyMS.UpdatePuppy(&upDog)
	puppyMS.DeletePuppy(pups[1].pid)

	fmt.Fprintln(out, puppyMS)
}

// using the sync map store + methods
func usingSyncMap(pups []Puppy) {
	// initialise the sync store
	var puppySS Storer = &SyncStore{}
	for _, n := range pups {
		pup := n
		_, e := puppySS.CreatePuppy(&pup)
		if e != nil {
			fmt.Fprintln(out, e)
		}
	}
	if p, e := puppySS.ReadPuppy(pups[1].pid); p != nil {
		fmt.Fprintln(out, *p)
	} else {
		fmt.Fprintln(out, e)
	}

	upDog := pups[0]
	upDog.colour = "red"
	puppySS.UpdatePuppy(&upDog)
	puppySS.DeletePuppy(pups[1].pid)

	for _, pup := range pups {
		if p, e := puppySS.ReadPuppy(pup.pid); p != nil {
			fmt.Fprintln(out, *p)
		} else {
			fmt.Fprintln(out, e)
		}
	}
}

func main() {
	pups := []Puppy{
		{99, "poodle", "blue", "$10.99"},
		{100, "lab", "orange", "$9.99"},
		{101, "cat", "striped", "$99.99"},
	}
	usingNormMap(pups)
	usingSyncMap(pups)
}
