package goboot

// A provided simple json response message for sending to a client
type JsonResponse struct {
	// Optional text field, can be used to communicate 
	// additional error information
	Text			string				`json:"text,omitempty"`
	// Optional text field, echo back from requests
	RequestId	string				`json:"request_id,omitempty"`
	// The payload data, a varible type of data 
	Payload		interface{}   `json:"payload"`
}