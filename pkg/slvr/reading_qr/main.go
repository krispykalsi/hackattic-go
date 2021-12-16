package reading_qr

import (
	"crypto/tls"
	"github.com/krispykalsi/hackattic/pkg/slvr/utils"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"image"
	"image/png"
	"io"
	"log"
	"net/http"
)

type readingQr struct{}

func New() *readingQr {
	return &readingQr{}
}

func (r readingQr) Solve(data []byte) []byte {
	p := &problem{}
	utils.FromJson(data, p)

	img := downloadImage(p.ImageUrl)
	code := readQr(img)

	s := solution{Code: code}
	return utils.ToJson(s)
}

func downloadImage(url string) image.Image {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Couldn't download img: %v", err)
	}
	defer func(r io.ReadCloser) {
		closeErr := r.Close()
		if closeErr != nil {
			log.Printf("Couldn't close response body: %v", err)
		}
	}(resp.Body)

	img, err := png.Decode(resp.Body)
	if err != nil {
		log.Fatalf("Couldn't decode img: %v", err)
	}

	return img
}

func readQr(img image.Image) string {
	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		log.Fatalf("Couldn't img to binary bitmap: %v", err)
	}
	qrReader := qrcode.NewQRCodeReader()
	result, err := qrReader.Decode(bmp, nil)
	if err != nil {
		log.Fatalf("Couldn't parse qr code: %v", err)
	}
	return result.GetText()
}
