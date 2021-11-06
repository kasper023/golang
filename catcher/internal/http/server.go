package http

import (
	"catcher/internal/models"
	"catcher/internal/store"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type Server struct {
	ctx         context.Context
	idleConnsCh chan struct{}
	store       store.Store

	Address string
}

func NewServer(ctx context.Context, address string, store store.Store) *Server {
	return &Server{
		ctx:         ctx,
		idleConnsCh: make(chan struct{}),
		store:       store,

		Address: address,
	}
}

func (s *Server) basicHandler() chi.Router {
	r := chi.NewRouter()

	// REST
	r.Post("/payments", func(w http.ResponseWriter, r *http.Request) {
		payment := new(models.Payment)
		if err := json.NewDecoder(r.Body).Decode(payment); err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		s.store.Payments().Create(r.Context(), payment)
	})
	r.Get("/payments", func(w http.ResponseWriter, r *http.Request) {
		payments, err := s.store.Payments().All(r.Context())
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		render.JSON(w, r, payments)
	})
	r.Get("/payments/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		payment, err := s.store.Payments().ByID(r.Context(), id)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		render.JSON(w, r, payment)
	})
	r.Put("/payments", func(w http.ResponseWriter, r *http.Request) {
		payment := new(models.Payment)
		if err := json.NewDecoder(r.Body).Decode(payment); err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		s.store.Payments().Update(r.Context(), payment)
	})
	r.Delete("/payments/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		s.store.Payments().Delete(r.Context(), id)
	})

	return r
}

func (s *Server) Run() error {
	srv := &http.Server{
		Addr:         s.Address,
		Handler:      s.basicHandler(),
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 30,
	}
	go s.ListenCtxForGT(srv)

	log.Println("[HTTP] Server running on", s.Address)
	return srv.ListenAndServe()
}

func (s *Server) ListenCtxForGT(srv *http.Server) {
	<-s.ctx.Done() // блокируемся, пока контекст приложения не отменен

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("[HTTP] Got err while shutting down^ %v", err)
	}

	log.Println("[HTTP] Proccessed all idle connections")
	close(s.idleConnsCh)
}

func (s *Server) WaitForGracefulTermination() {
	// блок до записи или закрытия канала
	<-s.idleConnsCh
}
