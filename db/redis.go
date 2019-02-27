package db

const (
	KeyUsernameSet = "guard:account:usernames"
)

func KeyUsernameSalt(username string) (key string) {
	return "guard:account:" + username + ":salt"
}
