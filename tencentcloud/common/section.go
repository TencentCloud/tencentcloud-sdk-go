package common

import "sync"

type sections struct {
	contains map[string]*section
	mu       sync.RWMutex
}

func (ss *sections) section(name string) *section {
	ss.mu.Lock()
	defer ss.mu.Unlock()

	s, ok := ss.contains[name]
	if !ok {
		s = new(section)
		ss.contains[name] = s
	}
	return s
}

type section struct {
	content map[string]*value
	mu      sync.RWMutex
}

func (s *section) key(name string) *value {
	s.mu.Lock()
	defer s.mu.Unlock()

	v, ok := s.content[name]
	if !ok {
		v = new(value)
		s.content[name] = v
	}
	return v
}
