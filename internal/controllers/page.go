package controllers

import (
	"github.com/DaveLaj/budget-tracker/internal/templates"
	"github.com/gin-gonic/gin"
)

func ShowPage(c *gin.Context) {
	budgetPage := templates.Page()

	err := budgetPage.Render(c.Request.Context(), c.Writer)
	if err != nil {
		c.String(500, "Error rendering template: %v", err)
	}
}
