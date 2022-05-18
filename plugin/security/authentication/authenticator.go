package authentication

import "context"

type Authenticator interface {
	Authenticate(ctx context.Context)
}
