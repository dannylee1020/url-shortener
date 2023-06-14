package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os/exec"
)

func (app *application) RedirectHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	shortURL := params.ByName("url")
	fmt.Println(shortURL)

	urlData, err := app.models.URL.QueryWithShort(shortURL)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, urlData, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	openWebPage(urlData.LongURL)

}

func openWebPage(url string) {
	cmd := exec.Command("open", url)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
