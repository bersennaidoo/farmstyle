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
				ErrorResponse(problems.InvalidBody(problems.ProblemJson{
					Detail: errVal.Reason,
				}))(w, r)
				return
			case *openapi3filter.SecurityRequirementsError:
				// Allow this for now ( optional securities appear to be an issue )
				log.Printf("errVal %s", errVal)
				break
			default:
				ErrorResponse(problems.InvalidRequest(problems.ProblemJson{}))(w, r)
				return
			}
		}

		// All good, carry on...
		next.ServeHTTP(w, r)
	}))
}
