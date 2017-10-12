package app

import "net/http"

// Mount mounts handlers to mux
func Mount(mux *http.ServeMux) {
	mux.HandleFunc("/", index) // list all new
	mux.Handle("/upload/", http.StripPrefix("/upload", http.FileServer(http.Dir("upload"))))
	//mux.HandleFunc("/news/", newsView)
	mux.Handle("/news/", http.StripPrefix("/news", http.HandlerFunc(newsView)))
	mux.HandleFunc("/register", adminRegister)
	// mux.Handle("/news/", http.StripPrefix("/news", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	id := r.URL.Path[1:]
	// 	newsView(id).ServeHTTP(w, r)
	// })))

	adminMux := http.NewServeMux()
	adminMux.HandleFunc("/login", adminLogin)
	adminMux.HandleFunc("/listTest", adminListTest)
	adminMux.HandleFunc("/create", adminCreate)
	adminMux.HandleFunc("/createTest", adminCreateTest)
	adminMux.HandleFunc("/edit", adminEdit)

	adminMux.HandleFunc("/editTest", adminEditTest)
	adminMux.HandleFunc("/runTest", adminRunTest)

	mux.Handle("/admin/", http.StripPrefix("/admin", onlyAdmin(adminMux)))
	// /news/:path
	// /admin/login/
	// /admin/list
	// /admin/create
	// /admin/edit

}

func onlyAdmin(h http.Handler) http.Handler {
	return h
}
