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
	FindBySlug(slug string) (Campaign, error)
	FindById(ID int) (Campaign, error)
	Updated(campaign Campaign) (Campaign, error)
	MarkAllCampaignImegesIsPrimaryFalse(campaignID int) (bool, error)
	SaveCampaignImages(campaignImages CampaignImages) (CampaignImages, error)
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
		totalPages = 1
	}
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

func (r *repository) Updated(campaign Campaign) (Campaign, error) {
	err := r.db.Save(&campaign).Error
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
	formatter := CampaignsFormatter(campaigns)
	paginate.Rows = formatter
	return &paginate, nil
}

func (r *repository) FindAll(paginate helper.Pagination) (*helper.Pagination, error) {
	var campaigns []Campaign
	err := r.db.Scopes(pagination(campaigns, &paginate, r.db)).Preload("CampaignImages", "is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return nil, err
	}
	formatter := CampaignsFormatter(campaigns)
	paginate.Rows = formatter
	return &paginate, nil
}

func (r *repository) FindBySlug(slug string) (Campaign, error) {
	campaign := Campaign{}
	err := r.db.Where("slug = ?", slug).Preload("CampaignImages").Preload("User").Find(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (r *repository) FindById(ID int) (Campaign, error) {
	campaign := Campaign{}
	err := r.db.Where("id = ?", ID).Find(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (r *repository) SaveCampaignImages(campaignImages CampaignImages) (CampaignImages, error) {
	err := r.db.Create(&campaignImages).Error
	if err != nil {
		return campaignImages, err
	}
	return campaignImages, nil
}

func (r *repository) MarkAllCampaignImegesIsPrimaryFalse(campaignID int) (bool, error) {
	err := r.db.Model(&CampaignImages{}).Where("campaign_id = ?", campaignID).Update("is_primary", false).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
