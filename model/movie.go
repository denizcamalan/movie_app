package model

import "github.com/jinzhu/gorm"

type Movies struct{
	gorm.Model					`gorm:"size:255;not null;unique" json:"id"`
	Name 			string		`gorm:"size:255;not null;unique" json:"name"`
	Description 	string		`gorm:"size:255;not null;unique" json:"description"`
	MovieType 		string		`gorm:"size:255;not null;unique" json:"movie_type"`
}

type ExMovies struct{
	Name 			string		`json:"name"`
	Description 	string		`json:"description"`
	MovieType 		string		`json:"movie_type"`
}

