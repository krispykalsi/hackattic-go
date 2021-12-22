package server

type jwtPayload struct {
	Append string `json:"append,omitempty"`
	Exp    int64  `json:"exp,omitempty"`
	Nbf    int64  `json:"nbf,omitempty"`
}
