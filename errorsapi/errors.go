package errorsapi

import "errors"

var (
	// ErrInputEmpty input is empty
	ErrInputEmpty error = errors.New("Input cannot be empty")
	// ErrFieldEmpty field is empty
	ErrFieldEmpty error = errors.New("this field cannot be empty")
	//ErrInvalidBody if tried to unmarshall from json and failes at it
	ErrInvalidBody = errors.New("The sent data is not as expected")

	// ErrConnectionDB couldnt connect to database
	ErrConnectionDB = errors.New("Failed to connect to database")

	// ErrNotStartTransactionDB error by commit transaction
	ErrNotStartTransactionDB = errors.New("Transaction not started")

	// ErrFindTableDB couldn't find to query executed
	ErrFindTableDB = errors.New("Failed couldn't find to query executed")

	// ErrNotFoundTableDB No results found in table
	ErrNotFoundTableDB = errors.New("No results found in table")

	// ErrNotInsertTableDB return err when not insert
	ErrNotInsertTableDB = errors.New("Could not enter required information")

	// ErrNotUpdate error ao realizar o update
	ErrNotUpdate = errors.New("Could not perform requested update")

	// ErrTransactionCommitDB error by commit transaction
	ErrTransactionCommitDB = errors.New("Transaction not executable")

	// ErrLogin the username or password is not valid
	ErrLogin = errors.New("Failed to authenticate username")

	// ErrToken invalid token
	ErrToken = errors.New("Failed validate token")

	// ErrAlreadyExist dupe data
	ErrAlreadyExist = errors.New("Data already exists")

	// ErrChangePassword error changing user password
	ErrChangePassword = errors.New("Can't update user password")

	// ErrDateTimeExpires error to date time expired
	ErrDateTimeExpires = errors.New("Time to update expired")

	// ErrURLRequest error to http new reques
	ErrURLRequest = errors.New("Cannot the request")

	// ErrFinancialNotFound return error to financial not found
	ErrFinancialNotFound = errors.New("financial not found")

	// ErrFinancialOperation return error to fianancial operation not performed
	ErrFinancialOperation = errors.New("Operation not performed")

	// ErrDataBaseConnect return error connection to the database
	ErrDataBaseConnect = errors.New("couldnt connect to database")

	// ErrNotWorkingDay return error working day
	ErrNotWorkingDay = errors.New("Today is not working day")

	// ErrNotReceivedContent return error no content received
	ErrNotReceivedContent = errors.New("Not received content")

	// ErrDupAccount return error the account duplicated
	ErrDupAccount = errors.New("ops, the account duplicated")

	// ErrAddInformationOnDB return error the insert
	ErrAddInformationOnDB = errors.New("Error, could not create the data")

	// ErrDocNotFound return err to document no found
	ErrDocNotFound = errors.New("Document not found")

	// ErrServerMaintenance error is for when the server is in maintenance
	ErrServerMaintenance = errors.New("Server is in maintenance")
	// ErrTransactionAmountIsNegative error when the amount is positive
	ErrTransactionAmountIsNegative = errors.New("The value must be negative")
	// ErrTransactionAmountIsPositive error when the amount is negative
	ErrTransactionAmountIsPositive = errors.New("The value must be positive")
)
