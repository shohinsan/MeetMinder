package tokenrepo

type TokenRepository interface {
	// CreateToken takes a payload and returns a token
	CreateToken(payload map[string]interface{}) (string, error)

	// ParseToken takes a token and returns a payload
	ParseToken(token string) (map[string]interface{}, error)
}
