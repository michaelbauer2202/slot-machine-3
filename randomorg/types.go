package randomorg

type request struct {
	Key string `json:"apiKey"`
	N   int    `json:"n"`
	Min int    `json:"min"`
	Max int    `json:"max"`
}

type Response struct {
	Res Result `json:"result"`
}

type Result struct {
	Rand random `json:"random"`
}

type random struct {
	Data []int `json:"data"`
}

func (resp *Response) Ints() []int {
	return resp.Res.Rand.Data
}
