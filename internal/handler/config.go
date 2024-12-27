package handler

import (
	"fmt"
	"maps"
	"net/http"
	"strings"

	"github.com/ankorstore/yokai/config"
	"github.com/ankorstore/yokai/fxhttpserver"
	"github.com/ankorstore/yokai/log"
	"github.com/ankorstore/yokai/trace"
	"github.com/labstack/echo/v4"
)

type ExampleHandler struct {
	config *config.Config
}

var _ fxhttpserver.Handler = (*ExampleHandler)(nil)

func NewExampleHandler(config *config.Config) *ExampleHandler {
	return &ExampleHandler{
		config: config,
	}
}

func (h *ExampleHandler) Handle() echo.HandlerFunc {
	// helpers
	fmt.Printf("name: %s", h.config.AppName())               // name: app
	fmt.Printf("description: %s", h.config.AppDescription()) // description: app description
	fmt.Printf("env: %s", h.config.AppEnv())                 // env: dev
	fmt.Printf("version: %s", h.config.AppVersion())         // version: 0.1.0
	fmt.Printf("debug: %v", h.config.AppDebug())             // debug: false

	return func(c echo.Context) error {
		// example of correlated trace span
		ctx, span := trace.CtxTracerProvider(c.Request().Context()).Tracer("example tracer").Start(c.Request().Context(), "example span")
		defer span.End()

		// example of correlated log
		log.CtxLogger(ctx).Info().Msg("in example handler")

		ormConfig := strings.Builder{}
		for k, v := range maps.All(h.config.GetStringMapString("modules.orm")) {
			ormConfig.WriteString(fmt.Sprintf("%s: %s\n", k, v))
		}

		return c.JSON(http.StatusOK, map[string]string{
			"app name":         h.config.AppName(),
			"config file used": h.config.ConfigFileUsed(),
			"app version":      h.config.AppVersion(),
			"app env":          h.config.AppEnv(),
			"orm.config":       ormConfig.String(),
			"driver":           h.config.GetString("modules.orm.driver"),
		})
	}
}
