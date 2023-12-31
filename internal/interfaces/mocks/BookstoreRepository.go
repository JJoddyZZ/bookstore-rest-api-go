// Code generated by mockery v2.35.4. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	models "github.com/jzavala-globant/bookstore-rest-api-go/internal/models"
)

// BookstoreRepository is an autogenerated mock type for the BookstoreRepository type
type BookstoreRepository struct {
	mock.Mock
}

// ListBooks provides a mock function with given fields: _a0
func (_m *BookstoreRepository) ListBooks(_a0 context.Context) ([]models.Book, error) {
	ret := _m.Called(_a0)

	var r0 []models.Book
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]models.Book, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []models.Book); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Book)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewBookstoreRepository creates a new instance of BookstoreRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBookstoreRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *BookstoreRepository {
	mock := &BookstoreRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
