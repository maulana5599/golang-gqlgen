package graph

import (
	"github.com/maulana5599/golang-gqlgen/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos []*model.Todo
	link  []*model.Link
	user  []*model.User
}
