package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

const saveTimeout = 10e9

type URLMap struct {
	urls map[string]string
	mu   sync.RWMutex
}

func NewURLMap() *URLMap {
	return &URLMap{urls: make(map[string]string)}

}

func (m *URLMap) Set(key, url string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.urls[key] = url
	return
}

func (m *URLMap) Get(key string) (url string) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	url = m.urls[key]
	return
}

func (m *URLMap) WriteTo(w io.Writer) error {
	m.mu.RLock()
	defer m.mu.RUnlock()
	e := json.NewEncoder(w)
	return e.Encode(m.urls)
}

func (m *URLMap) ReadFrom(r io.Reader) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	e := json.NewDecoder(r)
	return e.Decode(&m.urls)
}

type URLStore struct {
	urls     *URLMap
	count    int
	mu       sync.Mutex
	filename string
	dirty    chan bool
}

func NewURLStore(filename string) *URLStore {
	s := &URLStore{
		urls:     NewURLMap(),
		filename: filename,
		dirty:    make(chan bool, 1),
	}
	if err := s.load(); err != nil {
		log.Print("UrlStore: ", err)
	}
	go s.saveLoop()
	return s
}

func (s *URLStore) Put(url string) (key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for {
		key = genKey(s.count)
		s.count++
		if u := s.urls.Get(key); u == "" {
			break
		}
	}
	s.urls.Set(key, url)
	select {
	case s.dirty <- true:
	default:
	}
	return
}

func (s *URLStore) Get(key string) (url string) {
	return s.urls.Get(key)
}

func (s *URLStore) load() error {
	f, err := os.Open(s.filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return s.urls.ReadFrom(f)
}

func (s *URLStore) save() error {
	f, err := os.OpenFile(s.filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	log.Print("URLStore: saving...")
	return s.urls.WriteTo(f)
}

func (s *URLStore) saveLoop() {
	for {
		<-s.dirty
		if err := s.save(); err != nil {
			log.Print("URLStore: ", err)
		}
		time.Sleep(saveTimeout)
	}
}
