package user

type Option func(u *_User)

func New(opts ...Option) User {
	var u _User
	for _, opt := range opts {
		opt(&u)
	}
	return &u
}

func WithUsername(username string) Option {
	return func(u *_User) {
		u.username = username
	}
}

func WithPassword(password string) Option {
	return func(u *_User) {
		u.password = password
	}
}

func WithPermissions(permissions ...string) Option {
	return func(u *_User) {
		u.permissions = permissions
	}
}

func WithLocked() Option {
	return func(u *_User) {
		u.locked = true
	}
}

func WithPasswordHasExpired() Option {
	return func(u *_User) {
		u.passwordHasExpired = true
	}
}
