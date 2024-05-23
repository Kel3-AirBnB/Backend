package middlewares

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func RemoveTrailingSlash(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Ambil alamat URL dari permintaan
		url := c.Request().URL.Path

		// Hapus trailing slash dari alamat URL
		cleanURL := strings.TrimRight(url, "/")

		// Jika alamat URL telah diubah, redirect ke versi yang bersih
		if cleanURL != url {
			return c.Redirect(http.StatusMovedPermanently, cleanURL)
		}

		// Lanjutkan ke handler berikutnya
		return next(c)
	}
}
