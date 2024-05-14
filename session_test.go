// FILEPATH: /C:/Users/nikaf/source/go/session-manager-redis/session_test.go

package sessionmanagerredis

import (
	"fmt"
	"testing"
	"time"
)

func TestSessionManager_PutAndGet(t *testing.T) {
	sessionManager := New()

	// Test putting a value into the storage
	err := sessionManager.Storage.Put("key1", "value1")
	if err != nil {
		t.Errorf("Failed to put value into storage: %v", err)
	}

	// Test getting the value from the storage
	value, err := sessionManager.Storage.Get("key1")
	if err != nil {
		t.Errorf("Failed to get value from storage: %v", err)
	}
	if value != "value1" {
		t.Errorf("Expected value 'value1', got %s", value)
	}
}

func TestSessionManager_PutAndGet_InvalidKey(t *testing.T) {
	sessionManager := New()

	// Test putting a value with an empty key into the storage
	err := sessionManager.Storage.Put("", "value1")
	if err == nil {
		t.Error("Expected an error when putting a value with an empty key")
	}

	// Test getting a value with an empty key from the storage
	_, err = sessionManager.Storage.Get("")
	if err == nil {
		t.Error("Expected an error when getting a value with an empty key")
	}
}

func TestSessionManager_PutAndGet_NonexistentKey(t *testing.T) {
	sessionManager := New()

	// Test getting a value with a nonexistent key from the storage
	_, err := sessionManager.Storage.Get("nonexistent")
	if err == nil {
		t.Error("Expected an error when getting a value with a nonexistent key")
	}
}

func TestSessionManager_Concurrency(t *testing.T) {
	sessionManager := New()
	for i := 0; i < 100; i += 100 {
		_ = sessionManager.Storage.Put(fmt.Sprint("key ", i), fmt.Sprint("val ", i))
	}

	go func() {
		for i := 0; i < 100; i++ {
			_ = sessionManager.Storage.Put(fmt.Sprint("key ", i), fmt.Sprint("val ", i))
		}
	}()
	go func() {
		for i := 100; i < 200; i++ {
			_ = sessionManager.Storage.Put(fmt.Sprint("key ", i), fmt.Sprint("val ", i))
		}
	}()
	go func() {
		for i := 200; i < 300; i++ {
			_ = sessionManager.Storage.Put(fmt.Sprint("key ", i), fmt.Sprint("val ", i))
		}
	}()
	go func() {
		<-time.After(1 * time.Microsecond)
		i := 0
		s, err := sessionManager.Storage.Get(fmt.Sprint("key ", i))
		if err != nil {
			t.Errorf("No key %s in Storage ", fmt.Sprint("key ", i))
		}
		if s != fmt.Sprint("val ", i) {
			t.Errorf("Value %s is not correct, should be %s", s, fmt.Sprint("val ", i))
		}
	}()
	go func() {
		<-time.After(1 * time.Microsecond)
		i := 100
		s, err := sessionManager.Storage.Get(fmt.Sprint("key ", i))
		if err != nil {
			t.Errorf("No key %s in Storage ", fmt.Sprint("key ", i))
		}
		if s != fmt.Sprint("val ", i) {
			t.Errorf("Value %s is not correct, should be %s", s, fmt.Sprint("val ", i))
		}
	}()
	go func() {
		<-time.After(1 * time.Microsecond)
		i := 200
		s, err := sessionManager.Storage.Get(fmt.Sprint("key ", i))
		if err != nil {
			t.Errorf("No key %s in Storage ", fmt.Sprint("key ", i))
		}
		if s != fmt.Sprint("val ", i) {
			t.Errorf("Value %s is not correct, should be %s", s, fmt.Sprint("val ", i))
		}
	}()
	go func() {
		<-time.After(1 * time.Microsecond)
		i := 1
		s, err := sessionManager.Storage.Get(fmt.Sprint("key ", i))
		if err != nil {
			t.Errorf("No key %s in Storage ", fmt.Sprint("key ", i))
		}
		if s != fmt.Sprint("val ", i) {
			t.Errorf("Value %s is not correct, should be %s", s, fmt.Sprint("val ", i))
		}
	}()
	go func() {
		<-time.After(1 * time.Microsecond)
		i := 10
		s, err := sessionManager.Storage.Get(fmt.Sprint("key ", i))
		if err != nil {
			t.Errorf("No key %s in Storage ", fmt.Sprint("key ", i))
		}
		if s != fmt.Sprint("val ", i) {
			t.Errorf("Value %s is not correct, should be %s", s, fmt.Sprint("val ", i))
		}
	}()
	<-time.After(50 * time.Millisecond)
	// for i := 0; i < 300; i += 100 {
	// 	s, err := sessionManager.Storage.Get(fmt.Sprint("key ", i))
	// 	if err != nil {
	// 		t.Errorf("No key %s in Storage ", fmt.Sprint("key ", i))
	// 	}
	// 	if s != fmt.Sprint("val ", i) {
	// 		t.Errorf("Value %s is not correct, should be %s", s, fmt.Sprint("val ", i))
	// 	}
	// 	// time.After(1 * time.Microsecond)
	// }

}
