package hashrepo

type HashRepository interface {
	// Hash takes a string and returns a hashed version of it
	Hash(val string) (string, error)

	// Compare takes a hash and a string and returns an error if they don't match
	Compare(hash, val string) error
}
