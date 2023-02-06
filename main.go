package main

import (
	"errors"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	lecho "github.com/ziflex/lecho/v3"
)

func main() {
	l := zerolog.New(os.Stderr).With().Timestamp().Str("service", "echo-logger-test").Logger()

	e := echo.New()
	e.HideBanner = true
	e.Logger = lecho.From(l)

	e.GET("/something", somethingHandler())

	if err := e.Start(":8093"); err != nil {
		l.Fatal().Err(err).Msg("echo is stillborn")
	}
}

func somethingHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		err := errors.New("this is an internal error")
		return echo.NewHTTPError(http.StatusBadRequest, "oh no somethng happened").SetInternal(err)
	}
}
