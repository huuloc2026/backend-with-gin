package response

const (
	// General
	ErrorCodeSuccess       = 2000
	ErrorCodeBadRequest    = 4000
	ErrorCodeUnauthorized  = 4010
	ErrorCodeForbidden     = 4030
	ErrorCodeNotFound      = 4040
	ErrorCodeConflict      = 4090
	ErrorCodeInternalError = 5000

	// Auth Errors
	ErrorCodeInvalidToken     = 4011
	ErrorCodeExpiredToken     = 4012
	ErrorCodeInvalidLogin     = 4013
	ErrorCodePermissionDenied = 4031

	// User Errors
	ErrorCodeUserNotFound      = 4041
	ErrorCodeUserAlreadyExists = 4091
	ErrorCodeInvalidUserInput  = 4001

	// Product Errors
	ErrorCodeProductNotFound      = 4042
	ErrorCodeProductOutOfStock    = 4002
	ErrorCodeInvalidProductInput  = 4003
	ErrorCodeProductAlreadyExists = 4092
)

// message
var msg = map[int]string{
	// General Messages
	ErrorCodeSuccess:       "Success",
	ErrorCodeBadRequest:    "Bad Request",
	ErrorCodeUnauthorized:  "Unauthorized",
	ErrorCodeForbidden:     "Forbidden",
	ErrorCodeNotFound:      "Resource Not Found",
	ErrorCodeConflict:      "Conflict",
	ErrorCodeInternalError: "Internal Server Error",

	// Auth Messages
	ErrorCodeInvalidToken:     "Invalid authentication token",
	ErrorCodeExpiredToken:     "Token has expired",
	ErrorCodeInvalidLogin:     "Invalid username or password",
	ErrorCodePermissionDenied: "You do not have permission to access this resource",

	// User Messages
	ErrorCodeUserNotFound:      "User not found",
	ErrorCodeUserAlreadyExists: "User already exists",
	ErrorCodeInvalidUserInput:  "Invalid user input",

	// Product Messages
	ErrorCodeProductNotFound:      "Product not found",
	ErrorCodeProductOutOfStock:    "Product is out of stock",
	ErrorCodeInvalidProductInput:  "Invalid product input",
	ErrorCodeProductAlreadyExists: "Product already exists",
}
