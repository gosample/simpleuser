package storage

import (
	"github.com/yaronsumel/simpleuser/server/storage/handlers/mongo"
	"github.com/yaronsumel/simpleuser/server/user"
)

// returns mongodb as default
func NewHandler() Interface {
	return mongo.NewHandler()
}

type Interface interface {
	Insert(user *user.Object) error
}
