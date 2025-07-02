package handler

import (
	"context"
	"database/sql"

	"github.com/hesselingleon/go-map-clone/internal/models"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
)

func CoordinateToTileUsecase(db *sql.DB) usecase.Interactor {

	u := usecase.NewInteractor(func(ctx context.Context, input models.CoordinateInput, output *models.CoordinateOutput) error {

		return nil
	})

	u.SetTitle("Calculating")
	u.SetDescription("")
	u.SetExpectedErrors(status.InvalidArgument)

	return u
}
