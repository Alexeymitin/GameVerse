package stat

import (
	"gameverse/configs"
	"gameverse/pkg/middleware"
	"gameverse/pkg/response"
	"net/http"
	"time"
)

const (
	GroupByDay   = "day"
	GroupByMonth = "month"
)

type StatHandler struct {
	StatRepository *StatRepository
}

type StatHandlerDeps struct {
	StatRepository *StatRepository
	Config         *configs.Config
}

func NewStatHandler(router *http.ServeMux, deps StatHandlerDeps) {
	handler := &StatHandler{
		StatRepository: deps.StatRepository,
	}

	router.Handle("GET /stat", middleware.IsAuth(handler.GetStat(), deps.Config))
}

func (h *StatHandler) GetStat() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		from, err := time.Parse("2006-01-02", req.URL.Query().Get("from"))
		if err != nil {
			http.Error(w, "Invalid 'from' date format. Use YYYY-MM-DD.", http.StatusBadRequest)
			return
		}

		to, err := time.Parse("2006-01-02", req.URL.Query().Get("to"))
		if err != nil {
			http.Error(w, "Invalid 'to' date format. Use YYYY-MM-DD.", http.StatusBadRequest)
			return
		}

		by := req.URL.Query().Get("by")
		if by != GroupByDay && by != GroupByMonth {
			http.Error(w, "Invalid 'by' parameter. Use 'day' or 'month'.", http.StatusBadRequest)
			return
		}

		stats := h.StatRepository.GetStats(by, from, to)
		response.Json(w, stats, http.StatusOK)
	}
}
