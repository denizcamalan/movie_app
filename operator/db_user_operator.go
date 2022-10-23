package operator

import (
	"errors"
	"html"
	"log"
	"strings"

	"github.com/denizcamalan/movie_app/config"
	"github.com/denizcamalan/movie_app/model"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type RegisterManager interface{
	GetUserByID(uid uint) (model.Users,error)
	SaveUser(name,password string) (model.Users, error)
	LoginCheck(username string, password string) (string,error)
	BeforeSave() error
}

type RegisterModel struct{
	user 	model.Users
	db		*gorm.DB
}

func NewRegiterModel() (*RegisterModel){
	var models RegisterModel
	db, err := config.DB_Connect()
	if err != nil { 
		log.Println(err) 
		return nil
	}
	db.Debug().DropTableIfExists(&model.Users{})
	db.AutoMigrate(&model.Users{})
	log.Println("created" + db.NewScope(&model.Users{}).TableName())
	models.db = db
	return &models
}

func verify_password(password,hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (r *RegisterModel) GetUserByID(uid uint) (model.Users,error) {

	if err := r.db.Where("ID = ?", uid).First(&r.user).Error; err != nil {
		log.Println("GetUserById",err)
		return r.user,errors.New("user not found")
	}
	log.Println(r.user)

	prepareGive(r.user)

	return r.user,nil
}

func prepareGive(user model.Users){
	user.Password = ""
}

func (r *RegisterModel) LoginCheck(username string, password string) (string,error) {


	if err := r.db.Model(r.user).Where("username = ?", username).First(&r.user).Error; err != nil {
		log.Println("LoginCheck",err)
		return "", err
	}

	if err := verify_password(password, r.user.Password); err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		log.Println("verify_password",err)
		return "", err
	}

	token,err := GenerateToken(uint(r.user.ID))
	if err != nil {
		return "",err
	}

	return token,nil
}

func (r *RegisterModel) SaveUser(name,password string) (model.Users, error) {

	if err := r.db.Model(r.user).Create(&model.Users{ Username: name, Password: password}).Error ; err != nil {
		log.Println("SaveUser",err)
		return r.user, err
	}
	return r.user, nil

}

func (r *RegisterModel) BeforeSave() error {

	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.user.Password),bcrypt.DefaultCost)
	if err != nil {
		log.Println("BeforeSave",err)
		return err
	}
	r.user.Password = string(hashedPassword)

	r.user.Username = html.EscapeString(strings.TrimSpace(r.user.Username))

	return nil
}