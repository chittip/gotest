package app

import "net/http"
import "github.com/chittip/gonews/pkg/view"
import "log"
import "github.com/chittip/gonews/pkg/model"

func newsView(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[1:]
	log.Println(id)
	n, err := model.GetNews(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	view.News(w, n)
}

// type news struct //1:45
// type newsView string //id

// func (id newsView) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	log.Println(id)
// 	view.News(w, nil)
// }
