package client

import (
	"api/api"
	"api/model"
	"context"
)

func (cli *Client) CreateAuthor(ctx context.Context, lastName string, firstName, middleName string) error {
	req := model.CreateUpdateAuthorRequest{
		AuthorBase: model.AuthorBase{
			LastName:   lastName,
			FirstName:  firstName,
			MiddleName: middleName,
		},
	}
	_, err := cli.post(ctx, api.AuthorsPath, req)
	if err != nil {
		return err
	}

	return nil
}
