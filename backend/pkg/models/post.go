package models

import (
	"time"

	"github.com/fatihsen-dev/go-fullstack-social-media/pkg/config"
	"gorm.io/gorm"
)

type Post struct{
	ID 			uint   		 `json:"id"`
	Title     	string 		 `json:"title"`
	Subtitle  	string 		 `json:"subtitle"`
	Description string 		 `json:"description"`
	Owner 	 	uint	 		 `json:"owner"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time 		 `json:"updated_at"`
  	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func init(){
	config.Connect()
	DB = config.GetDB()
	DB.AutoMigrate(&Post{})
}