package singleton

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestGetInstanceLazy(t *testing.T) {
	for i := 0; i < 1000; i++ {
		repository1, repository2, repository3, repository4 := createRepository()
		assert.Same(t, repository1, repository2)
		assert.Same(t, repository1, repository3)
		assert.Same(t, repository1, repository4)
	}
}

func createRepository() (repository1, repository2, repository3, repository4 ParamRepository) {
	var wait sync.WaitGroup
	wait.Add(2)
	go func() {
		repository1 = GetInstanceLazy()
		wait.Done()
	}()
	go func() {
		repository2 = GetInstanceLazy()
		wait.Done()
	}()

	repository3 = GetInstanceLazy()
	repository4 = GetInstanceLazy()

	wait.Wait()
	return
}
