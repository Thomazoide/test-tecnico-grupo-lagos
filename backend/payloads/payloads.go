package payloads

type ResponsePayload struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Error      bool   `json:"error"`
	Data       any    `json:"data,omitempty"`
}

type ProductsPayload struct {
	Count      int       `json:"count"`
	Page       int       `json:"page"`
	Page_count int       `json:"page_count"`
	Page_size  int       `json:"page_size"`
	Products   []Product `json:"products"`
	Skip       int       `json:"skip"`
}

type ProductByCodePayload struct {
	Status         int     `json:"status"`
	Status_verbose string  `json:"status_verbose"`
	Code           string  `json:"code"`
	Product        Product `json:"product"`
}

type Product struct {
	Brands           string `json:"brands"`
	Code             string `json:"code"`
	Image_url        string `json:"image_url"`
	Nutriscore_grade string `json:"nutriscore_grade"`
	Product_name     string `json:"product_name"`
}
