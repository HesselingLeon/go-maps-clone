package handler

import (
	"math"
	"testing"
)

func Test_calculateTileTopLeftCorner(t *testing.T) {
	tests := []struct {
		name                     string
		column                   float64
		row                      float64
		numberOfTileForZoomLevel float64
		expected_lon             float64
		expected_lat             float64
	}{
		{
			name:                     "calculation_correct",
			column:                   5,
			row:                      8,
			numberOfTileForZoomLevel: math.Exp2(12),
			expected_lon:             -179.560546875,
			expected_lat:             84.99010018023479,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lon, lat := calculateTileTopLeftCorner(tt.column, tt.row, tt.numberOfTileForZoomLevel)

			if math.Abs(lon-tt.expected_lon) > 0.00001 {
				t.Errorf("calculateTileTopLeftCorner(): lon= %v, want %v", lon, tt.expected_lon)

			}

			if math.Abs(lat-tt.expected_lat) > 0.00001 {
				t.Errorf("calculateTileTopLeftCorner(): lat=%v, want %v", lat, tt.expected_lat)
			}

		})
	}
}
