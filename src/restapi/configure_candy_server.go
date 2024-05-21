// This file is safe to edit. Once it exists it will not be overwritten

package restapi

/*
#cgo CFLAGS: -I../internal/cow
#include "cow.h"
*/
import "C"
import (
	"CandyServer/internal/prices"
	"crypto/tls"
	"fmt"
	"net/http"
	"unsafe"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"CandyServer/restapi/operations"
)

//go:generate swagger generate server --target ../../src --name CandyServer --spec ../swagger/swagger.yaml --principal interface{}

func configureFlags(api *operations.CandyServerAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.CandyServerAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.BuyCandyHandler = operations.BuyCandyHandlerFunc(func(params operations.BuyCandyParams) middleware.Responder {
		price, ok := prices.Prices[*params.Order.CandyType]
		if *params.Order.Money < 0 || *params.Order.CandyCount < 0 || !ok {
			var resp string
			if *params.Order.Money < 0 {
				resp += "wrong amount of money  "
			}
			if *params.Order.CandyCount < 0 {
				resp += "wrong count of candy | "
			}
			if !ok {
				resp += "wrong candy type"
			}
			res := operations.NewBuyCandyBadRequest()
			res.SetPayload(&operations.BuyCandyBadRequestBody{Error: resp})
			return res
		}
		change := (*params.Order.Money) - price*(*params.Order.CandyCount)
		if change < 0 {
			resp := fmt.Sprintf("You neeed %d more money!", -change)
			res := operations.NewBuyCandyPaymentRequired()
			res.SetPayload(&operations.BuyCandyPaymentRequiredBody{Error: resp})
			return res
		}
		res := operations.NewBuyCandyCreated()
		goString := "Thank you!"
		cString := C.CString(goString)
		defer C.free(unsafe.Pointer(cString))
		responseC := C.ask_cow(cString)
		responseGo := C.GoString(responseC)
		res.SetPayload(&operations.BuyCandyCreatedBody{
			Change: change,
			Thanks: responseGo,
		})
		return res
	})

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
