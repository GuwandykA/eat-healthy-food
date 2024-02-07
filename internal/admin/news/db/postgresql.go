package db

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"saglyk-backend/internal/admin/news"
	"saglyk-backend/pkg/logging"
)

type repository struct {
	client *pgxpool.Pool
	logger *logging.Logger
}

func NewRepository(client *pgxpool.Pool, logger *logging.Logger) news.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}

func (r repository) GetAllData(ctx context.Context, limit, page int) (news.DataNewsDTO, error) {

	ns := []news.GetNewsDTO{}
	var (
		dt   news.DataNewsDTO
		err  error
		rows pgx.Rows
	)

	qC := `select count(*) from news;`
	r.client.QueryRow(ctx, qC).Scan(&dt.Count)

	q := ` select 
				id,	img, title, content, "date"
			 from news
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
		var n news.GetNewsDTO
		errN := rows.Scan(
			&n.Id, &n.Image,
			&n.Title, &n.Content,
			&n.Date,
		)
		if errN != nil {
			r.logger.Error("errN :::", errN)
			continue
		}

		ns = append(ns, n)
	}
	dt.News = ns
	return dt, nil
}

func (r repository) GetOneData(ctx context.Context, id int) (news.GetNewsDTO, error) {

	var (
		n news.GetNewsDTO
	)

	q := ` select 
				id,	img, title, content, "date"
			 from news where id = $1 ; `

	err := r.client.QueryRow(ctx, q, id).Scan(
		&n.Id, &n.Image,
		&n.Title, &n.Content,
		&n.Date,
	)
	if err != nil {
		return n, err
	}
	return n, nil
}

func (r repository) AddData(ctx context.Context, req news.AddNewsDTO) (news.GetNewsDTO, error) {
	var (
		Id int
		n  news.GetNewsDTO
	)

	q := `INSERT INTO news (img, title, content, "date") VALUES ($1, $2, $3, $4) RETURNING  id;`

	err := r.client.QueryRow(ctx, q, req.Image, req.Title, req.Content, req.Date).Scan(
		&Id,
	)
	n, err = r.GetOneData(context.Background(), Id)

	if err != nil && err.Error() != "no rows in result set" {
		return n, err
	}

	return n, nil
}

func (r repository) UpdateData(ctx context.Context, req news.AddNewsDTO, id int) (news.GetNewsDTO, error) {
	var (
		q string
		n news.GetNewsDTO
	)

	q = ` update news set img = $1, title = $2, content = $3, "date" = $4 where id = $5;`

	_, err := r.client.Exec(ctx, q, req.Image, req.Title, req.Content, req.Date, id)
	n, err = r.GetOneData(context.Background(), id)
	if err != nil && err.Error() != "no rows in result set" {
		return n, err
	}

	return n, nil
}

func (r repository) DeleteData(ctx context.Context, uuid int) error {
	var (
		q string
	)

	q = ` delete from news where id = $1 ;`
	_, err := r.client.Exec(ctx, q, uuid)

	if err != nil && err.Error() != "no rows in result set" {
		return err
	}

	return nil
}
