package proxy

import "errors"

type Server struct {
	data map[string]string
}

func (s *Server) Save(record Record, reply *bool) error {
	if s.data == nil {
		s.data = make(map[string]string)
	}
	s.data[record.Key] = record.Value
	*reply = true
	return nil
}
func (s *Server) Get(key string, reply *string) error {
	val, ok := s.data[key]
	if !ok {
		*reply = ""
		return errors.New("DB has no key" + key)
	}
	*reply = val
	return nil
}
