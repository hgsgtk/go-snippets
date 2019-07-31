package mockexample_test

import (
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/hgsgtk/go-snippets/testing-codes/mockexample"
	"github.com/hgsgtk/go-snippets/testing-codes/mockexample/mock_service"
	"testing"
)

func TestTripService_Run_CustomerBan(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockCustomer := mock_service.NewMockCustomerRepository(ctrl)
	mockTrip := mock_service.NewMockTripRepository(ctrl)

	argCode := "customer_01"
	mockCustomer.EXPECT().GetByCode(argCode).
		AnyTimes().Return(mockexample.Customer{ID: 1}, nil)
	mockTrip.EXPECT().GetByCustomerID(1, "reserved").
		AnyTimes().Return(mockexample.Trip{}, nil)

	s := mockexample.TripService{Customer: mockCustomer, Trip: mockTrip}

	got, gotErr := s.Run(argCode)
	if gotErr != nil {
		t.Fatalf(
			"TripService.Run(`%s`) got unexpected error %#v, want no error",
			argCode, gotErr,
		)
	}
	want := mockexample.OutputTrip{Active: false}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("TripService.Run(`%s`) differs: (-got +want)\n%s", argCode, diff)
	}
}
