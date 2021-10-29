package Utils

type Payload struct {
	Data   [][3]string `json:"data"`
	Errors []error     `json:"errors"`
}
