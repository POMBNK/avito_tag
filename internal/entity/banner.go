package entity

type Banner struct {
	// Content Содержимое баннера
	Content *map[string]interface{} `json:"content,omitempty"`

	// FeatureId Идентификатор фичи
	FeatureId *int `json:"feature_id,omitempty"`

	// IsActive Флаг активности баннера
	IsActive *bool `json:"is_active,omitempty"`

	// TagIds Идентификаторы тэгов
	TagIds *[]int `json:"tag_ids,omitempty"`
}
