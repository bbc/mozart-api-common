package storage

import "github.com/bbc/mozart-api-common/Godeps/_workspace/src/github.com/stretchr/testify/mock"

type MockStorage struct {
	mock.Mock
}

func (s *MockStorage) Get(key string) (string, *Error) {
	args := s.Called(key)
	return args.String(0), args.Get(1).(*Error)
}

func (s *MockStorage) Set(key string, data string) *Error {
	args := s.Called(key, data)
	return args.Get(0).(*Error)
}
