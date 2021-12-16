package slvr

type problem struct {
	ImageUrl string `json:"image_url"`
}

type solution struct {
	Code string `json:"code"`
}

type readingQr struct{}

func (r readingQr) Solve(data []byte) []byte {
	p := &problem{}
	fromJson(data, p)
	code := p.ImageUrl
	s := solution{Code: code}
	return toJson(s)
}
