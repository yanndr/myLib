package client

import (
	"api/api"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"path"
)

func (cli *Client) CreateAuthor(ctx context.Context, lastName string, firstName, middleName string) error {
	req := api.CreateUpdateAuthorRequest{
		AuthorBase: api.AuthorBase{
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

func (cli *Client) GetAuthors(ctx context.Context, lastname string) ([]api.Author, error) {
	queryParam := fmt.Sprintf("?lastname=%v", url.PathEscape(lastname))
	apiResponse, err := cli.get(ctx, api.AuthorsPath+queryParam)
	if err != nil {
		return nil, err
	}

	var authors []api.Author
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

func (cli *Client) UpdateAuthor(ctx context.Context, id int64, lastname, firstname, middleName string, modifiedLastname, modifiedFirstname, modifiedMiddleName bool) error {
	patchRequest := api.PatchAuthorRequest{
		ModifiedLastName:   modifiedLastname,
		ModifiedFirstName:  modifiedFirstname,
		ModifiedMiddleName: modifiedMiddleName,
		AuthorBase: api.AuthorBase{
			LastName:   lastname,
			FirstName:  firstname,
			MiddleName: middleName,
		},
	}
	_, err := cli.patch(ctx, path.Join(api.AuthorsPath, fmt.Sprintf("%v", id)), patchRequest)
	if err != nil {
		return err
	}

	return nil
}
