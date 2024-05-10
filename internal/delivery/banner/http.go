package banner

import (
	"context"
	"github.com/POMBNK/avitotag/internal/pkg/apierror"
	"net/http"

	"github.com/POMBNK/avitotag/internal/service"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	service service.BannerService
}

func New(service service.BannerService) *Server {
	return &Server{service: service}
}

func (s *Server) Register(engine *chi.Mux) http.Handler {
	return HandlerFromMux(
		NewStrictHandlerWithOptions(s, nil,
			StrictHTTPServerOptions{
				ResponseErrorHandlerFunc: apierror.ResponseErrorHandler(),
			}), engine)
}

func (s *Server) GetBanner(ctx context.Context, request GetBannerRequestObject) (GetBannerResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) PostBanner(ctx context.Context, request PostBannerRequestObject) (PostBannerResponseObject, error) {

	banner := ToEntity(request)
	bannerID, err := s.service.CreateBannerWithEntities(ctx, banner)
	if err != nil {
		errS := err.Error()
		errResp := PostBanner400JSONResponse{
			Error: &errS,
		}
		return errResp, err
	}

	return PostBanner201JSONResponse{
		BannerId: &bannerID,
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
