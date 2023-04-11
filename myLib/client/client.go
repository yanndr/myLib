package client

import (
	"api/api"
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

func (cli *Client) get(ctx context.Context, path string) (api.RawResponse, error) {
	return cli.doRequest(ctx, http.MethodGet, path, nil)
}

func (cli *Client) post(ctx context.Context, path string, data any) (api.RawResponse, error) {
	return cli.doRequest(ctx, http.MethodPost, path, data)
}

func (cli *Client) put(ctx context.Context, path string, data any) (api.RawResponse, error) {
	return cli.doRequest(ctx, http.MethodPut, path, data)
}

func (cli *Client) patch(ctx context.Context, path string, data any) (api.RawResponse, error) {
	return cli.doRequest(ctx, http.MethodPatch, path, data)
}

func (cli *Client) delete(ctx context.Context, path string) (api.RawResponse, error) {
	return cli.doRequest(ctx, http.MethodDelete, path, nil)
}

func (cli *Client) doRequest(ctx context.Context, method, path string, data any) (api.RawResponse, error) {
	var body io.Reader
	if data != nil {
		bodyData, err := json.Marshal(data)
		if err != nil {
			return api.RawResponse{}, err
		}
		body = bytes.NewBuffer(bodyData)
	}
	req, err := http.NewRequest(method, cli.BaseAddress+path, body)

	apiResponse, err := cli.do(ctx, req)
	if err != nil {
		return api.RawResponse{}, err
	}
	if apiResponse.Status == api.ErrorStatus {

		switch apiResponse.ErrorCode {
		case api.DuplicateErrorCode:
			return api.RawResponse{}, DuplicateResourceErr
		case api.NotFoundErrorCode:
			return api.RawResponse{}, ResourceNotFoundErr
		}
		switch apiResponse.StatusCode {
		case http.StatusNotFound:
			return api.RawResponse{}, ResourceNotFoundErr
		default:
			return api.RawResponse{}, fmt.Errorf(apiResponse.Error)
		}
	}
	return apiResponse, nil
}

func (cli *Client) do(ctx context.Context, req *http.Request) (api.RawResponse, error) {
	req = req.WithContext(ctx)
	resp, err := cli.Client.Do(req)
	if err != nil {
		return api.RawResponse{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return api.RawResponse{}, err
	}

	apiResponse := api.RawResponse{}
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return api.RawResponse{}, err
	}
	return apiResponse, nil

}

func (cli *Client) GetVersion(ctx context.Context) (string, error) {
	apiResponse, err := cli.get(ctx, "/")
	if err != nil {
		return "", err
	}

	var version api.Information
	err = json.Unmarshal(apiResponse.Content, &version)
	if err != nil {
		return "", err
	}

	return version.APIVersion, nil
}
