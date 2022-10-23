package operator

import (
	"log"

	"github.com/denizcamalan/movie_app/config"
	"github.com/denizcamalan/movie_app/model"
	"github.com/jinzhu/gorm"
)

type MovieModelMeneger interface{
	ListAll() ([]model.Movies, error)
	AddMovie(name,description,movie_type string) error
	SelectMovie(id uint) (model.Movies, error)
	UpdateMovie(id uint, name, description, movie_type string) error
	DeleteMovie(id uint) error
}

type MovieModel struct{
	db			*gorm.DB
	movie		model.Movies
	arr_movie	[]model.Movies
}

func NewMovieModel() (*MovieModel){
	var models MovieModel
	db, err := config.DB_Connect()
	if err != nil {return nil}
	log.Println("Connection error")
	db.Debug().DropTableIfExists(&model.Movies{})
	db.AutoMigrate(&model.Movies{})
	log.Println("created" + db.NewScope(model.Movies{}).TableName())
	models.db = db
	return &models
}

// Add items to movie from products config
func (mm *MovieModel) AddMovie(name,description,movie_type string) error{

	mm.movie.Name = name
	mm.movie.Description = description
	mm.movie.MovieType = movie_type

	if err := mm.db.Model(mm.movie).Create(&mm.movie).Error; err != nil {
		return err
	}
	log.Println("Data is added by database")
	return nil
}

// List all of the selected items into the movie
func (mm *MovieModel) ListAll() ([]model.Movies, error) {

	if err := mm.db.Model(mm.movie).Find(&mm.movie).Error; err != nil {
		return nil,err
	}
	
	log.Println("Data is listed by database")
	return mm.arr_movie,nil

}

// select item from config by item's id
func (mm *MovieModel) SelectMovie(id uint) (model.Movies, error){

	if err := mm.db.Model(mm.movie).Where("movie_id <> ?", id).Find(&mm.movie).Error; err != nil {
		return mm.movie,err
	}
	log.Println("Data is selected by database")
	return mm.movie,nil
}

// update movie table 
func (mm *MovieModel) UpdateMovie(id uint, name, description, movie_type string) error {

	if err := mm.db.Model(mm.movie).Model(&mm.movie).Where("movie_id = ?", id).
		Update(model.Movies{Name: name, Description: description,MovieType: movie_type}).Error; err != nil {
		log.Println(err)
		return err
	}

	log.Println("Data is updated by database")
	return nil
}

// delete items from movie table
func (mm *MovieModel) DeleteMovie(id uint) error{

	mm.movie.ID = id
	if err := mm.db.Model(mm.movie).Delete(&mm.movie).Error; err != nil{
		log.Println(err)
		return err
	}

	log.Println("Data is deleted by database")
	return nil
}


