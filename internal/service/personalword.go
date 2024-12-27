package service

import (
	"context"

	"github.com/ankorstore/yokai/config"
	"gorm.io/gorm"

	"github.com/mohammadbnei/vos-personal-words/internal/model"
	"github.com/mohammadbnei/vos-personal-words/proto"
)

// ExampleService is the gRPC server service for ExampleService.
type PersonalService struct {
	proto.UnimplementedPersonalWordServiceServer
	config *config.Config
	db     *gorm.DB
}

// NewPersonalService returns a new [PersonalService] instance.
func NewPersonalService(cfg *config.Config) *PersonalService {
	return &PersonalService{
		config: cfg,
	}
}

func (ps *PersonalService) SaveWord(ctx context.Context, req *proto.SaveWordRequest) (*proto.SaveWordResponse, error) {
	var word model.Word

	word.Word = req.GetWord()
	word.UserID = req.GetUserId()

	return &proto.SaveWordResponse{}, nil
}
