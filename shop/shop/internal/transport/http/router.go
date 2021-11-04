package http

import (
	"shop/internal/transport/http/handler"

	"github.com/go-chi/chi/v5"
)

func configureRouter(c *chi.Mux, manager *handler.Manager) {
	c.Route("/categories", func(r chi.Router) {
		r.Get("/", manager.GetAllCategories)
		r.Get("/{id}", manager.GetCategoryById)
		r.Get("/{id}/products", manager.GetProductsBySubCategoryId)
		r.Get("/{id}/products/{productID}", manager.GetProductById)
	})
}
