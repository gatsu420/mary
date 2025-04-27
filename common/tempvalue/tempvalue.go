package tempvalue

type value struct {
	userID string
}

func NewValue() *value {
	return &value{}
}

func (v *value) SetUserID(userID string) {
	v.userID = userID
}

func (v *value) GetUserID() string {
	return v.userID
}
