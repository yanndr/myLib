package client

import (
	"api/api"
	"api/model"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"path"
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

func (cli *Client) GetAuthors(ctx context.Context, lastname string) ([]model.Author, error) {
	queryParam := fmt.Sprintf("?lastname=%v", url.PathEscape(lastname))
	apiResponse, err := cli.get(ctx, api.AuthorsPath+queryParam)
	if err != nil {
		return nil, err
	}

	var authors []model.Author
	err = json.Unmarshal(apiResponse.Content, &authors)
	if err != nil {
		return nil, err
	}

	return authors, nil
}

func (cli *Client) DeleteAuthor(ctx context.Context, id int64) error {
	_, err := cli.delete(ctx, path.Join(api.AuthorsPath, fmt.Sprint(id)))
	if err != nil {
		return err
	}

	return nil
}
