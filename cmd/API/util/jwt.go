package util

func VerifyJWT(token string) bool {
	if token == "" {
		return false
	}
	return true
}
