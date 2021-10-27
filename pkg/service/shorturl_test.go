package service

import (
	"errors"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"github.com/piklin/test_ozon_fintech/pkg/repository"
	mock_repository "github.com/piklin/test_ozon_fintech/pkg/repository/mocks"
)

func TestService_GetFullURL(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockShortURL, shortURL string)

	testTable := []struct {
		name            string
		input           string
		mockBehavior    mockBehavior
		expectedFullURL string
		expectedError   error
	}{
		{
			name:  "OK",
			input: "K4kf_deoA1",
			mockBehavior: func(r *mock_repository.MockShortURL, shortURL string) {
				r.EXPECT().SearchShortURL(shortURL).Return("https://www.google.com/", nil)
			},
			expectedFullURL: "https://www.google.com/",
			expectedError:   nil,
		},
		{
			name:  "Repository error",
			input: "K4kf_deoA1",
			mockBehavior: func(r *mock_repository.MockShortURL, shortURL string) {
				r.EXPECT().SearchShortURL(shortURL).Return("", errors.New("Repository error"))
			},
			expectedFullURL: "",
			expectedError:   errors.New("Repository error"),
		},
		{
			name:  "Full URL does not exists",
			input: "K4kf_deoA1",
			mockBehavior: func(r *mock_repository.MockShortURL, shortURL string) {
				r.EXPECT().SearchShortURL(shortURL).Return("", nil)
			},
			expectedFullURL: "",
			expectedError:   nil,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mock := mock_repository.NewMockShortURL(c)
			testCase.mockBehavior(mock, testCase.input)

			repository := &repository.Repository{ShortURL: mock}
			service := NewService(repository)

			// perform
			fullURL, error := service.GetFullURL(testCase.input)

			// check
			assert.Equal(t, testCase.expectedFullURL, fullURL)
			assert.Equal(t, testCase.expectedError, error)
		})
	}
}
