package http

import (
	"TelegramAPI/output/telegram"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (T *Api) CheckConnection(w http.ResponseWriter, r *http.Request) {

	var (
		logger = T.logging
	)
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Checked success"))
	if err != nil {
		logger.Errorln("Unable to write status", err)

	}
	return

}
func (T *Api) PostMessage(w http.ResponseWriter, r *http.Request) {
	var (
		logger = T.logging
	)

	jsonMessage, err := io.ReadAll(r.Body)
	if err != nil || len(jsonMessage) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(jsonMessage, &T.Message)
	if err != nil {
		logger.Errorln("Unable to unmarshal", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}
	fmt.Printf("%+v", T.Message)
	text, chn := InputApi(T).Export()
	tgrm := telegram.NewTelegram("5203368767:AAH5fdcTNV2Zj41Cp3SbTXYYWftY7H5eze0", chn)
	_, err = tgrm.Write([]byte(text))
	if err != nil {
		logger.Errorln(err)

	}
	w.WriteHeader(http.StatusOK)

}
