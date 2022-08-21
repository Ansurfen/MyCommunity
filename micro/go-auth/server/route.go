package server

import "go-auth/routes"

func (s *SEngine) UseRouter() *SEngine {

	s.POST("/jwt/release", routes.ReleaseJWT)
	s.POST("/jwt/parser", routes.ParserToken)

	s.POST("/crypto/encode/aes", routes.EncodeWithAES)
	s.POST("/crypto/encode/md5", routes.EncodeWithMD5)
	s.POST("/crypto/encode/sha256", routes.EncodeWithSHA256)
	
	s.POST("/crypto/decode/aes", routes.DecodeWithAES)

	return s
}
