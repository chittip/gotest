package app

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/chittip/gonews/pkg/model"
	"github.com/chittip/gonews/pkg/view"
)

func adminLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		userID, err := model.Login(username, password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println(userID)
		http.Redirect(w, r, "/admin/list", http.StatusSeeOther)
		return
	}
	view.AdminLogin(w, nil)
}

func adminRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		err := model.Register(username, password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		return
	}
	view.AdminRegister(w, nil)
}

func adminList(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		action := r.FormValue("action")
		id := r.FormValue("id")
		if action == "delete" {
			err := model.DeleteNews(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		http.Redirect(w, r, "/admin/list", http.StatusSeeOther)
		return
	}
	list, err := model.ListNews()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	view.AdminList(w, &view.AdminListData{
		List: list,
	})
}

func adminEdit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	n, err := model.GetNews(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodPost {
		n.Title = r.FormValue("title")
		n.Detail = r.FormValue("detail")

		if file, fileHeader, err := r.FormFile("image"); err == nil {
			defer file.Close()
			//uuid.New().String()
			//fileName := time.Now().Format(time.RFC3339) + "-" + fileHeader.Filename
			fileName := time.Now().Format("2006-01-02-15-04-5") + "-" + fileHeader.Filename
			fp, err := os.Create("upload/" + fileName)
			if err == nil {
				io.Copy(fp, file)
			} else {
				fmt.Println(err)
			}
			fp.Close()
			n.Image = "/upload/" + fileName
		}

		err = model.UpdateNews(n)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/admin/list", http.StatusSeeOther)
		return
	}

	view.AdminEdit(w, n)
}

func adminCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		n := model.News{
			Title:  r.FormValue("title"),
			Detail: r.FormValue("detail"),
		}

		if file, fileHeader, err := r.FormFile("image"); err == nil {
			defer file.Close()
			//uuid.New().String()
			//fileName := time.Now().Format(time.RFC3339) + "-" + fileHeader.Filename
			fileName := time.Now().Format("2006-01-02-15-04-5") + "-" + fileHeader.Filename
			fp, err := os.Create("upload/" + fileName)
			if err == nil {
				io.Copy(fp, file)
			} else {
				fmt.Println(err)
			}
			fp.Close()
			n.Image = "/upload/" + fileName
		}
		err := model.CreateNews(n)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/admin/create", http.StatusSeeOther)
		return
	}
	view.AdminCreate(w, nil)
}
