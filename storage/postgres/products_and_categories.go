package postgres

import (
	"context"
	pc "go_catalog_service/genproto/product_categories_service"
	"go_catalog_service/storage"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type productCategoriesRepo struct {
	db *pgxpool.Pool
}

func NewProductCategoriesRepo(db *pgxpool.Pool) storage.ProductCategoriesRepoI {
	return &productCategoriesRepo{
		db: db,
	}
}

func (c *productCategoriesRepo) CreateProductCategories(ctx context.Context, req *pc.CreateProductCategories) (*pc.GetProductCategories, error) {

	id := uuid.NewString()

	_, err := c.db.Exec(ctx, `
		INSERT INTO product_categories (
			id,
			product_id,
			category_id
		) VALUES ($1,$2,$3)`,
		id,
		req.ProductId,
		req.CategoryId)
	if err != nil {
		log.Println("error while creating product_categories")
		return nil, err
	}

	productCategories, err := c.GetProductCategoriesById(ctx, &pc.ProductCategoriesPrimaryKey{Id: id})
	if err != nil {
		log.Println("error while getting productCategories by id")
		return nil, err
	}
	return productCategories, nil
}

func (c *productCategoriesRepo) UpdateProductCategories(ctx context.Context, req *pc.UpdateProductCategories) (*pc.GetProductCategories, error) {

	_, err := c.db.Exec(ctx, `
	UPDATE product_categories SET
		product_id = $1,
		category_id = $2
		WHERE id = $3
	`,
		req.ProductId,
		req.CategoryId,
		req.Id)
	if err != nil {
		log.Println("error while updating productCategories")
		return nil, err
	}

	productCategories, err := c.GetProductCategoriesById(ctx, &pc.ProductCategoriesPrimaryKey{Id: req.Id})
	if err != nil {
		log.Println("error while getting productCategories by id")
		return nil, err
	}
	return productCategories, nil
}

func (c *productCategoriesRepo) GetProductCategoriesById(ctx context.Context, id *pc.ProductCategoriesPrimaryKey) (*pc.GetProductCategories, error) {
	var (
		productCategories pc.GetProductCategories
	)

	query := `SELECT
				id,
				product_id,
				category_id
			FROM product_categories
			WHERE id = $1`

	rows := c.db.QueryRow(ctx, query, id.Id)

	if err := rows.Scan(
		&productCategories.Id,
		&productCategories.ProductId,
		&productCategories.CategoryId); err != nil {
		return &productCategories, err
	}

	return &productCategories, nil
}

func (c *productCategoriesRepo) DeleteProductCategories(ctx context.Context, id *pc.ProductCategoriesPrimaryKey) (emptypb.Empty, error) {

	_, err := c.db.Exec(ctx, `
	DELETE 
	FROM product_categories 
	WHERE id = $1
		`,
		id.Id)

	if err != nil {
		return emptypb.Empty{}, err
	}
	return emptypb.Empty{}, nil
}
