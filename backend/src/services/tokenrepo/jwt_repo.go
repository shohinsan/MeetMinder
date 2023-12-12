package tokenrepo

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtRepo struct {
	signingMethod jwt.SigningMethod
	key           []byte
	lifetime      time.Duration
}

// Constructor for JWTTokenRepo
func NewJWTTokenRepo(signingMethod string, secretKey []byte, lifetime time.Duration) TokenRepository {
	return &jwtRepo{
		signingMethod: jwt.GetSigningMethod(signingMethod),
		key:           secretKey,
		lifetime:      lifetime,
	}
}

// mapToClaims is a helper function to convert a map[string]interface{} to jwt.MapClaims.
func mapToClaims(payload map[string]interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{}

	for key, value := range payload {
		claims[key] = value
	}

	return claims
}

// CreateToken implements TokenRepository.
func (m *jwtRepo) CreateToken(payload map[string]interface{}) (string, error) {
	// Set up claims
	claims := mapToClaims(payload)
	claims["exp"] = time.Now().Add(m.lifetime).Unix()

	// Create token
	token := jwt.NewWithClaims(m.signingMethod, claims)
	return token.SignedString(m.key)
}

// ParseToken implements TokenRepository.
func (m *jwtRepo) ParseToken(token string) (map[string]interface{}, error) {
	// Parse token
	signingMethods := []string{m.signingMethod.Alg()}

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return m.key, nil
	}, jwt.WithValidMethods(signingMethods), jwt.WithExpirationRequired())

	// Check if token is valid
	if err != nil {
		return nil, err
	}

	// Extract claims
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, jwt.ErrInvalidKey
	}

	return claims, nil
}
