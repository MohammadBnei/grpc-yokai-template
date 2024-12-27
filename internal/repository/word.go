package repository

import (
	"context"

	"github.com/samber/oops"
	"gorm.io/gorm"

	"github.com/mohammadbnei/vos-personal-words/internal/model"
)

type WordRepository struct {
	db *gorm.DB
}

func NewWordRepository(db *gorm.DB) *WordRepository {
	return &WordRepository{
		db: db,
	}
}

func (r *WordRepository) FindById(ctx context.Context, id uint) (*model.Word, error) {
	errBuilder := errBuilder.
		WithContext(ctx)

	var word model.Word

	res := r.db.WithContext(ctx).Preload("Category").First(&word, id)
	if res.Error != nil {
		return nil, errBuilder.Errorf("error finding word by id: %v", res.Error)
	}

	return &word, nil
}

func (r *WordRepository) Find(ctx context.Context, word model.Word) ([]model.Word, error) {
	errBuilder := errBuilder.
		WithContext(ctx)

	var words []*model.Word

	res := r.db.WithContext(ctx).Preload("Category").Find(&words)
	if res.Error != nil {
		return nil, errBuilder.Errorf("error finding words: %v", res.Error)
	}

	return convertPointerSliceToValueSlice(words), nil
}

func (r *WordRepository) Create(ctx context.Context, word model.Word) error {
	errBuilder := errBuilder.
		WithContext(ctx)

	res := r.db.WithContext(ctx).Create(word)

	return errBuilder.Errorf("error creating word: %v", res.Error)
}

func (r *WordRepository) Update(ctx context.Context, id uint, newWord model.Word) error {
	errBuilder := errBuilder.
		WithContext(ctx)

	res := r.db.WithContext(ctx).Model(&model.Word{}).Where("id = ?", id).Updates(newWord)
	if res.Error != nil {
		return errBuilder.Errorf("error updating word: %v", res.Error)
	}

	return nil
}

func (r *WordRepository) Delete(ctx context.Context, id uint) error {
	errBuilder := errBuilder.
		WithContext(ctx)

	res := r.db.WithContext(ctx).Delete(&model.Word{}, id)
	if res.Error != nil {
		return errBuilder.Errorf("error deleting word: %v", res.Error)
	}

	return nil
}

var errBuilder = oops.In("repository").
	Tags("database", "sql")

func convertPointerSliceToValueSlice[T any](ptrSlice []*T) []T {
	result := make([]T, len(ptrSlice))
	for i, ptr := range ptrSlice {
		result[i] = *ptr
	}
	return result
}
