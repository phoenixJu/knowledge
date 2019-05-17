package main

import (
	"encoding/base64"
	"fmt"
	"personal/utils"
)

func main() {
	data := `[{"ts":7.51,"te":10.33,"s":"I had a dream","ws":null},{"ts":10.33,"te":13.87,"s":"That you were calling me","ws":null},{"ts":13.87,"te":16.95,"s":"Caught in between","ws":null},{"ts":16.95,"te":20.52,"s":"And floating down stream","ws":null},{"ts":20.52,"te":33.62,"s":"I won't drop you if you won't drop me","ws":null},{"ts":33.62,"te":36.98,"s":"I had a dream","ws":null},{"ts":36.98,"te":40.19,"s":"Where I became a fiend","ws":null},{"ts":40.19,"te":43.4,"s":"the villain was me","ws":null},{"ts":43.4,"te":46.73,"s":"But how can I be redeemed","ws":null},{"ts":46.73,"te":59.86,"s":"I won't drop you if you won't drop me","ws":null},{"ts":59.86,"te":62.16,"s":"I won't leave you downtown","ws":null},{"ts":62.16,"te":63.89,"s":"I won't leave you down","ws":null},{"ts":63.89,"te":65.42,"s":"If you say so","ws":null},{"ts":65.42,"te":73.38,"s":"I can take you home","ws":null},{"ts":73.38,"te":76.58,"s":"I could not wake","ws":null},{"ts":76.58,"te":79.87,"s":"This dream would be my fate","ws":null},{"ts":79.87,"te":83.13,"s":"Floating away","ws":null},{"ts":83.13,"te":113.63,"s":"but you would never stay","ws":null},{"ts":113.63,"te":116.64,"s":"There's a coffin with my name dear","ws":null},{"ts":116.64,"te":119.97,"s":"I could see it beckoning","ws":null},{"ts":119.97,"te":123.35,"s":"What's my name here","ws":null},{"ts":123.35,"te":126.89,"s":"When I'm living out a dream","ws":null},{"ts":126.89,"te":129.88,"s":"There's a coffin with my name here","ws":null},{"ts":129.88,"te":133.23,"s":"I could see it beckoning","ws":null},{"ts":133.23,"te":143.86,"s":"What's my name here","ws":null},{"ts":143.86,"te":145.1,"s":"I won't leave you downtown","ws":null},{"ts":145.1,"te":184.87,"s":"I won't leave you down","ws":null},{"ts":184.87,"te":190,"s":"","ws":null}]`
	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)
	sDec, _ := base64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()
	uEnc := utils.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := base64.StdEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
