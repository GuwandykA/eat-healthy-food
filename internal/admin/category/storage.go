package category

import "context"

type Repository interface {
	GetAllData(ctx context.Context) ([]CategoryDTO, error)
	GetOneData(ctx context.Context, req ReqIdDTO) ([]SubCategoryDTO, error)

	GetProductData(ctx context.Context, req ReqIdDTO) (SubCategoryDTO, error)

	AddData(ctx context.Context, dto CategoryIdDTO) error
	UpdateData(ctx context.Context, dto CategoryIdDTO) error
	DeleteData(ctx context.Context, uuid int) error
	// sub-category
	AddSubData(ctx context.Context, dto SubCategoryIdDTO) error
	UpdateSubData(ctx context.Context, dto SubCategoryIdDTO) error
	DeleteSubData(ctx context.Context, uuid int) error
}
