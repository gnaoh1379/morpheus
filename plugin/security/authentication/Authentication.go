package authentication

import (
	"github.com/nguyenhaihoang/morpheus/plugin/security/user"
)

type Authentication interface {
	GetUser() user.User
	IsAuthenticated() bool
}
