package persistence

import (
	"time"

	"github.com/DaveLaj/budget-tracker/internal/database"
)

type Record struct {
	ID     int
	Name   string
	Date   time.Time
	Amount string
}

func GetAllRecords() (*[]Record, error) {

	result, err := database.BudgetTrackerPool.Query("SELECT id, name, date, amount FROM budget_records")
	if err != nil {
		return &[]Record{}, err
	}
	defer result.Close()

	var records []Record

	for result.Next() {
		var record Record
		err = result.Scan(&record.ID, &record.Name, &record.Date, &record.Amount)
		if err != nil {
			return &[]Record{}, err
		}
		records = append(records, record)
	}

	if result.Err() != nil {
		return &[]Record{}, result.Err()
	}

	return &records, nil
}
