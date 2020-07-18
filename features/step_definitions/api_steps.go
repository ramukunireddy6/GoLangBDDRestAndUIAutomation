package step_definitions

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func SendRequest(request string) *http.Response {
	readEndpoint := strings.HasPrefix(request, "ENV")
	readENV := strings.TrimPrefix(request, "ENV:")

	if readEndpoint {
		request = os.Getenv(readENV)
	}

	req, _ := http.Get(os.Getenv("BASE_URL") + request)

	return req
}

func ValidateResponse(response int, resp *http.Response) error {
	actualCode := resp.StatusCode
	if expectCode := (response); actualCode != expectCode {
		log.Fatalln("actual status code :", actualCode)
		log.Fatalln("expected status code :", expectCode)
	} else {
		fmt.Println("SCENARIO IS SUCCESS")
	}

	return nil
}
