module github.com/a2-ito/go-echo-onion-sample

go 1.19

require github.com/labstack/echo v3.3.10+incompatible

require interactor v0.0.0
require router v0.0.0
require handler v0.0.0
require repository v0.0.0

require (
	github.com/jinzhu/gorm v1.9.16 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/labstack/gommon v0.3.1 // indirect
	github.com/mattn/go-colorable v0.1.11 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/nanamen/go-echo-rest-sample v0.0.0-20180902044854-6c7e89a4fa2a // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.1 // indirect
	golang.org/x/crypto v0.0.0-20211117183948-ae814b36b871 // indirect
	golang.org/x/net v0.0.0-20211112202133-69e39bad7dc2 // indirect
	golang.org/x/sys v0.0.0-20211103235746-7861aae1554b // indirect
	golang.org/x/text v0.3.6 // indirect
)

replace interactor => ./interactor
replace router => ./presentation/http/echo/router
replace handler => ./presentation/http/echo/handler
replace repository => ./domain/repository
