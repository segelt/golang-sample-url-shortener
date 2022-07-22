package hashutility

import (
	"fmt"
	"gobasictinyurl/src/helpers"
	"gobasictinyurl/src/middlewares"
	"gobasictinyurl/src/models"
	"gobasictinyurl/src/persistence"
	"net/http"
)

type Handlers struct {
	// internalstorage map[string]string
	// lock sync.RWMutex
}

func NewHashPage() (h *Handlers) {
	hashPage := new(Handlers)

	// hashPage.internalstorage = make(map[string]string)
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

		var urlEntry models.UrlEntry
		record := persistence.Instance.Where("ID = ?", endpoint).First(&urlEntry)

		if record.Error == nil {
			w.WriteHeader(http.StatusAccepted)
			w.Write([]byte(fmt.Sprintf("value: %s -- retrieved from map.", urlEntry.Value)))
		} else {
			nextseq := helpers.GetNextHashSeq(endpoint)

			//TODO : Handling errors on creation
			persistence.Instance.Create(&models.UrlEntry{ID: endpoint, Value: nextseq, UserID: "not implemented yet"})

			w.WriteHeader(http.StatusAccepted)
			w.Write([]byte(fmt.Sprintf("value: %s -- added to map.", nextseq)))
		}
	}

}

func (h *Handlers) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", middlewares.MultipleMiddleware(h.parseHashRequest, middlewares.LogMiddleware, middlewares.Auth))
}
