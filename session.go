package sessionmanagerredis

import (
	"sync"

	"errors"
)

type SessionManager struct {
	Storage Storage
	mu      *sync.Mutex
}

func New() *SessionManager {
	return &SessionManager{Storage: NewMemStorage(), mu: &sync.Mutex{}}
}

type Storage interface {
	Put(string, string) error
	Get(string) (string, error)
}

type MemStorage struct {
	data map[string]string
	mu   *sync.Mutex
}

func NewMemStorage() *MemStorage {
	return &MemStorage{data: make(map[string]string), mu: &sync.Mutex{}}
}

func (s *MemStorage) Put(key, value string) error {
	if key == "" {
		return errors.New("key should not be empty")
	}
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = value

	return nil
}

func (s *MemStorage) Get(key string) (string, error) {
	if key == "" {
		return "", errors.New("key should not be empty")
	}
	s.mu.Lock()
	defer s.mu.Unlock()

	if v, ok := s.data[key]; ok {
		return v, nil
	} else {
		return "", errors.New("key not found")
	}

}
