package bithavocid

import (
    "fmt"
)

type MappedErrorList map[string]map[string]string
type FlatErrorList map[string]string

func (mappedErrors MappedErrorList) flatenize() FlatErrorList {
    flat := make(FlatErrorList)
    if mappedErrors != nil {
        for fieldName, e := range mappedErrors {
            if e != nil {
                flat[fieldName] = e["msg"]
            }
        }
    }
    return flat
}

type IdError struct {
    FirstMessage string
    Messages FlatErrorList
}

func (err IdError) Error() string {
    return err.FirstMessage
}

func (messages FlatErrorList) ToError() IdError {
    err:= IdError {
        Messages: messages,
    }
    for fieldName, message := range messages {
        err.FirstMessage = fmt.Sprintf("%s -> %s", fieldName, message)
        break
    }
    return err
}

func (result baseResult) checkErrors() error {
    errors := result.Messages.flatenize()
    if len(errors) > 0 {
        return errors.ToError()
    }
    return nil
}

