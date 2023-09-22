package main

import (
	"errors"
	"fmt"
	"net/http"
)

func main() {
	// di dependency injection 의존성 주입
	// ioc inversion of control
	l := loggerAdapter(logOutPut)
	ds := newSimpleDataStore()
	logi := newSimpleLogin(l, ds)
	c := newController(l, logi)
	http.HandleFunc("/hello", c.handleGreeting)
	http.ListenAndServe(":8080", nil)
}

type handler interface {
	ServerHttp(http.ResponseWriter, *http.Request)
}

type handlerFunc func(http.ResponseWriter, *http.Request)

func (f handlerFunc) ServeHttp(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}

func logOutPut(msg string) {
	fmt.Println(msg)
}

type simpleDataStore struct {
	userData map[string]string
}

func (sds simpleDataStore) userNameForID(userId string) (string, bool) {
	name, ok := sds.userData[userId]
	return name, ok
}

func newSimpleDataStore() simpleDataStore {
	return simpleDataStore{
		userData: map[string]string{
			"1": "joy",
			"2": "world",
			"3": "hello",
		},
	}
}

type dataStore interface {
	userNameForID(userID string) (string, bool)
}

type logger interface {
	Log(msg string)
}

type loggerAdapter func(msg string)

func (lg loggerAdapter) Log(msg string) {
	lg.Log(msg)
}

type simpleLogic struct {
	l  logger
	ds dataStore
}

func (sl simpleLogic) SayHello(userId string) (string, error) {
	sl.l.Log("in say hello for" + userId)
	name, ok := sl.ds.userNameForID(userId)
	if !ok {
		return "", errors.New("unknown user")
	}
	return "hello," + name, nil
}

func newSimpleLogin(l logger, ds dataStore) simpleLogic {
	return simpleLogic{
		l:  l,
		ds: ds,
	}
}

type logic interface {
	SayHello(userID string) (string, error)
}

type controller struct {
	l    logger
	logi logic
}

func (c controller) handleGreeting(w http.ResponseWriter, r *http.Request) {
	c.l.Log("in say hello")
	userId := r.URL.Query().Get("user_id")
	msg, err := c.logi.SayHello(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(msg))
}

func newController(l logger, logi logic) controller {
	return controller{
		l:    l,
		logi: logi,
	}
}
