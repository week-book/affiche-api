package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"runtime"
	"sync/atomic"
	"time"

	appdb "github.com/week-book/affiche-api/internal/db"
)

const requiredMigrationVersion = 1

type HealthHandler struct {
	DB                 *sql.DB
	lastBeatNs         atomic.Int64
	GoroutineThreshold int
	HeartbeatStale     time.Duration
}

func NewHealthHandler(db *sql.DB) *HealthHandler {
	h := &HealthHandler{
		DB:                 db,
		GoroutineThreshold: 2000,
		HeartbeatStale:     15 * time.Second,
	}
	h.Touch()
	return h
}

func (h *HealthHandler) Touch() {
	h.lastBeatNs.Store(time.Now().UnixNano())
}

func (h *HealthHandler) Liveness(w http.ResponseWriter, r *http.Request) {
	g := runtime.NumGoroutine()
	last := time.Unix(0, h.lastBeatNs.Load())
	now := time.Now()

	status := "ok"
	code := http.StatusOK
	details := map[string]any{
		"goroutines": g,
		"last_beat":  last.Format(time.RFC3339Nano),
		"now":        now.Format(time.RFC3339Nano),
	}

	if h.GoroutineThreshold > 0 && g > h.GoroutineThreshold {
		status = "error"
		details["reason"] = "too_many_goroutines"
		code = http.StatusInternalServerError
	}

	if h.HeartbeatStale > 0 && now.Sub(last) > h.HeartbeatStale {
		status = "error"
		details["reason"] = "stale_heartbeat"
		details["stale_by_seconds"] = now.Sub(last).Seconds()
		code = http.StatusInternalServerError
	}

	writeJSON(w, code, map[string]any{
		"status":  status,
		"details": details,
	})
}

func (h *HealthHandler) Readiness(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	errs := map[string]string{}

	if err := h.DB.PingContext(ctx); err != nil {
		errs["db"] = err.Error()
	}

	if err := appdb.CheckMigrations(ctx, h.DB, requiredMigrationVersion); err != nil {
		errs["migrations"] = err.Error()
	}

	if len(errs) > 0 {
		writeJSON(w, http.StatusServiceUnavailable, map[string]any{
			"status": "unavailable",
			"errors": errs,
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"status": "ready",
	})
}

func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}
