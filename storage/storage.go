package storage

import (
	"context"
	ct "go_catalog_service/genproto/catalog_service"
)

type StorageI interface {
	CloseDB()
	Category() CategoryRepoI
}

type CategoryRepoI interface {
	CreateCategory(ctx context.Context, req *ct.CreateCategory) (resp *ct.GetCategory, err error)
	UpdateCategory(ctx context.Context, req *ct.UpdateCategory) (resp *ct.GetCategory, err error)
	GetListCategory(ctx context.Context, req *ct.GetListCategoryRequest) (resp *ct.GetListCategoryResponse, err error)
	GetCategoryById(ctx context.Context, id *ct.CategoryPrimaryKey) (resp *ct.GetCategory, err error)
	DeleteCategory(ctx context.Context, id *ct.CategoryPrimaryKey) (err error)
}
