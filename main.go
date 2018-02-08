package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sparkymat/reactor"
	"github.com/sparkymat/reductor/config"
)

func main() {
	appConfig := config.Load()

	if _, err := os.Stat("./public/js"); err != nil && os.IsNotExist(err) {
		panic(err)
	}

	if _, err := os.Stat("./public/css"); err != nil && os.IsNotExist(err) {
		panic(err)
	}

	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./public/js"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./public/css"))))

	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		log.Printf("Handling incoming request: [IP]%v [Path]%v", request.RemoteAddr, request.URL.Path)

		r := reactor.New(appConfig.PageTitle, appConfig.Body.AppId)
		for _, cssInclude := range appConfig.Head.CssIncludes {
			r.AddCustomCssLink(cssInclude)
		}
		r.MapCssFolder("public/css", "/css")
		r.MapJavascriptFolder("public/js", "/js")

		response.Write([]byte(r.Html().String()))
	})

	log.Printf("Listening on http://0.0.0.0:%d/", appConfig.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", appConfig.Port), nil))
}
