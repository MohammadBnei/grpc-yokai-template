package model

import (
	"fmt"
	"strconv"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"

	"github.com/mohammadbnei/vos-template/proto"
)

type Word struct {
	gorm.Model
	Word       string     `gorm:"not null;uniqueIndex:word_unique_constraint"`
	Categories []Category `gorm:"many2many:category_words"`
	UserID     string     `gorm:"not null;uniqueIndex:word_unique_constraint"`
}

// ToProto convert a Word model to its protobuf representation.
func (w Word) ToProto() *proto.Word {
	categoryNames := make([]string, len(w.Categories))

	for i, category := range w.Categories {
		categoryNames[i] = category.Name
	}

	return &proto.Word{
		Id:            fmt.Sprintf("%d", w.ID),
		Word:          w.Word,
		UserId:        w.UserID,
		CategoryNames: categoryNames,
		CreatedAt:     timestamppb.New(w.CreatedAt),
		UpdatedAt:     timestamppb.New(w.UpdatedAt),
	}
}

// FromProto convert a protobuf Word message to its corresponding model.
func (w *Word) FromProto(protoWord *proto.Word) error {
	id, err := strconv.ParseUint(protoWord.GetId(), 10, 64)
	if err != nil {
		return errBuilder.Errorf("error parsing ID: %v", err)
	}

	w.ID = uint(id)
	w.Word = protoWord.GetWord()
	w.UserID = protoWord.GetUserId()

	categories := make([]Category, len(protoWord.GetCategoryNames()))

	for i, categoryName := range protoWord.GetCategoryNames() {
		categories[i] = Category{Name: categoryName, UserID: w.UserID}
	}

	w.Categories = categories

	w.CreatedAt = protoWord.GetCreatedAt().AsTime()
	w.UpdatedAt = protoWord.GetUpdatedAt().AsTime()

	return nil
}
