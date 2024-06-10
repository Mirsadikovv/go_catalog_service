package postgres

import (
	"context"
	"database/sql"
	ct "go_catalog_service/genproto/catalog_service"
	"go_catalog_service/pkg"
	"go_catalog_service/storage"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/jackc/pgx/v4/pgxpool"
)

type categoryRepo struct {
	db *pgxpool.Pool
}

func NewCategoryRepo(db *pgxpool.Pool) storage.CategoryRepoI {
	return &categoryRepo{
		db: db,
	}
}

func (c *categoryRepo) CreateCategory(ctx context.Context, req *ct.CreateCategory) (*ct.GetCategory, error) {

	id := uuid.NewString()
	slug := slug.Make(req.NameEn)
	if req.ParentId == "" {
		req.ParentId = id
	}

	_, err := c.db.Exec(ctx, `
		INSERT INTO category (
			id,
			slug,
			name_uz,
			name_ru,
			name_en,
			active,
			order_no,
			parent_id
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8
		)`,
		id,
		slug,
		string(req.NameUz),
		string(req.NameRu),
		string(req.NameEn),
		req.Active,
		req.OrderNo,
		string(req.ParentId))
	if err != nil {
		log.Println("error while creating category")
		return nil, err
	}

	category, err := c.GetCategoryById(ctx, &ct.CategoryPrimaryKey{Id: id})
	if err != nil {
		log.Println("error while getting category by id")
		return nil, err
	}
	return category, nil
}

func (c *categoryRepo) UpdateCategory(ctx context.Context, req *ct.UpdateCategory) (*ct.GetCategory, error) {

	id := uuid.NewString()
	slug := slug.Make(req.NameEn)
	if req.ParentId == "" {
		req.ParentId = id
	}

	_, err := c.db.Exec(ctx, `
		UPDATE category SET
		slug = $1,
		name_uz = $2,
		name_ru = $3,
		name_en = $4,
		active = $5,
		order_no = $6,
		parent_id = $7,
		updated_at = NOW()
		WHERE id = $8
		`,
		slug,
		req.NameUz,
		req.NameRu,
		req.NameEn,
		req.Active,
		req.OrderNo,
		req.ParentId,
		id)
	if err != nil {
		log.Println("error while updating category")
		return nil, err
	}

	category, err := c.GetCategoryById(ctx, &ct.CategoryPrimaryKey{Id: id})
	if err != nil {
		log.Println("error while getting category by id")
		return nil, err
	}
	return category, nil
}

func (c *categoryRepo) GetListCategory(ctx context.Context, req *ct.GetListCategoryRequest) (*ct.GetListCategoryResponse, error) {
	categories := ct.GetListCategoryResponse{}
	filter := ""
	offest := (req.Offset - 1) * req.Limit
	if req.Search != "" {
		filter = ` AND slug ILIKE '%` + req.Search + `%' `
	}
	query := `SELECT
				id,
				slug,
				name_uz,
				name_ru,
				name_en,
				active,
				order_no,
				parent_id
			FROM category
			WHERE TRUE ` + filter + `
			OFFSET $1 LIMIT $2
`
	rows, err := c.db.Query(ctx, query, offest, req.Limit)

	if err != nil {
		log.Println("error while getting category by id")
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var (
			category  ct.GetCategory
			parent_id sql.NullString
		)
		if err = rows.Scan(&category.Id,
			&category.Slug,
			&category.NameUz,
			&category.NameRu,
			&category.NameEn,
			&category.Active,
			&category.OrderNo,
			&parent_id); err != nil {
			return &categories, err
		}
		category.ParentId = pkg.NullStringToString(parent_id)
		categories.Categories = append(categories.Categories, &category)
	}

	err = c.db.QueryRow(ctx, `SELECT count(*) from category WHERE TRUE `+filter+``).Scan(&categories.Count)
	if err != nil {
		return &categories, err
	}

	return &categories, nil
}

func (c *categoryRepo) GetCategoryById(ctx context.Context, id *ct.CategoryPrimaryKey) (*ct.GetCategory, error) {
	var (
		category   ct.GetCategory
		created_at sql.NullString
		updated_at sql.NullString
	)

	query := `SELECT
				id,
				slug,
				name_uz,
				name_ru,
				name_en,
				active,
				order_no,
				parent_id,
				created_at,
				updated_at
			FROM category
			WHERE id = $1`

	rows := c.db.QueryRow(ctx, query, id.Id)

	if err := rows.Scan(&category.Id,
		&category.Slug,
		&category.NameUz,
		&category.NameRu,
		&category.NameEn,
		&category.Active,
		&category.OrderNo,
		&category.ParentId,
		&created_at,
		&updated_at); err != nil {
		return &category, err
	}

	category.CreatedAt = pkg.NullStringToString(created_at)
	category.UpdatedAt = pkg.NullStringToString(updated_at)

	return &category, nil
}

func (c *categoryRepo) DeleteCategory(ctx context.Context, id *ct.CategoryPrimaryKey) (emptypb.Empty, error) {

	_, err := c.db.Exec(ctx, `
		UPDATE category SET
		deleted_at = NOW()
		WHERE id = $1
		`,
		id.Id)

	if err != nil {
		return emptypb.Empty{}, err
	}
	return emptypb.Empty{}, nil
}
