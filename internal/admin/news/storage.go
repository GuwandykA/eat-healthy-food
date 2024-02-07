package news

import "context"

type Repository interface {
	GetAllData(ctx context.Context, limit, page int) (DataNewsDTO, error)
	GetOneData(ctx context.Context, id int) (GetNewsDTO, error)
	AddData(ctx context.Context, dto AddNewsDTO) (GetNewsDTO, error)
	UpdateData(ctx context.Context, dto AddNewsDTO, id int) (GetNewsDTO, error)
	DeleteData(ctx context.Context, id int) error
}
