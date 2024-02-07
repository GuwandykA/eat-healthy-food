package db

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"saglyk-backend/internal/admin/banner"
	"saglyk-backend/pkg/logging"
)

type repository struct {
	client *pgxpool.Pool
	logger *logging.Logger
}

func NewRepository(client *pgxpool.Pool, logger *logging.Logger) banner.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}

func (r repository) GetAllData(ctx context.Context, limit, page int) (banner.DataBannerDTO, error) {

	ns := []banner.GetBannerDTO{}
	var (
		dt   banner.DataBannerDTO
		err  error
		rows pgx.Rows
	)

	qC := `select count(*) from banner;`
	r.client.QueryRow(ctx, qC).Scan(&dt.Count)

	q := ` select 
				id,	img
			 from banner
		 order by created_at desc  `

	if limit == 0 || page == 0 {
		q += ";"
		rows, err = r.client.Query(ctx, q)
	} else {
		q += "limit $1 offset $2;"
		rows, err = r.client.Query(ctx, q, limit, (page-1)*limit)
	}

	defer rows.Close()
	if err != nil {
		return dt, nil
	}

	for rows.Next() {
		var n banner.GetBannerDTO
		errN := rows.Scan(
			&n.Id, &n.Image,
		)
		if errN != nil {
			r.logger.Error("errN :::", errN)
			continue
		}

		ns = append(ns, n)
	}
	dt.Categories = ns
	return dt, nil
}

func (r repository) GetOneData(ctx context.Context, id int) (banner.GetBannerDTO, error) {

	var (
		n banner.GetBannerDTO
	)

	q := ` select 
				id, img
			 from banner where id = $1 ; `

	err := r.client.QueryRow(ctx, q, id).Scan(
		&n.Id, &n.Image,
	)
	if err != nil {
		return n, err
	}
	return n, nil
}

func (r repository) AddData(ctx context.Context, req banner.AddBannerDTO) (banner.GetBannerDTO, error) {
	var (
		n banner.GetBannerDTO
	)

	q := `INSERT INTO banner (img) VALUES ($1) RETURNING  id, img;`

	err := r.client.QueryRow(ctx, q, req.Image).Scan(
		&n.Id, &n.Image,
	)
	if err != nil && err.Error() != "no rows in result set" {
		return n, err
	}

	return n, nil
}

func (r repository) UpdateData(ctx context.Context, req banner.AddBannerDTO, id int) (banner.GetBannerDTO, error) {
	var (
		q string
		n banner.GetBannerDTO
	)

	q = ` update banner set img = $1 where id = $2 RETURNING  id, img;`

	err := r.client.QueryRow(ctx, q, req.Image, id).Scan(
		&n.Id, &n.Image,
	)
	if err != nil && err.Error() != "no rows in result set" {
		return n, err
	}

	return n, nil
}

func (r repository) DeleteData(ctx context.Context, uuid int) error {
	var (
		q string
	)

	q = ` delete from banner where id = $1 ;`
	_, err := r.client.Exec(ctx, q, uuid)

	if err != nil && err.Error() != "no rows in result set" {
		return err
	}

	return nil
}
