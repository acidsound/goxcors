package main

import (
	"appengine"
	"appengine/urlfetch"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Context interface {
	// Debugf formats its arguments according to the format, analogous to fmt.Printf,
	// and records the text as a log message at Debug level.
	Debugf(format string, args ...interface{})

	// Infof is like Debugf, but at Info level.
	Infof(format string, args ...interface{})

	// Warningf is like Debugf, but at Warning level.
	Warningf(format string, args ...interface{})

	// Errorf is like Debugf, but at Error level.
	Errorf(format string, args ...interface{})

	// Criticalf is like Debugf, but at Critical level.
	Criticalf(format string, args ...interface{})
}

func init() {
	http.HandleFunc("/cors", getCrossDomainRequest)
	http.HandleFunc("/jsonp", getJSONPRequest)
}

func runProxy(client *http.Client, r *http.Request, c Context) string {
	query := r.URL.Query()
	var req *http.Request
	var err error
	req, err = http.NewRequest(query["method"][0], query["url"][0], r.Body)
	req.Header.Add("user-agent", "Mozilla/5.0 (X11; Linux x86_64; rv:5.0) Gecko/20100101 Firefox/5.0)")
	for _, v := range query["header"] {
		kv := strings.Split(v, "|")
		if len(kv) < 2 {
			c.Errorf("%s : malformed header. headers must be seperated by the string \"|\"", v)
			return fmt.Sprintf("{'err':'%s : malformed header. headers must be seperated by the string \"|\"'}", v)
		}
		if strings.ToLower(kv[0]) == "user-agent" {
			c.Infof("set user-agent to %s", kv[1])
			req.Header.Del("user-agent")
		}
		req.Header.Add(kv[0], kv[1])
	}

	resp, err := client.Do(req)
	if err != nil {
		c.Errorf("Response Error : %q ", err)
		return fmt.Sprintf("{'err':'%q'}", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Errorf("Response Error : %q ", err)
		return fmt.Sprintf("{'err':'%q'}", err)
	}
	c.Infof("%s", body)
	return fmt.Sprintf("%s", body)
}

/* curl -s -A "Mozilla/5.0 (X11; Linux x86_64; rv:5.0) Gecko/20100101 Firefox/5.0" "http://translate.google.com/translate_a/t?client=p&sl=&tl=ko&text=compensation"
 * ::TEST
 * curl "http://localhost:8080/jsonp?method=POST&callback=call&url=http%3A%2F%2Ftranslate.google.com%2Ftranslate_a%2Ft%3Fclient%3Dx%26sl%3D%26tl%3Den%26text%3D%25EC%2597%25AC%25EB%259F%25AC%25EB%25B6%2584%25EC%259D%25B4%2520%25EB%25AA%25B0%25EB%259E%2590%25EB%258D%2598%2520%25EA%25B5%25AC%25EA%25B8%2580%2520%25EB%25B2%2588%25EC%2597%25AD%25EA%25B8%25B0"
 * curl -d '{"device":"Mozilla","options":"enable_pre_space","requests":[{"writing_guide":{"writing_area_width":360,"writing_area_height":567},"ink":[[[98,118,141,192,200],[159,149,139,138,190],[0,1,2,3,4]],[[89,112,133,236],[254,248,242,217],[5,6,7,8]],[[202],[230],[9]],[[202,198,187,189,194],[230,248,322,384,424],[10,11,12,13,14]]],"language":"ko"}]}' "http://localhost:8080/cors?method=POST&header=Content-Type|application/json&url=https%3A%2F%2Fwww.google.com%2Finputtools%2Frequest%3Fime%3Dhandwriting%26app%3Dmobilesearch%26cs%3D1%26oe%3DUTF-8"
 * http://www.google.com/inputtools/request?ime=transliteration_en_ru&num=5&cp=0&cs=0&ie=utf-8&oe=utf-8&text=prosto
[*/
func getCrossDomainRequest(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	client := urlfetch.Client(c)

	c.Infof("Request Type: CORS, method: %s", r.URL.Query()["method"][0])
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With")

	body := runProxy(client, r, c)
	fmt.Fprintf(w, "%s", body)
}

func getJSONPRequest(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	client := urlfetch.Client(c)

	c.Infof("Request Type: JSONP, method: %s", r.URL.Query()["method"][0])
	body := runProxy(client, r, c)
	fmt.Fprintf(w, "%s(%s)", r.URL.Query()["callback"][0], body)
}
