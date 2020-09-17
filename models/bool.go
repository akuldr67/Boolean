package models

type Boolean struct {
	ID    uint   `json:"id" gorm:"primary_key"`
	Key   string `json:"key"`
	Value bool   `json:"value"`
}

func (b *Boolean) TableName() string {
	return "booleans"
}
