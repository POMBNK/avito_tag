// Package tag provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package tag

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
	"github.com/oapi-codegen/runtime"
	strictnethttp "github.com/oapi-codegen/runtime/strictmiddleware/nethttp"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Получение всех баннеров c фильтрацией по фиче и/или тегу
	// (GET /banner)
	GetBanner(w http.ResponseWriter, r *http.Request, params GetBannerParams)
	// Создание нового баннера
	// (POST /banner)
	PostBanner(w http.ResponseWriter, r *http.Request, params PostBannerParams)
	// Удаление баннера по идентификатору
	// (DELETE /banner/{id})
	DeleteBannerId(w http.ResponseWriter, r *http.Request, id int, params DeleteBannerIdParams)
	// Обновление содержимого баннера
	// (PATCH /banner/{id})
	PatchBannerId(w http.ResponseWriter, r *http.Request, id int, params PatchBannerIdParams)
	// Получение баннера для пользователя
	// (GET /user_banner)
	GetUserBanner(w http.ResponseWriter, r *http.Request, params GetUserBannerParams)
}

// Unimplemented server implementation that returns http.StatusNotImplemented for each endpoint.

type Unimplemented struct{}

// Получение всех баннеров c фильтрацией по фиче и/или тегу
// (GET /banner)
func (_ Unimplemented) GetBanner(w http.ResponseWriter, r *http.Request, params GetBannerParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Создание нового баннера
// (POST /banner)
func (_ Unimplemented) PostBanner(w http.ResponseWriter, r *http.Request, params PostBannerParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Удаление баннера по идентификатору
// (DELETE /banner/{id})
func (_ Unimplemented) DeleteBannerId(w http.ResponseWriter, r *http.Request, id int, params DeleteBannerIdParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Обновление содержимого баннера
// (PATCH /banner/{id})
func (_ Unimplemented) PatchBannerId(w http.ResponseWriter, r *http.Request, id int, params PatchBannerIdParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Получение баннера для пользователя
// (GET /user_banner)
func (_ Unimplemented) GetUserBanner(w http.ResponseWriter, r *http.Request, params GetUserBannerParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// GetBanner operation middleware
func (siw *ServerInterfaceWrapper) GetBanner(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetBannerParams

	// ------------- Optional query parameter "feature_id" -------------

	err = runtime.BindQueryParameter("form", true, false, "feature_id", r.URL.Query(), &params.FeatureId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "feature_id", Err: err})
		return
	}

	// ------------- Optional query parameter "tag_id" -------------

	err = runtime.BindQueryParameter("form", true, false, "tag_id", r.URL.Query(), &params.TagId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "tag_id", Err: err})
		return
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", r.URL.Query(), &params.Limit)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "limit", Err: err})
		return
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", r.URL.Query(), &params.Offset)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "offset", Err: err})
		return
	}

	headers := r.Header

	// ------------- Optional header parameter "token" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("token")]; found {
		var Token string
		n := len(valueList)
		if n != 1 {
			siw.ErrorHandlerFunc(w, r, &TooManyValuesForParamError{ParamName: "token", Count: n})
			return
		}

		err = runtime.BindStyledParameterWithOptions("simple", "token", valueList[0], &Token, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false})
		if err != nil {
			siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "token", Err: err})
			return
		}

		params.Token = &Token

	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetBanner(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PostBanner operation middleware
func (siw *ServerInterfaceWrapper) PostBanner(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params PostBannerParams

	headers := r.Header

	// ------------- Optional header parameter "token" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("token")]; found {
		var Token string
		n := len(valueList)
		if n != 1 {
			siw.ErrorHandlerFunc(w, r, &TooManyValuesForParamError{ParamName: "token", Count: n})
			return
		}

		err = runtime.BindStyledParameterWithOptions("simple", "token", valueList[0], &Token, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false})
		if err != nil {
			siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "token", Err: err})
			return
		}

		params.Token = &Token

	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostBanner(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeleteBannerId operation middleware
func (siw *ServerInterfaceWrapper) DeleteBannerId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params DeleteBannerIdParams

	headers := r.Header

	// ------------- Optional header parameter "token" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("token")]; found {
		var Token string
		n := len(valueList)
		if n != 1 {
			siw.ErrorHandlerFunc(w, r, &TooManyValuesForParamError{ParamName: "token", Count: n})
			return
		}

		err = runtime.BindStyledParameterWithOptions("simple", "token", valueList[0], &Token, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false})
		if err != nil {
			siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "token", Err: err})
			return
		}

		params.Token = &Token

	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteBannerId(w, r, id, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PatchBannerId operation middleware
func (siw *ServerInterfaceWrapper) PatchBannerId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params PatchBannerIdParams

	headers := r.Header

	// ------------- Optional header parameter "token" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("token")]; found {
		var Token string
		n := len(valueList)
		if n != 1 {
			siw.ErrorHandlerFunc(w, r, &TooManyValuesForParamError{ParamName: "token", Count: n})
			return
		}

		err = runtime.BindStyledParameterWithOptions("simple", "token", valueList[0], &Token, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false})
		if err != nil {
			siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "token", Err: err})
			return
		}

		params.Token = &Token

	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PatchBannerId(w, r, id, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetUserBanner operation middleware
func (siw *ServerInterfaceWrapper) GetUserBanner(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetUserBannerParams

	// ------------- Required query parameter "tag_id" -------------

	if paramValue := r.URL.Query().Get("tag_id"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "tag_id"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "tag_id", r.URL.Query(), &params.TagId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "tag_id", Err: err})
		return
	}

	// ------------- Required query parameter "feature_id" -------------

	if paramValue := r.URL.Query().Get("feature_id"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "feature_id"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "feature_id", r.URL.Query(), &params.FeatureId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "feature_id", Err: err})
		return
	}

	// ------------- Optional query parameter "use_last_revision" -------------

	err = runtime.BindQueryParameter("form", true, false, "use_last_revision", r.URL.Query(), &params.UseLastRevision)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "use_last_revision", Err: err})
		return
	}

	headers := r.Header

	// ------------- Optional header parameter "token" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("token")]; found {
		var Token string
		n := len(valueList)
		if n != 1 {
			siw.ErrorHandlerFunc(w, r, &TooManyValuesForParamError{ParamName: "token", Count: n})
			return
		}

		err = runtime.BindStyledParameterWithOptions("simple", "token", valueList[0], &Token, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false})
		if err != nil {
			siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "token", Err: err})
			return
		}

		params.Token = &Token

	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetUserBanner(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/banner", wrapper.GetBanner)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/banner", wrapper.PostBanner)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/banner/{id}", wrapper.DeleteBannerId)
	})
	r.Group(func(r chi.Router) {
		r.Patch(options.BaseURL+"/banner/{id}", wrapper.PatchBannerId)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/user_banner", wrapper.GetUserBanner)
	})

	return r
}

type GetBannerRequestObject struct {
	Params GetBannerParams
}

type GetBannerResponseObject interface {
	VisitGetBannerResponse(w http.ResponseWriter) error
}

type GetBanner200JSONResponse []struct {
	// BannerId Идентификатор баннера
	BannerId *int `json:"banner_id,omitempty"`

	// Content Содержимое баннера
	Content *map[string]interface{} `json:"content,omitempty"`

	// CreatedAt Дата создания баннера
	CreatedAt *time.Time `json:"created_at,omitempty"`

	// FeatureId Идентификатор фичи
	FeatureId *int `json:"feature_id,omitempty"`

	// IsActive Флаг активности баннера
	IsActive *bool `json:"is_active,omitempty"`

	// TagIds Идентификаторы тэгов
	TagIds *[]int `json:"tag_ids,omitempty"`

	// UpdatedAt Дата обновления баннера
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (response GetBanner200JSONResponse) VisitGetBannerResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetBanner401Response struct {
}

func (response GetBanner401Response) VisitGetBannerResponse(w http.ResponseWriter) error {
	w.WriteHeader(401)
	return nil
}

type GetBanner403Response struct {
}

func (response GetBanner403Response) VisitGetBannerResponse(w http.ResponseWriter) error {
	w.WriteHeader(403)
	return nil
}

type GetBanner500JSONResponse struct {
	Error *string `json:"error,omitempty"`
}

func (response GetBanner500JSONResponse) VisitGetBannerResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type PostBannerRequestObject struct {
	Params PostBannerParams
	Body   *PostBannerJSONRequestBody
}

type PostBannerResponseObject interface {
	VisitPostBannerResponse(w http.ResponseWriter) error
}

type PostBanner201JSONResponse struct {
	// BannerId Идентификатор созданного баннера
	BannerId *int `json:"banner_id,omitempty"`
}

func (response PostBanner201JSONResponse) VisitPostBannerResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type PostBanner400JSONResponse struct {
	Error *string `json:"error,omitempty"`
}

func (response PostBanner400JSONResponse) VisitPostBannerResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type PostBanner401Response struct {
}

func (response PostBanner401Response) VisitPostBannerResponse(w http.ResponseWriter) error {
	w.WriteHeader(401)
	return nil
}

type PostBanner403Response struct {
}

func (response PostBanner403Response) VisitPostBannerResponse(w http.ResponseWriter) error {
	w.WriteHeader(403)
	return nil
}

type PostBanner500JSONResponse struct {
	Error *string `json:"error,omitempty"`
}

func (response PostBanner500JSONResponse) VisitPostBannerResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type DeleteBannerIdRequestObject struct {
	Id     int `json:"id"`
	Params DeleteBannerIdParams
}

type DeleteBannerIdResponseObject interface {
	VisitDeleteBannerIdResponse(w http.ResponseWriter) error
}

type DeleteBannerId204Response struct {
}

func (response DeleteBannerId204Response) VisitDeleteBannerIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(204)
	return nil
}

type DeleteBannerId400JSONResponse struct {
	Error *string `json:"error,omitempty"`
}

func (response DeleteBannerId400JSONResponse) VisitDeleteBannerIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type DeleteBannerId401Response struct {
}

func (response DeleteBannerId401Response) VisitDeleteBannerIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(401)
	return nil
}

type DeleteBannerId403Response struct {
}

func (response DeleteBannerId403Response) VisitDeleteBannerIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(403)
	return nil
}

type DeleteBannerId404Response struct {
}

func (response DeleteBannerId404Response) VisitDeleteBannerIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type DeleteBannerId500JSONResponse struct {
	Error *string `json:"error,omitempty"`
}

func (response DeleteBannerId500JSONResponse) VisitDeleteBannerIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type PatchBannerIdRequestObject struct {
	Id     int `json:"id"`
	Params PatchBannerIdParams
	Body   *PatchBannerIdJSONRequestBody
}

type PatchBannerIdResponseObject interface {
	VisitPatchBannerIdResponse(w http.ResponseWriter) error
}

type PatchBannerId200Response struct {
}

func (response PatchBannerId200Response) VisitPatchBannerIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type PatchBannerId400JSONResponse struct {
	Error *string `json:"error,omitempty"`
}

func (response PatchBannerId400JSONResponse) VisitPatchBannerIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type PatchBannerId401Response struct {
}

func (response PatchBannerId401Response) VisitPatchBannerIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(401)
	return nil
}

type PatchBannerId403Response struct {
}

func (response PatchBannerId403Response) VisitPatchBannerIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(403)
	return nil
}

type PatchBannerId404Response struct {
}

func (response PatchBannerId404Response) VisitPatchBannerIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type PatchBannerId500JSONResponse struct {
	Error *string `json:"error,omitempty"`
}

func (response PatchBannerId500JSONResponse) VisitPatchBannerIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type GetUserBannerRequestObject struct {
	Params GetUserBannerParams
}

type GetUserBannerResponseObject interface {
	VisitGetUserBannerResponse(w http.ResponseWriter) error
}

type GetUserBanner200JSONResponse map[string]interface{}

func (response GetUserBanner200JSONResponse) VisitGetUserBannerResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetUserBanner400JSONResponse struct {
	Error *string `json:"error,omitempty"`
}

func (response GetUserBanner400JSONResponse) VisitGetUserBannerResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type GetUserBanner401Response struct {
}

func (response GetUserBanner401Response) VisitGetUserBannerResponse(w http.ResponseWriter) error {
	w.WriteHeader(401)
	return nil
}

type GetUserBanner403Response struct {
}

func (response GetUserBanner403Response) VisitGetUserBannerResponse(w http.ResponseWriter) error {
	w.WriteHeader(403)
	return nil
}

type GetUserBanner404Response struct {
}

func (response GetUserBanner404Response) VisitGetUserBannerResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type GetUserBanner500JSONResponse struct {
	Error *string `json:"error,omitempty"`
}

func (response GetUserBanner500JSONResponse) VisitGetUserBannerResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Получение всех баннеров c фильтрацией по фиче и/или тегу
	// (GET /banner)
	GetBanner(ctx context.Context, request GetBannerRequestObject) (GetBannerResponseObject, error)
	// Создание нового баннера
	// (POST /banner)
	PostBanner(ctx context.Context, request PostBannerRequestObject) (PostBannerResponseObject, error)
	// Удаление баннера по идентификатору
	// (DELETE /banner/{id})
	DeleteBannerId(ctx context.Context, request DeleteBannerIdRequestObject) (DeleteBannerIdResponseObject, error)
	// Обновление содержимого баннера
	// (PATCH /banner/{id})
	PatchBannerId(ctx context.Context, request PatchBannerIdRequestObject) (PatchBannerIdResponseObject, error)
	// Получение баннера для пользователя
	// (GET /user_banner)
	GetUserBanner(ctx context.Context, request GetUserBannerRequestObject) (GetUserBannerResponseObject, error)
}

type StrictHandlerFunc = strictnethttp.StrictHTTPHandlerFunc
type StrictMiddlewareFunc = strictnethttp.StrictHTTPMiddlewareFunc

type StrictHTTPServerOptions struct {
	RequestErrorHandlerFunc  func(w http.ResponseWriter, r *http.Request, err error)
	ResponseErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: StrictHTTPServerOptions{
		RequestErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		},
		ResponseErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		},
	}}
}

func NewStrictHandlerWithOptions(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc, options StrictHTTPServerOptions) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: options}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
	options     StrictHTTPServerOptions
}

// GetBanner operation middleware
func (sh *strictHandler) GetBanner(w http.ResponseWriter, r *http.Request, params GetBannerParams) {
	var request GetBannerRequestObject

	request.Params = params

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetBanner(ctx, request.(GetBannerRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetBanner")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetBannerResponseObject); ok {
		if err := validResponse.VisitGetBannerResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// PostBanner operation middleware
func (sh *strictHandler) PostBanner(w http.ResponseWriter, r *http.Request, params PostBannerParams) {
	var request PostBannerRequestObject

	request.Params = params

	var body PostBannerJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.PostBanner(ctx, request.(PostBannerRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostBanner")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(PostBannerResponseObject); ok {
		if err := validResponse.VisitPostBannerResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// DeleteBannerId operation middleware
func (sh *strictHandler) DeleteBannerId(w http.ResponseWriter, r *http.Request, id int, params DeleteBannerIdParams) {
	var request DeleteBannerIdRequestObject

	request.Id = id
	request.Params = params

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteBannerId(ctx, request.(DeleteBannerIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteBannerId")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(DeleteBannerIdResponseObject); ok {
		if err := validResponse.VisitDeleteBannerIdResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// PatchBannerId operation middleware
func (sh *strictHandler) PatchBannerId(w http.ResponseWriter, r *http.Request, id int, params PatchBannerIdParams) {
	var request PatchBannerIdRequestObject

	request.Id = id
	request.Params = params

	var body PatchBannerIdJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.PatchBannerId(ctx, request.(PatchBannerIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PatchBannerId")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(PatchBannerIdResponseObject); ok {
		if err := validResponse.VisitPatchBannerIdResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetUserBanner operation middleware
func (sh *strictHandler) GetUserBanner(w http.ResponseWriter, r *http.Request, params GetUserBannerParams) {
	var request GetUserBannerRequestObject

	request.Params = params

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetUserBanner(ctx, request.(GetUserBannerRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetUserBanner")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetUserBannerResponseObject); ok {
		if err := validResponse.VisitGetUserBannerResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xZ3U4bRxR+ldVcb4LT0BtfppWqtFITqepVqazBO4ZJ9y8zYxRkWQKbNpFCg1r1olLV",
	"VKR9gMXgxIF4eYUzb1SdGRN72V2wIZAW+cre2d3zf77vzGyL8LARkWqLKK58RqoEdqGvN6AHA73pwB4k",
	"MIShWUqhR1yyxoTkUUiq5M7tyu0KabskillIY06q5K5ZcklM1apEqQvLNAyZwL8rTOFPFDNBFY/C+x6p",
	"ki+YumefwJcEDZhiQpLqdy3iMVkXPFZWGbyCFA6hD0MHEjiAdzCAISTEJRxvrzLqGSEhDdALFf3AQuIS",
	"WV9lAUW97AkNYuMh9QIe1k6eUOsxLkoleLhC2m23ZSU+bjKxPhbYYFQ1BatxLyP1lJG/wwGaqDsw0Fsw",
	"gENIdAdSveHgpX4Kg7FGHiq2wkS5SkVXLq6uA33YN/GZVp3PA67O0vYHDDDsujOD0KjRkOxMqS/1lt7S",
	"m9Avlvu9SwSTcRRKZgrqk0oFf+pRqFhoCorGsc/rpqQWHkkU2prQxhULzIuxwMpT3IqxZYnhnSWqmXYo",
	"iq2bMczzOMqk/sMJ3Uo0mXta5S6kqFRvwGsT5BT6eWXjCm4t2XZdIlVnicgoYLXRtessEcWeqMk75hJv",
	"NIU/sW6u2mMnouVHrK6MD4JRxbwaVQXR+c1EI3H0JqTwBg6MkQO9k7e3EYkARRCPKnZL8YDl282dbKwP",
	"1E8u4bJG64qvsQKR/8ARJLCPKHKIcqEHQ0j1Jv4vTfByFPmMhijbdqWcxVj93NEd/TPsjyD0fU3mLR+t",
	"UCHoOl43Y+/8VEAKe+gE9OAILbh4OtoF1ZA1CR/JmvHgK3xqsXKnwMK/IIUjvQ1vjHGJQaUjve2gWZiB",
	"ng0QDE6egKEVdndmYdg2fUQRBw5sPnUXjiFBeZ/OCBpZrGBCRGIiXeXBygcHfoWh7uqO3jCJGeodTE2q",
	"n8EA9rBCHAN9yLc2TyhCNoOAivWxz1391OYVHe2ZN37McbNTN02BITH6Ev0TvgBvHTiG9KRhMFIL+BQM",
	"RiShuxiiOJIF9Pwwkv81fjaE8LjJpLoXeeuXSOtNQOo5eJaDZ3F7Yu1wwTyb3nZuurhziYq62FCRIVJE",
	"8X1Iz580pgGfzyyPW0T9aAj4J/Th0GQSMRDLZqifY9eMPMaLOYNcGYPsTo5p6KYdFArLDF8d7dgWWtxr",
	"2xD6TLE8NXxu1i053Pfy9GBgH/eBY9A3m5lsA15oczNFc3xcesogymKBb7+MHXB0V2/CMfT1M0yNo7sm",
	"WWaQm3fu5Tt38dwEwAEc6Z0TlklGMoeQwFtbgv9vAPh7XE8WADLdY4dDGJTSr50OqaqvFoyHuDyHgPmE",
	"+n5CDZu+T5dRsLXtCibWEhVXMsGW6Lq2ibZE/6Um3Ere1pMt/JxprphpbhazvDx97AR9u5vJwk/ZoNmU",
	"TNTO/z7wrWSi7AzizGPzqWnmFTajIcKCzO/McNqd+UjwAVjuAh8NmpLVfCpVTbA1bj7UZDU3aNNXpNqg",
	"vsxTx/i0KdEdrHkDkrqL84PeNsXyAntgqLdMQ72zh0z6RcFm/2wCLo/1zGRsymj6cXy2Zpued7/85sHX",
	"tyDFzMEeFjm8Lhm5ru+AqKDlM3BUloU5GVzftuOGcULBgfWp/cbI7dLaa7fb/wYAAP//8ZJEs5YeAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
