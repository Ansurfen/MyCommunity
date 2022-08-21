package server

import . "go-auth/middlewares"

func (s *SEngine) UseMiddleWare() *SEngine {
	s.Use(Cors())
	return s
}
