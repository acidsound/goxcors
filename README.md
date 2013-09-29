# goxcors
CORS/JSONP proxy server to avoid Same-Origin-Policy on browser

## how to run
	git clone http://github.com/acidsound/goxcors && cd goxcors
	export GOPATH=`pwd`
	mkdir -p github.com/gorilla
	cd github.com/gorilla
	git clone git://github.com/gorilla/mux.git	
	git clone git://github.com/gorilla/context.git	
	cd ../..
	dev_appserver.py .

## example
	curl "http://localhost:8080/jsonp?method=POST&callback=call&url=http%3A%2F%2Ftranslate.google.com%2Ftranslate_a%2Ft%3Fclient%3Dx%26sl%3D%26tl%3Den%26text%3D%25EC%2597%25AC%25EB%259F%25AC%25EB%25B6%2584%25EC%259D%25B4%2520%25EB%25AA%25B0%25EB%259E%2590%25EB%258D%2598%2520%25EA%25B5%25AC%25EA%25B8%2580%2520%25EB%25B2%2588%25EC%2597%25AD%25EA%25B8%25B0"

## usage
* by CORS
	http://localhost:8080/cors?method=GET&url=...
* by JSONP (for IE)
	http://localhost:8080/jsonp?method=POST&callback=call&url=...
