package tempvalue

type value struct {
	userID string
}

var val = &value{}

func SetUserID(userID string) {
	val.userID = userID
}

func GetUserID() string {
	return val.userID
}
