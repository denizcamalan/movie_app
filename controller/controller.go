package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/denizcamalan/movie_app/model"
	"github.com/denizcamalan/movie_app/operator"
	"github.com/denizcamalan/movie_app/repo"
	"github.com/gin-gonic/gin"
)

var (
	movie_controller = repo.NewOperatorModel().DB_Operator()
	user_controller = repo.NewOperatorModel().Register_Operator()
)

type ApiController interface{
	Login(c *gin.Context)
	AddData(c *gin.Context)
	GetDataById(c *gin.Context)
	UpdateData(c *gin.Context)
	DeleteDataByID(c *gin.Context)
	GetAll(c *gin.Context)
	Register(c *gin.Context)
	CurrentUser(c *gin.Context)
}

type API struct{
	movie 			model.Movies
	ex_movie		model.ExMovies
	user			model.User
	ex_user			model.ExUser
}

// Login godoc
// @Summary create data User
// @Description get {object}
// @Tags User
// @Accept  json
// @Produce  json
// @Param User body model.ExUser true "Users"
// @Success 200 string	token
// @Failure 400 {object} model.Message
// @Failure 406 {object} model.Message
// @Router /login [post]
func (api *API) Login(c *gin.Context) {

	if err := c.ShouldBindJSON(&api.ex_user); err != nil {
		c.JSON(http.StatusBadRequest, model.Message{Message: err.Error()})
		return
	}

	api.user.Username = api.ex_user.Username
	api.user.Password = api.ex_user.Password

	token, err := user_controller.LoginCheck(api.user.Username, api.user.Password)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, model.Message{Message: "error : username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token":token})
}

// Register godoc
// @Summary create data User
// @Description get {object}
// @Tags User
// @Accept  json
// @Produce  json
// @Param User body model.ExUser true "Users"
// @Success 200 {object} model.Message
// @Failure 400 {object} model.Message
// @Failure 204 {object} model.Message
// @Router /register [post]
func (api *API) Register(c *gin.Context){
	
	if err := c.ShouldBindJSON(&api.ex_user); err != nil {
		c.JSON(http.StatusBadRequest, model.Message{Message: err.Error()})
		return
	}

	api.user.Username = api.ex_user.Username
	api.user.Password = api.ex_user.Password

	if _,err := user_controller.SaveUser(api.user.Username,api.user.Password); err != nil{
		c.JSON(http.StatusNoContent, model.Message{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.Message{Message: "Registration Success !"})

}

// listAll godoc
// @Summary get all items in the model.ExMovies list
// @Tags Movies
// @Accept  json
// @Produce json
// @Success 200 {array} []model.ExMovies
// @Failure 404 {object} model.Message
// @Router /list [get]
func (api *API) GetAll(c *gin.Context) {

	movies, err := movie_controller.ListAll()
	if err != nil{
		c.JSON(http.StatusNotFound,model.Message{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, movies)	
}

// AddData godoc
// @Summary create data Movies
// @Description get {object}
// @Tags Movies
// @Accept  json
// @Produce  json
// @Param Movies body model.ExMovies true "Movies"
// @Security ApiKeyAuth
// @Success 200 {object} model.ExMovies
// @Failure 400 {object} model.Message
// @Failure 404 {object} model.ExMovies
// @Router /admin/add [post]
func (api *API) AddData(c *gin.Context) {

	api.CurrentUser(c)
	
	if err := c.ShouldBindJSON(&api.ex_movie); err != nil{
		c.JSON(http.StatusBadRequest, model.Message{Message: err.Error()})
		return
	}

	api.movie.Name = api.ex_movie.Name
	api.movie.Description = api.ex_movie.Description
	api.movie.MovieType = api.ex_movie.MovieType

	if err := movie_controller.AddMovie(api.movie.Name,api.movie.Description,api.movie.MovieType); err != nil {
		c.JSON(http.StatusNotFound, model.Message{Message: err.Error()})
		return
    }
	c.JSON(http.StatusOK, model.Message{Message: api.movie.Name+" is added."})
}

// getDataById godoc
// @Summary show Movies by ID
// @Description get string by ID
// @Tags Movies
// @Accept  json
// @Produce  json
// @Param id path int true "model.ExMovies"
// @Security ApiKeyAuth
// @Success 200 {object} model.ExMovies
// @Failure 204 {object} model.Message
// @Router /admin/get/{id} [get]
func (api *API) GetDataById(c *gin.Context) {

	api.CurrentUser(c)

	intID, err := strconv.Atoi(c.Param("id"))
	if err != nil {	
		log.Println(err)
		return
	}

	api.movie.ID = uint(intID)

	strData, err:= movie_controller.SelectMovie(api.movie.ID)
	if  err != nil {
		c.JSON(http.StatusNoContent, model.Message{Message: err.Error()})
		return
	}
		c.JSON(http.StatusOK, strData)
}


// UpdateData godoc
// @Summary update Movies by ID
// @Tags Movies
// @Description update by json Movies
// @Accept  json
// @Produce json
// @Param id path int true "model.ExMovies"
// @Param Movies body 	 model.ExMovies true "Movies"
// @Security ApiKeyAuth
// @Success 200 {object} model.ExMovies
// @Failure 400 {object} model.Message
// @Failure 404 {object} model.Message
// @Router /admin/update/{id} [put]
func (api *API) UpdateData(c *gin.Context) {

	api.CurrentUser(c)

	intID, err := strconv.Atoi(c.Param("id"))
	if err != nil {	
		return
	}

	err = c.ShouldBindJSON(&api.ex_movie)
	if err != nil{
		c.JSON(http.StatusBadRequest, model.Message{Message: err.Error()})
		return
	}
	api.movie.Name = api.ex_movie.Name
	api.movie.Description = api.ex_movie.Description
	api.movie.MovieType = api.ex_movie.MovieType

	if err := movie_controller.UpdateMovie(uint(intID),api.movie.Name,api.movie.Description,api.movie.MovieType); err != nil {
		c.JSON(http.StatusNotFound, model.Message{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.Message{Message: api.movie.Name+" is updated."})
}

// deleteDataByID godoc
// @Summary delete a model.ExMovies item by ID
// @Description delete Movies by ID
// @Tags Movies
// @Accept  json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "model.ExMovies"
// @Success 200 {object} model.Message
// @Failure 404 {object} model.Message
// @Router /admin/delete/{id} [delete]
func (api *API) DeleteDataByID(c *gin.Context) {

	api.CurrentUser(c)

	intID, err := strconv.Atoi(c.Param("id"))
	if err != nil {	
		log.Println(err)
		return
	}

	movie, _ := movie_controller.SelectMovie(uint(intID))
	if err := movie_controller.DeleteMovie(uint(intID)); err != nil {
		c.JSON(http.StatusNotFound, model.Message{Message: err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, model.Message{Message: movie.Name+" is deleted."})
}

// get user godoc
// @Summary get user in the model.User
// @Tags User
// @Accept  json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} model.Message
// @Failure 400 {object} model.Message
// @Failure 442 {object} model.Message
// @Router /admin/user [get]
func (api *API) CurrentUser(c *gin.Context){

	user_id, err := operator.ExtractTokenID(c)
	
	if err != nil {
		c.JSON(http.StatusBadRequest,  model.Message{Message: err.Error()})
		return
	}
	
	u,err := user_controller.GetUserByID(uint(user_id))
	
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Message{Message: err.Error()})
		return
	}

	api.user = u

	c.JSON(http.StatusOK, model.Message{Message: "data : "+ u.Username})
}
