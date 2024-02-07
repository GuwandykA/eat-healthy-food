package banner

import "context"

type Repository interface {
	GetAllData(ctx context.Context, limit, page int) (DataBannerDTO, error)
	GetOneData(ctx context.Context, id int) (GetBannerDTO, error)
	AddData(ctx context.Context, dto AddBannerDTO) (GetBannerDTO, error)
	UpdateData(ctx context.Context, dto AddBannerDTO, id int) (GetBannerDTO, error)
	DeleteData(ctx context.Context, id int) error
}
