package model

type WebPage struct {
	URL      string `gorm:"column:url"`
	HTML     string `gorm:"column:html"`
	Title    string `gorm:"column:title"`
	Category string `gorm:"column:category"`
}

func (w WebPage) TableName() string {
	return "webpage"
}
