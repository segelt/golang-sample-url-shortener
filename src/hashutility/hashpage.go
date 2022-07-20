package hashutility

import "net/http"

func ParseHashRequest(w http.ResponseWriter, r *http.Request) {
	endpoint := r.URL.Query().Get("endpoint")

	if endpoint == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No endpoint url given"))
	}

	nextseq := GetNextHashSeq(endpoint)

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(nextseq))
}

func SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", ParseHashRequest)
}
