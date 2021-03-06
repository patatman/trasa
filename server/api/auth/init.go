package auth

import (
	"github.com/seknox/trasa/server/global"
	"github.com/seknox/trasa/server/models"
)

//InitStore initialises package state
func InitStore(state *global.State) {
	Store = authStore{state}
}

//InitStoreMock will init mock state of this package
func InitStoreMock() *authMock {
	lmock := new(authMock)
	Store = lmock
	return lmock
}

//Store is the package state variable which contains database connections
var Store adapter

type authStore struct {
	*global.State
}

type adapter interface {
	GetLoginDetails(trasaID, orgDomain string) (*models.UserWithPass, error)
	Logout(sessionID string) error
}
