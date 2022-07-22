package hashutility

import (
	"fmt"
	"gobasictinyurl/src/auth"
	"gobasictinyurl/src/helpers"
	"gobasictinyurl/src/middlewares"
	"gobasictinyurl/src/models"
	"gobasictinyurl/src/persistence"
	"net/http"
)

type Handlers struct{}

func NewHashPage() (h *Handlers) {
	hashPage := new(Handlers)
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

			userid := r.Context().Value("userid").(string)
			persistence.Instance.Create(&models.UrlEntry{ID: endpoint, Value: nextseq, UserID: userid})

			w.WriteHeader(http.StatusAccepted)
			w.Write([]byte(fmt.Sprintf("value: %s -- added to map.", nextseq)))
		}
	}
}

func (h *Handlers) GetEndpointsOfUser(w http.ResponseWriter, r *http.Request) {
	userid := r.Context().Value("userid").(string)

	user, err := auth.GetUserFromStorageEagerById(userid)

	if err != nil {
		panic(err)
	} else {
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(fmt.Sprintf("entry counts: %d.", len(user.UrlEntries))))
	}
}

func (h *Handlers) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", middlewares.MultipleMiddleware(h.parseHashRequest, middlewares.LogMiddleware, middlewares.Auth))
	mux.HandleFunc("/myentries", middlewares.MultipleMiddleware(h.GetEndpointsOfUser, middlewares.LogMiddleware, middlewares.Auth))
}
