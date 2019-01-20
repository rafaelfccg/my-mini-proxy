package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	// Setup proxy
	binanceUrl := "https://www.binance.com"

	// e.Use(middleware.Proxy(middleware.NewRoundRobinBalancer(targets)))
	e.Any("/*", func(c echo.Context) error {
		transport := &http.Transport{}
		client := &http.Client{
			Transport: transport,
		}
		println(c.Request().RequestURI)
		req, err := http.NewRequest(c.Request().Method, binanceUrl+c.Request().RequestURI, nil)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Could not create request!")
		}
		req.Header = c.Request().Header
		if c.Request().Body != nil {
			req.Body = c.Request().Body
		}

		response, err := client.Do(req)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Request failed!")
		}

		for k, headerValues := range response.Header {
			for _, value := range headerValues {
				c.Response().Header().Add(k, value)
			}
		}

		return c.Stream(response.StatusCode, response.Header.Get(echo.HeaderContentType), response.Body)
	})
	e.Logger.Fatal(e.Start(":8080"))
}
