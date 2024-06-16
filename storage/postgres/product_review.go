package postgres

import (
	"context"
	"database/sql"
	rv "go_catalog_service/genproto/review_service"
	"go_catalog_service/pkg"
	"go_catalog_service/storage"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type reviewRepo struct {
	db *pgxpool.Pool
}

func NewReviewRepo(db *pgxpool.Pool) storage.ReviewRepoI {
	return &reviewRepo{
		db: db,
	}
}

func (c *reviewRepo) CreateReview(ctx context.Context, req *rv.CreateReview) (*rv.GetReview, error) {

	id := uuid.NewString()

	_, err := c.db.Exec(ctx, `
		INSERT INTO product_reviews (
				id,
				customer_id,
				product_id,
				text,
				rating,
				order_id
		) VALUES ($1,$2,$3,$4,$5,$6
		)`,
		id,
		req.CustomerId,
		req.ProductId,
		req.Text,
		req.Rating,
		req.OrderId)
	if err != nil {
		log.Println("error while creating review")
		return nil, err
	}

	review, err := c.GetReviewById(ctx, &rv.ReviewPrimaryKey{Id: id})
	if err != nil {
		log.Println("error while getting review by id")
		return nil, err
	}
	return review, nil
}

func (c *reviewRepo) GetListReview(ctx context.Context, req *rv.GetListReviewRequest) (*rv.GetListReviewResponse, error) {
	reviews := rv.GetListReviewResponse{}
	var (
		created_at sql.NullString
	)
	filter := ""
	offest := (req.Offset - 1) * req.Limit
	if req.Search != "" {
		filter = ` AND text ILIKE '%` + req.Search + `%' `
	}
	query := `SELECT
				id,
				customer_id,
				product_id,
				text,
				rating,
				order_id,
				created_at
				FROM product_reviews
			WHERE product_id = $1 AND deleted_at IS NULL ` + filter + `
			OFFSET $2 LIMIT $3`
	rows, err := c.db.Query(ctx, query, req.ProductId, offest, req.Limit)
	if err != nil {
		log.Println("error while getting all reviews")
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var (
			review rv.GetReview
		)
		if err = rows.Scan(&review.Id,
			&review.CustomerId,
			&review.ProductId,
			&review.Text,
			&review.Rating,
			&review.OrderId,
			&created_at); err != nil {
			return &reviews, err
		}
		review.CreatedAt = pkg.NullStringToString(created_at)
		reviews.Reviews = append(reviews.Reviews, &review)

	}
	err = c.db.QueryRow(ctx, `SELECT count(*) from product_reviews WHERE product_id = $1 AND deleted_at IS NULL`+filter+``, req.ProductId).Scan(&reviews.Count)
	if err != nil {
		return &reviews, err
	}

	return &reviews, nil
}

func (c *reviewRepo) GetReviewById(ctx context.Context, id *rv.ReviewPrimaryKey) (*rv.GetReview, error) {
	var (
		review     rv.GetReview
		created_at sql.NullString
	)

	query := `SELECT
				id,
				customer_id,
				product_id,
				text,
				rating,
				order_id,
				created_at
				FROM product_reviews
			WHERE id = $1 AND deleted_at IS NULL`

	rows := c.db.QueryRow(ctx, query, id.Id)

	if err := rows.Scan(&review.Id,
		&review.CustomerId,
		&review.ProductId,
		&review.Text,
		&review.Rating,
		&review.OrderId,
		&created_at); err != nil {
		return &review, err
	}

	review.CreatedAt = pkg.NullStringToString(created_at)

	return &review, nil
}

func (c *reviewRepo) DeleteReview(ctx context.Context, id *rv.ReviewPrimaryKey) (emptypb.Empty, error) {

	_, err := c.db.Exec(ctx, `
		UPDATE product_reviews SET
		deleted_at = NOW()
		WHERE id = $1
		`,
		id.Id)

	if err != nil {
		return emptypb.Empty{}, err
	}
	return emptypb.Empty{}, nil
}
