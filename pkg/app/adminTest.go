package app

import (
	"net/http"

	"github.com/chittip/gotest/pkg/model"
	"github.com/chittip/gotest/pkg/view"
)

func adminListTest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		action := r.FormValue("action")
		id := r.FormValue("id")
		if action == "delete" {
			err := model.DeleteTest(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		http.Redirect(w, r, "/admin/list", http.StatusSeeOther)
		return
	}
	list, err := model.ListTest()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	view.AdminListTest(w, &view.AdminListDataTest{
		List: list,
	})
}

/*
	ID         bson.ObjectId `bson:"_id"`
	APIName    string
	URLPath    string
	MethodType string
	URLParam   string
	Body       string
	CreateAt   time.Time `bson:"createAt"`
	UpdatedAt  time.Time `bson:"updateAt"`
*/

func adminCreateTest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		n := model.Test{
			APIName:    r.FormValue("apiName"),
			URLPath:    r.FormValue("urlPath"),
			MethodType: r.FormValue("methodType"),
			URLParam:   r.FormValue("urlParam"),
			Body:       r.FormValue("body"),
			Expected:   r.FormValue("Expected"),
		}

		// if file, fileHeader, err := r.FormFile("image"); err == nil {
		// 	defer file.Close()
		// 	//uuid.New().String()
		// 	//fileName := time.Now().Format(time.RFC3339) + "-" + fileHeader.Filename
		// 	fileName := time.Now().Format("2006-01-02-15-04-5") + "-" + fileHeader.Filename
		// 	fp, err := os.Create("upload/" + fileName)
		// 	if err == nil {
		// 		io.Copy(fp, file)
		// 	} else {
		// 		fmt.Println(err)
		// 	}
		// 	fp.Close()
		// 	n.Image = "/upload/" + fileName
		// }
		err := model.CreateTest(n)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/admin/createTest", http.StatusSeeOther)
		return
	}
	view.AdminCreateTest(w, nil)
}

func adminRunTest(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	test, err := model.GetTest(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = model.RunTest(test)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//view.AdminEdit(w, n)
	http.Redirect(w, r, "/admin/listTest", http.StatusSeeOther)
	return
}

func adminEditTest(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	test, err := model.GetTest(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodPost {
		test.APIName = r.FormValue("apiName")
		test.URLPath = r.FormValue("urlPath")
		test.MethodType = r.FormValue("methodType")
		test.URLParam = r.FormValue("urlParam")
		test.Body = r.FormValue("body")
		test.Expected = r.FormValue("Expected")

		// if file, fileHeader, err := r.FormFile("image"); err == nil {
		// 	defer file.Close()
		// 	//uuid.New().String()
		// 	//fileName := time.Now().Format(time.RFC3339) + "-" + fileHeader.Filename
		// 	fileName := time.Now().Format("2006-01-02-15-04-5") + "-" + fileHeader.Filename
		// 	fp, err := os.Create("upload/" + fileName)
		// 	if err == nil {
		// 		io.Copy(fp, file)
		// 	} else {
		// 		fmt.Println(err)
		// 	}
		// 	fp.Close()
		// 	n.Image = "/upload/" + fileName
		// }

		err = model.UpdateTest(test)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/admin/listTest", http.StatusSeeOther)
		return
	}
	view.AdminEditTest(w, test)
}
