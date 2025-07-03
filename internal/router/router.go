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
	webservice := web.NewService(openapi31.NewReflector())

	webservice.OpenAPISchema().SetTitle("Leon Maps")
	webservice.OpenAPISchema().SetDescription("This app showcases a Google Maps clone.")
	webservice.OpenAPISchema().SetVersion("v0.1")

	webservice.Wrap(
		gzip.Middleware,
	)

	cordinateCalculationUseCase := handler.CoordinateToTileUsecase(db)
	webservice.Get("/tile/{z}/{x}/{y}", cordinateCalculationUseCase)

	webservice.Docs("/docs", swgui.New)

	return webservice
}
