package interfaces

import (
	"context"
	"net/http"

	"github.com/jzavala-globant/bookstore-rest-api-go/internal/models"
)

type BookstoreController interface {
	ListBooks(w http.ResponseWriter, r *http.Request)
}

type BookstoreService interface {
	ListBooks(context.Context) (models.APIResponse, error)
}

type BookstoreRepository interface {
	ListBooks(context.Context) (models.Book, error)
}
