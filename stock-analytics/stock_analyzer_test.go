package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetTopCompaniesByAveragePrice(t *testing.T) {
	tests := []struct {
		name    string
		stocks  []CompanyStock
		want    []string
		wantErr bool
	}{
		{
			name: "example from problem statement",
			stocks: []CompanyStock{
				{"Shopify", []float64{100, 102, 98, 105, 110, 108, 107}},
				{"Google", []float64{200, 198, 202, 210, 205, 207, 208}},
				{"Amazon", []float64{150, 148, 152, 153, 151, 149, 150}},
				{"Meta", []float64{120, 125, 130, 127, 129, 131, 128}},
				{"Netflix", []float64{300, 305, 310, 315, 320, 325, 330}},
			},
			want:    []string{"Netflix", "Google", "Amazon"},
			wantErr: false,
		},
		{
			name: "exactly 3 companies",
			stocks: []CompanyStock{
				{"Apple", []float64{100, 101, 102, 103, 104, 105, 106}},
				{"Microsoft", []float64{200, 201, 202, 203, 204, 205, 206}},
				{"Tesla", []float64{150, 151, 152, 153, 154, 155, 156}},
			},
			want:    []string{"Microsoft", "Tesla", "Apple"},
			wantErr: false,
		},
		{
			name: "tie breaking by name",
			stocks: []CompanyStock{
				{"Apple", []float64{100, 100, 100, 100, 100, 100, 100}},
				{"Microsoft", []float64{100, 100, 100, 100, 100, 100, 100}},
				{"Tesla", []float64{100, 100, 100, 100, 100, 100, 100}},
				{"Zebra", []float64{200, 200, 200, 200, 200, 200, 200}},
			},
			want:    []string{"Zebra", "Apple", "Microsoft"},
			wantErr: false,
		},
		{
			name: "less than 3 companies",
			stocks: []CompanyStock{
				{"Apple", []float64{100, 101, 102, 103, 104, 105, 106}},
				{"Microsoft", []float64{200, 201, 202, 203, 204, 205, 206}},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "company with wrong number of prices",
			stocks: []CompanyStock{
				{"Apple", []float64{100, 101, 102, 103, 104, 105}}, // only 6 prices
				{"Microsoft", []float64{200, 201, 202, 203, 204, 205, 206}},
				{"Tesla", []float64{150, 151, 152, 153, 154, 155, 156}},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty input",
			stocks: []CompanyStock{},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetTopCompaniesByAveragePrice(tt.stocks)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTopCompaniesByAveragePrice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTopCompaniesByAveragePrice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateAverage(t *testing.T) {
	tests := []struct {
		name   string
		prices []float64
		want   float64
	}{
		{
			name:   "normal prices",
			prices: []float64{100, 102, 98, 105, 110, 108, 107},
			want:   104.28571428571429, // (100+102+98+105+110+108+107)/7
		},
		{
			name:   "all same prices",
			prices: []float64{100, 100, 100, 100, 100, 100, 100},
			want:   100.0,
		},
		{
			name:   "empty slice",
			prices: []float64{},
			want:   0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateAverage(tt.prices); got != tt.want {
				t.Errorf("calculateAverage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkGetTopCompaniesByAveragePrice(b *testing.B) {
	// Create a large dataset for benchmarking
	stocks := make([]CompanyStock, 1000)
	for i := 0; i < 1000; i++ {
		stocks[i] = CompanyStock{
			Name:   fmt.Sprintf("Company%d", i),
			Prices: []float64{float64(i), float64(i+1), float64(i+2), float64(i+3), float64(i+4), float64(i+5), float64(i+6)},
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := GetTopCompaniesByAveragePrice(stocks)
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
} 
