package repository

import (
	"shop/internal/model"
)

type CategoryRepository interface {
	GetAllCategories() []*model.Category
	GetCategoryById(id int) ([]*model.SubCategory, error)
}

type CategoryRepositoryImpl struct {
	categories map[int]*model.Category
}

func NewCategoryRepositoryImpl(categories map[int]*model.Category) CategoryRepository {
	return &CategoryRepositoryImpl{categories: categories}
}

func (c CategoryRepositoryImpl) GetAllCategories() []*model.Category {
	categories := make([]*model.Category, 0, len(c.categories))
	for _, v := range c.categories {
		categories = append(categories, v)
	}
	return categories
}

func (c CategoryRepositoryImpl) GetCategoryById(id int) ([]*model.SubCategory, error) {
	for _, v := range c.categories {
		if v.ID == id {
			return v.SubCategories, nil
		}
	}
	return nil, ErrNotFound
}
