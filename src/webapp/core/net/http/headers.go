package http

const (
	// HeaderContentType represents a http content-type header, it's value is supposed to be a mime type
	HeaderContentType = "Content-Type"

	// HeaderTransferEncoding represents a http transfer-encoding header.
	HeaderTransferEncoding = "Transfer-Encoding"

	// HeaderAccept the Accept header
	HeaderAccept = "Accept"

	charsetKey = "charset"

	// DefaultMime the default fallback mime type
	DefaultMime = "application/octet-stream"
	// JSONMime the json mime type
	JSONMime = "application/json"
	// YAMLMime the yaml mime type
	YAMLMime = "application/x-yaml"
	// XMLMime the xml mime type
	XMLMime = "application/xml"
	// TextMime the text mime type
	TextMime = "text/plain"
	// HTMLMime the html mime type
	HTMLMime = "text/html"
	// MultipartFormMime the multipart form mime type
	MultipartFormMime = "multipart/form-data"
	// URLencodedFormMime the url encoded form mime type
	URLencodedFormMime = "application/x-www-form-urlencoded"
)
