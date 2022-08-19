package server

import . "MyCommunity/middlewares"

func (s *SEngine) UseMiddleWare() *SEngine {
	s.Use(Cors())
	return s
}
