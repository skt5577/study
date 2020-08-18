package singleton

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestGetInstance(t *testing.T) {
	for i := 0; i < 1000; i++ {
		repository1, repository2, repository3, repository4 := create()
		assert.Same(t, repository1, repository2)
		assert.Same(t, repository1, repository3)
		assert.Same(t, repository1, repository4)
	}
}

func create() (repository1, repository2, repository3, repository4 ParamRepository) {
	var wait sync.WaitGroup
	wait.Add(2)
	go func() {
		repository1 = GetInstance()
		wait.Done()
	}()
	go func() {
		repository2 = GetInstance()
		wait.Done()
	}()

	repository3 = GetInstance()
	repository4 = GetInstance()

	wait.Wait()
	return
}