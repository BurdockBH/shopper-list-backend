package shopper

import (
	"log"
	"shopper-list-backend/db"
	"shopper-list-backend/viewmodels"
)

func GetShopper() ([]viewmodels.Shopper, error) {
	query := "SELECT * FROM shopper"
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

	var shoppers []viewmodels.Shopper
	for rows.Next() {
		var shopper viewmodels.Shopper
		err := rows.Scan(&shopper.ID, &shopper.FirstName, &shopper.LastName)
		if err != nil {
			log.Fatal(err)
		}
		shoppers = append(shoppers, shopper)
	}

	if shoppers == nil {
		log.Fatal("No shoppers found")
	}

	return shoppers, nil
}
