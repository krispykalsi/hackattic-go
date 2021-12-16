package help_me_unpack

type solution struct {
	Int             int32   `json:"int"`
	UInt            uint    `json:"uint"`
	Short           int16   `json:"short"`
	Float           float64 `json:"float"`
	Double          float64 `json:"double"`
	BigEndianDouble float64 `json:"big_endian_double"`
}
