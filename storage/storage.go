package storage

import (
	"context"
	ct "go_catalog_service/genproto/catalog_service"
	pd "go_catalog_service/genproto/product_service"
	rv "go_catalog_service/genproto/review_service"

	"google.golang.org/protobuf/types/known/emptypb"
)

type StorageI interface {
	CloseDB()
	Category() CategoryRepoI
	Product() ProductRepoI
	Review() ReviewRepoI
}

type CategoryRepoI interface {
	CreateCategory(context.Context, *ct.CreateCategory) (*ct.GetCategory, error)
	UpdateCategory(context.Context, *ct.UpdateCategory) (*ct.GetCategory, error)
	GetListCategory(context.Context, *ct.GetListCategoryRequest) (*ct.GetListCategoryResponse, error)
	GetCategoryById(context.Context, *ct.CategoryPrimaryKey) (*ct.GetCategory, error)
	DeleteCategory(context.Context, *ct.CategoryPrimaryKey) (emptypb.Empty, error)
}

type ProductRepoI interface {
	CreateProduct(context.Context, *pd.CreateProduct) (*pd.GetProduct, error)
	UpdateProduct(context.Context, *pd.UpdateProduct) (*pd.GetProduct, error)
	GetListProduct(context.Context, *pd.GetListProductRequest) (*pd.GetListProductResponse, error)
	GetProductById(context.Context, *pd.ProductPrimaryKey) (*pd.GetProduct, error)
	DeleteProduct(context.Context, *pd.ProductPrimaryKey) (emptypb.Empty, error)
}

type ReviewRepoI interface {
	CreateReview(context.Context, *rv.CreateReview) (*rv.GetReview, error)
	GetListReview(context.Context, *rv.GetListReviewRequest) (*rv.GetListReviewResponse, error)
	DeleteReview(context.Context, *rv.ReviewPrimaryKey) (emptypb.Empty, error)
}
