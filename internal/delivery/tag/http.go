package tag

import (
	"context"
	"fmt"
	"github.com/POMBNK/avitotag/internal/pkg/apierror"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Server struct {
	engine *chi.Mux
}

func New(engine *chi.Mux) *Server {
	return &Server{engine: engine}
}

func (s *Server) Register() http.Handler {
	return HandlerFromMux(
		NewStrictHandlerWithOptions(s, nil,
			StrictHTTPServerOptions{RequestErrorHandlerFunc: apierror.ResponseErrorHandler()}), s.engine)
}

func (s *Server) GetBanner(ctx context.Context, request GetBannerRequestObject) (GetBannerResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) PostBanner(ctx context.Context, request PostBannerRequestObject) (PostBannerResponseObject, error) {
	fmt.Println(request)
	id := 666
	return PostBanner201JSONResponse{
		BannerId: &id,
	}, nil
}

func (s *Server) DeleteBannerId(ctx context.Context, request DeleteBannerIdRequestObject) (DeleteBannerIdResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) PatchBannerId(ctx context.Context, request PatchBannerIdRequestObject) (PatchBannerIdResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) GetUserBanner(ctx context.Context, request GetUserBannerRequestObject) (GetUserBannerResponseObject, error) {
	//TODO implement me
	panic("implement me")
}
