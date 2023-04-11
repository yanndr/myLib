package controllers

import (
	"api/api"
	"api/internal/services"
	"errors"
)

func handleServiceError(err error) (api.Response, error) {
	if errors.As(err, &services.DuplicateErr{}) {
		return api.Response{}, api.NewDuplicateErr(err.Error())
	} else if errors.As(err, &services.ValidationErr{}) {
		return api.Response{}, api.NewBadFormatErr(err.Error())
	} else if errors.As(err, &services.NotFoundErr{}) {
		return api.Response{}, api.NewNotFoundErr(err.Error())
	}
	return api.Response{}, err
}
