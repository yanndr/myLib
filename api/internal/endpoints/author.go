package endpoints

import (
	"api/internal/controllers"
	"api/internal/services"
	"net/http"
	"path"
)

func newAuthorsEndpoint(parentPath string, authorSvc services.AuthorService) *Route {
	c := controllers.AuthorController{
		AuthorService: authorSvc,
		BasePath:      path.Join(parentPath, AuthorsPath),
	}
	return &Route{
		Pattern: AuthorsPath,
		Actions: map[string]EndpointHandler{
			http.MethodGet:  notImplementedHandler,
			http.MethodPost: c.Create,
		},
		SubRoutes: []*Route{&authorEndpoint},
	}
}

var authorEndpoint = Route{
	Pattern: "/{id}",
	Actions: map[string]EndpointHandler{
		http.MethodGet:    notImplementedHandler,
		http.MethodPut:    notImplementedHandler,
		http.MethodPatch:  notImplementedHandler,
		http.MethodDelete: notImplementedHandler,
	},
	SubRoutes: []*Route{&authorBooksEndpoint},
}

var authorBooksEndpoint = Route{
	Pattern: BooksPath,
	Actions: map[string]EndpointHandler{
		http.MethodGet: notImplementedHandler,
	},
}
