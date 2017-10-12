package view

import (
	"net/http"

	"github.com/chittip/gotest/pkg/model"
)

// IndexData ..
type IndexData struct {
	List []*model.News
}

// Index renders index view
func Index(w http.ResponseWriter, data *IndexData) {
	render(tpIndex, w, data)
}

// News renders index view
func News(w http.ResponseWriter, data *model.News) {
	render(tpNews, w, data)
}

// AdminLogin renders admin login
func AdminLogin(w http.ResponseWriter, data interface{}) {
	render(tpAdminLogin, w, data)
}

// AdminRegister renders admin login
func AdminRegister(w http.ResponseWriter, data interface{}) {
	render(tpAdminRegister, w, data)
}

// AdminListData ..
type AdminListData struct {
	List []*model.News
}

// AdminListDataTest is struct for send AdminListTest
type AdminListDataTest struct {
	List []*model.Test
}

// AdminCreate renders admin login
func AdminCreate(w http.ResponseWriter, data interface{}) {
	render(tpAdminLCreate, w, data)
}

// AdminEdit renders admin login
func AdminEdit(w http.ResponseWriter, data *model.News) {
	render(tpAdminLEdit, w, data)
}

// AdminList renders admin login
func AdminList(w http.ResponseWriter, data *AdminListData) {
	render(tpAdminList, w, data)
}

// AdminListTest renders admin login
func AdminListTest(w http.ResponseWriter, data *AdminListDataTest) {
	render(tpAdminListTest, w, data)
}

// AdminCreateTest ...
func AdminCreateTest(w http.ResponseWriter, data interface{}) {
	render(tpAdminLCreateTest, w, data)
}

// AdminCreateTest ...
func AdminEditTest(w http.ResponseWriter, data interface{}) {
	render(tpAdminLEditTest, w, data)
}
