package service

import (
	"shop/internal/model"
	"shop/internal/repository"
)

type ProductService interface {
	GetProductsBySubCategoryId(categoryID int) ([]*model.Product, error)
	GetProductById(productID int) (*model.Product, error)
}

type ProductServiceImpl struct {
	pr repository.ProductRepository
}

func NewProductServiceImpl(pr repository.ProductRepository) ProductService {
	return &ProductServiceImpl{pr: pr}
}

func (p ProductServiceImpl) GetProductById(productID int) (*model.Product, error) {
	return p.pr.GetProductById(productID)
}

func (p ProductServiceImpl) GetProductsBySubCategoryId(categoryID int) ([]*model.Product, error) {
	return p.pr.GetProductsBySubCategoryId(categoryID)
}
