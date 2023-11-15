package controller

import (
	"net/http"
	"web/lib"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	// re := regexp.MustCompile("/public/*")

	// if re.Match([]byte(r.URL.Path)) {
	// file, err := os.ReadFile("." + r.URL.Path)
	// if err != nil {
	// 	fmt.Println(err)
	// 	http.Redirect(w, r, "/error", http.StatusInternalServerError)
	// 	return
	// }

	// w.Write(file)
	// http.ServeFile(w, r, "./public/file.html") // Remark: Alternative
	// return
	// }

	// Note: Handle unrecognized routes
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/404", http.StatusPermanentRedirect)
		return
	}

	// token := r.URL.Query().Get("token")

	// if token != "" {
	// 	ctx := context.WithValue(r.Context(), lib.Success{}, "Login Success!")
	// 	r = r.WithContext(ctx)
	// }

	lib.Render(w, r, "home.page.gohtml", nil)
	// fmt.Fprintln(w, "Something went wrong!")
}
