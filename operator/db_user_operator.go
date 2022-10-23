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
	GetUserByID(uid uint) (model.User,error)
	SaveUser(name,password string) (*model.User, error)
	LoginCheck(username string, password string) (string,error)
	BeforeSave() error
}

type RegisterModel struct{
	user 	model.User
	db		*gorm.DB
}

func NewRegiterModel() (*RegisterModel){
	var models RegisterModel
	db, err := config.DB_Connect()
	if err != nil { 
		log.Println(err) 
		return nil
	}
	db.Debug().DropTableIfExists(&model.User{})
	db.AutoMigrate(&model.User{})
	log.Println("created" + db.NewScope(&model.User{}).TableName())
	models.db = db
	return &models
}

func verify_password(password,hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (r *RegisterModel) GetUserByID(uid uint) (model.User,error) {

	if err := r.db.Model(r.user).First(&r.user,uid).Error; err != nil {
		log.Println("GetUserById",err)
		return r.user,errors.New("user not found")
	}
	r.user.ID = uid
	r.PrepareGive()
	
	return r.user,nil
}

func (r *RegisterModel) PrepareGive(){
	r.user.Password = ""
}

func (r *RegisterModel) LoginCheck(username string, password string) (string,error) {


	if err := r.db.Model(r.user).Where("username <> ?", username).Take(&r.user).Error; err != nil {
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

func (r *RegisterModel) SaveUser(name,password string) (*model.User, error) {

	if err := r.db.Model(r.user).Create(&r.user).Error ; err != nil {
		log.Println("SaveUser",err)
		return &r.user, err
	}
	return &r.user, nil

}

func (r *RegisterModel) BeforeSave() error {

	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.user.Password),bcrypt.DefaultCost)
	if err != nil {
		log.Println("BeforeSave",err)
		return err
	}
	r.user.Password = string(hashedPassword)

	//remove spaces in username 
	r.user.Username = html.EscapeString(strings.TrimSpace(r.user.Username))

	return nil
}