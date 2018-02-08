package main

import (
	"log"
	"net/http"

	"github.com/sparkymat/reactor"
	"github.com/sparkymat/reductor/config"
)

func main() {
	appConfig := config.Load()

	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./public/js"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./public/css"))))

	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		log.Printf("Handling incoming request: [IP]%v [Path]%v", request.RemoteAddr, request.URL.Path)

		r := reactor.New(appConfig.AppName, appConfig.Body.AppId)
		for _, cssInclude := range appConfig.Head.CssIncludes {
			r.AddCustomCssLink(cssInclude)
		}
		r.MapCssFolder("public/css", "/css")
		r.MapJavascriptFolder("public/js", "/js")

		response.Write([]byte(r.Html().String()))
	})

	log.Printf("Listening on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
