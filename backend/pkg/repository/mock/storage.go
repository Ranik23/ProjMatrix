package mock

import (
	"context"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func NewMockRepository() *MockRepository {
	repo := &MockRepository{}

	repo.On("Get", mock.Anything, "123").Return([]byte("default_data"), 42.0, nil)
	repo.On("Save", mock.Anything, "123", 42.0, []byte("default_data")).Return(nil)

	return repo
}

func (r *MockRepository) Get(ctx context.Context, id string) ([]byte, float64, error) {
	ret := r.Mock.Called(ctx, id)

	data, _ := ret.Get(0).([]byte)
	time, _ := ret.Get(1).(float64)
	err := ret.Error(2)

	return data, time, err
}

func (r *MockRepository) Save(ctx context.Context, id string, time float64, data []byte) error {
	ret := r.Mock.Called(ctx, id, time, data)

	return ret.Error(0)
}
