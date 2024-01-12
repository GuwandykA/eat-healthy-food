package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"saglyk-backend/internal/admin/category"
	"saglyk-backend/pkg/logging"
)

type repository struct {
	client *pgxpool.Pool
	logger *logging.Logger
}

func NewRepository(client *pgxpool.Pool, logger *logging.Logger) category.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}

func (r repository) GetAllData(ctx context.Context) ([]category.CategoryDTO, error) {

	var (
		ns []category.CategoryDTO
	)

	q := ` select 
				cs.Id, cs.name,
				count(ss.name),
				cs.img, cs.created_at
			 from categories cs 
		left join products ps on ps.cat_id = cs.Id
		group by cs.Id
		order by created_at desc ; `

	rows, err := r.client.Query(ctx, q)
	defer rows.Close()

	if err != nil {
		return ns, nil
	}

	for rows.Next() {
		var n category.CategoryDTO
		errN := rows.Scan(
			&n.Id, &n.Name,
			&n.Image, &n.Count,
			&n.CreatedAt,
		)
		if errN != nil {
			fmt.Println("errN :::", errN)
			continue
		}

		ns = append(ns, n)
	}
	return ns, nil
}

func (r repository) GetOneData(ctx context.Context, req category.ReqIdDTO) ([]category.SubCategoryDTO, error) {

	var (
		ns []category.SubCategoryDTO
	)

	q := `select 
				ps.id, cs.name as cat_name,
				ps.name, ps.img,
			 from products ps
		  left join categories cs on cs.id =  sc.cat_id
		  where  ps.cat_id = $1
		 order by sc.created_at desc ; `

	rows, err := r.client.Query(ctx, q, req.Id)
	defer rows.Close()

	if err != nil {
		return ns, nil
	}

	for rows.Next() {
		var n category.SubCategoryDTO
		errN := rows.Scan(
			&n.Id, &n.CatName,
			&n.Name, &n.Image,
			&n.Item,
			&n.CreatedAt,
		)
		if errN != nil {
			fmt.Println("errN :::", errN)
			continue
		}

		ns = append(ns, n)
	}

	return ns, nil
}

func (r repository) GetOneProduct(ctx context.Context, req category.ReqIdDTO) (category.SubCategoryDTO, error) {

	var (
		n category.SubCategoryDTO
	)

	q := `select 
				ps.id, cs.name as cat_name,
				ps.name, ps.img,
				ps.item, ps.created_at
			 from products ps
		  left join categories cs on cs.id =  sc.cat_id
		  where  ps.id = $1; `

	err := r.client.QueryRow(ctx, q, req.Id).Scan(
		&n.Id, &n.CatName,
		&n.Name, &n.Image,
		&n.Item, &n.CreatedAt,
	)
	if err != nil {
		fmt.Println("errN :::", err)

	}

	return n, nil
}

func (r repository) AddData(ctx context.Context, req category.CategoryIdDTO) error {

	q := `INSERT INTO categories (name, img) VALUES ($1, $2);`

	_, err := r.client.Exec(ctx, q, req.Name, req.Image)
	if err != nil && err.Error() != "no rows in result set" {
		return err
	}

	return nil
}

func (r repository) UpdateData(ctx context.Context, req category.CategoryIdDTO) error {
	var (
		q string
	)

	q = ` update categories set name = $1, img = $2 where Id = $3;`

	_, err := r.client.Exec(ctx, q, req.Name, req.Image, req.Id)

	if err != nil && err.Error() != "no rows in result set" {
		return err
	}

	return nil
}

func (r repository) DeleteData(ctx context.Context, Id string) error {
	var (
		q string
	)

	q = ` delete from categories where Id = $1 ;`
	_, err := r.client.Exec(ctx, q, Id)

	if err != nil && err.Error() != "no rows in result set" {
		return err
	}

	return nil
}

///////////////////// sub-categories ///////////

func (r repository) AddSubData(ctx context.Context, req category.SubCategoryIdDTO) error {

	q := `INSERT INTO products (cat_id, name, img, item) VALUES ($1, $2, $3, $4);`

	_, err := r.client.Exec(ctx, q, req.CatId, req.Name, req.Image, req.Item)
	if err != nil && err.Error() != "no rows in result set" {
		return err
	}

	return nil
}

func (r repository) UpdateSubData(ctx context.Context, req category.SubCategoryIdDTO) error {
	var (
		q string
	)

	q = ` update products set cat_id = $1, name = $2 , img = $3 , item = $4 where Id = $5;`

	_, err := r.client.Exec(ctx, q, req.CatId, req.Name, req.Image, req.Item, req.Id)

	if err != nil && err.Error() != "no rows in result set" {
		return err
	}

	return nil
}

func (r repository) DeleteSubData(ctx context.Context, Id string) error {
	var (
		q string
	)

	q = ` delete from products where Id = $1 ;`
	_, err := r.client.Exec(ctx, q, Id)

	if err != nil && err.Error() != "no rows in result set" {
		return err
	}

	return nil
}
