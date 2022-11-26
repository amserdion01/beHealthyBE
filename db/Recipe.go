package db

import (
	_ "github.com/jinzhu/gorm"
	"github.com/lib/pq"
	dt "gorm.io/datatypes"
)

type Recipe struct {
	ID          string         `gorm:"column:id;type:uuid;default:uuid_generate_v4();primary_key"`
	Name        string         `gorm:"column:name;type:text;not null"`
	Author      string         `gorm:"column:author;type:varchar;not null"`
	AuthorID    string         `gorm:"column:author_id;type:text;not null"`
	Ingredients dt.JSONMap     `gorm:"column:ingredients;type:json;not null"`
	Details     string         `gorm:"column:details;type:text;not null"`
	Portions    int            `gorm:"column:portions;type:int4;not null"`
	Preparation dt.Time        `gorm:"column:preparation;type:time;not null"`
	Cooking     dt.Time        `gorm:"column:cooking;type:time;not null"`
	Tools       pq.StringArray `gorm:"column:tools;type:text[];not null"`
}

//  ["a", "b"]
//  {"a": "500ml", "b": "300ml"}
