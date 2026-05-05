package assets

import "embed"

// Static embeds the entire static directory (CSS, images, etc.)
//go:embed all:static
var Static embed.FS
