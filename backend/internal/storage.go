package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowHomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func ShowRegistrationForm(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

func ShowLoginForm(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func ShowCart(c *gin.Context) {
	c.HTML(http.StatusOK, "cart.html", nil)
}

func ShowCatalog(c *gin.Context) {
	c.HTML(http.StatusOK, "catalog.html", nil)
}

func ShowFavourites(c *gin.Context) {
	c.HTML(http.StatusOK, "favourites.html", nil)
}
