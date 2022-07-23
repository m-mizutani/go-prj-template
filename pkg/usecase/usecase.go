package usecase

import (
	"github.com/m-mizutani/go-prj-template/pkg/infra"
	"github.com/m-mizutani/go-prj-template/pkg/model"
)

type Usecase struct {
	clients *infra.Clients
}

func New(clients *infra.Clients) *Usecase {
	return &Usecase{
		clients: clients,
	}
}

func (x *Usecase) CreateUser(ctx *model.Context, req *model.CreateUserRequest) (*model.CreateUserResponse, error) {
	user := model.NewUser(req.Name)

	if err := x.clients.DB().SaveUser(ctx, user); err != nil {
		return nil, err
	}

	return &model.CreateUserResponse{
		ID: user.ID,
	}, nil
}
