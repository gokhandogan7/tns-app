package internal

// JSONValue is a convenience alias to remember that a HandleJson
// function return value is encoded as json
type JSONValue interface{}

// HTTPResponse convenience struct to allow setting a custom
// http status code
type HTTPResponse struct {
	Status int
	Body   JSONValue
}
