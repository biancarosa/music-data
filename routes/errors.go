package routes

const (
	validationMessage   string = "Missing required parameters."
	internalServerError string = "Internal server error."
)

type validationError struct {
	queryParameters []string
}

type validationResponse struct {
	Body struct {
		Message         string   `json:"message"`
		QueryParameters []string `json:"queryParameters"`
	}
}

type errorResponse struct {
	Body struct {
		Message string `json:"message"`
	}
}

func (err validationError) BuildResponse() validationResponse {
	var resp validationResponse
	resp.Body.Message = validationMessage
	resp.Body.QueryParameters = err.queryParameters
	return resp
}
