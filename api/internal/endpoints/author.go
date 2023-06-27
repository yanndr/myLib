package endpoints

import (
	"api/api"
	"api/internal/controllers"
	"api/internal/services"
	"net/http"
	"path"
)

func newAuthorsEndpoint(parentPath string, authorSvc services.AuthorService) *Route {
	c := controllers.AuthorController{
		AuthorService: authorSvc,
		BasePath:      path.Join(parentPath, api.AuthorsPath),
	}
	return &Route{
		Pattern: api.AuthorsPath,
		Actions: map[string]EndpointHandler{
			http.MethodGet:  c.GetAll,
			http.MethodPost: c.Create,
		},
		SubRoutes: []*Route{newAuthorEndpoint(c)},
	}
}

func newAuthorEndpoint(c controllers.AuthorController) *Route {
	return &Route{
		Pattern: "/{id}",
		Actions: map[string]EndpointHandler{
			http.MethodGet:    c.Get,
			http.MethodPut:    c.Update,
			http.MethodPatch:  c.PartialUpdate,
			http.MethodDelete: c.Delete,
		},
		SubRoutes: []*Route{&authorBooksEndpoint},
	}
}

var authorBooksEndpoint = Route{
	Pattern: api.BooksPath,
	Actions: map[string]EndpointHandler{
		http.MethodGet: notImplementedHandler,
	},
}
