package search

import "fmt"

type Builder struct {
	Defaults *Search
}

var DefaultBuilder = Builder{}

func Build(fn func(s *Search) (*Search, error)) (*Search, error) {
	return DefaultBuilder.Build(fn)
}
func (b Builder) Build(fn func(s *Search) (*Search, error)) (*Search, error) {
	errCh := make(chan error, 2)
	var err error
	var s *Search
	if b.Defaults != nil {
		s = b.Defaults.Clone()
	} else {
		s = NewSearch()
	}

	s, err = b.try(fn, s, errCh)
	if err != nil {
		return nil, err
	}
	err = <-errCh
	if err != nil {
		return nil, err
	}
	return s, nil
}
func (b Builder) try(fn func(s *Search) (*Search, error), s *Search, errCh chan error) (*Search, error) {
	defer func(errCh chan error) {
		if err := recover(); err != nil {
			if rErr, ok := err.(error); ok {
				errCh <- rErr
			} else {
				errCh <- fmt.Errorf("%s", err)
			}
			close(errCh)
		} else {
			close(errCh)
		}
	}(errCh)
	s, err := fn(s)
	if err != nil {
		errCh <- err
	}
	return s, err
}
