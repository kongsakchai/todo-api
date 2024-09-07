package todo

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type handlerTestcase struct {
	mockContext
	title        string
	params       map[string]string
	payload      string
	stID         int64
	stResp       Todo
	stErr        error
	expectStatus int
	expectResp   any
}

func TestGetTodosHandler(t *testing.T) {
	testcases := []handlerTestcase{
		{
			title:        "should response all todo when storage not error",
			stResp:       Todo{ID: 1, Title: "Title 1", Description: "Description 1", Done: false},
			expectStatus: 200,
			expectResp:   []Todo{{ID: 1, Title: "Title 1", Description: "Description 1", Done: false}},
		},
		{
			title:        "should response error when storage error",
			stErr:        errors.New("error storage"),
			expectStatus: 500,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.title, func(t *testing.T) {
			mock := &mockStorage{}
			mock.todo = tc.stResp
			mock.err = tc.stErr
			handler := NewHandler(mock)

			ctx := mockContext{}
			handler.Todos(&ctx)

			assert.Equal(t, ctx.status, tc.expectStatus)
			assert.Equal(t, ctx.response, tc.expectResp)
		})
	}
}

func TestGetTodoHandler(t *testing.T) {
	testcases := []handlerTestcase{
		{
			title:        "should response error when id param empty",
			params:       map[string]string{"id": ""},
			expectStatus: 400,
		},
		{
			title:        "should response error when id param not integer",
			params:       map[string]string{"id": "abc"},
			expectStatus: 400,
		},
		{
			title:        "should response todo when storage not error",
			params:       map[string]string{"id": "1"},
			stResp:       Todo{ID: 1, Title: "Title 1", Description: "Description 1", Done: false},
			expectStatus: 200,
			expectResp:   Todo{ID: 1, Title: "Title 1", Description: "Description 1", Done: false},
		},
		{
			title:        "should response error when storage error",
			params:       map[string]string{"id": "1"},
			stErr:        errors.New("error storage"),
			expectStatus: 500,
		},
		{
			title:        "should response success but have message todo not found when todo not found",
			params:       map[string]string{"id": "1"},
			stErr:        ErrNotFound,
			expectStatus: 200,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.title, func(t *testing.T) {
			mock := &mockStorage{}
			mock.todo = tc.stResp
			mock.err = tc.stErr
			handler := NewHandler(mock)

			ctx := mockContext{params: tc.params}
			handler.Todo(&ctx)

			assert.Equal(t, ctx.status, tc.expectStatus)
			assert.Equal(t, ctx.response, tc.expectResp)
		})
	}
}

func TestCreateTodoHandler(t *testing.T) {
	testcases := []handlerTestcase{
		{
			title:        "should response error when payload invalid",
			payload:      `{"title": "Title 1"`,
			expectStatus: 400,
		},
		{
			title:        "should response success when storage not error",
			payload:      `{"title": "Title 2"}`,
			stID:         1,
			expectStatus: 201,
			expectResp:   Todo{ID: 1, Title: "Title 2", Description: "", Done: false},
		},
		{
			title:        "should response error when storage error",
			payload:      `{"title": "Title 3"}`,
			stErr:        errors.New("error storage"),
			expectStatus: 500,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.title, func(t *testing.T) {
			mock := &mockStorage{}
			mock.err = tc.stErr
			mock.id = tc.stID
			handler := NewHandler(mock)

			ctx := mockContext{payload: string(tc.payload)}
			handler.Create(&ctx)

			assert.Equal(t, ctx.status, tc.expectStatus)
			assert.Equal(t, ctx.response, tc.expectResp)
		})
	}
}

func TestUpdateTodoHandler(t *testing.T) {
	testcases := []handlerTestcase{
		{
			title:        "should response error when id param empty",
			params:       map[string]string{"id": ""},
			expectStatus: 400,
		},
		{
			title:        "should response error when id param not integer",
			params:       map[string]string{"id": "abc"},
			expectStatus: 400,
		},
		{
			title:        "should response error when payload invalid",
			params:       map[string]string{"id": "1"},
			payload:      `{"title": "Title 1"`,
			expectStatus: 400,
		},
		{
			title:        "should response success when storage not error",
			params:       map[string]string{"id": "1"},
			payload:      `{"id":2,"title": "Title 2"}`,
			expectResp:   Todo{ID: 1, Title: "Title 2", Description: "", Done: false},
			expectStatus: 200,
		},
		{
			title:        "should response error when storage error",
			params:       map[string]string{"id": "1"},
			payload:      `{"id":2,"title": "Title 3"}`,
			stErr:        errors.New("error storage"),
			expectStatus: 500,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.title, func(t *testing.T) {
			mock := &mockStorage{}
			mock.err = tc.stErr
			handler := NewHandler(mock)

			ctx := mockContext{params: tc.params, payload: string(tc.payload)}
			handler.Update(&ctx)

			assert.Equal(t, ctx.status, tc.expectStatus)
			assert.Equal(t, ctx.response, tc.expectResp)
		})
	}
}

func TestDeleteTodoHandler(t *testing.T) {

	testcases := []handlerTestcase{
		{
			title:        "should response error when id param empty",
			params:       map[string]string{"id": ""},
			expectStatus: 400,
		},
		{
			title:        "should response error when id param not integer",
			params:       map[string]string{"id": "abc"},
			expectStatus: 400,
		},
		{
			title:        "should response success when storage not error",
			params:       map[string]string{"id": "1"},
			expectStatus: 200,
		},
		{
			title:        "should response error when storage error",
			params:       map[string]string{"id": "1"},
			stErr:        errors.New("error storage"),
			expectStatus: 500,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.title, func(t *testing.T) {
			mock := &mockStorage{}
			mock.err = tc.stErr
			handler := NewHandler(mock)

			ctx := mockContext{params: tc.params}
			handler.Delete(&ctx)

			assert.Equal(t, ctx.status, tc.expectStatus)
		})
	}
}
