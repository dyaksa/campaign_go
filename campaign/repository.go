package campaign

import (
	"campaignproject/helper"
	"math"

	"gorm.io/gorm"
)

type Repository interface {
	Save(campaign Campaign) (Campaign, error)
	FindByUserId(ID int, paginate helper.Pagination) (*helper.Pagination, error)
	FindAll(paginate helper.Pagination) (*helper.Pagination, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func pagination(value interface{}, paginate *helper.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)
	paginate.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(paginate.TotalRows) / float64(paginate.Limit)))
	paginate.TotalPages = totalPages
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(paginate.GetOffset()).Limit(paginate.GetLimit()).Order(paginate.GetSort())
	}
}

func (r *repository) Save(campaign Campaign) (Campaign, error) {
	err := r.db.Create(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (r *repository) FindByUserId(ID int, paginate helper.Pagination) (*helper.Pagination, error) {
	var campaigns []Campaign
	err := r.db.Scopes(pagination(campaigns, &paginate, r.db)).Preload("CampaignImages", "is_primary = 1").Where("user_id = ?", ID).Find(&campaigns).Error
	if err != nil {
		return nil, err
	}
	formatter := CreateListFormatter(campaigns)
	paginate.Rows = formatter
	return &paginate, nil
}

func (r *repository) FindAll(paginate helper.Pagination) (*helper.Pagination, error) {
	var campaigns []Campaign
	err := r.db.Scopes(pagination(campaigns, &paginate, r.db)).Preload("CampaignImages").Find(&campaigns).Error
	if err != nil {
		return nil, err
	}
	formatter := CreateListFormatter(campaigns)
	paginate.Rows = formatter
	return &paginate, nil
}
