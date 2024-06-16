package service

import (
	"context"
	"go_catalog_service/config"
	"go_catalog_service/genproto/review_service"

	"go_catalog_service/grpc/client"
	"go_catalog_service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ReviewService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
}

func NewReviewService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *ReviewService {
	return &ReviewService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (f *ReviewService) Create(ctx context.Context, req *review_service.CreateReview) (*review_service.GetReview, error) {

	f.log.Info("---CreateReview--->>>", logger.Any("req", req))

	resp, err := f.strg.Review().CreateReview(ctx, req)
	if err != nil {
		f.log.Error("---CreateReview--->>>", logger.Error(err))
		return &review_service.GetReview{}, err
	}

	return resp, nil
}

func (f *ReviewService) GetList(ctx context.Context, req *review_service.GetListReviewRequest) (*review_service.GetListReviewResponse, error) {
	f.log.Info("---GetListReview--->>>", logger.Any("req", req))

	resp, err := f.strg.Review().GetListReview(ctx, req)
	if err != nil {
		f.log.Error("---GetListReview--->>>", logger.Error(err))
		return &review_service.GetListReviewResponse{}, err
	}

	return resp, nil
}

func (f *ReviewService) Delete(ctx context.Context, req *review_service.ReviewPrimaryKey) (*emptypb.Empty, error) {

	f.log.Info("---DeleteReview--->>>", logger.Any("req", req))

	_, err := f.strg.Review().DeleteReview(ctx, req)
	if err != nil {
		f.log.Error("---DeleteReview--->>>", logger.Error(err))
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
