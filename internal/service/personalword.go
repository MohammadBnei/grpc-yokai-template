package service

import (
	"context"
	"slices"

	"github.com/ankorstore/yokai/config"
	"github.com/ankorstore/yokai/log"
	"github.com/samber/oops"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"github.com/mohammadbnei/vos-template/internal/model"
	"github.com/mohammadbnei/vos-template/proto"
)

// ExampleService is the gRPC server service for ExampleService.
type PersonalService struct {
	proto.UnimplementedPersonalWordServiceServer
	config *config.Config
	db     *gorm.DB
}

// NewPersonalService returns a new [PersonalService] instance.
func NewPersonalService(cfg *config.Config, db *gorm.DB) *PersonalService {
	return &PersonalService{
		config: cfg,
		db:     db,
	}
}

// SaveWord saves a word for a user.
//
// The word is saved with the given word and user id.
func (ps *PersonalService) SaveWord(ctx context.Context, req *proto.SaveWordRequest) (*proto.SaveWordResponse, error) {
	errBuilder := errBuilder.
		WithContext(ctx)

	switch "" {
	case req.GetWord():
		return nil, status.Error(codes.InvalidArgument, "word cannot be empty")
	case req.GetUserId():
		return nil, status.Error(codes.InvalidArgument, "user id cannot be empty")
	}

	word := model.Word{
		Word:   req.GetWord(),
		UserID: req.GetUserId(),
	}

	if err := ps.db.WithContext(ctx).FirstOrCreate(&word, word).Error; err != nil {
		log.CtxLogger(ctx).Error().Err(errBuilder.Errorf("error creating word: %v", err)).Msg("error creating word")

		return nil, status.Error(codes.Unknown, "error creating word")
	}

	listName := req.GetListName()
	if listName == "" {
		listName = "default"
	}

	list := model.List{Name: listName, UserID: req.GetUserId()}

	if err := ps.db.WithContext(ctx).Preload("Words").FirstOrCreate(&list, list).Error; err != nil {
		log.CtxLogger(ctx).Error().Err(errBuilder.Errorf("error fetching list: %v", err)).Msg("error fetching list")

		return nil, status.Error(codes.Unknown, "error fetching list")
	}

	if !slices.ContainsFunc(list.Words, func(w model.Word) bool { return w.ID == word.ID }) {
		list.Words = append(list.Words, word)
	}

	if err := ps.db.WithContext(ctx).Save(&list).Error; err != nil {
		log.CtxLogger(ctx).Error().Err(errBuilder.Errorf("error saving list: %v", err)).Msg("error saving list")

		return nil, status.Error(codes.Unknown, "error saving list")
	}

	return &proto.SaveWordResponse{
		List: list.ToProto(),
		Word: word.ToProto(),
	}, nil
}

var errBuilder = oops.In("service").
	Tags("personalword")
