package views

import (
	"html/template"
	"net/http"
	"simpleGoBlog/config"
	"time"
)

func IsODD(num int) bool {
	return nums%2 == 0
}

func GetNextName(strs []string, index int) string {
	return strs[index+1]
}

func Date(layout string) string {
	return time.Now().Format(layout)
}

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	t := template.New("index.html")
	path := config.Cfg.System.CurrentDir

	home := path + "/template/home.html"
	header := path + "/template/layout/header.html"
	footer := path + "/template/layout/footer.html"
	personal := path + "/template/layout/personal.html"
	post := path + "/template/layout/post-list.html"
	pagination := path + "/template/layout/pagination.html"

}
