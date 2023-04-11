package controllers

import (
	"api/internal/model"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/golang/mock/gomock"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

var c *AuthorController
var m *MockAuthorService

func setupTest(t testing.TB) func(t testing.TB) {
	log.Println("setup test")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m = NewMockAuthorService(ctrl)

	c = &AuthorController{
		AuthorService: m,
		BasePath:      "/v1/authors",
	}

	return func(t testing.TB) {
		log.Println("teardown test")

	}
}

func TestAuthorController_Create(t *testing.T) {
	authorRequest := model.CreateUpdateAuthorRequest{AuthorBase: model.AuthorBase{LastName: "test"}}
	tests := []struct {
		name       string
		body       any
		want       model.APIResponse
		wantErr    bool
		wantSvcErr bool
	}{
		{name: "Success", body: authorRequest, want: model.NewCreatedResponse("/v1/authors/1"), wantErr: false},
		{name: "service exception", body: authorRequest, want: model.APIResponse{}, wantErr: true, wantSvcErr: true},
		{name: "Nil input", body: nil, want: model.APIResponse{}, wantErr: true},
		{name: "bad input", body: struct{ test string }{test: "hello"}, want: model.APIResponse{}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			teardownTest := setupTest(t)
			defer teardownTest(t)

			if !tt.wantErr {
				m.EXPECT().Create(gomock.Any(), gomock.Any()).Return(int64(1), nil).Times(1)
			}

			if tt.wantSvcErr {
				m.EXPECT().Create(gomock.Any(), gomock.Any()).Return(int64(-1), fmt.Errorf("expected error")).Times(1)
			}

			var body io.Reader
			if tt.body != nil {
				bodyData, _ := json.Marshal(tt.body)
				body = bytes.NewBuffer(bodyData)
			}

			r := httptest.NewRequest(http.MethodGet, "/", body)

			got, err := c.Create(r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}
