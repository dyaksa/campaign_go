package transaction

import (
	"campaignproject/helper"
	"math"

	"gorm.io/gorm"
)

type Repository interface {
	GetByCampaignID(input GetCampaignID, paginate helper.Pagination) (*helper.Pagination, error)
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
	if totalPages < 0 {
		totalPages = 0
	}
	paginate.TotalPages = totalPages
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(paginate.GetOffset()).Limit(paginate.GetLimit()).Order(paginate.GetSort())
	}
}

func (r *repository) GetByCampaignID(input GetCampaignID, paginate helper.Pagination) (*helper.Pagination, error) {
	var transactions []Transaction
	err := r.db.Scopes(pagination(transactions, &paginate, r.db)).Preload("User").Where("campaign_id = ?", input.ID).Find(&transactions).Error
	if err != nil {
		return &paginate, err
	}
	campaignFormatter := FormatCampaignTransactions(transactions)
	paginate.Rows = campaignFormatter
	return &paginate, nil
}
