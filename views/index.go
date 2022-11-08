package views

import (
	"net/http"
	"simpleGoBlog/common"
	"simpleGoBlog/config"
	"simpleGoBlog/models"
)

// type IndexData struct {
// 	Title string `json:"title"`
// 	Desc  string `json:"desc"`
// }

// func IsODD(num int) bool {
// 	return num%2 == 0
// }

// func GetNextName(strs []string, index int) string {
// 	return strs[index+1]
// }

// func Date(layout string) string {
// 	return time.Now().Format(layout)
// }

// func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
// 	t := template.New("index.html")
// 	path := config.Cfg.System.CurrentDir

// 	home := path + "/template/home.html"
// 	header := path + "/template/layout/header.html"
// 	footer := path + "/template/layout/footer.html"
// 	personal := path + "/template/layout/personal.html"
// 	post := path + "/template/layout/post-list.html"
// 	pagination := path + "/template/layout/pagination.html"

// 	t.Funcs(template.FuncMap{"isODD": IsODD, "getNextName": GetNextName, "date": Date})
// 	t, err := t.ParseFiles(path+"/template/index.html", home, header, footer, personal, post, pagination)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	//页面上涉及到的所有的数据，必须有定义
// 	var categorys = []models.Category{
// 		{
// 			Cid:  1,
// 			Name: "go",
// 		},
// 	}
// 	var posts = []models.PostMore{
// 		{
// 			Pid:          1,
// 			Title:        "go博客",
// 			Content:      "内容",
// 			UserName:     "码神",
// 			ViewCount:    123,
// 			CreateAt:     "2022-02-20",
// 			CategoryId:   1,
// 			CategoryName: "go",
// 			Type:         0,
// 		},
// 	}
// 	var hr = &models.HomeResponse{
// 		Viewer:    config.Cfg.Viewer,
// 		Categorys: categorys,
// 		Posts:     posts,
// 		Total:     1,
// 		Page:      1,
// 		Pages:     []int{1},
// 		PageEnd:   true,
// 	}
// 	t.Execute(w, hr)
// }

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index
	//页面上涉及到的所有的数据，必须有定义
	var categorys = []models.Category{
		{
			Cid:  1,
			Name: "go",
		},
	}
	var posts = []models.PostMore{
		{
			Pid:          1,
			Title:        "go博客",
			Content:      "内容",
			UserName:     "码神",
			ViewCount:    123,
			CreateAt:     "2022-02-20",
			CategoryId:   1,
			CategoryName: "go",
			Type:         0,
		},
	}
	var hr = &models.HomeResponse{
		Viewer:    config.Cfg.Viewer,
		Categorys: categorys,
		Posts:     posts,
		Total:     1,
		Page:      1,
		Pages:     []int{1},
		PageEnd:   true,
	}

	index.WriteData(w, hr)
}
