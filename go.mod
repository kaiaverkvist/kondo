module github.com/kaiaverkvist/kondo

go 1.19

require (
	github.com/kelindar/binary v1.0.17
	github.com/labstack/echo/v4 v4.9.0
	github.com/labstack/gommon v0.4.0
	github.com/olahol/melody v1.1.1
	github.com/stretchr/testify v1.8.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/mattn/go-colorable v0.1.11 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.1 // indirect
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5 // indirect
	golang.org/x/net v0.0.0-20211015210444-4f30a5c0130f // indirect
	golang.org/x/sys v0.0.0-20211103235746-7861aae1554b // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/time v0.0.0-20201208040808-7e3f01d25324 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// This exists until the pull request to Melody merges.
replace github.com/olahol/melody v1.1.1 => github.com/kaiaverkvist/melody v0.0.0-20221001162103-ff254fae60c1
