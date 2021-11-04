package handler

import (
	"encoding/json"
	"net/http"
	"shop/internal/model"
	"shop/internal/repository"
	"shop/internal/service"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type CategoryHandler struct {
	cs service.CategoryService
}

func NewCategoryHandler(cs service.CategoryService) *CategoryHandler {
	return &CategoryHandler{cs: cs}
}

func (c CategoryHandler) GetAllCategories(rw http.ResponseWriter, r *http.Request) {
	rootCategories := c.cs.GetAllCategories()
	mappedRootCategories := make([]*model.RootCategory, 0, len(rootCategories))
	for _, v := range rootCategories {
		mappedRootCategories = append(mappedRootCategories, &model.RootCategory{ID: v.ID, Name: v.Name})
	}
	if err := json.NewEncoder(rw).Encode(mappedRootCategories); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (c CategoryHandler) GetCategoryById(rw http.ResponseWriter, r *http.Request) {
	id, parseErr := strconv.Atoi(chi.URLParam(r, "id"))
	if parseErr != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	subCategories, subErr := c.cs.GetCategoryById(id)
	if subErr != nil {
		if subErr == repository.ErrNotFound {
			rw.WriteHeader(http.StatusNotFound)
		} else {
			rw.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	rootSubCategories := make([]*model.RootSubCategory, 0, len(subCategories))
	for _, v := range subCategories {
		rootSubCategories = append(rootSubCategories, &model.RootSubCategory{ID: v.ID, Name: v.Name})
	}
	if err := json.NewEncoder(rw).Encode(rootSubCategories); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
}
