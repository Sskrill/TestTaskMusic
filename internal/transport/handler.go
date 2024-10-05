package transport

import (
	"github.com/Sskrill/TestTaskMusic/internal/domain"
	"github.com/Sskrill/TestTaskMusic/pkg/customLogger"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

type Service interface {
	AddSong(song domain.Song) error
	EditSong(id int, song domain.UpdateSong) error
	GetSongDetails(songName, performerName string) (domain.Song, error)
	DeleteSong(id int) error
	GetSongText(nameSong, performerName string) ([]string, error)
	GetSongsByFilters(songFilters *domain.FiltersForSong) ([]*domain.Song, error)
}
type Handler struct {
	service Service
	logger  customLogger.CustomLogger
}

func NewHandler(service Service, logger customLogger.CustomLogger) *Handler {
	return &Handler{service: service, logger: logger}
}
func (h *Handler) InitRouter() *mux.Router {
	r := mux.NewRouter()
	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8084/swagger/doc.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)
	song := r.PathPrefix("/song").Subrouter()
	{
		song.HandleFunc("/add", h.addSong).Methods(http.MethodPost)
		song.HandleFunc("/edit/{id}", h.editSong).Methods(http.MethodPut)
		song.HandleFunc("/details/{song_name}/{performer_name}", h.getDetailsSong).Methods(http.MethodGet)
		song.HandleFunc("/delete/{id}", h.deleteSong).Methods(http.MethodDelete)
		song.HandleFunc("/text/{song_name}/{performer_name}", h.getSongText).Methods(http.MethodGet)
		song.HandleFunc("/filters", h.getSongsByFilters).Methods(http.MethodGet)
	}
	return r
}
