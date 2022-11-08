package requests

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type OTPRequestBody struct {
	CSRF   string `json:"_csrf"`
	Mobile string `json:"mobile"`
}

type OTPVerifyRequestBody struct {
	CSRF string `json:"_csrf"`
	OTP  string `json:"otp"`
}

type OTPRequestRespone struct {
	StatusCode int    `json:"statusCode"`
	TID        string `json:"tid"`
	SID        string `json:"sid"`
}

const Cookie = "__SW=_kewALH-vYeYRQDfwlbXjl5rh0YGC9BB; _device_id=7ce6899b-c97c-7123-3e1b-c823836afa0d; WZRK_G=c9452692805b4e4c888aa191a72bb18b; _gcl_au=1.1.121360058.1664760648; _ga=GA1.2.918831839.1664760649; userLocation=%7B%22lat%22%3A%2222.57215%22%2C%22lng%22%3A%2288.411976%22%2C%22address%22%3A%22IA%20Block%2C%20Sector%20III%2C%20Salt%20Lake%20City%2C%20Kolkata%2C%20West%20Bengal%20700106%2C%20India%22%2C%22area%22%3A%22Salt%20Lake%20City%22%2C%22id%22%3A%2210498134%22%7D; swgy_logout_clear=1; _guest_tid=8a6de086-91fd-47f6-9084-3f1b424cf718; _sid=3kt7e0d7-79f5-43ca-b18b-23edd1af6cf9; fontsLoaded=1; WZRK_S_W86-ZZK-WR6Z=%7B%22p%22%3A1%2C%22s%22%3A1667864991%2C%22t%22%3A1667864990%7D; _gid=GA1.2.771349870.1667864992; _gat_0=1"

func SendOTP(number string) (tid string, sid string, err error) {
	client := resty.New()

	if len(number) != 10 {
		return "", "", fmt.Errorf("invalid number")
	}

	//Create the request body from the user entered number
	OTPBody := OTPRequestBody{
		CSRF:   "dL7ppcABBhuz-cQQ1ul8xPGl3sg6ERkpj6ao5vcs",
		Mobile: number,
	}

	var OTPResponse OTPRequestRespone

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:106.0) Gecko/20100101 Firefox/106.0").
		SetHeader("Cookie", Cookie).
		SetBody(&OTPBody).
		SetResult(&OTPResponse).
		Post("https://www.swiggy.com/dapi/auth/sms-otp")

	if err != nil || resp.StatusCode() != 200 || OTPResponse.StatusCode == 1 {

		return "", "", fmt.Errorf("failed to get OTP")
	}

	return OTPResponse.TID, OTPResponse.SID, nil

}

func VerifyOTP(tid string, sid string, otp string) (string, error) {
	client := resty.New()

	if len(otp) != 6 {
		return "", fmt.Errorf("invalid otp")
	}

	//Create the request body from the user entered OTP
	OTPBody := OTPVerifyRequestBody{
		CSRF: "dL7ppcABBhuz-cQQ1ul8xPGl3sg6ERkpj6ao5vcs",
		OTP:  otp,
	}

	var OTPResponse OTPRequestRespone

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:106.0) Gecko/20100101 Firefox/106.0").
		//The guest uuid helps identify which number this is the OTP for
		SetHeader("Cookie", "__SW=_kewALH-vYeYRQDfwlbXjl5rh0YGC9BB;_guest_tid="+tid+";_sid="+sid).
		SetBody(&OTPBody).
		SetResult(&OTPResponse).
		Post("https://www.swiggy.com/dapi/auth/otp-verify")

	if err != nil || resp.StatusCode() != 200 || OTPResponse.StatusCode == 1 {

		return "", fmt.Errorf("failed to verify OTP")
	}

	//Find the _session_tid cookie from the response cookies.
	respCookies := resp.Cookies()

	for _, cookie := range respCookies {
		if cookie.Name == "_session_tid" {
			return cookie.Value, nil
		}
	}

	return "", fmt.Errorf("OTP verified but authenticated failed")

}
