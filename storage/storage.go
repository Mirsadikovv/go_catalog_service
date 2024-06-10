package storage

import (
	"context"
	ct "go_catalog_service/genproto/catalog_service"

	"google.golang.org/protobuf/types/known/emptypb"
)

type StorageI interface {
	CloseDB()
	Category() CategoryRepoI
}

type CategoryRepoI interface {
	CreateCategory(context.Context, *ct.CreateCategory) (*ct.GetCategory, error)
	UpdateCategory(context.Context, *ct.UpdateCategory) (*ct.GetCategory, error)
	GetListCategory(context.Context, *ct.GetListCategoryRequest) (*ct.GetListCategoryResponse, error)
	GetCategoryById(context.Context, *ct.CategoryPrimaryKey) (*ct.GetCategory, error)
	DeleteCategory(context.Context, *ct.CategoryPrimaryKey) (emptypb.Empty, error)
}
