package shopping_list

import (
	"fmt"
	"log"
	"shopper-list-backend/db"
	"shopper-list-backend/viewmodels"
)

func AddShopperList(u *viewmodels.Shopping_list) ([]viewmodels.Shopping_list, error) {
	query := "SELECT s.first_name, i.name, COUNT(*) AS item_count FROM shopping_list_table sl JOIN shopper s ON sl.shopper_id = s.id JOIN items i ON sl.item_id = i.id WHERE sl.shopper_id = ? AND sl.item_id = ? GROUP BY sl.shopper_id, sl.item_id;"
	st, err := db.DB.Prepare(query)
	if err != nil {
		log.Print(err)
	}
	defer st.Close()

	rows, err := st.Query(u.ShopperID, u.ItemID)
	if err != nil {
		log.Print(err)
	}
	defer rows.Close()

	var totalCount int
	var name string
	var itemName string

	for rows.Next() {
		if err := rows.Scan(&name, &itemName, &totalCount); err != nil {
			log.Print(err)
			return nil, err
		}
		break
	}

	if totalCount >= 1 {
		log.Printf("User %v cannot have more than 1 %v", name, itemName)
		return nil, fmt.Errorf("user %v cannot have more than 1 %v", name, itemName)
	}

	query = "SELECT i.name, COUNT(*) AS item_count FROM shopping_list_table sl JOIN items i ON sl.item_id = i.id WHERE sl.item_id = ? GROUP BY sl.item_id;"
	st, err = db.DB.Prepare(query)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer st.Close()

	rows, err = st.Query(u.ItemID)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	for rows.Next() {
		if err := rows.Scan(&itemName, &totalCount); err != nil {
			log.Print(err)
			return nil, err
		}
		break
	}

	if totalCount >= 3 {
		log.Printf("Cannot have more than 3 %vs", itemName)
		return nil, fmt.Errorf("cannot have more than 3 %vs", itemName)
	}

	query = "INSERT INTO shopping_list_table (shopper_id, item_id) VALUES (?, ?)"
	st, err = db.DB.Prepare(query)
	if err != nil {
		log.Print(err)
	}

	_, err = st.Exec(u.ShopperID, u.ItemID)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	query = "SELECT * FROM shopping_list_table"

	st, err = db.DB.Prepare(query)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	rows, err = st.Query()
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer rows.Close()

	var shopping_list []viewmodels.Shopping_list

	for rows.Next() {
		var item viewmodels.Shopping_list
		err := rows.Scan(&item.ID, &item.ShopperID, &item.ItemID)
		if err != nil {
			log.Print(err)
		}
		shopping_list = append(shopping_list, item)
	}

	if shopping_list == nil {
		log.Print("No items found")
		return nil, err
	}

	return shopping_list, nil
}

func GetShoppingList() ([]viewmodels.Shopping_list, error) {
	query := "SELECT * FROM shopping_list_table"

	st, err := db.DB.Prepare(query)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	rows, err := st.Query()
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer rows.Close()

	var shopping_list []viewmodels.Shopping_list

	for rows.Next() {
		var item viewmodels.Shopping_list
		err := rows.Scan(&item.ID, &item.ShopperID, &item.ItemID)
		if err != nil {
			log.Print(err)
		}
		shopping_list = append(shopping_list, item)
	}

	if shopping_list == nil {
		log.Print("No items found")
		return nil, err
	}

	return shopping_list, nil
}

func DeleteShoppingLists() error {
	query := "DELETE FROM shopping_list_table"

	st, err := db.DB.Prepare(query)
	if err != nil {
		log.Print(err)
		return err
	}

	_, err = st.Exec()
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}
