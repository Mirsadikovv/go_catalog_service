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
	// Create(ctx context.Context, req *ct.CreateCategory) (resp *ct.GetCategory, err error)
	GetListCategory(ctx context.Context, req *ct.GetListCategoryRequest) (resp *ct.GetListCategoryResponse, err error)
	// GetAll(ctx context.Context, req *ct.GetListCategoryRequest) (resp *ct.GetListCategoryResponse, err error)
	// Update(ctx context.Context, req *ct.Category) (resp *ct.Category, err error)
	// Delete(ctx context.Context, req *ct.Category) error
}
