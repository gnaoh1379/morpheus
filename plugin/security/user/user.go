package user

type User interface {
	GetUsername() string
	GetPassword() string
	GetPermissions() []string
	IsActivated() bool
	IsLocked() bool
	PasswordHasExpired() bool
}
