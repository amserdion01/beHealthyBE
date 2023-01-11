package db

import (
	_ "github.com/jinzhu/gorm"
)

type Recipe struct {
	ID          string `gorm:"column:id;type:uuid;default:uuid_generate_v4();primary_key"`
	Name        string `gorm:"column:name;type:text;not null"`
	Author      string `gorm:"column:author;type:varchar;not null"`
	AuthorID    string `gorm:"column:author_id;type:text;not null"`
	Ingredients string `gorm:"column:ingredients;type:text;not null"`
	Details     string `gorm:"column:details;type:text;not null"`
	Portions    int    `gorm:"column:portions;type:int4;not null"`
	Preparation int    `gorm:"column:preparation;type:int4;not null"`
	Cooking     int    `gorm:"column:cooking;type:int4;not null"`
	Tools       string `gorm:"column:tools;type:text;not null"`
	ImageURL    string `gorm:"column:image_url;type:text;not null"`
	VideoURL    string `gorm:"column:video_url;type:text"`
}

//  ["a", "b"]
//  {"a": "500ml", "b": "300ml"}
