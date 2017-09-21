package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/pkg/rock/rocks/examples/authenticator"
)

var a authenticator.Authenticator

// ContosoUser represents *your* internal User type.
type ContosoUser struct {
	ID        int
	Username  string
	Password  string
	ManagerID int
	Groups    []string
}

func (u *ContosoUser) Username() string {
	return u.Username
}

func (u *ContosoUser) Password() string {
	return u.Password
}

// InfrastructureChange represents some business requirement
type InfrastructureChange struct {
	auth        authenticator.Authenticator
	Change      string
	Requester   ContosoUser
	RequestTime time.Time
}

// New returns a new InfrastuctureChange
func New(auth authenticator.Authenticator, requester ContosoUser, change string) *InfrastructureChange {

	return &InfrastructureChange{
		auth:        auth,
		Requester:   requester,
		Change:      string,
		RequestTime: time.Now(),
	}

}

// Execute is the business requirement that needs the Rock
func (ic *InfrastructureChange) Execute() error {

	if v := ic.auth.Authenticate(ic.Requester, u, p); v != nil {

		return errors.New("Authentication Failed")
	}
	fmt.Println("Performing requested change:", ic.Change)

}

// userFromContoso is an adapter function that turns an company (Contoso) user into
// the user type expected by the Authenticator rock
func userFromContoso(c ContosoUser) authenticator.User {
	var u authenticator.User
	u.Username = c.Username
	u.Password = c.Password
	return u
}
