package test

import (
	"GinHello/initRouter"
	"bytes"
	"gopkg.in/go-playground/assert.v1"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

func TestUserSave(t *testing.T) {
	username := "list"
	router := initRouter.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/user/"+username, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "用户"+username+"已经保存", w.Body.String())
}

func TestUserSaveByQuery(t *testing.T) {
	username := "list"
	age := 18
	router := initRouter.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/user?name="+username+"&age="+strconv.Itoa(age), nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "用户"+username+",年龄:"+strconv.Itoa(age)+"已经保存", w.Body.String())
}

func TestUserSaveByDafaultQuery(t *testing.T){
	username := "list"
	router := initRouter.SetupRouter()
	w := httptest.NewRecorder()
	req,_ := http.NewRequest(http.MethodGet,"/user?name="+username,nil)
	router.ServeHTTP(w,req)
	assert.Equal(t,http.StatusOK,w.Code)
	assert.Equal(t,"用户"+username+",年龄:20已经保存",w.Body.String())
}

func TestUserPostForm(t *testing.T){
	value := url.Values{}
	value.Add("email","libaocai@gmail.com")
	value.Add("assword","1234")
	value.Add("password-again","1234")
	w := httptest.NewRecorder()
	req,_ := http.NewRequest(http.MethodPost,"/user/register",bytes.NewBufferString(value.Encode()))
	req.Header.Add("Content-Type","application/x-wwww-form-urlencoded;param=value")
	router.ServeHTTP(w,req)
	assert.Equal(t,http.StatusOK,w.Code)
}

func TestUserPostFormEmailErrorAndPasswordError(t *testing.T) {
	value := url.Values{}
	value.Add("email", "youngtxhui@qq.com")
	value.Add("password", "12345")
	value.Add("password-again", "12345")
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/user/register", bytes.NewBufferString(value.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; param=value")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUserLogin(t *testing.T) {
	email := "youngxhui@qq.com"
	value := url.Values{}
	value.Add("email", email)
	value.Add("password", "1234")
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/user/login", bytes.NewBufferString(value.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; param=value")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, strings.Contains(w.Body.String(), email), true)
}
