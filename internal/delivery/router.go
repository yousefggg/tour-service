package delivery

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"

	jwtlib "github.com/yousefggg/common-lib/pkg/jwt"
	_ "github.com/yousefggg/tour-service/docs"

	"github.com/yousefggg/tour-service/internal/delivery/handler"
	tourMiddleware "github.com/yousefggg/tour-service/internal/delivery/middleware"
)

type Router struct {
	tourHandler    *handler.TourHandler
	bookingHandler *handler.BookingHandler
	jwtManager     *jwtlib.TokenManager
}

func NewRouter(
	tourHandler *handler.TourHandler,
	bookingHandler *handler.BookingHandler,
	jwtManager *jwtlib.TokenManager,
) *Router {
	return &Router{
		tourHandler:    tourHandler,
		bookingHandler: bookingHandler,
		jwtManager:     jwtManager,
	}
}

func (r *Router) Setup() http.Handler {
	router := chi.NewRouter()

	router.Get("/swagger/*", httpSwagger.WrapHandler)

	router.Route("/api/v1", func(api chi.Router) {

		api.Get("/tours", r.tourHandler.GetAllTours)
		api.Get("/tours/{id}", r.tourHandler.GetTourByID)

		api.Group(func(protected chi.Router) {

			protected.Use(tourMiddleware.JWTMiddleware(r.jwtManager))

			protected.Post("/bookings", r.bookingHandler.CreateBooking)
			protected.Get("/bookings", r.bookingHandler.GetUserBookings)
			protected.Get("/bookings/{id}", r.bookingHandler.GetBookingByID)

			protected.Route("/admin", func(admin chi.Router) {

				admin.Post("/tours", r.tourHandler.CreateTour)
				admin.Put("/tours/{id}", r.tourHandler.UpdateTour)
				admin.Delete("/tours/{id}", r.tourHandler.DeleteTour)
			})
		})
	})

	return router
}