package main

import (
	"context"
	"kubecoverage/size"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	handler := http.NewServeMux()
	server := http.Server{Addr: ":8080", Handler: handler}

	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("http req %+v", r.URL)
		inStr := r.URL.Query().Get("size")
		inInt, err := strconv.Atoi(inStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		outStr := size.Size(inInt)
		outBytes := []byte(outStr)
		w.Write(outBytes)
	})

	go func() {
		c := make(chan os.Signal)
		//signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		signal.Notify(c)
		s := <-c
		log.Printf("signal=%d action=server.Close", s)
		err := server.Shutdown(context.Background())
		if err != nil {
			log.Fatal(err)
		}
	}()

	log.Printf("starting server on %s", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
