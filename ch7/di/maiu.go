package main

import (
	"errors"
	"fmt"
	"net/http"
)

func main() {
	l := loggerAdapter(logOutput)
	ds := newSimpleDataStore()
	myLogic := newSimpleLogic(l, ds)
	c := newController(l, myLogic)
	http.HandleFunc("/hello", c.HandlerGreeting)
	http.ListenAndServe(":8080", nil)
}

type logger interface {
	Log(msg string)
}

func logOutput(msg string) {
	fmt.Println(msg)
}

type loggerAdapter func(msg string)

func (l loggerAdapter) Log(msg string) {
	l(msg)
}

type dataStore interface {
	UserNameForUserID(userId string) (string, bool)
}

type simpleDataStore struct {
	userData map[string]string
}

func (sd simpleDataStore) UserNameForUserID(userId string) (string, bool) {
	name, ok := sd.userData[userId]
	return name, ok
}

func newSimpleDataStore() simpleDataStore {
	return simpleDataStore{
		userData: map[string]string{
			"1": "warmOil",
			"2": "angryOil",
			"3": "joy",
		},
	}
}

type logic interface {
	SayHello(userId string) (string, error)
}

type simpleLogic struct {
	l  logger
	ds dataStore
}

func (sl simpleLogic) SayHello(userId string) (string, error) {
	sl.l.Log("in say hello for " + userId)
	name, ok := sl.ds.UserNameForUserID(userId)
	if !ok {
		return "", errors.New("unknown userId:" + userId)
	}

	return "hello, " + name, nil
}

func newSimpleLogic(l logger, ds dataStore) simpleLogic {
	return simpleLogic{
		l:  l,
		ds: ds,
	}
}

type controller struct {
	l     logger
	logic logic
}

func (c controller) HandlerGreeting(w http.ResponseWriter, r *http.Request) {
	c.l.Log("in say ya")
	userId := r.URL.Query().Get("user_id")
	msg, err := c.logic.SayHello(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(msg))
}

func newController(l logger, logic logic) controller {
	return controller{
		l:     l,
		logic: logic,
	}
}
