package templates

import _ "embed"

//go:embed index.html
var IndexHTML string

//go:embed packages.html
var Packages string

//go:embed result.html
var Result string
