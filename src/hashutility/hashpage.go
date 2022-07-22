package hashutility

import (
	"fmt"
	"gobasictinyurl/src/helpers"
	"gobasictinyurl/src/middlewares"
	"net/http"
	"sync"
)

type Handlers struct {
	internalstorage map[string]string
	lock            sync.RWMutex
}

func NewHashPage() (h *Handlers) {
	hashPage := new(Handlers)

	hashPage.internalstorage = make(map[string]string)
	// hashPage.lock = sync.RWMutex{}
	return hashPage
}

func (h *Handlers) parseHashRequest(w http.ResponseWriter, r *http.Request) {
	endpoint := r.URL.Query().Get("endpoint")

	if endpoint == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No endpoint url given"))
	} else {
		// Check if the hash of this endpoint is already defined

		h.lock.RLock()
		val, ok := h.internalstorage[endpoint]
		h.lock.RUnlock()

		if ok {
			w.WriteHeader(http.StatusAccepted)
			w.Write([]byte(fmt.Sprintf("value: %s -- retrieved from map.", val)))
		} else {
			nextseq := helpers.GetNextHashSeq(endpoint)
			// nextseq := getNextHashSeq(endpoint)
			h.lock.Lock()
			defer h.lock.Unlock()
			h.internalstorage[endpoint] = nextseq

			w.WriteHeader(http.StatusAccepted)
			w.Write([]byte(fmt.Sprintf("value: %s -- added to map.", nextseq)))
		}
	}

}

func (h *Handlers) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", middlewares.MultipleMiddleware(h.parseHashRequest, middlewares.Auth))
}
