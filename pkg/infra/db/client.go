package db

import "github.com/m-mizutani/go-prj-template/pkg/model"

type Client struct {
	dbName string
}

func New(dbName string) *Client {
	return &Client{dbName: dbName}
}

func (x *Client) SaveUser(ctx *model.Context, user *model.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	// DBに保存するコード
	return nil
}
