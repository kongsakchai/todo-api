package todo

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTodosHandler(t *testing.T) {
	type testcase struct {
		title          string
		stResp         Todo
		stErr          error
		expectedStatus int
		expectedResp   []Todo
	}

	testcases := []testcase{
		{
			title:          "should response all todo when storage not error",
			stResp:         Todo{ID: 1, Title: "Title 1", Description: "Description 1", Done: false},
			expectedStatus: 200,
			expectedResp:   []Todo{{ID: 1, Title: "Title 1", Description: "Description 1", Done: false}},
		},
		{
			title:          "should response error when storage error",
			stErr:          errors.New("error storage"),
			expectedStatus: 500,
			expectedResp:   nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.title, func(t *testing.T) {
			mock := &mockStorage{}
			mock.todo = tc.stResp
			mock.err = tc.stErr
			handler := NewHandler(mock)

			ctx := mockContext[[]Todo]{}
			handler.Todos(&ctx)

			assert.Equal(t, tc.expectedStatus, ctx.status)
			assert.Equal(t, tc.expectedResp, ctx.response)
		})
	}
}

func TestGetTodoHandler(t *testing.T) {
	type testcase struct {
		title          string
		params         map[string]string
		stResp         Todo
		stErr          error
		expectedStatus int
		expectedResp   Todo
	}

	testcases := []testcase{
		{
			title:          "should response error when id param empty",
			params:         map[string]string{"id": ""},
			expectedStatus: 400,
		},
		{
			title:          "should response error when id param not integer",
			params:         map[string]string{"id": "abc"},
			expectedStatus: 400,
		},
		{
			title:          "should response todo when storage not error",
			params:         map[string]string{"id": "1"},
			stResp:         Todo{ID: 1, Title: "Title 1", Description: "Description 1", Done: false},
			expectedStatus: 200,
			expectedResp:   Todo{ID: 1, Title: "Title 1", Description: "Description 1", Done: false},
		},
		{
			title:          "should response error when storage error",
			params:         map[string]string{"id": "1"},
			stErr:          errors.New("error storage"),
			expectedStatus: 500,
		},
		{
			title:          "should response success but have message todo not found when todo not found",
			params:         map[string]string{"id": "1"},
			stErr:          ErrNotFound,
			expectedStatus: 200,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.title, func(t *testing.T) {
			mock := &mockStorage{}
			mock.todo = tc.stResp
			mock.err = tc.stErr
			handler := NewHandler(mock)

			ctx := mockContext[Todo]{params: tc.params}
			handler.Todo(&ctx)

			assert.Equal(t, tc.expectedStatus, ctx.status)
			assert.Equal(t, tc.expectedResp, ctx.response)
		})
	}
}

func TestCreateTodoHandler(t *testing.T) {
	type testcase struct {
		title          string
		payload        string
		stID           int64
		stErr          error
		expectedStatus int
		expectedResp   Todo
	}

	testcases := []testcase{
		{
			title:          "should response error when payload invalid",
			payload:        `{"title": "Title 1"`,
			expectedStatus: 400,
		},
		{
			title:          "should response success when storage not error",
			payload:        `{"title": "Title 2"}`,
			stID:           1,
			expectedStatus: 201,
		},
		{
			title:          "should response error when storage error",
			payload:        `{"title": "Title 3"}`,
			stErr:          errors.New("error storage"),
			expectedStatus: 500,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.title, func(t *testing.T) {
			mock := &mockStorage{}
			mock.err = tc.stErr
			mock.id = tc.stID
			handler := NewHandler(mock)

			ctx := mockContext[Todo]{payload: string(tc.payload)}
			handler.Create(&ctx)

			assert.Equal(t, tc.expectedStatus, ctx.status)
			assert.Equal(t, tc.expectedResp, ctx.response)
		})
	}
}

func TestUpdateTodoHandler(t *testing.T) {
	type testcase struct {
		title          string
		params         map[string]string
		payload        string
		stErr          error
		expectedStatus int
		expectedResp   Todo
	}

	testcases := []testcase{
		{
			title:          "should response error when id param empty",
			params:         map[string]string{"id": ""},
			expectedStatus: 400,
		},
		{
			title:          "should response error when id param not integer",
			params:         map[string]string{"id": "abc"},
			expectedStatus: 400,
		},
		{
			title:          "should response error when payload invalid",
			params:         map[string]string{"id": "1"},
			payload:        `{"title": "Title 1"`,
			expectedStatus: 400,
		},
		{
			title:          "should response success when storage not error",
			params:         map[string]string{"id": "1"},
			payload:        `{"id":2,"title": "Title 2"}`,
			expectedResp:   Todo{ID: 1, Title: "Title 2", Description: "", Done: false},
			expectedStatus: 200,
		},
		{
			title:          "should response error when storage error",
			params:         map[string]string{"id": "1"},
			payload:        `{"id":2,"title": "Title 3"}`,
			stErr:          errors.New("error storage"),
			expectedStatus: 500,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.title, func(t *testing.T) {
			mock := &mockStorage{}
			mock.err = tc.stErr
			handler := NewHandler(mock)

			ctx := mockContext[Todo]{params: tc.params, payload: string(tc.payload)}
			handler.Update(&ctx)

			assert.Equal(t, tc.expectedStatus, ctx.status)
			assert.Equal(t, tc.expectedResp, ctx.response)
		})
	}
}

func TestDeleteTodoHandler(t *testing.T) {
	type testcase struct {
		title          string
		params         map[string]string
		stErr          error
		expectedStatus int
	}

	testcases := []testcase{
		{
			title:          "should response error when id param empty",
			params:         map[string]string{"id": ""},
			expectedStatus: 400,
		},
		{
			title:          "should response error when id param not integer",
			params:         map[string]string{"id": "abc"},
			expectedStatus: 400,
		},
		{
			title:          "should response success when storage not error",
			params:         map[string]string{"id": "1"},
			expectedStatus: 200,
		},
		{
			title:          "should response error when storage error",
			params:         map[string]string{"id": "1"},
			stErr:          errors.New("error storage"),
			expectedStatus: 500,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.title, func(t *testing.T) {
			mock := &mockStorage{}
			mock.err = tc.stErr
			handler := NewHandler(mock)

			ctx := mockContext[any]{params: tc.params}
			handler.Delete(&ctx)

			assert.Equal(t, tc.expectedStatus, ctx.status)
		})
	}
}
