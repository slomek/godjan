package godjan

import (
	"fmt"
	"net/http"
)

func New() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi there!")
	})
}

func AddPlugin(p Plugin) {
	path := fmt.Sprintf("/%s", p.Name())
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, path)
	})
	for _, m := range p.Models() {
		p := fmt.Sprintf("/%s/%s", p.Name(), m.ID())
		http.HandleFunc(p, func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, p)
		})
	}
}

func Start() {
	http.ListenAndServe(":3000", nil)
}
