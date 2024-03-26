package api

import (
	"context"

	"github.com/mtfedev/contacts-menager/types"
	"go.mongodb.org/mongo-driver/bson"
)

type Dropper interface {
	Drop(context.Context) error
}
type UserDO interface {
	Dropper
	GetUserByEmail(context.Context, string) (*types.User, error)
	GetUserByID(context.Context, string) (*types.User, error)
	GetUsers(context.Context) ([]*types.User, error)
	InsertUser(context.Context, *types.User) (*types, error)
	DeleteUser(context.Context, string) error
	UpdateUser(ctx context.Context, filter bson.M, params types.UpateUserParams) error
}

//  - add new contact
//delete one contact
//get one contact
//get all contacts
//modify one contact
