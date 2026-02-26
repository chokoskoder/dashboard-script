package errors

import (
	"context"
	"errors"
	"fmt"
	"math"
	"time"
)

//max retries logic
//exponential backoff logic will come here

type RetryConfig struct {
	MaxRetries int
	BaseDelay  time.Duration
}

var DefaultRetryPolicy = RetryConfig{
	MaxRetries: 3,
	BaseDelay:  1 * time.Second,
}

func RunWithRetry(ctx context.Context , operation func() error) error { //doesnt this func() decalration mean that the function can only return error??
	var lastErr error

	//run a loopo for the retry mechansim
	for attempt := 0 ; attempt <= DefaultRetryPolicy.MaxRetries; attempt++{
		err := operation()//run the operation

		if err == nil {
			return nil
		}
		lastErr = err

		var etlError *ETLError
		if errors.As(err,&etlError){
			//if the error says dont retry , stop immediately -> but why are we passing a non retry function to this retry function in the first place ? 
			//from my understanding we are not going to call this on each and every function , but like a function in the service directlory where a lot of small function will be working inside it and thus it will be checked here
			if !etlError.Retryable{
				return err
			} else if etlError.Severity == "FATAL" {
				//we need to check here if the error is fatal and we dont want to continue 
				return err
			} else if etlError.Retryable && etlError.Severity == "CRITICAL"{
				continue
			}

			if attempt == DefaultRetryPolicy.MaxRetries {
				break
			}

			backoff := time.Duration(math.Pow(2 , float64(attempt))) * DefaultRetryPolicy.BaseDelay
			fmt.Printf("Attempt %d failed: %v. Retrying in %v...\n", attempt+1, err, backoff)
			
			select {
			case <- time.After(backoff):
				continue
			case <- ctx.Done():
				return ctx.Err()
			}
		}

	}
	return fmt.Errorf("operation failed after %d attempts. Last error: %w", DefaultRetryPolicy.MaxRetries, lastErr)

	
}