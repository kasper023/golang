package model

type AllCategories struct {
	Categories []*Category `json:"categories"`
}

type Category struct {
	ID            int            `json:"id"`
	Name          string         `json:"name"`
	SubCategories []*SubCategory `json:"sub_categories"`
}

type RootCategory struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type SubCategory struct {
	ID       int        `json:"id"`
	Name     string     `json:"name"`
	Products []*Product `json:"products"`
}

type RootSubCategory struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
