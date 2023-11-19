package item

import (
	"log"
	"shopper-list-backend/db"
	"shopper-list-backend/viewmodels"
)

func GetItem() ([]viewmodels.Item, error) {
	query := "SELECT * FROM items"
	st, err := db.DB.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer st.Close()

	rows, err := st.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var items []viewmodels.Item

	for rows.Next() {
		var item viewmodels.Item
		err := rows.Scan(&item.ID, &item.Name, &item.Price, &item.Description, &item.Quantity)
		if err != nil {
			log.Fatal(err)
		}
		items = append(items, item)
	}

	if items == nil {
		log.Fatal("No shoppers found")
	}

	return items, nil
}
