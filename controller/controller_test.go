package controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/denizcamalan/movie_app/controller"
	"github.com/denizcamalan/movie_app/model"
	"github.com/denizcamalan/movie_app/repo"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	movie_controller = repo.NewOperatorModel().DB_Operator()
	user_controller = repo.NewOperatorModel().Register_Operator()
	api_controller = &controller.API{}
	token string
)

func TestGetAll(t *testing.T){

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		api_controller.GetAll(c)

		assert.Equal(t, http.StatusOK, w.Code)
}

func TestRegister(t *testing.T){
	var user model.ExUser = model.ExUser{Username: "deniz", Password: "123"}
	
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	var b bytes.Buffer
	err := json.NewEncoder(&b).Encode(user)
	if err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodPost, "/movie_archive/login", &b)
	c.Request = req

	api_controller.Register(c)

	var got gin.H
    err = json.Unmarshal(w.Body.Bytes(), &got)
    if err != nil {
        t.Fatal(err)
    }

	assert.Equal(t, gin.H(gin.H{"message":"Registration Success!"}), got)

	assert.Equal(t, http.StatusOK, w.Code)

}
func TestLogin(t *testing.T) {
	
	var user model.ExUser = model.ExUser{Username: "deniz", Password: "123"}
	user_controller.SaveUser(user.Username,user.Password)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	var b bytes.Buffer
	err := json.NewEncoder(&b).Encode(user)
	if err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodPost, "/movie_archive/login", &b)
	c.Request = req

	api_controller.Login(c)

	token, _ = user_controller.LoginCheck(user.Username, user.Password)

	var got gin.H
    err = json.Unmarshal(w.Body.Bytes(), &got)
    if err != nil {
        t.Fatal(err)
    }

	assert.Equal(t,  gin.H(gin.H{"token":token}), got)

	assert.Equal(t, http.StatusOK, w.Code)

}
