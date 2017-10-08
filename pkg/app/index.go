package app

import (
	"log"
	"net/http"

	"github.com/chittip/gonews/pkg/model"
	"github.com/chittip/gonews/pkg/view"
)

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	list, err := model.ListNews()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println(list)
	view.Index(w, &view.IndexData{
		List: list,
	})
}
