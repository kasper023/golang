package http

import (
	"encoding/json"
	"net/http"
	"os"
	"shop/internal/model"
	"shop/internal/repository"
	"shop/internal/service"
	"shop/internal/transport/http/handler"

	"github.com/go-chi/chi/v5"
)

func initDB() (map[int]*model.Category, error) {
	f, fErr := os.Open("C:/Users/User/Desktop/Fall 2021/golang/shop/shop/db.json")
	if fErr != nil {
		return nil, fErr
	}
	categories := make([]*model.Category, 0)
	allCategories := &model.AllCategories{Categories: categories}
	if decErr := json.NewDecoder(f).Decode(allCategories); decErr != nil {
		return nil, decErr
	}
	dbMap := make(map[int]*model.Category)
	for _, v := range allCategories.Categories {
		dbMap[v.ID] = v
	}
	return dbMap, nil
}

func initSubDB(db map[int]*model.Category) map[int]*model.SubCategory {
	subDBMap := make(map[int]*model.SubCategory)
	for _, v := range db {
		for _, cv := range v.SubCategories {
			subDBMap[cv.ID] = cv
		}
	}
	return subDBMap
}

func initProductDB(db map[int]*model.SubCategory) map[int]*model.Product {
	prMap := make(map[int]*model.Product)
	for _, v := range db {
		for _, pv := range v.Products {
			prMap[pv.ID] = pv
		}
	}
	return prMap
}

func StartServer(addr string, ch chan error) {
	db, err := initDB()
	if err != nil {
		ch <- err
		return
	}
	subDB := initSubDB(db)
	productDB := initProductDB(subDB)

	pr := repository.NewProductRepositoryImpl(subDB, productDB)
	cr := repository.NewCategoryRepositoryImpl(db)

	ps := service.NewProductServiceImpl(pr)
	cs := service.NewCategoryServiceImpl(cr)

	prodHandler := handler.NewProductHandler(ps)
	catHandler := handler.NewCategoryHandler(cs)

	manager := handler.NewManager(prodHandler, catHandler)

	router := chi.NewRouter()

	configureRouter(router, manager)

	ch <- http.ListenAndServe(":8000", router)
}
