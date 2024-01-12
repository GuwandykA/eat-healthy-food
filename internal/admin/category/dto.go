package category

import "time"

type DataDTO struct {
	SubCategories []SubCategoryDTO `json:"subCat"`
	Count         int              `json:"count"`
}

type CategoryDTO struct {
	Id        int       `json:"Id"`
	Name      string    `json:"name"`
	Image     string    `json:"image"`
	Count     int       `json:"count"`
	CreatedAt time.Time `json:"createdAt"`
}

type SubCategoryDTO struct {
	Id        int         `json:"Id"`
	CatName   string      `json:"catName"`
	Name      string      `json:"name"`
	Image     string      `json:"imagePath"`
	Item      interface{} `json:"item"`
	CreatedAt time.Time   `json:"createdAt"`
}

type CategoryIdDTO struct {
	Id    int    `json:"Id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

type SubCategoryIdDTO struct {
	Id    int         `json:"Id"`
	CatId int         `json:"catId"`
	Name  string      `json:"name"`
	Image string      `json:"image"`
	Item  interface{} `json:"item"`
}

type ReqIdDTO struct {
	Id int `json:"Id"`
}
