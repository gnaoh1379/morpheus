package user

type _User struct {
	username           string
	password           string
	permissions        []string
	activated          bool
	locked             bool
	passwordHasExpired bool
}

func (u _User) GetUsername() string {
	return u.username
}

func (u _User) GetPassword() string {
	return u.password
}

func (u _User) GetPermissions() []string {
	if u.permissions == nil {
		return []string{}
	}
	return u.permissions
}

func (u _User) IsActivated() bool {
	return u.activated
}

func (u _User) IsLocked() bool {
	return u.locked
}

func (u _User) PasswordHasExpired() bool {
	return u.passwordHasExpired
}
