package operator

import (
	log "github.com/siruspen/logrus"

	"github.com/denizcamalan/movie_app/config"
	"github.com/denizcamalan/movie_app/model"
	"github.com/jinzhu/gorm"
)

var db = config.DB_Connect()
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
	db.Debug().DropTableIfExists(&model.Movies{})
	db.AutoMigrate(&model.Movies{})
	log.Infof("created %s",db.NewScope(model.Movies{}).TableName())
	models.db = db
	return &models
}

// Add items to movie from products config
func (mm *MovieModel) AddMovie(name,description,movie_type string) error{

	if err := mm.db.Model(mm.movie).Create(&model.Movies{Name: name, Description: description,MovieType: movie_type}).Error; err != nil {
		return err
	}
	log.Info("Data is added by database")
	return nil
}

// List all of the selected items into the movie
func (mm *MovieModel) ListAll() ([]model.Movies, error) {

	if err := mm.db.Model(mm.movie).Find(&mm.arr_movie).Error; err != nil {
		return nil,err
	}
	
	log.Info("Data is listed by database")
	return mm.arr_movie,nil

}

// select item from config by item's id
func (mm *MovieModel) SelectMovie(id uint) (model.Movies, error){

	if err := mm.db.Model(mm.movie).Where("ID = ?", id).Find(&mm.movie).Error; err != nil {
		return mm.movie,err
	}
	log.Info("Data is selected by database")
	return mm.movie,nil
}

// update movie table 
func (mm *MovieModel) UpdateMovie(id uint, name, description, movie_type string) error {

	if err := mm.db.Model(mm.movie).Model(&mm.movie).Where("ID = ?", id).
		Update(model.Movies{Name: name, Description: description,MovieType: movie_type}).Error; err != nil {
		log.Error(err)
		return err
	}

	log.Info("Data is updated by database")
	return nil
}

// delete items from movie table
func (mm *MovieModel) DeleteMovie(id uint) error{

	mm.movie.ID = id
	if err := mm.db.Model(mm.movie).Delete(&mm.movie).Error; err != nil{
		log.Error(err)
		return err
	}

	log.Info("Data is deleted by database")
	return nil
}


