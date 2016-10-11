package godjan

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var router = *mux.Router
var ds DataSource

func init() {
	ds = &MemoryDataSource{}
	r := mux.NewRouter()
}

func New() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi there!")
	})
	ds.Init()
}

func AddPlugin(p Plugin) {
	path := fmt.Sprintf("/%s", p.Name())
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, path)
	})
	for _, m := range p.Models() {
		path = fmt.Sprintf("/%s/%s", p.Name(), m.DbID())
		ds.Save(m, SampleData{"123", 432})
		http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("data: %v", ds.Get(m, "123"))
			fmt.Fprintf(w, path)
		})

		path = fmt.Sprintf("/%s/%s/", p.Name(), m.DbID())
		http.HandleFunc(path)
	}
}

func Start() {
	http.ListenAndServe(":3000", nil)
}
