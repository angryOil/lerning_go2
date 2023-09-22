package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	//serverListening()
	useMuxHandleFunc()
}

func requestTimeHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r)
		fmt.Println("requesty ime for %s: %v", r.URL.Path, time.Now().Sub(start))
	})
}

func useMuxHandleFunc() {
	hi := http.NewServeMux()
	hi.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		hiHandler{}.ServeHTTP(w, r)
	})
	bye := http.NewServeMux()
	bye.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		byeHandler{}.ServeHTTP(w, r)
	})

	requestedHi := requestTimeHandler(hi)
	requestedBye := requestTimeHandler(bye)

	mux := http.NewServeMux()
	mux.Handle("/hi/", http.StripPrefix("/hi", requestedHi))
	//mux.Handle("/hi/", http.StripPrefix("/hi", requestTimeHandler(hi)))
	mux.Handle("/bye/", http.StripPrefix("/bye", requestedBye))
	//mux.Handle("/bye/", http.StripPrefix("/bye", bye))
	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}
	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

type byeHandler struct {
}

func (bh byeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		w.Write([]byte("무명님 안녕히 가세요~~ have a good day"))
		return
	}
	w.Write([]byte(fmt.Sprintf("%s 님 안녕히....  오늘도 좋은날 보내세요!\nhave a good day", name)))
}

type hiHandler struct {
}

func (hh hiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("?")
	name := r.URL.Query().Get("name")
	if name == "" {
		w.Write([]byte("어... 이름없는 분 안녕하세요? 고랭 열심히 합시다"))
		return
	}
	w.Write([]byte(fmt.Sprintf("%s go lang 열시미 합시다!!", name)))
}

func serverListening() {
	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      helloHandler{},
	}
	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

type helloHandler struct {
}

func (hh helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("고랭 공부가 잘되게 해주세요 "))
}
