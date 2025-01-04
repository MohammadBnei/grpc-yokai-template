package model

import (
	"fmt"
	"strconv"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"

	"github.com/mohammadbnei/vos-template/proto"
)

type List struct {
	gorm.Model
	Name   string `gorm:"not null;uniqueIndex:list_unique_constraint"`
	Words  []Word `gorm:"many2many:word_lists"`
	UserID string `gorm:"not null;uniqueIndex:list_unique_constraint"`
}

// ToProto convert a List model to its protobuf representation.
func (l List) ToProto() *proto.List {
	words := make([]string, len(l.Words))

	for i, word := range l.Words {
		words[i] = word.Word
	}

	return &proto.List{
		Id:        fmt.Sprintf("%d", l.ID),
		Name:      l.Name,
		Words:     words,
		UserId:    l.UserID,
		CreatedAt: timestamppb.New(l.CreatedAt),
		UpdatedAt: timestamppb.New(l.UpdatedAt),
	}
}

// FromProto convert a protobuf List message to its corresponding model.
func (l *List) FromProto(protoList *proto.List) error {
	id, err := strconv.ParseUint(protoList.GetId(), 10, 64)
	if err != nil {
		return errBuilder.Errorf("error parsing ID: %v", err)
	}

	l.ID = uint(id)
	l.Name = protoList.GetName()
	l.UserID = protoList.GetUserId()

	words := make([]Word, len(protoList.GetWords()))

	for i, word := range protoList.GetWords() {
		words[i] = Word{Word: word, UserID: l.UserID}
	}

	l.Words = words

	l.CreatedAt = protoList.GetCreatedAt().AsTime()
	l.UpdatedAt = protoList.GetUpdatedAt().AsTime()

	return nil
}
