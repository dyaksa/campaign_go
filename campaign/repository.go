package campaign

import "gorm.io/gorm"

type Repository interface {
	Save(campaign Campaign) (Campaign, error)
	Get(campaign Campaign) Campaign
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Save(campaign Campaign) (Campaign, error) {
	err := r.db.Create(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (r *repository) Get(campaign Campaign) Campaign {
	var data []Campaign
	r.db.Find(&data)
	return campaign
}
