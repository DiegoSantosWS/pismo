package apianswer

import (
	"encoding/json"
	"log"
	"net/http"
	"pismo/errorsapi"
)

// AnswerRequest answers the request as needed using the framework echo
func AnswerRequest(w http.ResponseWriter, r *http.Request, in interface{}, err error) {
	statusCode := http.StatusBadRequest
	if err == nil {
		statusCode = http.StatusOK
	}
	switch err {
	case errorsapi.ErrInvalidBody:
		statusCode = http.StatusPreconditionFailed
	case errorsapi.ErrAddInformationOnDB, errorsapi.ErrConnectionDB, errorsapi.ErrNotStartTransactionDB, errorsapi.ErrConnectionDB, errorsapi.ErrNotStartTransactionDB, errorsapi.ErrServerMaintenance:
		statusCode = http.StatusServiceUnavailable
	case errorsapi.ErrLogin, errorsapi.ErrToken, errorsapi.ErrLogin, errorsapi.ErrToken:
		statusCode = http.StatusUnauthorized
	case errorsapi.ErrAlreadyExist, errorsapi.ErrAlreadyExist:
		statusCode = http.StatusConflict
	case errorsapi.ErrNotFoundTableDB, errorsapi.ErrDateTimeExpires, errorsapi.ErrNotWorkingDay, errorsapi.ErrNotFoundTableDB, errorsapi.ErrDateTimeExpires:
		statusCode = http.StatusNotAcceptable
	case errorsapi.ErrChangePassword, errorsapi.ErrChangePassword:
		statusCode = http.StatusForbidden
	case errorsapi.ErrFinancialNotFound, errorsapi.ErrDocNotFound:
		statusCode = http.StatusNotFound
	}

	res, err := json.Marshal(in)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if _, err := w.Write(res); err != nil {
		log.Println(err)
	}
}
