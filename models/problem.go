package models

import (
	"fmt"
	"gorm.io/gorm"
)

type Problem struct {
	gorm.Model
	Identity   string `gorm:"column:identity;type:varchar(36);" json:"identity"`
	CategoryId string `gorm:"column:category_id;type:varchar(255);" json:"category-id"`
	Title      string `gorm:"column:title;type:varchar(255)" json:"title"`
	Content    string `gorm:"column:content;type:text" json:"content"`
	MaxRuntime int    `gorm:"column:max_runtime;type:int(11);" json:"max-runtime"`
	MaxMem     int    `gorm:"column:max_mem;type:int(11);" json:"max_mem"`
}

func (table *Problem) TableName() string {
	return "problem"
}

func GetProblemList() {
	data := make([]*Problem, 0)

	DB.Find(&data)
	for _, v := range data {
		fmt.Printf("Problem ==> %v\n", v)
	}
}
