module github.com/gomorpheus/morpheus-go-sdk

go 1.17

require github.com/go-resty/resty/v2 v2.7.0

require golang.org/x/net v0.0.0-20211029224645-99673261e6eb // indirect

retract v0.1.6

// voodoo
//replace github.com/go-resty/resty => gopkg.in/resty.v1 v1.12.0
