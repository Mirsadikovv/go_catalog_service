package service

import (
	"context"
	"go_catalog_service/config"
	"go_catalog_service/genproto/product_categories_service"
	"go_catalog_service/genproto/product_service"

	"go_catalog_service/grpc/client"
	"go_catalog_service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ProductService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
}

func NewProductService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *ProductService {
	return &ProductService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (f *ProductService) Create(ctx context.Context, req *product_service.CreateProduct) (*product_service.GetProduct, error) {

	f.log.Info("---CreateProduct--->>>", logger.Any("req", req))

	resp, err := f.strg.Product().CreateProduct(ctx, req)
	if err != nil {
		f.log.Error("---CreateProduct--->>>", logger.Error(err))
		return resp, err
	}
	req2 := &product_categories_service.CreateProductCategories{}
	req2.ProductId = resp.Id
	req2.CategoryId = req.CategoryId
	resp2, err := f.strg.ProductCategories().CreateProductCategories(ctx, req2)
	if err != nil {
		f.log.Error("---CreateProduct--->>>", logger.Error(err))
	}
	println(resp2)

	return resp, nil
}
func (f *ProductService) Update(ctx context.Context, req *product_service.UpdateProduct) (*product_service.GetProduct, error) {

	f.log.Info("---UpdateProduct--->>>", logger.Any("req", req))

	resp, err := f.strg.Product().UpdateProduct(ctx, req)
	if err != nil {
		f.log.Error("---UpdateProduct--->>>", logger.Error(err))
		return &product_service.GetProduct{}, err
	}

	return resp, nil
}

func (f *ProductService) GetList(ctx context.Context, req *product_service.GetListProductRequest) (*product_service.GetListProductResponse, error) {
	f.log.Info("---GetListProduct--->>>", logger.Any("req", req))

	resp, err := f.strg.Product().GetListProduct(ctx, req)
	if err != nil {
		f.log.Error("---GetListProduct--->>>", logger.Error(err))
		return &product_service.GetListProductResponse{}, err
	}

	return resp, nil
}

func (f *ProductService) GetByID(ctx context.Context, id *product_service.ProductPrimaryKey) (*product_service.GetProduct, error) {
	f.log.Info("---GetProduct--->>>", logger.Any("req", id))

	resp, err := f.strg.Product().GetProductById(ctx, id)
	if err != nil {
		f.log.Error("---GetProduct--->>>", logger.Error(err))
		return &product_service.GetProduct{}, err
	}

	return resp, nil
}

func (f *ProductService) Delete(ctx context.Context, req *product_service.ProductPrimaryKey) (*emptypb.Empty, error) {

	f.log.Info("---DeleteProduct--->>>", logger.Any("req", req))

	_, err := f.strg.Product().DeleteProduct(ctx, req)
	if err != nil {
		f.log.Error("---DeleteProduct--->>>", logger.Error(err))
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
