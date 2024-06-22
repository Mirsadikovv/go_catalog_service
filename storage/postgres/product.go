package postgres

import (
	"context"
	"database/sql"
	"fmt"
	pd "go_catalog_service/genproto/product_service"
	"go_catalog_service/pkg"
	"go_catalog_service/storage"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/jackc/pgx/v4/pgxpool"
)

type productRepo struct {
	db *pgxpool.Pool
}

func NewProductRepo(db *pgxpool.Pool) storage.ProductRepoI {
	return &productRepo{
		db: db,
	}
}

func (c *productRepo) CreateProduct(ctx context.Context, req *pd.CreateProduct) (*pd.GetProduct, error) {

	id := uuid.NewString()
	slug := slug.Make(req.NameEn)

	tx, err := c.db.Begin(ctx)
	if err != nil {
		log.Println("error in transaction create product", err)
	}
	resp, err := tx.Exec(ctx, `
		INSERT INTO product (
			id,
			slug,
			name_uz,
			name_ru,
			name_en,
			description_uz,
			description_ru,
			description_en,
			active,
			order_no,
			in_price,
			out_price,
			left_count,
			discount_percent
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14
		)`,
		id,
		slug,
		req.NameUz,
		req.NameRu,
		req.NameEn,
		req.DescriptionUz,
		req.DescriptionRu,
		req.DescriptionEn,
		req.Active,
		req.OrderNo,
		req.InPrice,
		req.OutPrice,
		req.LeftCount,
		req.DiscountPercent)

	if err != nil {
		tx.Rollback(ctx)
		log.Println("error while creating product", resp)
		return nil, err
	}
	pc_id := uuid.NewString()
	resp2, err := tx.Exec(ctx, `
	INSERT INTO product_categories (
		id,
		product_id,
		category_id
	) VALUES ($1,$2,$3)`,
		pc_id,
		id,
		req.CategoryId)
	if err != nil {
		tx.Rollback(ctx)
		log.Println("error while creating product_categories", resp2)
		return nil, err
	}
	err = tx.Commit(ctx)
	if err != nil {
		tx.Rollback(ctx)
		log.Println("error while creating product_categories transaction")
	}

	product, err := c.GetProductById(ctx, &pd.ProductPrimaryKey{Id: id})
	if err != nil {
		log.Println("error while getting product by id")
		return nil, err
	}

	return product, nil
}

func (c *productRepo) UpdateProduct(ctx context.Context, req *pd.UpdateProduct) (*pd.GetProduct, error) {

	slug := slug.Make(req.NameEn)

	tx, err := c.db.Begin(ctx)
	if err != nil {
		log.Println("error in transaction update product", err)
	}
	resp1, err := tx.Exec(ctx, `
	UPDATE product SET
		slug = $1,
		name_uz = $2,
		name_ru = $3,
		name_en = $4,
		description_uz = $5,
		description_ru = $6,
		description_en = $7,
		active = $8,
		order_no = $9,
		in_price = $10,
		out_price = $11,
		left_count = $12,
		discount_percent = $13,
		updated_at = NOW()
		WHERE id = $14
	`,
		slug,
		req.NameUz,
		req.NameRu,
		req.NameEn,
		req.DescriptionUz,
		req.DescriptionRu,
		req.DescriptionEn,
		req.Active,
		req.OrderNo,
		req.InPrice,
		req.OutPrice,
		req.LeftCount,
		req.DiscountPercent,
		req.Id)
	if err != nil {
		tx.Rollback(ctx)
		log.Println("error while updating product", resp1)
		return nil, err
	}
	resp2, err := tx.Exec(ctx, `
	UPDATE product_categories SET
		product_id = $1,
		category_id = $2
		WHERE product_id = $3
	`,
		req.Id,
		req.CategoryId,
		req.Id)
	if err != nil {
		tx.Rollback(ctx)
		log.Println("error while updating productCategories", resp2)
		return nil, err
	}
	err = tx.Commit(ctx)
	if err != nil {
		tx.Rollback(ctx)
		log.Println("error while updating product_categories transaction", err)
	}

	product, err := c.GetProductById(ctx, &pd.ProductPrimaryKey{Id: req.Id})
	if err != nil {
		log.Println("error while getting product by id")
		return nil, err
	}
	return product, nil
}

func (c *productRepo) GetListProduct(ctx context.Context, req *pd.GetListProductRequest) (*pd.GetListProductResponse, error) {
	products := pd.GetListProductResponse{}
	var (
		created_at sql.NullString
		updated_at sql.NullString
	)
	// filter_by_description := ""
	filter_by_name := ""
	offest := (req.Offset - 1) * req.Limit
	if req.Search != "" {
		filter_by_name = fmt.Sprintf(`AND (name_uz ILIKE '%%%v%%' OR name_ru ILIKE '%%%v%%' OR name_uz ILIKE '%%%v%%')`, req.Search, req.Search, req.Search)
		// filter_by_description = fmt.Sprintf(`AND (description_uz ILIKE '%%%v%%' OR description_ru ILIKE '%%%v%%' OR description_uz ILIKE '%%%v%%')`, req.Search, req.Search, req.Search)
	}
	query := `SELECT
				id,
				slug,
				name_uz,
				name_ru,
				name_en,
				description_uz,
				description_ru,
				description_en,
				active,
				order_no,
				in_price,
				out_price,
				left_count,
				discount_percent,
				created_at,
				updated_at
			FROM product
			WHERE TRUE AND deleted_at is null ` + filter_by_name + `
			OFFSET $1 LIMIT $2
`
	rows, err := c.db.Query(ctx, query, offest, req.Limit)

	if err != nil {
		log.Println("error while getting all products")
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var (
			product pd.GetProduct
		)
		if err = rows.Scan(&product.Id,
			&product.Slug,
			&product.NameUz,
			&product.NameRu,
			&product.NameEn,
			&product.DescriptionUz,
			&product.DescriptionRu,
			&product.DescriptionEn,
			&product.Active,
			&product.OrderNo,
			&product.InPrice,
			&product.OutPrice,
			&product.LeftCount,
			&product.DiscountPercent,
			&created_at,
			&updated_at,
		); err != nil {
			return &products, err
		}
		product.CreatedAt = pkg.NullStringToString(created_at)
		product.UpdatedAt = pkg.NullStringToString(updated_at)
		products.Products = append(products.Products, &product)
	}

	err = c.db.QueryRow(ctx, `SELECT count(*) from product WHERE TRUE AND deleted_at is null `+filter_by_name+``).Scan(&products.Count)
	if err != nil {
		return &products, err
	}

	return &products, nil
}

func (c *productRepo) GetProductById(ctx context.Context, id *pd.ProductPrimaryKey) (*pd.GetProduct, error) {
	var (
		product    pd.GetProduct
		created_at sql.NullString
		updated_at sql.NullString
	)

	query := `SELECT
				id,
				slug,
				name_uz,
				name_ru,
				name_en,
				description_uz,
				description_ru,
				description_en,
				active,
				order_no,
				in_price,
				out_price,
				left_count,
				discount_percent,
				created_at,
				updated_at
			FROM product
			WHERE id = $1 AND deleted_at IS NULL`

	rows := c.db.QueryRow(ctx, query, id.Id)

	if err := rows.Scan(&product.Id,
		&product.Slug,
		&product.NameUz,
		&product.NameRu,
		&product.NameEn,
		&product.DescriptionUz,
		&product.DescriptionRu,
		&product.DescriptionEn,
		&product.Active,
		&product.OrderNo,
		&product.InPrice,
		&product.OutPrice,
		&product.LeftCount,
		&product.DiscountPercent,
		&created_at,
		&updated_at); err != nil {
		return &product, err
	}

	product.CreatedAt = pkg.NullStringToString(created_at)
	product.UpdatedAt = pkg.NullStringToString(updated_at)

	return &product, nil
}

func (c *productRepo) DeleteProduct(ctx context.Context, id *pd.ProductPrimaryKey) (emptypb.Empty, error) {

	_, err := c.db.Exec(ctx, `
		UPDATE category SET
		active = false,
		deleted_at = NOW()
		WHERE id = $1
		`,
		id.Id)

	if err != nil {
		return emptypb.Empty{}, err
	}
	return emptypb.Empty{}, nil
}
