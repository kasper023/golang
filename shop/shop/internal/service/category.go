package service

import (
	"shop/internal/model"
	"shop/internal/repository"
)

type CategoryService interface {
	GetAllCategories() []*model.Category
	GetCategoryById(id int) ([]*model.SubCategory, error)
}

type CategoryServiceImpl struct {
	cr repository.CategoryRepository
}

func NewCategoryServiceImpl(cr repository.CategoryRepository) CategoryService {
	return &CategoryServiceImpl{cr: cr}
}

func (c CategoryServiceImpl) GetAllCategories() []*model.Category {
	return c.cr.GetAllCategories()
}

func (c CategoryServiceImpl) GetCategoryById(id int) ([]*model.SubCategory, error) {
	return c.cr.GetCategoryById(id)
}
