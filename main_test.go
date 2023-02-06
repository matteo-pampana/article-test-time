package main

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestService_CreateItem(t *testing.T) {
	tests := []struct {
		name              string
		storeExpectations func(m *MockStore)
		testTime          time.Time
		wantErr           error
	}{
		{
			name: "too fast creation",
			storeExpectations: func(m *MockStore) {
				m.EXPECT().GetLastItem().Return(
					&Item{
						CreatedAt: time.Date(2023, 02, 06, 10, 10, 1, 0, time.UTC),
					}, nil,
				)
			},
			testTime: time.Date(2023, 02, 06, 10, 10, 10, 0, time.UTC),
			wantErr:  errTooFast,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockStore := NewMockStore(ctrl)
			tt.storeExpectations(mockStore)

			s := NewService(mockStore)
			s.now = func() time.Time {
				return tt.testTime
			}
			err := s.CreateItem(Item{})
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
