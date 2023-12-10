package tokenrepo

import "errors"

type testTokenRepo struct{}

// Constructor for testTokenRepo
func NewTestTokenRepo() TokenRepository {
	return &testTokenRepo{}
}

func (m *testTokenRepo) CreateToken(payload map[string]interface{}) (string, error) {
	if payload["error"] != nil {
		return "", errors.New("error creating token")
	}

	return "token", nil
}

func (m *testTokenRepo) ParseToken(token string) (map[string]interface{}, error) {
	if token == "error" {
		return nil, errors.New("error parsing token")
	}

	claims := make(map[string]interface{})
	claims["id"] = 1
	claims["username"] = "test"

	return claims, nil
}
