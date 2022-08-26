package http

import (
	"net/http"
)

func NewRouter(T *Api) {
	var (
		logger = T.logging
	)

	http.HandleFunc("/PostMessage", T.PostMessage)
	http.HandleFunc("/CheckConnection", T.CheckConnection)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {

		logger.Fatalln("Unable to create server", err)

	}

}
