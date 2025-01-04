package model

import (
	"fmt"
	"strconv"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"

	"github.com/mohammadbnei/vos-template/proto"
)

type Category struct {
	gorm.Model
	Name   string `gorm:"not null;uniqueIndex:category_unique_constraint"`
	UserID string `gorm:"not null;uniqueIndex:category_unique_constraint"`
	Color  string
}

// ToProto convert a Category model to its protobuf representation.
func (c Category) ToProto() *proto.Category {
	return &proto.Category{
		Id:        fmt.Sprintf("%d", c.ID), // Convert ID to string for protobuf,
		Name:      c.Name,
		UserId:    c.UserID,
		Color:     c.Color,
		CreatedAt: timestamppb.New(c.CreatedAt),
		UpdatedAt: timestamppb.New(c.UpdatedAt),
	}
}

// FromProto convert a protobuf Category message to its corresponding model.
func (c *Category) FromProto(protoCategory *proto.Category) error {
	id, err := strconv.ParseUint(protoCategory.GetId(), 10, 64)
	if err != nil {
		return errBuilder.Errorf("error parsing ID: %v", err)
	}

	c.ID = uint(id)
	c.Name = protoCategory.GetName()
	c.UserID = protoCategory.GetUserId()
	c.Color = protoCategory.GetColor()

	c.CreatedAt = protoCategory.GetCreatedAt().AsTime()
	c.UpdatedAt = protoCategory.GetUpdatedAt().AsTime()

	return nil
}
