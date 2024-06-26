package grpc

import (
	"go_catalog_service/config"
	"go_catalog_service/genproto/catalog_service"
	"go_catalog_service/genproto/product_categories_service"
	"go_catalog_service/genproto/product_service"
	"go_catalog_service/genproto/review_service"

	"go_catalog_service/grpc/client"
	"go_catalog_service/grpc/service"
	"go_catalog_service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvc client.ServiceManagerI) (grpcServer *grpc.Server) {

	grpcServer = grpc.NewServer()

	catalog_service.RegisterCategoryServiceServer(grpcServer, service.NewCategoryService(cfg, log, strg, srvc))
	product_service.RegisterProductServiceServer(grpcServer, service.NewProductService(cfg, log, strg, srvc))
	review_service.RegisterReviewServiceServer(grpcServer, service.NewReviewService(cfg, log, strg, srvc))
	product_categories_service.RegisterProductCategoriesServiceServer(grpcServer, service.NewProductCategoriesService(cfg, log, strg, srvc))

	reflection.Register(grpcServer)
	return
}
