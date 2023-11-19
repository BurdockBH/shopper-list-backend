package router

import (
	"net/http"
	"shopper-list-backend/service/item"
	"shopper-list-backend/service/shopper"
	"shopper-list-backend/service/shopping_list"
)

func InitializeRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/shopper/get", shopper.GetShoppers)
	router.HandleFunc("/items/get", item.GetItems)
	router.HandleFunc("/shopping-list/add", shopping_list.AddShopperList)
	router.HandleFunc("/shopping-list/get", shopping_list.GetShoppingList)
	router.HandleFunc("/shopping-list/delete", shopping_list.DeleteShoppingList)

	return router
}
