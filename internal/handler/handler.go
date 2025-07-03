package handler

import (
	"context"
	"database/sql"
	"math"

	"github.com/hesselingleon/go-map-clone/internal/models"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
)

func CoordinateToTileUsecase(db *sql.DB) usecase.Interactor {

	u := usecase.NewInteractor(func(ctx context.Context, input models.CoordinateInput, output *models.CoordinateOutput) error {

		numberOfTileForZoomLevel := math.Pow(2, float64(input.Zoom))

		lon_deg_nw, lat_deg_nw := calculateTileTopLeftCorner(float64(input.Column), float64(input.Row), numberOfTileForZoomLevel)

		lon_deg_se, lat_deg_se := calculateTileTopLeftCorner(float64(input.Column+1), float64(input.Row+1), numberOfTileForZoomLevel)

		output.Lon_deg_nw = lon_deg_nw
		output.Lat_deg_nw = lat_deg_nw
		output.Lon_deg_se = lon_deg_se
		output.Lat_deg_se = lat_deg_se

		return nil
	})

	u.SetTitle("Calculating")
	u.SetDescription("")
	u.SetExpectedErrors(status.InvalidArgument)

	return u
}

func calculateTileTopLeftCorner(column, row, numberOfTileForZoomLevel float64) (float64, float64) {
	lon_deg_nw := float64(column)/numberOfTileForZoomLevel*360.0 - 180.0
	lat_rad_nw := math.Atan(math.Sinh(math.Pi * (1 - 2*float64(row)/numberOfTileForZoomLevel)))
	lat_deg_nw := lat_rad_nw * 180.0 / math.Pi
	return lon_deg_nw, lat_deg_nw
}
