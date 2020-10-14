package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

var cookie = "sid=0dda3596-8222-4c73-b977-cdda08c4fa22; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1602402861; ec=6bdtZMWQ-1602407956215-5da14262279e2-1545715071; FSSBBIl1UgzbN7NO=5mJKnOQ4eOUuhBFZywRPJR24_wMCW3XuT7lJhibSPrySiJsHZgorJ3JM_VISPeH7EsG6dNcFWp2rTdVBTiXT_xA; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1602642247; FSSBBIl1UgzbN7NP=5UgRf1m5D3V3qqqmTxiaxFG4pEtTgqCgs7C5S_6Re8ULI18.EsJs4VpS9.90ZoF7Jyv.jovGswIBWO8CJjJmTgshF3iwP_AxhpEkAhKbHwYZwCChXgkObinmWOQ3U6Rw0Jiw9mmJKD5TEvKkWpwWypz1ze9Dw7cPbN0UlzDBq7NbPnOLUvvv6QVKZB.v6Z_xzOIa2jQ.yiMXshmqDRShDeoiGQwQpNWDheIfsGVPK8Fb9D8oGXyMv6MEvliSchfixRE_83RV0nZpX1N_w4rGHbO; _efmdata=NbBCqZ0ViOqJD5nY1wqWA%2FqRhFVJkAPsHA3ZyIbn2eZAGh1t%2F4gH9eq2%2Bl8%2F0ioehh5ky77O2lyba%2Bpzlqb4A1rL%2BgIRug9l5bLyd9gx8vU%3D; _exid=YKvLc10gGVUs89fagceHuikNAqM2TMkGMa9ejg8iJFqKISvahXAf84F%2BiJekZw5muLeQ6%2FTFeRDbf12qNxywhw%3D%3D"
var rateLimiter = time.Tick(time.Millisecond * 100)
//Fetch is a function to get the content encoding with UTF8 of the target URL.
func Fetch(url string) ([]byte, error) {
	<- rateLimiter
	client := &http.Client{}

	url = strings.Replace(url, "http://", "https://", 1)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Fetch: fail to new http request, %v\n", err)
	}
	//"cookie"不能是"Cookie"
	req.Header.Add("cookie", cookie)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
	//req.Header.Add("Referer", url)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		resp, err = client.Do(req)
		if err != nil {
			panic(err)
		}
		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("Error: status code %v", resp.StatusCode)
		}
	}
	log.Printf("resp header = %v", resp.Header)
	log.Printf("cookie = %s", resp.Cookies())
	//cookie = strings.Join(resp.Cookies()[0].Unparsed, " ")
	//推断编码方式
	bodyReader := bufio.NewReader(resp.Body)
	e := determinEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())

	return ioutil.ReadAll(utf8Reader)
}

func determinEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("fetcher error : %v.\n", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
