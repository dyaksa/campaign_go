package transaction

import "campaignproject/user"

type GetCampaignID struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}
