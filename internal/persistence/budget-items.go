package persistence

import (
	"fmt"
	"time"

	"github.com/DaveLaj/budget-tracker/internal/database"
)

type Item struct {
	Name        string
	Description string
	Amount      int
	Date        time.Time
}

func InsertItem(item Item) error {
	_, err := database.BudgetTrackerPool.Exec("INSERT INTO items (name, description, amount, date) VALUES (?, ?, ?, ?)", item.Name, item.Description, item.Amount, item.Date)
	if err != nil {
		return fmt.Errorf("Error Inserting Item into Database: %w", err)
	}

	return nil
}

func GetAllItems() (Item, error) {
	return Item{}, nil
}
