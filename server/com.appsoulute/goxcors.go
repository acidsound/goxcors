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
	r.HandleFunc("/post", getCrossDomainRequest)
	r.HandleFunc("/postp", getJSONPRequest)
	http.Handle("/", r)
}

/* curl -s -A "Mozilla/5.0 (X11; Linux x86_64; rv:5.0) Gecko/20100101 Firefox/5.0" "http://translate.google.com/translate_a/t?client=p&sl=&tl=ko&text=compensation" */
func getCrossDomainRequest(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	client := urlfetch.Client(c)

	c.Infof("Request Type: CORS")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With")

	reqURL := r.URL.Query()["url"][0]
	c.Infof("reqURL : %s", reqURL)
	req, err := http.NewRequest("POST", reqURL, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:5.0) Gecko/20100101 Firefox/5.0)")

	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err :=ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Errorf("ioutil error Get %s", err)
		fmt.Fprintf(w, "{'err':'%s'}", err)
		return
	}
	c.Infof("%s", body)
	fmt.Fprintf(w, "%s", body)
}

func getJSONPRequest(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	client := urlfetch.Client(c)

	c.Infof("Request Type: JSONP")
	reqURL := r.URL.Query()["url"][0]
	reqCallBack := r.URL.Query()["callback"][0]
	c.Infof("reqURL : %s", reqURL)
	req, err := http.NewRequest("POST", reqURL, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:5.0) Gecko/20100101 Firefox/5.0)")

	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err :=ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Errorf("ioutil error Get %s", err)
		fmt.Fprintf(w, "{'err':'%s'}", err)
		return
	}
	c.Infof("%s", body)
	fmt.Fprintf(w, "%s(%s)", reqCallBack, body)
}