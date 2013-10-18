# goxcors
CORS/JSONP proxy server to avoid Same-Origin-Policy on browser

## how to run
	git clone http://github.com/acidsound/goxcors && cd goxcors
	export GOPATH=`pwd`
	dev_appserver.py .

## example
	curl "http://localhost:8080/jsonp?method=POST&callback=call&url=http%3A%2F%2Ftranslate.google.com%2Ftranslate_a%2Ft%3Fclient%3Dx%26sl%3D%26tl%3Den%26text%3D%25EC%2597%25AC%25EB%259F%25AC%25EB%25B6%2584%25EC%259D%25B4%2520%25EB%25AA%25B0%25EB%259E%2590%25EB%258D%2598%2520%25EA%25B5%25AC%25EA%25B8%2580%2520%25EB%25B2%2588%25EC%2597%25AD%25EA%25B8%25B0"

	curl -d '{"device":"Mozilla","options":"enable_pre_space","requests":[{"writing_guide":{"writing_area_width":360,"writing_area_height":567},"ink":[[[98,118,141,192,200],[159,149,139,138,190],[0,1,2,3,4]],[[89,112,133,236],[254,248,242,217],[5,6,7,8]],[[202],[230],[9]],[[202,198,187,189,194],[230,248,322,384,424],[10,11,12,13,14]]],"language":"ko"}]}' "http://localhost:8080/cors?method=POST&header=Content-Type|application/json&url=https%3A%2F%2Fwww.google.com%2Finputtools%2Frequest%3Fime%3Dhandwriting%26app%3Dmobilesearch%26cs%3D1%26oe%3DUTF-8"

## usage
* by CORS
	http://localhost:8080/cors?method=GET&url=...
* by JSONP (for IE)
	http://localhost:8080/jsonp?method=POST&callback=call&url=...
* with header (seperated by pipe(|) chars)
	http://localhost:8080/cors?header=Content-type|application/json&header=....

## demo
http://goxcors.appspot.com
