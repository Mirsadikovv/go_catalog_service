package service

import (
	"context"
	"go_catalog_service/config"
	"go_catalog_service/genproto/catalog_service"
	"go_catalog_service/grpc/client"
	"go_catalog_service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
)

type CategoryService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
}

// mustEmbedUnimplementedCategoryServiceServer implements catalog_service.CategoryServiceServer.
func (f *CategoryService) mustEmbedUnimplementedCategoryServiceServer() {
	panic("unimplemented")
}

func NewCategoryService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *CategoryService {
	return &CategoryService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (f *CategoryService) Create(ctx context.Context, req *catalog_service.CreateCategory) (resp *catalog_service.GetCategory, err error) {

	f.log.Info("---CreateCategory--->>>", logger.Any("req", req))

	resp, err = f.strg.Category().CreateCategory(ctx, req)
	if err != nil {
		f.log.Error("---CreateCategory--->>>", logger.Error(err))
		return &catalog_service.GetCategory{}, err
	}

	return resp, nil
}
func (f *CategoryService) Update(ctx context.Context, req *catalog_service.UpdateCategory) (resp *catalog_service.GetCategory, err error) {

	f.log.Info("---UpdateCategory--->>>", logger.Any("req", req))

	resp, err = f.strg.Category().UpdateCategory(ctx, req)
	if err != nil {
		f.log.Error("---UpdateCategory--->>>", logger.Error(err))
		return &catalog_service.GetCategory{}, err
	}

	return resp, nil
}

func (f *CategoryService) GetList(ctx context.Context, req *catalog_service.GetListCategoryRequest) (resp *catalog_service.GetListCategoryResponse, err error) {
	f.log.Info("---GetListCategory--->>>", logger.Any("req", req))

	resp, err = f.strg.Category().GetListCategory(ctx, req)
	if err != nil {
		f.log.Error("---GetListCategory--->>>", logger.Error(err))
		return &catalog_service.GetListCategoryResponse{}, err
	}

	return resp, nil
}

func (f *CategoryService) GetByID(ctx context.Context, id *catalog_service.CategoryPrimaryKey) (resp *catalog_service.GetCategory, err error) {
	f.log.Info("---GetCategory--->>>", logger.Any("req", id))

	resp, err = f.strg.Category().GetCategoryById(ctx, id)
	if err != nil {
		f.log.Error("---GetCategory--->>>", logger.Error(err))
		return &catalog_service.GetCategory{}, err
	}

	return resp, nil
}

func (f *CategoryService) Delete(ctx context.Context, req *catalog_service.CategoryPrimaryKey) (empty *catalog_service.Empty, err error) {

	f.log.Info("---DeleteCategory--->>>", logger.Any("req", req))

	err = f.strg.Category().DeleteCategory(ctx, req)
	if err != nil {
		f.log.Error("---DeleteCategory--->>>", logger.Error(err))
		return nil, err
	}

	return nil, nil
}
