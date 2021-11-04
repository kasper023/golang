package repository

import "shop/internal/model"

type ProductRepository interface {
	GetProductsBySubCategoryId(categoryID int) ([]*model.Product, error)
	GetProductById(productID int) (*model.Product, error)
}

type ProductRepositoryImpl struct {
	subCategories map[int]*model.SubCategory
	allProducts   map[int]*model.Product
}

func NewProductRepositoryImpl(subCategories map[int]*model.SubCategory, allProducts map[int]*model.Product) ProductRepository {
	return &ProductRepositoryImpl{subCategories: subCategories, allProducts: allProducts}
}

func (p ProductRepositoryImpl) GetProductById(productID int) (*model.Product, error) {
	if pr, ok := p.allProducts[productID]; ok {
		return pr, nil
	}
	return nil, ErrNotFound
}

func (p ProductRepositoryImpl) GetProductsBySubCategoryId(categoryID int) ([]*model.Product, error) {
	if cr, ok := p.subCategories[categoryID]; ok {
		return cr.Products, nil
	}
	return nil, ErrNotFound
}
