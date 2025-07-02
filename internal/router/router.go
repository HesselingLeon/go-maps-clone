package router

import (
	"database/sql"

	"github.com/hesselingleon/go-map-clone/internal/handler"
	"github.com/swaggest/openapi-go/openapi31"
	"github.com/swaggest/rest/response/gzip"
	"github.com/swaggest/rest/web"
	swgui "github.com/swaggest/swgui/v5emb"
)

func NewRouter(db *sql.DB) *web.Service {
	s := web.NewService(openapi31.NewReflector())

	s.OpenAPISchema().SetTitle("Leon Maps")
	s.OpenAPISchema().SetDescription("This app showcases a Google Maps clone.")
	s.OpenAPISchema().SetVersion("v0.1")

	s.Wrap(
		gzip.Middleware,
	)

	cordinateCalculationUseCase := handler.CoordinateToTileUsecase(db)
	s.Get("/tile/{x}/{y}/{z}", cordinateCalculationUseCase)

	s.Docs("/docs", swgui.New)

	return s
}
