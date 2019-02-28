package db

const (
	KeyUsernameSet = "guard:account:usernames"
)

func KeyUsernameToken(username string) (key string) {
	return "guard:account:" + username + ":token"
}
