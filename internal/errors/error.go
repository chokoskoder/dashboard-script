package errors

import (
	"fmt"
)

//we need specific errors to our ETL script

type ETLStage string

const (
	StageExtract	ETLStage = "EXTRACT"
	StageLoad		ETLStage = "LOAD"
	StageTransform	ETLStage = "TRANSFORM"
	StageGeneral	ETLStage = "GENERAL"
)
//Enum for etl stage

type Severity string
const (
	SeverityFatal    Severity = "FATAL"    // Stop the entire pipeline immediately
	SeverityWarning  Severity = "WARNING"  // Log and continue (skip this record)
	SeverityCritical Severity = "CRITICAL" // Retry, then fail if persistent
)

type ErrorCode string
const (
	//updating ErrorCodes:
	CodeDBConn 				ErrorCode = "DB_CONNECTION_FAILED"
	CodeRecordAlreadyExists ErrorCode = "RECORD_ALREADY_PROCESSED"
	CodeFileParse			ErrorCode = "FILE_PARSING_ERROR"
	CodeInternal			ErrorCode = "INTERNAL_SYSTEM_ERROR"
	//generic errors ? -> there can be no generic errors which we can :
	//should we enable retries for generic errors ??
	CodeTransientFailure	ErrorCode = "TRANSIENT_SYSTEM_FAILURE"
	//what comes under internal ?

)

//ENum for severity , 3 levels -> retry mechanism will be related to this

type ETLError struct {
	Code		string					//This is something we will define on our own
	Message		string					//Human readable message
	Stage		ETLStage
	Severity	Severity
	Retryable 	bool					//IMPORTANT -> we need to make the script Idempotent
	RecordID	string					
	Metadata	map[string]interface{}	//extra content which will help in debugging
	Err			error					//the original underlying error
}

//this is an insider method kinda thing which will override throwing errors using the Error() function now whenever we use error related to ETLError struct we will print this 
func (e *ETLError) Error() string{
	return fmt.Sprintf("[%s] %s: %s" , e.Stage , e.Code , e.Message)
}

//Unwrap function , never seen it before
func (e *ETLError) Unwrap() error {
	return e.Err
}
//ok so from my understanding what we need is a retry mechanism for situations like when db connection fails OR we dont have a network connection.
//we dont need a data validation error because that is already implemented at the place we are getting our data from
//we need errors for when we cant connect to our source -> this seems similar to the first error why do we need a different one like this ?
//parsing errors , if we cant parse the files we will not be able to work with them

func NewIdempotencyError(recordID string, originalErr error) *ETLError {
	return &ETLError{
		Code:      string(CodeRecordAlreadyExists), // Cast to string if struct expects string
		Message:   "Record has already been processed",
		Stage:     StageLoad,       // It usually happens during the Write/Load phase
		Severity:  SeverityWarning, // WARNING = Don't wake up the engineer, just log it
		Retryable: false,           // NEVER retry a duplicate, it will fail forever
		RecordID:  recordID,
		Err:       originalErr,
	}
}

// NewDBConnectionError handles your single DB connection issues
func NewDBConnectionError(op string, originalErr error) *ETLError {
	return &ETLError{
		Code:      string(CodeDBConn),
		Message:   "Database connection failed during: " + op,
		Stage:     StageGeneral,    // Could be extract or load
		Severity:  SeverityFatal,   // If DB is down, stop the script
		Retryable: true,            // DB might come back up, so we can retry
		Err:       originalErr,
	}
}

func NewTransientFailure(stage ETLStage , message string , originalErr error) *ETLError{
	return &ETLError{
		Code : string(CodeTransientFailure),
		Message: message,
		Stage: StageGeneral,
		Severity: SeverityCritical, //what is the difference in not being able to connect to db and having basic network issues ?? this error is for timeouts , generic errors which can 

	}
}
