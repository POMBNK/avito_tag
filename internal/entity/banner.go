package entity

import "time"

type Banner struct {

	// ID Идентификатор баннера
	ID int `json:"id,omitempty"`

	// Name Название баннера
	Name string `json:"name,omitempty"`

	// CreatedAt Дата создания
	CreatedAt time.Time `json:"created_at,omitempty"`

	// UpdatedAt Дата обновления
	UpdatedAt *time.Time `json:"updated_at,omitempty"`

	// Content Содержимое баннера
	Content map[string]interface{} `json:"content,omitempty"`

	// FeatureId Идентификатор фичи
	FeatureId *int `json:"feature_id,omitempty"`

	// IsActive Флаг активности баннера
	IsActive *bool `json:"is_active,omitempty"`

	// TagIds Идентификаторы тэгов
	TagIds []int `json:"tag_ids,omitempty"`
}

type BannerSearchParams struct {
	FeatureId *int
	TagId     *int
	Limit     *int
	Offset    *int
}
