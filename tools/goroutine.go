package tools

import (
	"errors"
	"strings"
	"sync"
)

type TaskFunc func() error

func RunConcurrentTasks(tasks []TaskFunc) error {
	errChan := make(chan error, len(tasks))
	var wg sync.WaitGroup
	for _, task := range tasks {
		wg.Add(1)
		go func(ec chan error, f func() error) {
			defer wg.Done()
			if err := f(); err != nil {
				ec <- err
			}
		}(errChan, task)
	}

	wg.Wait()

	close(errChan)

	var errs []string
	for err := range errChan {
		errs = append(errs, err.Error())
	}
	if len(errs) > 0 {
		return errors.New(strings.Join(errs, "\n"))
	}
	return nil
}
