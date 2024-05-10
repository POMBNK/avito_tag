package banner

import "github.com/POMBNK/avitotag/internal/entity"

func ToEntity(banner PostBannerRequestObject) entity.Banner {
	return entity.Banner{
		Content:   banner.Body.Content,
		FeatureId: banner.Body.FeatureId,
		IsActive:  banner.Body.IsActive,
		TagIds:    banner.Body.TagIds,
	}
}
