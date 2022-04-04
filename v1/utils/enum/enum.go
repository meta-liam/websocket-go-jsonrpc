package enum

// device
const Auto string = "auto"
const PC string = "pc"
const APP string = "app"
const IOS string = "ios"
const Android string = "android"

const JsonrpcVersion string = "2.0"

const (
	Jsonrpc         = "jsonrpc"
	Jsonrpc_method  = "method"
	Jsonrpc_params  = "params"
	Jsonrpc_id      = "id"
	Jsonrpc_service = "service"
	Jsonrpc_result  = "result"
	Jsonrpc_error   = "error"
	Jsonrpc_type    = "type"
	Jsonrpc_client  = "client"
)

// JSON-RPC 错误代码
const (
	// Invalid JSON was received by the server.
	// An error occurred on the server while parsing the JSON text.
	ParseError int = -32700
	// The JSON sent is not a valid Request object.
	InvalidRequest int = -32600
	// The method does not exist / is not available.
	MethodNotFound int = -32601
	// Invalid method parameter(s).
	InvalidParams int = -32602
	// Internal JSON-RPC error.
	InternalError int = -32603
	// Reserved for implementation-defined server-errors.
	ServerError int = -32099
)
