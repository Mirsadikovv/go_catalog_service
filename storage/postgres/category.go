package postgres

import (
	"context"
	"database/sql"
	"fmt"
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
	var parentId sql.NullString
	if req.ParentId == "" {
		parentId = sql.NullString{Valid: false} // Устанавливаем NULL значение
	} else {
		parentId = sql.NullString{String: req.ParentId, Valid: true} // Устанавливаем значение
	}
	// if req.ParentId == "" {
	// 	req.ParentId = id
	// }

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
		parentId)
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

	slug := slug.Make(req.NameEn)

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
		req.Id)
	if err != nil {
		log.Println("error while updating category")
		return nil, err
	}

	category, err := c.GetCategoryById(ctx, &ct.CategoryPrimaryKey{Id: req.Id})
	if err != nil {
		log.Println("error while getting category by id")
		return nil, err
	}
	return category, nil
}

func (c *categoryRepo) GetListCategory(ctx context.Context, req *ct.GetListCategoryRequest) (*ct.GetListCategoryResponse, error) {
	categories := ct.GetListCategoryResponse{}
	filter_by_name := ""
	offest := (req.Offset - 1) * req.Limit
	if req.Search != "" {
		filter_by_name = fmt.Sprintf(` AND (name_uz ILIKE '%%%v%%' OR name_ru ILIKE '%%%v%%' OR name_en ILIKE '%%%v%%')`, req.Search, req.Search, req.Search)
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
			WHERE parent_id IS NULL AND deleted_at IS NULL ` + filter_by_name + `
			OFFSET $1 LIMIT $2
`
	rows, err := c.db.Query(ctx, query, offest, req.Limit)

	if err != nil {
		log.Println("error while getting all categories")
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
		primaryKey := &ct.CategoryPrimaryKey{
			Id: category.Id,
		}

		childCategories, err := c.GetCategoryById(ctx, primaryKey)
		if err != nil {
			return &categories, err
		}
		category.ChildCategories = childCategories.ChildCategories
		categories.Categories = append(categories.Categories, &category)
	}

	if err = rows.Err(); err != nil {
		log.Println("error while getting all categories", err)
	}

	err = c.db.QueryRow(ctx, `SELECT count(*) from category WHERE TRUE AND deleted_at is null`+filter_by_name+``).Scan(&categories.Count)
	if err != nil {
		return &categories, err
	}

	return &categories, nil
}

func (c *categoryRepo) GetCategoryById(ctx context.Context, id *ct.CategoryPrimaryKey) (*ct.GetCategory, error) {
	var (
		category   ct.GetCategory
		parent_id  sql.NullString
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
			WHERE id = $1 AND deleted_at IS NULL`

	rows := c.db.QueryRow(ctx, query, id.Id)

	if err := rows.Scan(&category.Id,
		&category.Slug,
		&category.NameUz,
		&category.NameRu,
		&category.NameEn,
		&category.Active,
		&category.OrderNo,
		&parent_id,
		&created_at,
		&updated_at); err != nil {
		return &category, err
	}
	query1 := `SELECT
				id,
				slug,
				name_uz,
				name_ru,
				name_en,
				active,
				order_no,
				parent_id
			FROM category
			WHERE parent_id = $1`
	rows1, err := c.db.Query(ctx, query1, id.Id)

	if err != nil {
		log.Println("error while getting all child categories")
		return nil, err
	}
	defer rows1.Close()
	for rows1.Next() {
		var (
			childCategory   ct.GetCategory
			child_parent_id sql.NullString
		)
		if err = rows1.Scan(&childCategory.Id,
			&childCategory.Slug,
			&childCategory.NameUz,
			&childCategory.NameRu,
			&childCategory.NameEn,
			&childCategory.Active,
			&childCategory.OrderNo,
			&child_parent_id); err != nil {
			return &childCategory, err
		}
		childCategory.ParentId = pkg.NullStringToString(child_parent_id)
		category.ChildCategories = append(category.ChildCategories, &childCategory)
	}
	if err = rows1.Err(); err != nil {
		log.Println("error while getting child categories", err)
	}
	category.ParentId = pkg.NullStringToString(parent_id)
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
