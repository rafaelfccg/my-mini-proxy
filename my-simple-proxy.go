package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const (
	binanceUrl  = "https://www.binance.com/"
	binancePath = "/binance/"

	bitblueUrl  = "https://www.bitblue.com/"
	bitbluePath = "/bitblue/"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	// Setup proxy
	e.Any(binancePath+"*", makeHandler(binanceUrl, binancePath))
	e.Any(bitbluePath+"*", makeHandler(bitblueUrl, bitbluePath))
	e.Any("/*", makeHandler(binanceUrl, "/"))
	e.Logger.Fatal(e.Start(":8080"))
}

func makeHandler(url, path string) echo.HandlerFunc {
	return func(c echo.Context) error {
		transport := &http.Transport{}
		client := &http.Client{
			Transport: transport,
		}
		requestUri := c.Request().RequestURI[len(path):]
		println(url + requestUri)
		req, err := http.NewRequest(c.Request().Method, url+requestUri, nil)
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
	}
}
