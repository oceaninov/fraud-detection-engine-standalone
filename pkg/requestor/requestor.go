package requestor

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"io"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
)

// Endpoint struct defines the structure for API endpoints configuration
type Endpoint struct {
	Path        string        // API endpoints path
	Method      string        // HTTP method (GET, POST, etc.)
	ConnTimeout time.Duration // Connection timeout duration
	CallTimeout time.Duration // Call timeout duration
	TLSTimeout  time.Duration // TLS handshake timeout duration
	KeepAlive   time.Duration // Keep-alive duration for the connection
}

// Client struct represents the client that will be used to make API requests
type Client struct {
	Namespace string              // Namespace for the client
	BaseUrl   string              // Base URL for the API
	Endpoints map[string]Endpoint // Map of endpoints configurations
	Logger    *zap.SugaredLogger  // Logger instance
}

// Request struct represents the details of an API request
type Request struct {
	EndpointName string      // Name of the endpoints
	Header       interface{} // Request headers
	Query        interface{} // Query parameters
	Body         interface{} // Request body
	Path         interface{} // Path parameters
}

// NewClient initializes and returns a new Client instance
func NewClient(namespace, baseurl string) *Client {
	c := new(Client)
	c.Namespace = namespace
	c.BaseUrl = baseurl
	c.Endpoints = make(map[string]Endpoint) // Initialize the Endpoints map
	return c
}

// SetLogger sets the logger for the client
func (c *Client) SetLogger(log *zap.SugaredLogger) {
	c.Logger = log
}

// SetEndpoints sets the endpoints for the client
func (c *Client) SetEndpoints(ep map[string]Endpoint) {
	c.Endpoints = ep
}

// body marshals the request body data into an io.Reader
func (c *Client) body(data interface{}) io.Reader {
	bodyBytes, err := json.Marshal(data)
	if err != nil {
		return nil // Return nil if there is an error in marshaling
	}
	return bytes.NewBuffer(bodyBytes)
}

// header converts the struct fields to a map of request headers
func (c *Client) header(data interface{}) map[string][]string {
	t := reflect.TypeOf(data)  // Get the type of the struct
	v := reflect.ValueOf(data) // Get the value of the struct

	headerData := make(map[string][]string) // Initialize the header data map

	// Iterate over the fields of the struct
	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get("header")  // Get the field tag value
		fieldValue := v.Field(i).Interface() // Get the field value

		fieldNil := fieldValue != nil && fmt.Sprintf("%v", fieldValue) != "<nil>"
		if fieldNil && fieldValue != "" {
			// Add the field value to the header data map
			headerData[tag] = []string{fmt.Sprintf("%v", fieldValue)}
		}
	}
	return headerData
}

// query converts the struct fields to URL query parameters
func (c *Client) query(data interface{}) string {
	t := reflect.TypeOf(data)  // Get the type of the struct
	v := reflect.ValueOf(data) // Get the value of the struct

	queryParams := url.Values{} // Initialize the URL values map

	// Iterate over the fields of the struct
	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get("query")   // Get the field tag value
		fieldValue := v.Field(i).Interface() // Get the field value

		fieldNil := fieldValue != nil && fmt.Sprintf("%v", fieldValue) != "<nil>"
		if fieldNil && fieldValue != "" {
			// Add the field value to the query parameters
			queryParams.Add(tag, fmt.Sprintf("%v", fieldValue))
		}
	}
	return queryParams.Encode() // Encode the query parameters
}

// path replaces path parameters in the given path template with the struct field values
func (c *Client) path(data interface{}, path string) string {
	t := reflect.TypeOf(data)  // Get the type of the struct
	v := reflect.ValueOf(data) // Get the value of the struct

	// Iterate over the fields of the struct
	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get("path")    // Get the field tag value
		fieldValue := v.Field(i).Interface() // Get the field value

		fieldNil := fieldValue != nil && fmt.Sprintf("%v", fieldValue) != "<nil>"
		if fieldNil && fieldValue != "" {
			// Replace the path parameter with the field value
			path = strings.Replace(path, ":"+tag, fmt.Sprintf("%v", fieldValue), 1)
		}
	}
	return path
}

// createHttpClient creates and returns an HTTP client with custom settings based on the endpoints configuration
func (c *Client) createHttpClient(endpoint Endpoint) *http.Client {
	client := new(http.Client)
	client.Timeout = endpoint.CallTimeout // Set the call timeout
	client.Transport = &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   endpoint.ConnTimeout, // Set the connection timeout
			KeepAlive: endpoint.KeepAlive,   // Set the keep-alive duration
		}).DialContext,
		MaxIdleConnsPerHost: 50,                                    // Set the maximum idle connections per host
		TLSHandshakeTimeout: endpoint.TLSTimeout,                   // Set the TLS handshake timeout
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true}, // Skip TLS verification
	}
	return client
}

// CreateRequest creates an HTTP request and returns the client, request, and any error encountered
func (c *Client) CreateRequest(request Request) (*http.Client, *http.Request, error) {
	var err error
	var req *http.Request
	var bodyBuffer io.Reader

	endpoint, ok := c.Endpoints[request.EndpointName] // Retrieve the endpoints configuration
	if !ok {
		return nil, nil, errors.New("endpoints not found")
	}

	// Construct the final URL path
	suffixTrimmedBaseUrl := strings.TrimSuffix(c.BaseUrl, "/")
	suffixTrimmedPath := strings.TrimSuffix(endpoint.Path, "/")
	finalPath := fmt.Sprintf("%s/%s", suffixTrimmedBaseUrl, suffixTrimmedPath)

	if request.Path != nil {
		finalPath = c.path(request.Path, finalPath) // Replace path parameters
	}

	if request.Query != nil {
		encodedQueryParams := c.query(request.Query) // Encode query parameters
		finalPath += fmt.Sprintf("?%s", encodedQueryParams)
	}

	if request.Body != nil {
		bodyBuffer = c.body(request.Body) // Marshal request body
		req, err = http.NewRequest(endpoint.Method, finalPath, bodyBuffer)
		if err != nil {
			return nil, nil, err
		}
	} else {
		req, err = http.NewRequest(endpoint.Method, finalPath, nil)
		if err != nil {
			return nil, nil, err
		}
	}

	if request.Header != nil {
		req.Header = c.header(request.Header) // Set request headers
	}

	return c.createHttpClient(endpoint), req, nil // Return the HTTP client and request
}
