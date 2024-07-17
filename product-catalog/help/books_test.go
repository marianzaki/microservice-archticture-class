package books

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestPostBook(t *testing.T) {
	// Mock the server
	r := SetUpRouter()
	r.POST("/books", GetBooks)

	// Prepare the request body
	bookId := xid.New().String()
	book := books_data.Book{
		ID:                bookId,
		Name:              "test_name",
		Author:            "test_author",
		YearOfPublication: 2020,
		Price:             2.99,
	}

	// Prepare the request
	jsonValue, _ := json.Marshal(book)
	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(jsonValue))

	// Send the request to the mocked server
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Make assertion
	assert.Equal(t, http.StatusCreated, w.Code)
}
