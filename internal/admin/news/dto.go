package news

import "time"

type DataNewsDTO struct {
	News  []GetNewsDTO `json:"news"`
	Count int          `json:"count"`
}

type GetNewsDTO struct {
	Id        int       `json:"Id"`
	Image     string    `json:"imagePath"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Date      string    `json:"date"`
	CreatedAt time.Time `json:"created_at"`
}

type AddNewsDTO struct {
	Image   string `json:"imagePath"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Date    string `json:"date"`
}
