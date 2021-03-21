package main

import "github.com/chanced/esutil/search"

func main() {
	s, err := search.Build(func(s *search.Search) (*search.Search, error) {
		return s.SetDocValueFields().SetExplain().SetFrom(), nil
	})
}
