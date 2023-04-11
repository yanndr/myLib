package client

import (
	"api/api"
	"api/model"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var ResourceNotFoundErr = fmt.Errorf("resource not found")
var DuplicateResourceErr = fmt.Errorf("duplicate resource found")

type Client struct {
	Client      *http.Client
	BaseAddress string
}

func NewClient(baseAddress string, client *http.Client) *Client {
	return &Client{
		Client:      client,
		BaseAddress: baseAddress,
	}
}

func (cli *Client) get(ctx context.Context, path string) (model.RawResponse, error) {
	return cli.doRequest(ctx, http.MethodGet, path, nil)
}

func (cli *Client) post(ctx context.Context, path string, data any) (model.RawResponse, error) {
	return cli.doRequest(ctx, http.MethodPost, path, data)
}

func (cli *Client) put(ctx context.Context, path string, data any) (model.RawResponse, error) {
	return cli.doRequest(ctx, http.MethodPut, path, data)
}

func (cli *Client) patch(ctx context.Context, path string, data any) (model.RawResponse, error) {
	return cli.doRequest(ctx, http.MethodPatch, path, data)
}

func (cli *Client) delete(ctx context.Context, path string) (model.RawResponse, error) {
	return cli.doRequest(ctx, http.MethodDelete, path, nil)
}

func (cli *Client) doRequest(ctx context.Context, method, path string, data any) (model.RawResponse, error) {
	var body io.Reader
	if data != nil {
		bodyData, err := json.Marshal(data)
		if err != nil {
			return model.RawResponse{}, err
		}
		body = bytes.NewBuffer(bodyData)
	}
	req, err := http.NewRequest(method, cli.BaseAddress+path, body)

	apiResponse, err := cli.do(ctx, req)
	if err != nil {
		return model.RawResponse{}, err
	}
	if apiResponse.Status == model.ErrorStatus {

		switch apiResponse.ErrorCode {
		case api.DuplicateErrorCode:
			return model.RawResponse{}, DuplicateResourceErr
		case api.NotFoundErrorCode:
			return model.RawResponse{}, ResourceNotFoundErr
		}
		switch apiResponse.StatusCode {
		case http.StatusNotFound:
			return model.RawResponse{}, ResourceNotFoundErr
		default:
			return model.RawResponse{}, fmt.Errorf(apiResponse.Error)
		}
	}
	return apiResponse, nil
}

func (cli *Client) do(ctx context.Context, req *http.Request) (model.RawResponse, error) {
	req = req.WithContext(ctx)
	resp, err := cli.Client.Do(req)
	if err != nil {
		return model.RawResponse{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.RawResponse{}, err
	}

	apiResponse := model.RawResponse{}
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return model.RawResponse{}, err
	}
	return apiResponse, nil

}

func (cli *Client) GetVersion(ctx context.Context) (string, error) {
	apiResponse, err := cli.get(ctx, "/")
	if err != nil {
		return "", err
	}

	var version model.APIInformation
	err = json.Unmarshal(apiResponse.Content, &version)
	if err != nil {
		return "", err
	}

	return version.APIVersion, nil
}
