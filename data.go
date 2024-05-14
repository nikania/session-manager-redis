package sessionmanagerredis

import (
	"crypto/rand"
	"encoding/base64"
	"sync"
	"time"
)

type Status int

const (
	Unchanged Status = iota
	Modified
)

type sessionData struct {
	deadline time.Time
	// status   Status
	// token    string
	values map[string]interface{}
	mu     sync.Mutex
}

func generateToken() (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
}
