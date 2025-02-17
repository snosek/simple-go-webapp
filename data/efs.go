package data

import (
	"embed"
)

//go:embed "products.json"
var Data embed.FS
