package facts

import (
	"fmt"
	"net/http"
)

func NewRouter(service *FactsService) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(404)
		fmt.Fprintln(w, "Page Not Found")
	})
	mux.HandleFunc("/fact", func(w http.ResponseWriter, _ *http.Request) {
		data, err := service.GetRandomFact()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintln(w, "Something goes wrong")
		}
		fmt.Fprintln(w, data)
	})
	return mux
}
