package goxcors

import (
	"appengine"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"appengine/urlfetch"
	"io/ioutil"
)

func init() {
	r := mux.NewRouter()
	r.HandleFunc("/cors", getCrossDomainRequest)
	r.HandleFunc("/jsonp", getJSONPRequest)
	http.Handle("/", r)
}

func runProxy(method, reqURL string, client *http.Client) string {
	req, err := http.NewRequest(method, reqURL, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:5.0) Gecko/20100101 Firefox/5.0)")

	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err :=ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Sprintf("{'err':'%q'}", err)
	}
	return fmt.Sprintf("%s",body)
}

/* curl -s -A "Mozilla/5.0 (X11; Linux x86_64; rv:5.0) Gecko/20100101 Firefox/5.0" "http://translate.google.com/translate_a/t?client=p&sl=&tl=ko&text=compensation" */
func getCrossDomainRequest(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	client := urlfetch.Client(c)

	c.Infof("Request Type: CORS, method: %s", r.URL.Query()["method"][0])
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With")

	reqURL := r.URL.Query()["url"][0]
	c.Infof("reqURL : %s", reqURL)
	body := runProxy(r.URL.Query()["method"][0], reqURL, client)
	c.Infof("%s", body)
	fmt.Fprintf(w, "%s", body)
}

func getJSONPRequest(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	client := urlfetch.Client(c)

	c.Infof("Request Type: JSONP, method: %s", r.URL.Query()["method"][0])

	reqURL := r.URL.Query()["url"][0]
	c.Infof("reqURL : %s", r.URL.Query()["url"][0])
	body := runProxy(r.URL.Query()["method"][0], reqURL, client)
	c.Infof("%s", body)
	fmt.Fprintf(w, "%s(%s)", r.URL.Query()["callback"][0], body)
}