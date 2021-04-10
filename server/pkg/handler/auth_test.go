package handler

import (
	"bytes"
	"errors"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/luckyshmo/api-example/models"
	"github.com/luckyshmo/api-example/pkg/service"
	service_mocks "github.com/luckyshmo/api-example/pkg/service/mocks"
	"github.com/magiconair/properties/assert"
)

func TestHandler_signUp(t *testing.T) {
	// Init Test Table
	type mockBehavior func(r *service_mocks.MockAuthorization, user models.User)

	tests := []struct {
		name                 string
		inputBody            string
		inputUser            models.User
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"username": "username", "name": "Test Name", "password": "qwerty"}`,
			inputUser: models.User{
				Username: "username",
				Name:     "Test Name",
				Password: "qwerty",
			},
			mockBehavior: func(r *service_mocks.MockAuthorization, user models.User) {
				r.EXPECT().CreateUser(user).Return(uuid.Nil, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":"` + fmt.Sprint(uuid.Nil) + `"}`,
		},
		{
			name:                 "Wrong Input",
			inputBody:            `{"username": "username"}`,
			inputUser:            models.User{},
			mockBehavior:         func(r *service_mocks.MockAuthorization, user models.User) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "Service Error",
			inputBody: `{"username": "username", "name": "Test Name", "password": "qwerty"}`,
			inputUser: models.User{
				Username: "username",
				Name:     "Test Name",
				Password: "qwerty",
			},
			mockBehavior: func(r *service_mocks.MockAuthorization, user models.User) {
				r.EXPECT().CreateUser(user).Return(uuid.Nil, errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)

			repo := service_mocks.NewMockAuthorization(c)
			test.mockBehavior(repo, test.inputUser)

			services := &service.Service{Authorization: repo}
			handler := Handler{services}

			// Init Endpoint
			r := gin.New()
			r.POST("/sign-up", handler.signUp)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up",
				bytes.NewBufferString(test.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestHandler_signIn(t *testing.T) {
	// Init Test Table
	type mockBehavior func(r *service_mocks.MockAuthorization, user signInInput)

	tests := []struct {
		name                 string
		inputBody            string
		inputUser            signInInput
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"username": "username", "password": "qwerty"}`,
			inputUser: signInInput{
				Username: "username",
				Password: "qwerty",
			},
			mockBehavior: func(r *service_mocks.MockAuthorization, user signInInput) {
				r.EXPECT().GenerateToken(user.Username, user.Password).Return("", nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"token":""}`,
		},
		{
			name:      "Ne ok",
			inputBody: `{"username": "username", "password": "qwerty"}`,
			inputUser: signInInput{
				Username: "username",
				Password: "qwerty",
			},
			mockBehavior: func(r *service_mocks.MockAuthorization, user signInInput) {
				r.EXPECT().GenerateToken(user.Username, user.Password).Return("", errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
		{
			name:                 "BINDJson error",
			inputBody:            `{"usaername": "username", "password": "qwerty"}`,
			inputUser:            signInInput{},
			mockBehavior:         func(r *service_mocks.MockAuthorization, user signInInput) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"Key: 'signInInput.Username' Error:Field validation for 'Username' failed on the 'required' tag"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)

			repo := service_mocks.NewMockAuthorization(c)
			test.mockBehavior(repo, test.inputUser)

			services := &service.Service{Authorization: repo}
			handler := Handler{services}

			// Init Endpoint
			r := gin.New()
			r.POST("/sign-in", handler.signIn)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-in",
				bytes.NewBufferString(test.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
