package handler_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	todo "github.com/govader/todo.go-vader.io/api"
	"github.com/govader/todo.go-vader.io/api/http/handler"
	echo "github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var now = time.Date(2019, time.July, 9, 22, 12, 56, 0, time.UTC)

type MockStore []todo.Task

func (mock *MockStore) Create(content string) (*todo.Task, error) {
	return &todo.Task{
		ID:        "9b66205a-0988-4f1e-894c-4563c14de404",
		Content:   content,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

func (mock *MockStore) Get(ID string) (*todo.Task, error) {
	return nil, nil
}

func (mock *MockStore) MarkAsDone(ID string) (*todo.Task, error) {
	return nil, nil
}

func (mock *MockStore) All() todo.TaskQuerier {
	return nil
}

func (mock *MockStore) Done() todo.TaskQuerier {
	return nil
}

func (mock *MockStore) NotDone() todo.TaskQuerier {
	return nil
}

func TestTaskHandler(t *testing.T) {
	mockStore := &MockStore{}

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"content":"make it happen"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &handler.TaskHandler{Store: mockStore}

	t.Run("Add", func(t *testing.T) {
		if assert.NoError(t, h.Add(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			assert.JSONEq(t, `{
				"id":"9b66205a-0988-4f1e-894c-4563c14de404",
				"content":"make it happen",
				"done":false,
				"created_at":"2019-07-09T22:12:56Z",
				"updated_at":"2019-07-09T22:12:56Z"
			}`, rec.Body.String())
		}
	})
}
