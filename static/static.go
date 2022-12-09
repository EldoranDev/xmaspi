package static

import "embed"

//go:embed *
var Files embed.FS

//go:embed index.html
var Index string
