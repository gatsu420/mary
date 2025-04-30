package tempvalue

type value struct {
	userID        string
	calledMethods map[string]struct{}
}

var val = &value{
	calledMethods: map[string]struct{}{},
}

func SetUserID(userID string) {
	val.userID = userID
}

func GetUserID() string {
	return val.userID
}

func SetCalledMethods(method string) {
	val.calledMethods[method] = struct{}{}
}

func GetCalledMethods() map[string]struct{} {
	return val.calledMethods
}

func FlushCalledMethods() {
	val.calledMethods = map[string]struct{}{}
}
