package hashrepo

import "golang.org/x/crypto/bcrypt"

type bcryptRepo struct {
	cost int
}

// Constructor for BcryptRepo
func NewBcryptRepo(cost int) HashRepository {
	return &bcryptRepo{
		cost: cost,
	}
}

// Compare implements HashRepository.
func (m *bcryptRepo) Compare(hash string, val string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(val))
}

// Hash implements HashRepository.
func (m *bcryptRepo) Hash(val string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(val), m.cost)
	return string(hash), err
}
