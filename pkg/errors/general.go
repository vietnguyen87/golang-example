package errors

const (
	// General Errors: 0 -> -49
	// Success indicates no error
	Success ErrorType = 200
	// Unknown error indicates unknown state or step
	Unknown ErrorType = 0
	// BadRequest error
	BadRequestErr ErrorType = 400
	// NotFound error
	NotFound ErrorType = 404
	// AuthenFailed error
	AuthenticationFailed ErrorType = 401
	// Internal server error
	InternalServerError ErrorType = 500
	// Call Internal API Error
	CallInternalAPIError ErrorType = -7
)
