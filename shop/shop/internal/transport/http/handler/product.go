package handler

import (
	"encoding/json"
	"net/http"
	"shop/internal/repository"
	"shop/internal/service"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	ps service.ProductService
}

func NewProductHandler(ps service.ProductService) *ProductHandler {
	return &ProductHandler{ps: ps}
}

func (p ProductHandler) GetProductsBySubCategoryId(rw http.ResponseWriter, r *http.Request) {
	id, parseErr := strconv.Atoi(chi.URLParam(r, "id"))
	if parseErr != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	products, prodErr := p.ps.GetProductsBySubCategoryId(id)
	if prodErr != nil {
		if prodErr == repository.ErrNotFound {
			rw.WriteHeader(http.StatusNotFound)
		} else {
			rw.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	if err := json.NewEncoder(rw).Encode(products); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (p ProductHandler) GetProductById(rw http.ResponseWriter, r *http.Request) {
	prID, convErr := strconv.Atoi(chi.URLParam(r, "productID"))
	if convErr != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	product, productErr := p.ps.GetProductById(prID)
	if productErr != nil {
		if productErr == repository.ErrNotFound {
			rw.WriteHeader(http.StatusNotFound)
		} else {
			rw.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	if err := json.NewEncoder(rw).Encode(product); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
	}
}
