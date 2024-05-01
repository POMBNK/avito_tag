package banner

import "github.com/POMBNK/avitotag/internal/entity"

func BannerDTOToEntity(banner *PostBannerJSONBody) *entity.Banner {
	return &entity.Banner{
		Content:   banner.Content,
		FeatureId: banner.FeatureId,
		IsActive:  banner.IsActive,
		TagIds:    banner.TagIds,
	}
}
