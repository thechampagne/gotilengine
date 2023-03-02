# gotilengine

[![](https://img.shields.io/github/v/tag/thechampagne/gotilengine?label=version)](https://github.com/thechampagne/gotilengine/releases/latest) [![](https://img.shields.io/github/license/thechampagne/gotilengine)](https://github.com/thechampagne/gotilengine/blob/main/LICENSE)

Go binding for **Tilengine** a 2D graphics engine with raster effects for retro/classic style game development.

### TODOs
- [ ] Raw interface without C types
- [ ] Wrapper

### Example
```go
package main

/*
#include <stdio.h>
#cgo LDFLAGS:-L./Tilengine -lTilengine
*/
// void MyCallback(int line);
import "C"
import (
	// "runtime/cgo"
	tln "github.com/system64MC/gotilengine"
)

//export MyCallback
func MyCallback(frame C.int) {
	tln.SetLayerPosition(0, int(frame), 0)
}

func main() {
	var a = "Initial"
	println(a)

	tln.Init(400, 240, 1, 0, 0)
	tln.CreateWindow("", 0)
	// var h = Tilengine.Affine{}
	// println(h.dx)
	tilemap := tln.LoadTilemap("Assets/Sonic/Sonic_md_fg1.tmx", "")
	tln.SetWindowTitle("Let's Go Tilengine!")
	tln.SetLayerTilemap(0, tilemap)
	tln.SetRasterCallback(tln.VideoCallback(C.MyCallback))

	for tln.ProcessWindow() {
		tln.DrawFrame(0)
	}

}

```

### References
 - [Tilengine](https://github.com/megamarc/Tilengine)

### License

This repo is released under the [MPL-2.0](https://github.com/thechampagne/gotilengine/blob/main/LICENSE).
