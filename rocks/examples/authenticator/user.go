package authenticator

import (
	"reflect"

	"github.com/pkg/rock/rocks"
)

func init() {
}

var version = "0.0.1"

type Authorizer interface {
	Authorize(u UserPassGrouper) bool
}

// A UserPasser is a type that can return a username and password
type UserPassGrouper interface {
	Username() string
	Password() string
	Groups() []string
}

func NewNoop() Authorizer {
	var a NoopAuthorizer
	var ai Authorizer
	a.Rock = rocks.New(reflect.TypeOf(a), reflect.TypeOf(ai), "User", "Brian Ketelsen", "MIT", version)
	return &a
}

type NoopAuthorizer struct {
	rocks.Rock
}

func (n *NoopAuthorizer) Authorize(u UserPassGrouper) bool {
	return true
}

func NewGroup(authorizedGroups []string) Authorizer {
	var a GroupAuthorizer
	a.AuthorizedGroups = authorizedGroups
	var ai Authorizer

	a.Rock = rocks.New(reflect.TypeOf(a), reflect.TypeOf(ai), "User", "Brian Ketelsen", "MIT", version)
	return &a
}

type GroupAuthorizer struct {
	rocks.Rock
	AuthorizedGroups []string
}

func (ga *GroupAuthorizer) includesGroup(g string) bool {

}

func (ga *GroupAuthorizer) Authenticate(u UserPassGrouper) bool {
	for _, g := range u.Groups() {
		ok := ga.includesGroup(g)
	}

}
