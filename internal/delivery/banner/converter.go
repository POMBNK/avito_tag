package banner

import "github.com/POMBNK/avitotag/internal/entity"

func ToCreateEntity(banner PostBannerRequestObject) entity.Banner {
	return entity.Banner{
		Content:   banner.Body.Content,
		FeatureId: banner.Body.FeatureId,
		IsActive:  banner.Body.IsActive,
		TagIds:    banner.Body.TagIds,
	}
}

func ToSearchParams(banner GetBannerRequestObject) entity.BannerSearchParams {
	return entity.BannerSearchParams{
		FeatureId: banner.Params.FeatureId,
		TagId:     banner.Params.TagId,
		Limit:     banner.Params.Limit,
		Offset:    banner.Params.Offset,
	}
}
