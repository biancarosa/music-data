package routes

const (
	validationMessage string = "Missing required parameters."
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

func (err validationError) BuildResponse() validationResponse {
	var resp validationResponse
	resp.Body.Message = validationMessage
	resp.Body.QueryParameters = err.queryParameters
	return resp
}
