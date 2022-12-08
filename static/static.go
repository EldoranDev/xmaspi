package static

import "embed"

//go:embed *.html
var Files embed.FS

//go:embed index.html
var Index string
