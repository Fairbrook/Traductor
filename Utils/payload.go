package Utils

type Payload struct {
	Data   [][3]string `json:"data"`
	Errors []error     `json:"errors"`
}

func ErrorsToArray(errors []error) []string {
	res := []string{}
	for _, e := range errors {
		res = append(res, e.Error())
	}
	return res
}
