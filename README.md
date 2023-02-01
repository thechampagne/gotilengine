# gotilengine

[![](https://img.shields.io/github/v/tag/thechampagne/gotilengine?label=version)](https://github.com/thechampagne/gotilengine/releases/latest) [![](https://img.shields.io/github/license/thechampagne/gotilengine)](https://github.com/thechampagne/gotilengine/blob/main/LICENSE)

Go binding for **Tilengine** a 2D graphics engine with raster effects for retro/classic style game development.

### TODOs
- [ ] Raw interface without C types
- [ ] Wrapper

### Example
```go
package main

// #include <stdlib.h>
import "C"
import (
	"unsafe"
	"github.com/thechampagne/gotilengine"
)

func main() {
	var foreground gotilengine.TLN_Tilemap

	tmx := (gotilengine.CString) (C.CString("assets/sonic/Sonic_md_fg1.tmx"))
	defer C.free(unsafe.Pointer(tmx))

	gotilengine.TLN_Init(400, 240, 1, 0, 0)
	foreground = gotilengine.TLN_LoadTilemap(tmx, (gotilengine.CString) (C.NULL))
	gotilengine.TLN_SetLayerTilemap(0, foreground)

	gotilengine.TLN_CreateWindow((gotilengine.CString) (C.NULL), 0)
	for gotilengine.TLN_ProcessWindow() != 0 {
		gotilengine.TLN_DrawFrame(0)
	}

	gotilengine.TLN_DeleteTilemap(foreground)
	gotilengine.TLN_Deinit()
}
```

### References
 - [Tilengine](https://github.com/megamarc/Tilengine)

### License

This repo is released under the [MPL-2.0](https://github.com/thechampagne/gotilengine/blob/main/LICENSE).
