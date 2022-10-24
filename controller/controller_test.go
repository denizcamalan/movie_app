package controller

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/denizcamalan/movie_app/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var api_controller = &API{}

func TestLogin(t *testing.T) {

    r := gin.Default()
	r.Group("/movie_archive")

	user := model.Users{Username: "deniz", Password: "123"}
	
	r.POST("/register",api_controller.Login)

	jsonValue, _ := json.Marshal(user)
	
    req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonValue))
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

	// token, _ := user_controller.LoginCheck(user.Username, user.Password)
	
	mockResponse := `{"Registration Success!"}`

    responseData, _ := ioutil.ReadAll(w.Body)
    assert.Equal(t, mockResponse, string(responseData))
    assert.Equal(t, http.StatusOK, w.Code)
}

