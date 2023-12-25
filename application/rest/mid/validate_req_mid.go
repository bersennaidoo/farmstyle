package mid

import (
	"context"
	"log"
	"net/http"

	"github.com/bersennaidoo/farmstyle/foundation/emsg"
	"github.com/getkin/kin-openapi/openapi3filter"
)

func ValidateRequestMiddleware(next http.Handler) http.Handler {
	return http.StripPrefix("/v1", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Find operation
		router := openapi3filter.NewRouter().WithSwaggerFromFile("openapi.yaml")
		route, pathParams, errOp := router.FindRoute(r.Method, r.URL)

		if errOp != nil {
			log.Printf("Operation not found for %s %s. Error: %s", r.Method, r.URL, errOp)
			next.ServeHTTP(w, r)
			return
		}

		// Validate request against operation
		requestValidationInput := &openapi3filter.RequestValidationInput{
			Request:    r,
			PathParams: pathParams,
			Route:      route,
		}

		something := context.TODO()

		if err := openapi3filter.ValidateRequest(something, requestValidationInput); err != nil {
			switch errVal := err.(type) {
			case *openapi3filter.RequestError:
				ErrorResponse(emsg.InvalidBody(emsg.ProblemJson{
					Detail: errVal.Reason,
				}))(w, r)
				return
			case *openapi3filter.SecurityRequirementsError:
				// Allow this for now ( optional securities appear to be an issue )
				log.Printf("errVal %s", errVal)
				break
			default:
				ErrorResponse(emsg.InvalidRequest(emsg.ProblemJson{}))(w, r)
				return
			}
		}

		// All good, carry on...
		next.ServeHTTP(w, r)
	}))
}
