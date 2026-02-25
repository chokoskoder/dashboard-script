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


)

type ErrorCode string
const (
    // System / Infrastructure Errors
	CodeTransientFailure ErrorCode	= "TRANSIENT_SYSTEM_FAILURE"
	CodeSourceConn       ErrorCode	= "SOURCE_CONNECTION_FAILED"
	CodeTargetConn       ErrorCode	= "TARGET_CONNECTION_FAILED"//our source and target is the same bruv what the helly
    
    // Data / Logic Errors
	//CodeDataValidation   ErrorCode	= "DATA_VALIDATION_FAILED" // we dont need a data validation error
	CodeTargetConstraint ErrorCode	= "TARGET_CONSTRAINT_VIOLATION"
	//this should be something like today it has been done once and we dont want to repeat even if we end up running the db IDEMPOTENCY
	CodeFileParse        ErrorCode	= "FILE_PARSING_ERROR"
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

func NewTargetConstraintError(message string , recordID string , originalErr error) *ETLError {
	e := &ETLError{}
	
	return e
}


//why am I 
