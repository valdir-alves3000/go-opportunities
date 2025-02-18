package repositories

import (
	"github.com/valdir-alves3000/go-opportunities/internal/core/schemas"
	"gorm.io/gorm"
)

type OpeningRepositoryImpl struct {
	db *gorm.DB
}

func NewOpeningRepository(db *gorm.DB) OpeningRepository {
	return &OpeningRepositoryImpl{db: db}
}

func (r *OpeningRepositoryImpl) Create(opening schemas.Opening) error {
	if err := r.db.Create(&opening).Error; err != nil {
		return err
	}
	return nil
}

func (r *OpeningRepositoryImpl) FindByID(id uint) (*schemas.Opening, error) {
	var opening schemas.Opening
	if err := r.db.First(&opening, id).Error; err != nil {
		return nil, err
	}
	return &opening, nil
}

func (r *OpeningRepositoryImpl) Update(opening schemas.Opening) error {
	if err := r.db.Save(&opening).Error; err != nil {
		return err
	}
	return nil
}

func (r *OpeningRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&schemas.Opening{}, id).Error
}

func (r *OpeningRepositoryImpl) FindAll(limit, offset int) ([]schemas.Opening, error) {
	var openings []schemas.Opening

	if err := r.db.Limit(limit).Offset(offset).Find(&openings).Error; err != nil {
		return nil, err
	}
	return openings, nil
}
