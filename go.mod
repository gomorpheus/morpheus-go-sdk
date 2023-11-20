module github.com/gomorpheus/morpheus-go-sdk

go 1.17

require github.com/go-resty/resty/v2 v2.10.0

require golang.org/x/net v0.17.0 // indirect

retract v0.1.6

// voodoo
//replace github.com/go-resty/resty => gopkg.in/resty.v1 v1.12.0
