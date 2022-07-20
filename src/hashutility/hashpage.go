package hashutility

import "net/http"

type Handlers struct {
	internalstorage map[string]bool
}

func NewHashPage() (h *Handlers) {
	hashPage := new(Handlers)

	hashPage.internalstorage = make(map[string]bool)
	return hashPage
}

func (h *Handlers) parseHashRequest(w http.ResponseWriter, r *http.Request) {
	endpoint := r.URL.Query().Get("endpoint")

	if endpoint == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No endpoint url given"))
	} else {
		nextseq := getNextHashSeq(endpoint)

		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(nextseq))
	}

}

func (h *Handlers) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", h.parseHashRequest)
}
