package rest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"training-frontend/package/log"
	"training-frontend/server/systems/helpers"
	"reflect"
	"strings"
	"time"

	"github.com/k0kubun/pp"
)

type Action int

const (
	BillSubmissionRequest Action = iota
	BillSubmissionAck
	BillCancellationRequest
	BilCancellationAck
	ControlNumberResponse
	ControlNumberAck
	PaymentResponse
	PaymentAck
	ReconcilationRequest
	ReconcilationAck
	ReconcilianResponse
)

type ResponseRequest struct {
	TimeStamp  time.Time `json:"timestamp"`
	StatusCode int32     `json:"status_code"`
	HttpStatus string    `json:"http_status"`
	Message    string    `json:"message"`
	Count      int32     `json:"count"`
	Action     Action    `json:"action"`
	Data       interface {
	} `json:"data,omitempty"`
}

func DecodeResponse(data []byte, value interface{}) (*ResponseRequest, error) {
	resp := &ResponseRequest{}
	err := json.Unmarshal(data, resp)
	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return nil, err
	}
	decode(resp.Data, value)
	resp.Data = value
	return resp, nil
}

// converts maps into struct
func decode(in, out interface{}) {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(in)
	json.NewDecoder(buf).Decode(out)
}

// EntityDecoder response decoder
func EntityDecoder(resp *Response, dtoModel interface{}) (helpers.Map, error) {

	r, err := DecodeResponse(resp.Body, &dtoModel)
	if err != nil {
		log.Errorf("Error decoding response %v", err)
	}

	data := helpers.Map{
		"data":       r.Data,
		"message":    r.Message,
		"statusCode": r.StatusCode,
		"count":      r.Count,
	}

	return data, err
}

func ResponseSuccessData(data interface{}) *ResponseRequest {
	val := reflect.ValueOf(data)
	msgDataResp := &ResponseRequest{
		TimeStamp:  time.Now(),
		StatusCode: http.StatusOK,
		HttpStatus: http.StatusText(http.StatusOK),
		Message:    "Success",
		Count:      int32(val.Len()),
		Data:       data,
	}
	return msgDataResp
}

func ResponseErrorData(data interface{}, errMsg string) *ResponseRequest {
	val := reflect.ValueOf(data)
	msgDataResp := &ResponseRequest{
		TimeStamp:  time.Now(),
		StatusCode: http.StatusInternalServerError,
		HttpStatus: http.StatusText(http.StatusInternalServerError),
		Message:    errMsg,
		Count:      int32(val.Len()),
		Data:       data,
	}
	return msgDataResp
}

func JsonResponseMessage(message, controllerName string, isCreated, isDeleted bool) string {
	if strings.Contains(strings.ToLower(message), strings.ToLower("Success")) {
		if isCreated {
			message = controllerName + " created successfully"
			return message
		} else if isDeleted {
			message = controllerName + " deleted successfully"
			return message
		}
		message = controllerName + " updated successfully"
		return message
	} else {
		if isCreated {
			message = "Error occurred. " + controllerName + " not created"
			return message
		} else if isDeleted {
			message = "Error occurred. " + controllerName + " not deleted"
			return message
		}
		message = "Error occurred. " + controllerName + " not updated"
		return message
	}
}

func Booleans() map[string]string {
	myBooleans := map[string]string{
		"True":  "True",
		"False": "False",
	}
	return myBooleans
}
