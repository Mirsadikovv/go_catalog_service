package postgres

import (
	"context"
	"database/sql"
	ct "go_catalog_service/genproto/catalog_service"
	"go_catalog_service/pkg"
	"go_catalog_service/storage"
	"log"

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

// func (c *categoryRepo) Create(ctx context.Context, req *ct.CreateCategory) (resp *ct.Category, err error) {

// 	resp = &ct.Category{}

// 	id := uuid.NewString()

// 	if req.ParentId == "" {
// 		req.ParentId = id
// 	}

// 	_, err = c.db.Exec(ctx, `
// 		INSERT INTO category (
// 			id,
// 			slug,
// 			name_uz,
// 			name_ru,
// 			name_en,
// 			active,
// 			order_no,
// 			parent_id
// 		) VALUES (
// 			$1,
// 			$2,
// 			$3,
// 			$4,
// 			$5,
// 			$6,
// 			$7,
// 			$8
// 		) `, id, req.Slug, req.NameUz, req.NameRu, req.NameEn, req.Active, req.OrderNo, req.ParentId)

// 	if err != nil {
// 		log.Println("error while creating category")
// 		return nil, err
// 	}

// 	category, err := c.GetByID(ctx, &ct.CategoryPrimaryKey{Id: id})
// 	if err != nil {
// 		log.Println("error while getting category by id")
// 		return nil, err
// 	}

// 	return category, nil
// }

// func (c *categoryRepo) GetListCategory(ctx context.Context, req *ct.CategoryPrimaryKey) (resp *ct.GetCategory, err error) {

// 	resp = &ct.GetCategory{}

// 	var ParentId sql.NullString

// 	err = c.db.QueryRow(ctx, `
// 		SELECT
// 			id,
// 			slug,
// 			name_uz,
// 			name_ru,
// 			name_en,
// 			active,
// 			order_no,
// 			parent_id
// 		FROM category
// 		WHERE id = $1
// 	`, req.Id).Scan(&resp.Id, &resp.Slug, &resp.NameUz, &resp.NameRu, &resp.NameEn, &resp.Active, &resp.OrderNo, &ParentId)

// 	if err != nil {
// 		log.Println("error while getting category by id")
// 		return nil, err
// 	}

// 	resp.ParentId = ParentId.String

// 	return resp, nil
// }

func (c *categoryRepo) GetListCategory(ctx context.Context, req *ct.GetListCategoryRequest) (resp *ct.GetListCategoryResponse, err error) {
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
			return resp, err
		}
		category.ParentId = pkg.NullStringToString(parent_id)
		categories.Categories = append(categories.Categories, &category)
	}

	err = c.db.QueryRow(ctx, `SELECT count(*) from category WHERE TRUE `+filter+``).Scan(&categories.Count)
	if err != nil {
		return resp, err
	}

	return &categories, nil
}
