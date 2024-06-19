package service

import (
	"context"
	"go_catalog_service/config"
	"go_catalog_service/genproto/product_categories_service"

	"go_catalog_service/grpc/client"
	"go_catalog_service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ProductCategoriesService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
}

func NewProductCategoriesService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *ProductCategoriesService {
	return &ProductCategoriesService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (f *ProductCategoriesService) Create(ctx context.Context, req *product_categories_service.CreateProductCategories) (*product_categories_service.GetProductCategories, error) {

	f.log.Info("---CreateProductCategories--->>>", logger.Any("req", req))

	resp, err := f.strg.ProductCategories().CreateProductCategories(ctx, req)
	if err != nil {
		f.log.Error("---CreateProductCategories--->>>", logger.Error(err))
		return &product_categories_service.GetProductCategories{}, err
	}

	return resp, nil
}
func (f *ProductCategoriesService) Update(ctx context.Context, req *product_categories_service.UpdateProductCategories) (*product_categories_service.GetProductCategories, error) {

	f.log.Info("---UpdateProductCategories--->>>", logger.Any("req", req))

	resp, err := f.strg.ProductCategories().UpdateProductCategories(ctx, req)
	if err != nil {
		f.log.Error("---UpdateProductCategories--->>>", logger.Error(err))
		return &product_categories_service.GetProductCategories{}, err
	}

	return resp, nil
}

func (f *ProductCategoriesService) GetByID(ctx context.Context, id *product_categories_service.ProductCategoriesPrimaryKey) (*product_categories_service.GetProductCategories, error) {
	f.log.Info("---GetProductCategories--->>>", logger.Any("req", id))

	resp, err := f.strg.ProductCategories().GetProductCategoriesById(ctx, id)
	if err != nil {
		f.log.Error("---GetProductCategories--->>>", logger.Error(err))
		return &product_categories_service.GetProductCategories{}, err
	}

	return resp, nil
}

func (f *ProductCategoriesService) Delete(ctx context.Context, req *product_categories_service.ProductCategoriesPrimaryKey) (*emptypb.Empty, error) {

	f.log.Info("---DeleteProductCategories--->>>", logger.Any("req", req))

	_, err := f.strg.ProductCategories().DeleteProductCategories(ctx, req)
	if err != nil {
		f.log.Error("---DeleteProductCategories--->>>", logger.Error(err))
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
