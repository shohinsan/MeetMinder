package hashrepo

import "errors"

type testHashRepo struct{}

// Constructor for testHashRepo
func NewTestHashRepo() HashRepository {
	return &testHashRepo{}
}

func (m *testHashRepo) Hash(val string) (string, error) {
	if val == "" {
		return "", errors.New("error hashing value")
	}
	return val, nil
}

func (m *testHashRepo) Compare(hash, val string) error {
	return nil
}
