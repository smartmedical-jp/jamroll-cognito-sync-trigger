package setting

var (
	userPoolID string
)

func GetUserPoolID() string {
	return userPoolID
}

func setUserPoolID(u string) {
	userPoolID = u
}
