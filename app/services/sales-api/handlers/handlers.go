// Package handlers manages the different versions of the API.
package handlers

import (
	"net/http"
	"os"

	"github.com/sudonite/service/app/services/sales-api/handlers/v1/testgrp"
	"github.com/sudonite/service/business/web/v1/mid"
	"github.com/sudonite/service/foundation/web"
	"go.uber.org/zap"
)

// APIMuxConfig contains all the mandatory systems required by handlers.
type APIMuxConfig struct {
	Build    string
	Shutdown chan os.Signal
	Log      *zap.SugaredLogger
}

// APIMux constructs a http.Handler with all application routes defined.
func APIMux(cfg APIMuxConfig) *web.App {
	app := web.NewApp(cfg.Shutdown, mid.Logger(cfg.Log), mid.Errors(cfg.Log), mid.Panics())

	app.Handle(http.MethodGet, "/test", testgrp.Test)

	return app
}
