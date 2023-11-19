package shopping_list

import (
	"encoding/json"
	"log"
	"net/http"
	shopping_list2 "shopper-list-backend/db/shopping-list"
	"shopper-list-backend/viewmodels"
)

func AddShopperList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var s viewmodels.Shopping_list

	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	r.Body.Close()

	shoppingList, err := shopping_list2.AddShopperList(&s)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response, _ := json.Marshal(viewmodels.ErrorResponse{
			Message: err.Error(),
		})
		w.Write(response)
		log.Println(err)
		return
	}

	jsonResponse, err := json.Marshal(shoppingList)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func GetShoppingList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	shoppingList, err := shopping_list2.GetShoppingList()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response, _ := json.Marshal(viewmodels.ErrorResponse{
			Message: err.Error(),
		})
		w.Write(response)
		log.Println(err)
		return
	}

	jsonResponse, err := json.Marshal(shoppingList)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func DeleteShoppingList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Methods", "DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	err := shopping_list2.DeleteShoppingLists()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response, _ := json.Marshal(viewmodels.ErrorResponse{
			Message: err.Error(),
		})
		w.Write(response)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(viewmodels.ErrorResponse{
		Message: "Deleted!",
	})
	w.Write(response)
}
