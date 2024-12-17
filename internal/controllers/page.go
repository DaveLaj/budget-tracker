package controllers

import (
	"github.com/DaveLaj/budget-tracker/internal/persistence"
	"github.com/DaveLaj/budget-tracker/internal/templates"
	"github.com/gin-gonic/gin"
)

func ShowPage(c *gin.Context) {
	records, err := persistence.GetAllRecords()
	if err != nil {
		c.String(500, "Error retrieving records: %v", err)
	}

	budgetPage := templates.Page(records)

	err = budgetPage.Render(c.Request.Context(), c.Writer)
	if err != nil {
		c.String(500, "Error rendering template: %v", err)
	}
}
