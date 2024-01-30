package gotilengine

/*
#cgo CFLAGS: -I${SRCDIR}
#cgo LDFLAGS: -lTilengine
#include "Tilengine.h"
void cRasterCallback(int line);
*/
import "C"
import (
	"runtime"
	"unsafe"
	// "runtime/cgo"
)

type (
	CString = *C.char
	CInt    = C.int
	CFloat  = C.float
	CBool   = C.bool
	CUint8  = C.uint8_t
	CUint16 = C.uint16_t
	CUint32 = C.uint32_t
)

const (
	CTrue  = C.true
	CFalse = C.false
)

const (
	TILENGINE_VER_MAJ        = 2
	TILENGINE_VER_MIN        = 13
	TILENGINE_VER_REV        = 2
	TILENGINE_HEADER_VERSION = ((TILENGINE_VER_MAJ << 16) | (TILENGINE_VER_MIN << 8) | TILENGINE_VER_REV)
)

type TileFlags = C.TLN_TileFlags

const (
	FLAG_NONE     TileFlags = 0
	FLAG_FLIPX              = 1 << 15
	FLAG_FLIPY              = 1 << 14
	FLAG_ROTATE             = 1 << 13
	FLAG_PRIORITY           = 1 << 12
	FLAG_MASKED             = 1 << 11
	FLAG_TILESET            = 7 << 8
	FLAG_PALETTE            = 7 << 5
)

type Blend = C.TLN_Blend

const (
	BLEND_NONE Blend = iota
	BLEND_MIX25
	BLEND_MIX50
	BLEND_MIX75
	BLEND_ADD
	BLEND_SUB
	BLEND_MOD
	BLEND_CUSTOM
	MAX_BLEND
	BLEND_MIX = BLEND_MIX50
)

type LayerType = C.TLN_LayerType

const (
	LAYER_NONE LayerType = iota
	LAYER_TILE
	LAYER_OBJECT
	LAYER_BITMAP
)

// typedef struct
// {
// 	float angle;
// 	float dx;
// 	float dy;
// 	float sx;
// 	float sy;
// }
// TLN_Affine;

// type Affine = C.TLN_Affine

type Affine struct {
	Angle float32
	Dx    float32
	Dy    float32
	Sx    float32
	Sy    float32
}

// typedef union Tile
// {
// 	uint32_t value;
// 	struct
// 	{
// 		uint16_t index;
// 		union
// 		{
// 			uint16_t flags;
// 			struct
// 			{
// 				uint8_t unused : 5;
// 				uint8_t palette : 3;
// 				uint8_t tileset : 3;
// 				bool masked : 1;
// 				bool priority : 1;
// 				bool rotated : 1;
// 				bool flipy : 1;
// 				bool flipx : 1;
// 			};
// 		};
// 	};
// }
// Tile;

type Tile uint32

// typedef struct
// {
// 	int index;
// 	int delay;
// }
// TLN_SequenceFrame;

// type SequenceFrame = C.TLN_SequenceFrame

type SequenceFrame struct {
	Index int32
	Delay int32
}

// typedef struct
// {
// 	int delay;
// 	uint8_t first;
// 	uint8_t count;
// 	uint8_t dir;
// }
// TLN_ColorStrip;

// type ColorStrip = C.TLN_ColorStrip

type ColorStrip struct {
	Delay int32
	First uint8
	Count uint8
	Dir   uint8
}

// typedef struct
// {
// 	char name[32];
// 	int num_frames;
// }
// TLN_SequenceInfo;

// type SequenceInfo = C.TLN_SequenceInfo

type SequenceInfo struct {
	name      [32]uint8
	NumFrames int32
}

// typedef struct
// {
// 	char name[64];
// 	int x;
// 	int y;
// 	int w;
// 	int h;
// }
// TLN_SpriteData;

// type SpriteData = C.TLN_SpriteData

type SpriteData struct {
	name [64]uint8
	X    int32
	Y    int32
	W    int32
	H    int32
}

// typedef struct
// {
// 	int w;
// 	int h;
// }
// TLN_SpriteInfo;

// type SpriteInfo = C.TLN_SpriteInfo

type SpriteInfo struct {
	W int32
	H int32
}

// typedef struct
// {
// 	uint16_t index;
// 	uint16_t flags;
// 	int row;
// 	int col;
// 	int xoffset;
// 	int yoffset;
// 	uint8_t color;
// 	uint8_t type;
// 	bool empty;
// }
// TLN_TileInfo;

// type TileInfo = C.TLN_TileInfo

type TileInfo struct {
	Index   uint16
	Flags   uint16
	Row     int32
	Col     int32
	Xoffset int32
	Yoffset int32
	Color   uint8
	Type    uint8
	Empty   uint8
}

// typedef struct
// {
// 	uint16_t id;
// 	uint16_t gid;
// 	uint16_t flags;
// 	int x;
// 	int y;
// 	int width;
// 	int height;
// 	uint8_t type;
// 	bool visible;
// 	char name[64];
// }
// TLN_ObjectInfo;

// type ObjectInfo = C.TLN_ObjectInfo

type ObjectInfo struct {
	Id      uint16
	Gid     uint16
	Flags   uint16
	X       int32
	Y       int32
	Width   int32
	Height  int32
	Type    uint8
	Visible uint8
	name    [64]uint8
}

// typedef struct
// {
// 	uint8_t	type;
// 	bool	priority;
// }
// TLN_TileAttributes;

// type TileAttributes = C.TLN_TileAttributes

type TileAttributes struct {
	Type     uint8
	Priority uint8
}

const (
	TLN_OVERLAY_NONE       = 0
	TLN_OVERLAY_SHADOWMASK = 0
	TLN_OVERLAY_APERTURE   = 0
	TLN_OVERLAY_SCANLINES  = 0
	TLN_OVERLAY_CUSTOM     = 0
)

type CRT = C.TLN_CRT

const (
	CRT_SLOT CRT = iota
	CRT_APERTURE
	CRT_SHADOW
)

// typedef struct
// {
// 	int16_t dx;
// 	int16_t dy;
// }
// TLN_PixelMap;

// type PixelMap = C.TLN_PixelMap

type PixelMap struct {
	dx int16
	dy int16
}

// typedef struct Engine*		 TLN_Engine;
type Engine struct {
	data C.TLN_Engine
}

// type Engine = C.TLN_Engine

// typedef union  Tile*		 TLN_Tile;
type TilePtr = *Tile

// type Tile = C.TLN_Tile

// typedef struct Tileset*		 TLN_Tileset;
type Tileset struct {
	data C.TLN_Tileset
}

// type Tileset = C.TLN_Tileset

// typedef struct Tilemap*		 TLN_Tilemap;
type Tilemap struct {
	data C.TLN_Tilemap
}

// type Tilemap = C.TLN_Tilemap

// typedef struct Palette*		 TLN_Palette;
type Palette struct {
	data C.TLN_Palette
}

// type Palette = C.TLN_Palette

// typedef struct Spriteset*	 TLN_Spriteset;
type Spriteset struct {
	data C.TLN_Spriteset
}

// type Spriteset = C.TLN_Spriteset

// typedef struct Sequence*	 TLN_Sequence;
type Sequence struct {
	data C.TLN_Sequence
}

// type Sequence = C.TLN_Sequence

// typedef struct SequencePack* TLN_SequencePack;
type SequencePack struct {
	data C.TLN_SequencePack
}

// type SequencePack = C.TLN_SequencePack

// typedef struct Bitmap*		 TLN_Bitmap;
type Bitmap struct {
	data C.TLN_Bitmap
}

// type Bitmap = C.TLN_Bitmap

// typedef struct ObjectList*	 TLN_ObjectList;
type ObjectList struct {
	data C.TLN_ObjectList
}

// type ObjectList = C.TLN_ObjectList

// typedef struct
// {
// 	TLN_Bitmap bitmap;
// 	uint16_t id;
// 	uint8_t	type;
// }
// TLN_TileImage;

// type TileImage = C.TLN_TileImage

type TileImage struct {
	Bitmap Bitmap
	Id     uint16
	Type   uint8
}

// typedef struct
// {
// 	int x;
// 	int y;
// 	int w;
// 	int h;
// 	uint32_t flags;
// 	TLN_Palette palette;
// 	TLN_Spriteset spriteset;
// 	int index;
// 	bool enabled;
// 	bool collision;
// }
// TLN_SpriteState;

// type SpriteState = C.TLN_SpriteState

type SpriteState struct {
	X         int32
	Y         int32
	W         int32
	H         int32
	Flags     uint32
	Palette   Palette
	SpriteSet Spriteset
	Index     int32
	Enabled   uint8
	Collision uint8
}

// typedef union SDL_Event SDL_Event;
type SDL_Event = C.SDL_Event

// typedef void(*TLN_VideoCallback)(int scanline);
type VideoCallback = C.TLN_VideoCallback

// typedef uint8_t(*TLN_BlendFunction)(uint8_t src, uint8_t dst);
type BlendFunction = C.TLN_BlendFunction

// typedef void(*TLN_SDLCallback)(SDL_Event*);
type SDLCallback = C.TLN_SDLCallback

type Player = C.TLN_Player

const (
	PLAYER1 Player = iota
	PLAYER2
	PLAYER3
	PLAYER4
)

type Input = C.TLN_Input

const (
	INPUT_NONE Input = iota
	INPUT_UP
	INPUT_DOWN
	INPUT_LEFT
	INPUT_RIGHT
	INPUT_BUTTON1
	INPUT_BUTTON2
	INPUT_BUTTON3
	INPUT_BUTTON4
	INPUT_BUTTON5
	INPUT_BUTTON6
	INPUT_START
	INPUT_QUIT
	INPUT_CRT

	INPUT_P1 = (PLAYER1 << 5)
	INPUT_P2 = (PLAYER2 << 5)
	INPUT_P3 = (PLAYER3 << 5)
	INPUT_P4 = (PLAYER4 << 5)

	INPUT_A = INPUT_BUTTON1
	INPUT_B = INPUT_BUTTON2
	INPUT_C = INPUT_BUTTON3
	INPUT_D = INPUT_BUTTON4
	INPUT_E = INPUT_BUTTON5
	INPUT_F = INPUT_BUTTON6
)

const (
	CWF_FULLSCREEN = (1 << 0)
	CWF_VSYNC      = (1 << 1)
	CWF_S1         = (1 << 2)
	CWF_S2         = (2 << 2)
	CWF_S3         = (3 << 2)
	CWF_S4         = (4 << 2)
	CWF_S5         = (5 << 2)
	CWF_NEAREST    = (1 << 6)
	CWF_NOVSYNC    = (1 << 7)
)

type Error = C.TLN_Error

const (
	ERR_OK Error = iota
	ERR_OUT_OF_MEMORY
	ERR_IDX_LAYER
	ERR_IDX_SPRITE
	ERR_IDX_ANIMATION
	ERR_IDX_PICTURE
	ERR_REF_TILESET
	ERR_REF_TILEMAP
	ERR_REF_SPRITESET
	ERR_REF_PALETTE
	ERR_REF_SEQUENCE
	ERR_REF_SEQPACK
	ERR_REF_BITMAP
	ERR_NULL_POINTER
	ERR_FILE_NOT_FOUND
	ERR_WRONG_FORMAT
	ERR_WRONG_SIZE
	ERR_UNSUPPORTED
	ERR_REF_LIST
	ERR_IDX_PALETTE
	MAX_ERR
)

type LogLevel = C.TLN_LogLevel

const (
	LOG_NONE LogLevel = iota
	LOG_ERRORS
	LOG_VERBOSE
)

type Layer int
type Sprite int

func boolToUint8(v bool) uint8 {
	var r uint8
	if v {
		r = 1
	}
	return r
}

func convertBool(a CBool) bool {
	return a == CTrue
}

type RasterCallback = func(line int32)

var myRastercallback RasterCallback

//export cRasterCallback
func cRasterCallback(line CInt) {
	if myRastercallback != nil {
		myRastercallback(int32(line))
	}
}

// {
// TLNAPI TLN_Engine TLN_Init (int hres, int vres, int numlayers, int numsprites, int numanimations);
func Init(hres int, vres int, numlayers int, numsprites int, numanimations int) Engine {
	runtime.LockOSThread()
	a := C.TLN_Init(CInt(hres), CInt(vres), CInt(numlayers), CInt(numsprites), CInt(numanimations))

	return Engine{data: a}
}

// TLNAPI void TLN_Deinit (void);
func Deinit() {
	C.TLN_Deinit()
}

// TLNAPI bool TLN_DeleteContext (TLN_Engine context);
func (context *Engine) DeleteContext() bool {
	if context.data == nil {
		return false
	}
	a := convertBool(C.TLN_DeleteContext(context.data))
	if a {
		context.data = nil
	}
	return a
}

// TLNAPI bool TLN_SetContext(TLN_Engine context);
func (context *Engine) SetContext() bool {
	if context.data == nil {
		return false
	}
	return convertBool(C.TLN_SetContext(context.data))
}

// TLNAPI TLN_Engine TLN_GetContext(void);
func GetContext() Engine {
	return Engine{data: C.TLN_GetContext()}
}

// TLNAPI void TLN_SetTargetFps(int fps);
func SetTargetFps(fps int) {
	C.TLN_SetTargetFps(CInt(fps))
}

// TLNAPI int TLN_GetTargetFps(void);
func GetTargetFps() int {
	return int(C.TLN_GetTargetFps())
}

// TLNAPI int TLN_GetWidth (void);
func GetWidth() int {
	return int(C.TLN_GetWidth())
}

// TLNAPI int TLN_GetHeight (void);
func GetHeight() int {
	return int(C.TLN_GetHeight())
}

// TLNAPI uint32_t TLN_GetNumObjects (void);
func GetNumObjects() uint32 {
	return uint32(C.TLN_GetNumObjects())
}

// TLNAPI uint32_t TLN_GetUsedMemory (void);
func GetUsedMemory() uint32 {
	return uint32(C.TLN_GetUsedMemory())
}

// TLNAPI uint32_t TLN_GetVersion (void);
func GetVersion() uint32 {
	return uint32(C.TLN_GetVersion())
}

// TLNAPI int TLN_GetNumLayers (void);
func GetNumLayers() int {
	return int(C.TLN_GetNumLayers())
}

// TLNAPI int TLN_GetNumSprites (void);
func GetNumSprites() int {
	return int(C.TLN_GetNumSprites())
}

// TLNAPI void TLN_SetBGColor (uint8_t r, uint8_t g, uint8_t b);
func SetBGColor(r uint8, g uint8, b uint8) {
	C.TLN_SetBGColor(CUint8(r), CUint8(g), CUint8(b))
}

// TLNAPI bool TLN_SetBGColorFromTilemap (TLN_Tilemap tilemap);
func (tilemap Tilemap) SetBGColorFromTilemap() bool {
	return convertBool(C.TLN_SetBGColorFromTilemap(tilemap.data))
}

// TLNAPI void TLN_DisableBGColor (void);
func DisableBGColor() {
	C.TLN_DisableBGColor()
}

// TLNAPI bool TLN_SetBGBitmap (TLN_Bitmap bitmap);
func (bitmap Bitmap) SetBGBitmap() bool {
	return convertBool(C.TLN_SetBGBitmap(bitmap.data))
}

// TLNAPI bool TLN_SetBGPalette (TLN_Palette palette);
func SetBGPalette(palette Palette) bool {
	return convertBool(C.TLN_SetBGPalette(palette.data))
}

// TLNAPI bool TLN_SetGlobalPalette(int index, TLN_Palette palette);
func SetGlobalPalette(index int, palette Palette) bool {
	return convertBool(C.TLN_SetGlobalPalette(CInt(index), palette.data))
}

// TLNAPI void TLN_SetRasterCallback (TLN_VideoCallback);
func setRasterCallback(cb VideoCallback) {
	C.TLN_SetRasterCallback(cb)
}

func SetRasterCallback(cb RasterCallback) {
	myRastercallback = cb
}

// TLNAPI void TLN_SetFrameCallback (TLN_VideoCallback);
func SetFrameCallback(cb VideoCallback) {
	C.TLN_SetFrameCallback(cb)
}

// TLNAPI void TLN_SetRenderTarget (uint8_t* data, int pitch);
func SetRenderTarget(data []CUint8, pitch int) {
	C.TLN_SetRenderTarget(&data[0], CInt(pitch))
}

// TLNAPI void TLN_UpdateFrame (int frame);
func UpdateFrame(frame int) {
	C.TLN_UpdateFrame(CInt(frame))
}

// TLNAPI void TLN_SetLoadPath (const char* path);
func SetLoadPath(path string) {
	C.TLN_SetLoadPath(C.CString(path))
}

// TLNAPI void TLN_SetCustomBlendFunction (TLN_BlendFunction);
func SetCustomBlendFunction(cb BlendFunction) {
	C.TLN_SetCustomBlendFunction(cb)
}

// TLNAPI void TLN_SetLogLevel(TLN_LogLevel log_level);
func SetLogLevel(log_level LogLevel) {
	C.TLN_SetLogLevel(log_level)
}

// TLNAPI bool TLN_OpenResourcePack(const char* filename, const char* key);
func OpenResourcePack(filename string, key string) bool {
	return convertBool(C.TLN_OpenResourcePack(C.CString(filename), C.CString(key)))
}

// TLNAPI void TLN_CloseResourcePack(void);
func CloseResourcePack() {
	C.TLN_CloseResourcePack()
}

// TLNAPI TLN_Palette TLN_GetGlobalPalette(int index);
func GetGlobalPalette(index int) Palette {
	return Palette{data: C.TLN_GetGlobalPalette(CInt(index))}
}

// }

// {
// TLNAPI void TLN_SetLastError (TLN_Error error);
func SetLastError(error Error) {
	C.TLN_SetLastError(error)
}

// TLNAPI TLN_Error TLN_GetLastError (void);
func GetLastError() Error {
	return C.TLN_GetLastError()
}

// TLNAPI const char *TLN_GetErrorString (TLN_Error error);
func GetErrorString(error Error) string {
	return C.GoString(C.TLN_GetErrorString(error))
}

// }

// , {
// TLNAPI bool TLN_CreateWindow (const char* overlay, int flags);
func CreateWindow(overlay string, flags int) bool {
	a := convertBool(C.TLN_CreateWindow(C.CString(overlay), CInt(flags)))
	setRasterCallback(VideoCallback(C.cRasterCallback))
	return a
}

// TLNAPI bool TLN_CreateWindowThread (const char* overlay, int flags);
func CreateWindowThread(overlay string, flags int) bool {
	a := convertBool(C.TLN_CreateWindowThread(C.CString(overlay), CInt(flags)))
	setRasterCallback(VideoCallback(C.cRasterCallback))
	return a
}

// TLNAPI void TLN_SetWindowTitle (const char* title);
func SetWindowTitle(title string) {
	C.TLN_SetWindowTitle(C.CString(title))
}

// TLNAPI bool TLN_ProcessWindow (void);
func ProcessWindow() bool {
	return convertBool(C.TLN_ProcessWindow())
}

// TLNAPI bool TLN_IsWindowActive (void);
func IsWindowActive() bool {
	return convertBool(C.TLN_IsWindowActive())
}

// TLNAPI bool TLN_GetInput (TLN_Input id);
func GetInput(id Input) bool {
	return convertBool(C.TLN_GetInput(id))
}

// TLNAPI void TLN_EnableInput (TLN_Player player, bool enable);
func EnableInput(player Player, enable bool) {
	// b := 0
	// if enable {
	// 	b = 1
	// }
	C.TLN_EnableInput(player, CBool(boolToUint8(enable)))
}

// TLNAPI void TLN_AssignInputJoystick (TLN_Player player, int index);
func AssignInputJoystick(player Player, index int) {
	C.TLN_AssignInputJoystick(player, CInt(index))
}

// TLNAPI void TLN_DefineInputKey (TLN_Player player, TLN_Input input, uint32_t keycode);
func DefineInputKey(player Player, input Input, keycode uint32) {
	C.TLN_DefineInputKey(player, input, CUint32(keycode))
}

// TLNAPI void TLN_DefineInputButton (TLN_Player player, TLN_Input input, uint8_t joybutton);
func DefineInputButton(player Player, input Input, joybutton uint8) {
	C.TLN_DefineInputButton(player, input, CUint8(joybutton))
}

// TLNAPI void TLN_DrawFrame (int frame);
func DrawFrame(frame int) {
	C.TLN_DrawFrame(CInt(frame))
}

// TLNAPI void TLN_WaitRedraw (void);
func WaitRedraw() {
	C.TLN_WaitRedraw()
}

// TLNAPI void TLN_DeleteWindow (void);
func DeleteWindow() {
	C.TLN_DeleteWindow()
}

// TLNAPI void TLN_EnableBlur (bool mode);
func EnableBlur(mode bool) {
	// b := 0
	// if mode {
	// 	b = 1
	// }
	C.TLN_EnableBlur(CBool(boolToUint8(mode)))
}

// TLNAPI void TLN_ConfigCRTEffect(TLN_CRT type, bool blur);
func ConfigCRTEffect(_type CRT, blur bool) {
	// b := 0
	// if blur {
	// 	b = 1
	// }
	C.TLN_ConfigCRTEffect(_type, CBool(boolToUint8(blur)))
}

// TLNAPI void TLN_EnableCRTEffect (int overlay, uint8_t overlay_factor, uint8_t threshold, uint8_t v0, uint8_t v1, uint8_t v2, uint8_t v3, bool blur, uint8_t glow_factor);
func EnableCRTEffect(overlay int, overlay_factor uint8, threshold uint8, v0 uint8, v1 uint8, v2 uint8, v3 uint8, blur bool, glow_factor uint8) {
	// b := 0
	// if blur {
	// 	b = 1
	// }
	C.TLN_EnableCRTEffect(CInt(overlay), CUint8(overlay_factor), CUint8(threshold), CUint8(v0), CUint8(v1), CUint8(v2), CUint8(v3), CBool(boolToUint8(blur)), CUint8(glow_factor))
}

// TLNAPI void TLN_DisableCRTEffect (void);
func DisableCRTEffect() {
	C.TLN_DisableCRTEffect()
}

// TLNAPI void TLN_SetSDLCallback(TLN_SDLCallback);
func SetSDLCallback(cb SDLCallback) {
	C.TLN_SetSDLCallback(cb)
}

// TLNAPI void TLN_Delay (uint32_t msecs);
func Delay(msecs uint32) {
	C.TLN_Delay(CUint32(msecs))
}

// TLNAPI uint32_t TLN_GetTicks (void);
func GetTicks() uint32 {
	return uint32(C.TLN_GetTicks())
}

// TLNAPI uint32_t TLN_GetAverageFps(void);
func GetAverageFps() uint32 {
	return uint32(C.TLN_GetAverageFps())
}

// TLNAPI int TLN_GetWindowWidth(void);
func GetWindowWidth() int {
	return int(C.TLN_GetWindowWidth())
}

// TLNAPI int TLN_GetWindowHeight(void);
func GetWindowHeight() int {
	return int(C.TLN_GetWindowHeight())
}

// TLNAPI int TLN_GetWindowScaleFactor(void);
func GetWindowScaleFactor() int {
	return int(C.TLN_GetWindowScaleFactor())
}

// TLNAPI void TLN_SetWindowScaleFactor(int factor);
func SetWindowScaleFactor(factor int) {
	C.TLN_SetWindowScaleFactor(CInt(factor))
}

// }

// {
// TLNAPI TLN_Spriteset TLN_CreateSpriteset (TLN_Bitmap bitmap, TLN_SpriteData* data, int num_entries);
func createSpriteset(bitmap Bitmap, data *C.TLN_SpriteData, num_entries int) Spriteset {
	return Spriteset{data: C.TLN_CreateSpriteset(bitmap.data, data, CInt(num_entries))}
}

func CreateSpriteset(bitmap Bitmap, data []SpriteData, num_entries int) Spriteset {
	return Spriteset{C.TLN_CreateSpriteset(bitmap.data, (*C.TLN_SpriteData)(unsafe.Pointer(&data[0])), CInt(num_entries))}
}

// TLNAPI TLN_Spriteset TLN_LoadSpriteset (const char* name);
func LoadSpriteset(name string) Spriteset {
	return Spriteset{data: C.TLN_LoadSpriteset(C.CString(name))}
}

// TLNAPI TLN_Spriteset TLN_CloneSpriteset (TLN_Spriteset src);
func (src Spriteset) CloneSpriteset() Spriteset {
	return Spriteset{data: C.TLN_CloneSpriteset(src.data)}
}

// TLNAPI bool TLN_GetSpriteInfo (TLN_Spriteset spriteset, int entry, TLN_SpriteInfo* info);
func (spriteset Spriteset) getSpriteInfo(entry int, info *C.TLN_SpriteInfo) bool {
	return convertBool(C.TLN_GetSpriteInfo(spriteset.data, CInt(entry), info))
}

func (spriteset Spriteset) GetSpriteInfo(entry int, info *SpriteInfo) bool {
	return convertBool(C.TLN_GetSpriteInfo(spriteset.data, CInt(entry), (*C.TLN_SpriteInfo)(unsafe.Pointer(info))))
}

// TLNAPI TLN_Palette TLN_GetSpritesetPalette (TLN_Spriteset spriteset);
func (spriteset Spriteset) GetSpritesetPalette() Palette {
	return Palette{data: C.TLN_GetSpritesetPalette(spriteset.data)}
}

// TLNAPI int TLN_FindSpritesetSprite (TLN_Spriteset spriteset, const char* name);
func (spriteset Spriteset) FindSpritesetSprite(name string) int {
	return int(C.TLN_FindSpritesetSprite(spriteset.data, C.CString(name)))
}

// TLNAPI bool TLN_SetSpritesetData (TLN_Spriteset spriteset, int entry, TLN_SpriteData* data, void* pixels, int pitch);
func (spriteset Spriteset) setSpritesetData(entry int, data *C.TLN_SpriteData, pixels unsafe.Pointer, pitch int) bool {
	return convertBool(C.TLN_SetSpritesetData(spriteset.data, CInt(entry), data, pixels, CInt(pitch)))
}

func (spriteset Spriteset) SetSpritesetData(entry int, data *SpriteData, pixels []uint8, pitch int) bool {
	return convertBool(C.TLN_SetSpritesetData(spriteset.data, CInt(entry), (*C.TLN_SpriteData)(unsafe.Pointer(data)), (unsafe.Pointer)(&pixels[0]), CInt(pitch)))
}

// TLNAPI bool TLN_DeleteSpriteset (TLN_Spriteset Spriteset);
func (spriteset *Spriteset) DeleteSpriteset() bool {
	a := convertBool(C.TLN_DeleteSpriteset(spriteset.data))
	if a {
		spriteset.data = nil
	}
	return a
}

// }

// {
// TLNAPI TLN_Tileset TLN_CreateTileset (int numtiles, int width, int height, TLN_Palette palette, TLN_SequencePack sp, TLN_TileAttributes* attributes);
func tLN_CreateTileset(numtiles CInt, width CInt, height CInt, palette Palette, sp SequencePack, attributes *C.TLN_TileAttributes) Tileset {
	return Tileset{data: C.TLN_CreateTileset(numtiles, width, height, palette.data, sp.data, attributes)}
}

func CreateTileset(numtiles, width, height int32, palette Palette, sp SequencePack, attributes []TileAttributes) Tileset {
	return tLN_CreateTileset(CInt(numtiles), CInt(width), CInt(height), palette, sp, (*C.TLN_TileAttributes)(unsafe.Pointer(&attributes[0])))
}

// TLNAPI TLN_Tileset TLN_CreateImageTileset(int numtiles, TLN_TileImage* images);
func CreateImageTileset(images []TileImage) Tileset {
	return Tileset{data: C.TLN_CreateImageTileset(CInt(len(images)), (*C.TLN_TileImage)(unsafe.Pointer(&images[0])))}
}

// func createImageTileset(numtiles int, images []TileImage) Tileset {
// 	return Tileset{data: C.TLN_CreateImageTileset(CInt(numtiles), (*C.TLN_TileImage)(unsafe.Pointer(&images[0])))}
// }

// TLNAPI TLN_Tileset TLN_LoadTileset (const char* filename);
func LoadTileset(filename string) Tileset {
	return Tileset{data: C.TLN_LoadTileset(C.CString(filename))}
}

// TLNAPI TLN_Tileset TLN_CloneTileset (TLN_Tileset src);
func (src Tileset) CloneTileset() Tileset {
	return Tileset{data: C.TLN_CloneTileset(src.data)}
}

// TLNAPI bool TLN_SetTilesetPixels (TLN_Tileset tileset, int entry, uint8_t* srcdata, int srcpitch);
func (tileset Tileset) SetTilesetPixels(entry int, srcdata []uint8, srcpitch int) bool {
	return convertBool(C.TLN_SetTilesetPixels(tileset.data, CInt(entry), (*CUint8)(&srcdata[0]), CInt(srcpitch)))
}

// TLNAPI int TLN_GetTileWidth (TLN_Tileset tileset);
func (tileset Tileset) GetTileWidth() int {
	return int(C.TLN_GetTileWidth(tileset.data))
}

// TLNAPI int TLN_GetTileHeight (TLN_Tileset tileset);
func (tileset Tileset) GetTileHeight() int {
	return int(C.TLN_GetTileHeight(tileset.data))
}

// TLNAPI int TLN_GetTilesetNumTiles(TLN_Tileset tileset);
func (tileset Tileset) GetTilesetNumTiles() int {
	return int(C.TLN_GetTilesetNumTiles(tileset.data))
}

// TLNAPI TLN_Palette TLN_GetTilesetPalette (TLN_Tileset tileset);
func (tileset Tileset) GetTilesetPalette() Palette {
	return Palette{data: C.TLN_GetTilesetPalette(tileset.data)}
}

// TLNAPI TLN_SequencePack TLN_GetTilesetSequencePack (TLN_Tileset tileset);
func (tileset Tileset) GetTilesetSequencePack() SequencePack {
	return SequencePack{data: C.TLN_GetTilesetSequencePack(tileset.data)}
}

// TLNAPI bool TLN_DeleteTileset (TLN_Tileset tileset);
func (tileset *Tileset) DeleteTileset() bool {
	a := convertBool(C.TLN_DeleteTileset(tileset.data))
	if a {
		tileset.data = nil
	}
	return a
}

// }

// {
// TLNAPI TLN_Tilemap TLN_CreateTilemap (int rows, int cols, TLN_Tile tiles, uint32_t bgcolor, TLN_Tileset tileset);
func CreateTilemap(rows, cols int, tiles []Tile, bgcolor uint32, tileset Tileset) Tilemap {
	return Tilemap{data: C.TLN_CreateTilemap(CInt(rows), CInt(cols), (C.TLN_Tile)(unsafe.Pointer(&tiles[0])), CUint32(bgcolor), tileset.data)}
}

// TLNAPI TLN_Tilemap TLN_LoadTilemap (const char* filename, const char* layername);
func LoadTilemap(filename string, layername string) Tilemap {
	(C.CString(filename))
	if layername == "" {
		return Tilemap{data: C.TLN_LoadTilemap(C.CString(filename), nil)}
	}
	return Tilemap{data: C.TLN_LoadTilemap(C.CString(filename), C.CString(layername))}
}

// TLNAPI TLN_Tilemap TLN_CloneTilemap (TLN_Tilemap src);
func (src Tilemap) CloneTilemap() Tilemap {
	return Tilemap{data: C.TLN_CloneTilemap(src.data)}
}

// TLNAPI int TLN_GetTilemapRows (TLN_Tilemap tilemap);
func (tilemap Tilemap) GetTilemapRows() int {
	return int(C.TLN_GetTilemapRows(tilemap.data))
}

// TLNAPI int TLN_GetTilemapCols (TLN_Tilemap tilemap);
func (tilemap Tilemap) GetTilemapCols() int {
	return int(C.TLN_GetTilemapCols(tilemap.data))
}

// TLNAPI bool TLN_SetTilemapTileset(TLN_Tilemap tilemap, TLN_Tileset tileset);
func (tilemap Tilemap) SetTilemapTileset(tileset Tileset) bool {
	return convertBool(C.TLN_SetTilemapTileset(tilemap.data, tileset.data))
}

// TLNAPI TLN_Tileset TLN_GetTilemapTileset (TLN_Tilemap tilemap);
func (tilemap Tilemap) GetTilemapTileset() Tileset {
	return Tileset{data: C.TLN_GetTilemapTileset(tilemap.data)}
}

// TLNAPI bool TLN_SetTilemapTileset2(TLN_Tilemap tilemap, TLN_Tileset tileset, int index);
func (tilemap Tilemap) SetTilemapTileset2(tileset Tileset, index int) bool {
	return convertBool(C.TLN_SetTilemapTileset2(tilemap.data, tileset.data, CInt(index)))
}

// TLNAPI TLN_Tileset TLN_GetTilemapTileset2(TLN_Tilemap tilemap, int index);
func (tilemap Tilemap) GetTilemapTileset2(index int) Tileset {
	return Tileset{data: C.TLN_GetTilemapTileset2(tilemap.data, CInt(index))}
}

// TLNAPI bool TLN_GetTilemapTile (TLN_Tilemap tilemap, int row, int col, TLN_Tile tile);
func (tilemap Tilemap) GetTilemapTile(row, col int, tile *Tile) bool {
	return convertBool(C.TLN_GetTilemapTile(tilemap.data, CInt(row), CInt(col), (C.TLN_Tile)(unsafe.Pointer(tile))))
}

// TLNAPI bool TLN_SetTilemapTile (TLN_Tilemap tilemap, int row, int col, TLN_Tile tile);
func (tilemap Tilemap) SetTilemapTile(row, col int, tile *Tile) bool {
	return convertBool(C.TLN_SetTilemapTile(tilemap.data, CInt(row), CInt(col), (C.TLN_Tile)(unsafe.Pointer(tile))))
}

// TLNAPI bool TLN_CopyTiles (TLN_Tilemap src, int srcrow, int srccol, int rows, int cols, TLN_Tilemap dst, int dstrow, int dstcol);
func (src Tilemap) CopyTiles(srcrow, srccol, rows, cols int, dst Tilemap, dstrow, dstcol int) bool {
	return convertBool(C.TLN_CopyTiles(src.data, CInt(srcrow), CInt(srccol), CInt(rows), CInt(cols), dst.data, CInt(dstrow), CInt(dstcol)))
}

// TLNAPI TLN_Tile TLN_GetTilemapTiles(TLN_Tilemap tilemap, int row, int col);
// func GetTilemapTiles(tilemap Tilemap, row, col int) *Tile {
// 	return (*Tile)(C.TLN_GetTilemapTiles(tilemap, CInt(row), CInt(col)))
// }

// TLNAPI bool TLN_DeleteTilemap (TLN_Tilemap tilemap);
func (tilemap *Tilemap) DeleteTilemap() bool {
	a := convertBool(C.TLN_DeleteTilemap(tilemap.data))
	if a {
		tilemap.data = nil
	}
	return a
}

// }

// {
// TLNAPI TLN_Palette TLN_CreatePalette (int entries);
func CreatePalette(entries int) Palette {
	return Palette{data: C.TLN_CreatePalette(CInt(entries))}
}

// TLNAPI TLN_Palette TLN_LoadPalette (const char* filename);
func LoadPalette(filename string) Palette {
	return Palette{data: C.TLN_LoadPalette(C.CString(filename))}
}

// TLNAPI TLN_Palette TLN_ClonePalette (TLN_Palette src);
func (src Palette) ClonePalette() Palette {
	return Palette{data: C.TLN_ClonePalette(src.data)}
}

// TLNAPI bool TLN_SetPaletteColor (TLN_Palette palette, int color, uint8_t r, uint8_t g, uint8_t b);
func (palette Palette) SetPaletteColor(color int, r, g, b uint8) bool {
	return convertBool(C.TLN_SetPaletteColor(palette.data, CInt(color), CUint8(r), CUint8(g), CUint8(b)))
}

// TLNAPI bool TLN_MixPalettes (TLN_Palette src1, TLN_Palette src2, TLN_Palette dst, uint8_t factor);
func (dst Palette) MixPalettes(src1 Palette, src2 Palette, factor uint8) bool {
	return convertBool(C.TLN_MixPalettes(src1.data, src2.data, dst.data, CUint8(factor)))
}

// TLNAPI bool TLN_AddPaletteColor (TLN_Palette palette, uint8_t r, uint8_t g, uint8_t b, uint8_t start, uint8_t num);
func (palette Palette) AddPaletteColor(r, g, b, start, num uint8) bool {
	return convertBool(C.TLN_AddPaletteColor(palette.data, CUint8(r), CUint8(g), CUint8(b), CUint8(start), CUint8(num)))
}

// TLNAPI bool TLN_SubPaletteColor (TLN_Palette palette, uint8_t r, uint8_t g, uint8_t b, uint8_t start, uint8_t num);
func (palette Palette) SubPaletteColor(r, g, b, start, num uint8) bool {
	return convertBool(C.TLN_SubPaletteColor(palette.data, CUint8(r), CUint8(g), CUint8(b), CUint8(start), CUint8(num)))
}

// TLNAPI bool TLN_ModPaletteColor (TLN_Palette palette, uint8_t r, uint8_t g, uint8_t b, uint8_t start, uint8_t num);
func (palette Palette) ModPaletteColor(r, g, b, start, num uint8) bool {
	return convertBool(C.TLN_ModPaletteColor(palette.data, CUint8(r), CUint8(g), CUint8(b), CUint8(start), CUint8(num)))
}

type Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

// TLNAPI uint8_t* TLN_GetPaletteData (TLN_Palette palette, int index);
func (palette Palette) GetPaletteData(index int) Color {
	c := (*Color)(unsafe.Pointer(C.TLN_GetPaletteData(palette.data, CInt(index))))
	return *c
}

// TLNAPI int TLN_GetPaletteNumColors (TLN_Palette palette);
func (palette Palette) GetPaletteNumColors() int {
	return int(C.TLN_GetPaletteNumColors(palette.data))
}

// TLNAPI bool TLN_DeletePalette (TLN_Palette palette);
func (palette *Palette) DeletePalette() bool {
	a := convertBool(C.TLN_DeletePalette(palette.data))
	if a {
		palette.data = nil
	}
	return a
}

// }

// {
// TLNAPI TLN_Bitmap TLN_CreateBitmap (int width, int height, int bpp);
func CreateBitmap(width, height, bpp int) Bitmap {
	return Bitmap{data: C.TLN_CreateBitmap(CInt(width), CInt(height), CInt(bpp))}
}

// TLNAPI TLN_Bitmap TLN_LoadBitmap (const char* filename);
func LoadBitmap(filename string) Bitmap {
	return Bitmap{data: C.TLN_LoadBitmap(C.CString(filename))}
}

// TLNAPI TLN_Bitmap TLN_CloneBitmap (TLN_Bitmap src);
func (src Bitmap) CloneBitmap() Bitmap {
	return Bitmap{data: C.TLN_CloneBitmap(src.data)}
}

// TLNAPI uint8_t* TLN_GetBitmapPtr (TLN_Bitmap bitmap, int x, int y);
func (bitmap Bitmap) GetBitmapPtr(x, y int) *CUint8 {
	return C.TLN_GetBitmapPtr(bitmap.data, CInt(x), CInt(y))
}

// TLNAPI int TLN_GetBitmapWidth (TLN_Bitmap bitmap);
func (bitmap Bitmap) GetBitmapWidth() int {
	return int(C.TLN_GetBitmapWidth(bitmap.data))
}

// TLNAPI int TLN_GetBitmapHeight (TLN_Bitmap bitmap);
func (bitmap Bitmap) GetBitmapHeight() int {
	return int(C.TLN_GetBitmapHeight(bitmap.data))
}

// TLNAPI int TLN_GetBitmapDepth (TLN_Bitmap bitmap);
func (bitmap Bitmap) GetBitmapDepth() int {
	return int(C.TLN_GetBitmapDepth(bitmap.data))
}

// TLNAPI int TLN_GetBitmapPitch (TLN_Bitmap bitmap);
func (bitmap Bitmap) GetBitmapPitch() int {
	return int(C.TLN_GetBitmapPitch(bitmap.data))
}

// TLNAPI TLN_Palette TLN_GetBitmapPalette (TLN_Bitmap bitmap);
func (bitmap Bitmap) GetBitmapPalette() Palette {
	return Palette{data: C.TLN_GetBitmapPalette(bitmap.data)}
}

// TLNAPI bool TLN_SetBitmapPalette (TLN_Bitmap bitmap, TLN_Palette palette);
func (bitmap Bitmap) SetBitmapPalette(palette Palette) bool {
	return convertBool(C.TLN_SetBitmapPalette(bitmap.data, palette.data))
}

// TLNAPI bool TLN_DeleteBitmap (TLN_Bitmap bitmap);
func (bitmap *Bitmap) DeleteBitmap() bool {
	a := convertBool(C.TLN_DeleteBitmap(bitmap.data))
	if a {
		bitmap.data = nil
	}
	return a
}

// }

// {
// TLNAPI TLN_ObjectList TLN_CreateObjectList(void);
func CreateObjectList() ObjectList {
	return ObjectList{data: C.TLN_CreateObjectList()}
}

// TLNAPI bool TLN_AddTileObjectToList(TLN_ObjectList list, uint16_t id, uint16_t gid, uint16_t flags, int x, int y);
func (list ObjectList) AddTileObjectToList(id, gid, flags uint16, x, y int) bool {
	return convertBool(C.TLN_AddTileObjectToList(list.data, CUint16(id), CUint16(gid), CUint16(flags), CInt(x), CInt(y)))
}

// TLNAPI TLN_ObjectList TLN_LoadObjectList(const char* filename, const char* layername);
func LoadObjectList(filename string, layername string) ObjectList {
	if layername == "" {
		return ObjectList{data: C.TLN_LoadObjectList(C.CString(filename), nil)}
	}
	return ObjectList{data: C.TLN_LoadObjectList(C.CString(filename), C.CString(layername))}
}

// TLNAPI TLN_ObjectList TLN_CloneObjectList(TLN_ObjectList src);
func (src ObjectList) CloneObjectList() ObjectList {
	return ObjectList{data: C.TLN_CloneObjectList(src.data)}
}

// TLNAPI int TLN_GetListNumObjects(TLN_ObjectList list);
func (list ObjectList) GetListNumObjects() int {
	return int(C.TLN_GetListNumObjects(list.data))
}

// TLNAPI bool TLN_GetListObject(TLN_ObjectList list, TLN_ObjectInfo* info);
func (list ObjectList) GetListObject(info *ObjectInfo) bool {
	return convertBool(C.TLN_GetListObject(list.data, (*C.TLN_ObjectInfo)(unsafe.Pointer(info))))
}

// func GetListObject(list ObjectList, info *ObjectInfo) bool {
// 	return convertBool(C.TLN_GetListObject(list, (*C.TLN_ObjectInfo)(unsafe.Pointer(info))))
// }

// TLNAPI bool TLN_DeleteObjectList(TLN_ObjectList list);
func (list *ObjectList) DeleteObjectList() bool {
	r := convertBool(C.TLN_DeleteObjectList(list.data))
	if r {
		list.data = nil
	}
	return r
}

// }

// {
// TLNAPI bool TLN_SetLayer (int nlayer, TLN_Tileset tileset, TLN_Tilemap tilemap);
func (layer Layer) SetLayer(tileset Tileset, tilemap Tilemap) bool {
	return convertBool(C.TLN_SetLayer(CInt(layer), tileset.data, tilemap.data))
}

// TLNAPI bool TLN_SetLayerTilemap(int nlayer, TLN_Tilemap tilemap);
func (layer Layer) SetLayerTilemap(tilemap Tilemap) bool {
	return convertBool(C.TLN_SetLayerTilemap(CInt(layer), tilemap.data))
}

// TLNAPI bool TLN_SetLayerBitmap(int nlayer, TLN_Bitmap bitmap);
func (layer Layer) SetLayerBitmap(bitmap Bitmap) bool {
	return convertBool(C.TLN_SetLayerBitmap(CInt(layer), bitmap.data))
}

// TLNAPI bool TLN_SetLayerPalette (int nlayer, TLN_Palette palette);
func (layer Layer) SetLayerPalette(palette Palette) bool {
	return convertBool(C.TLN_SetLayerPalette(CInt(layer), palette.data))
}

// TLNAPI bool TLN_SetLayerPosition (int nlayer, int hstart, int vstart);
func (layer Layer) SetLayerPosition(hstart, vstart int) bool {
	return convertBool(C.TLN_SetLayerPosition(CInt(layer), CInt(hstart), CInt(vstart)))
}

// TLNAPI bool TLN_SetLayerScaling (int nlayer, float xfactor, float yfactor);
func (layer Layer) SetLayerScaling(xfactor float32, yfactor float32) bool {
	return convertBool(C.TLN_SetLayerScaling(CInt(layer), CFloat(xfactor), CFloat(yfactor)))
}

// TLNAPI bool TLN_SetLayerAffineTransform (int nlayer, TLN_Affine *affine);
func (layer Layer) SetLayerAffineTransform(affine *Affine) bool {
	return convertBool(C.TLN_SetLayerAffineTransform(CInt(layer), (*C.TLN_Affine)(unsafe.Pointer(affine))))
}

// TLNAPI bool TLN_SetLayerTransform (int layer, float angle, float dx, float dy, float sx, float sy);
func (layer Layer) SetLayerTransform(angle, dx, dy, sx, sy float32) bool {
	return convertBool(C.TLN_SetLayerTransform(CInt(layer), CFloat(angle), CFloat(dx), CFloat(dy), CFloat(sx), CFloat(sy)))
}

// TLNAPI bool TLN_SetLayerPixelMapping (int nlayer, TLN_PixelMap* table);
func (layer Layer) SetLayerPixelMapping(table []PixelMap) bool {
	return convertBool(C.TLN_SetLayerPixelMapping(CInt(layer), (*C.TLN_PixelMap)(unsafe.Pointer(&table[0]))))
}

// TLNAPI bool TLN_SetLayerBlendMode (int nlayer, TLN_Blend mode, uint8_t factor);
func (layer Layer) SetLayerBlendMode(mode Blend, factor uint8) bool {
	return convertBool(C.TLN_SetLayerBlendMode(CInt(layer), mode, CUint8(factor)))
}

// TLNAPI bool TLN_SetLayerColumnOffset (int nlayer, int* offset);
func (layer Layer) SetLayerColumnOffset(offset []int32) bool {
	return convertBool(C.TLN_SetLayerColumnOffset(CInt(layer), (*CInt)(unsafe.Pointer(&offset[0]))))
}

// TLNAPI bool TLN_SetLayerClip (int nlayer, int x1, int y1, int x2, int y2);
func (layer Layer) SetLayerClip(x1, y1, x2, y2 int) bool {
	return convertBool(C.TLN_SetLayerClip(CInt(layer), CInt(x1), CInt(y1), CInt(x2), CInt(y2)))
}

// TLNAPI bool TLN_DisableLayerClip (int nlayer);
func (layer Layer) DisableLayerClip() bool {
	return convertBool(C.TLN_DisableLayerClip(CInt(layer)))
}

// TLNAPI bool TLN_SetLayerWindow(int nlayer, int x1, int y1, int x2, int y2, bool invert);
func (layer Layer) SetLayerWindow(x1, y1, x2, y2 int, invert bool) bool {
	return convertBool(C.TLN_SetLayerWindow(CInt(layer), CInt(x1), CInt(y1), CInt(x2), CInt(y2), CBool(boolToUint8(invert))))
}

// TLNAPI bool TLN_SetLayerWindowColor(int nlayer, uint8_t r, uint8_t g, uint8_t b, TLN_Blend blend);
func (layer Layer) SetLayerWindowColor(r, g, b uint8, blend Blend) bool {
	return convertBool(C.TLN_SetLayerWindowColor(CInt(layer), CUint8(r), CUint8(g), CUint8(b), blend))
}

// TLNAPI bool TLN_DisableLayerWindow(int nlayer);
func (layer Layer) DisableLayerWindow() bool {
	return convertBool(C.TLN_DisableLayerWindow(CInt(layer)))
}

// TLNAPI bool TLN_DisableLayerWindowColor(int nlayer);
func (layer Layer) DisableLayerWindowColor() bool {
	return convertBool(C.TLN_DisableLayerWindow(CInt(layer)))
}

// TLNAPI bool TLN_SetLayerMosaic (int nlayer, int width, int height);
func (layer Layer) SetLayerMosaic(width, height int) bool {
	return convertBool(C.TLN_SetLayerMosaic(CInt(layer), CInt(width), CInt(height)))
}

// TLNAPI bool TLN_DisableLayerMosaic (int nlayer);
func (layer Layer) DisableLayerMosaic() bool {
	return convertBool(C.TLN_DisableLayerMosaic(CInt(layer)))
}

// TLNAPI bool TLN_ResetLayerMode (int nlayer);
func (layer Layer) ResetLayerMode() bool {
	return convertBool(C.TLN_ResetLayerMode(CInt(layer)))
}

// TLNAPI bool TLN_SetLayerObjects(int nlayer, TLN_ObjectList objects, TLN_Tileset tileset);
func (layer Layer) SetLayerObjects(objects ObjectList, tileset Tileset) bool {
	return convertBool(C.TLN_SetLayerObjects(CInt(layer), objects.data, tileset.data))
}

// TLNAPI bool TLN_SetLayerPriority(int nlayer, bool enable);
func (layer Layer) SetLayerPriority(enable bool) bool {
	// e := 0
	// if enable {
	// 	e = 1
	// }
	return convertBool(C.TLN_SetLayerPriority(CInt(layer), CBool(boolToUint8(enable))))
}

// TLNAPI bool TLN_SetLayerParent(int nlayer, int parent);
func (layer Layer) SetLayerParent(parent Layer) bool {
	return convertBool(C.TLN_SetLayerParent(CInt(layer), CInt(parent)))
}

// TLNAPI bool TLN_DisableLayerParent(int nlayer);
func (layer Layer) DisableLayerParent() bool {
	return convertBool(C.TLN_DisableLayerParent(CInt(layer)))
}

// TLNAPI bool TLN_DisableLayer (int nlayer);
func (layer Layer) DisableLayer() bool {
	return convertBool(C.TLN_DisableLayer(CInt(layer)))
}

// TLNAPI bool TLN_EnableLayer(int nlayer);
func (layer Layer) EnableLayer() bool {
	return convertBool(C.TLN_EnableLayer(CInt(layer)))
}

// TLNAPI TLN_LayerType TLN_GetLayerType(int nlayer);
func (layer Layer) GetLayerType() LayerType {
	return C.TLN_GetLayerType(CInt(layer))
}

// TLNAPI TLN_Palette TLN_GetLayerPalette (int nlayer);
func (layer Layer) GetLayerPalette() Palette {
	return Palette{data: C.TLN_GetLayerPalette(CInt(layer))}
}

// TLNAPI TLN_Tileset TLN_GetLayerTileset(int nlayer);
func (layer Layer) GetLayerTileset() Tileset {
	return Tileset{data: C.TLN_GetLayerTileset(CInt(layer))}
}

// TLNAPI TLN_Tilemap TLN_GetLayerTilemap(int nlayer);
func (layer Layer) GetLayerTilemap() Tilemap {
	return Tilemap{data: C.TLN_GetLayerTilemap(CInt(layer))}
}

// TLNAPI TLN_Bitmap TLN_GetLayerBitmap(int nlayer);
func (layer Layer) GetLayerBitmap() Bitmap {
	return Bitmap{data: C.TLN_GetLayerBitmap(CInt(layer))}
}

// TLNAPI TLN_ObjectList TLN_GetLayerObjects(int nlayer);
func (layer Layer) GetLayerObjects() ObjectList {
	return ObjectList{data: C.TLN_GetLayerObjects(CInt(layer))}
}

// TLNAPI bool TLN_GetLayerTile (int nlayer, int x, int y, TLN_TileInfo* info);
func (layer Layer) GetLayerTile(x, y int, info *TileInfo) bool {
	return convertBool(C.TLN_GetLayerTile(CInt(layer), CInt(x), CInt(y), (*C.TLN_TileInfo)(unsafe.Pointer(info))))
}

// TLNAPI int TLN_GetLayerWidth (int nlayer);
func (layer Layer) GetLayerWidth() int {
	return int(C.TLN_GetLayerWidth(CInt(layer)))
}

// TLNAPI int TLN_GetLayerHeight (int nlayer);
func (layer Layer) GetLayerHeight() int {
	return int(C.TLN_GetLayerHeight(CInt(layer)))
}

// TLNAPI int TLN_GetLayerX(int nlayer);
func (layer Layer) GetLayerX() int {
	return int(C.TLN_GetLayerX(CInt(layer)))
}

// TLNAPI int TLN_GetLayerY(int nlayer);
func (layer Layer) GetLayerY() int {
	return int(C.TLN_GetLayerY(CInt(layer)))
}

// }

// {
// TLNAPI bool TLN_ConfigSprite (int nsprite, TLN_Spriteset spriteset, uint32_t flags);
func (sprite Sprite) ConfigSprite(spriteset Spriteset, flags uint32) bool {
	return convertBool(C.TLN_ConfigSprite(CInt(sprite), spriteset.data, CUint32(flags)))
}

// TLNAPI bool TLN_SetSpriteSet (int nsprite, TLN_Spriteset spriteset);
func (sprite Sprite) SetSpriteSet(spriteset Spriteset) bool {
	return convertBool(C.TLN_SetSpriteSet(CInt(sprite), spriteset.data))
}

// TLNAPI bool TLN_SetSpriteFlags (int nsprite, uint32_t flags);
func (sprite Sprite) SetSpriteFlags(flags uint32) bool {
	return convertBool(C.TLN_SetSpriteFlags(CInt(sprite), CUint32(flags)))
}

// TLNAPI bool TLN_EnableSpriteFlag(int nsprite, uint32_t flag, bool enable);
func (sprite Sprite) EnableSpriteFlag(flag uint32, enable bool) bool {
	// e := 0
	// if enable {
	// 	e = 1
	// }
	return convertBool(C.TLN_EnableSpriteFlag(CInt(sprite), CUint32(flag), CBool(boolToUint8(enable))))
}

// TLNAPI bool TLN_SetSpritePivot(int nsprite, float px, float py);
func (sprite Sprite) SetSpritePivot(px float32, py float32) bool {
	return convertBool(C.TLN_SetSpritePivot(CInt(sprite), CFloat(px), CFloat(py)))
}

// TLNAPI bool TLN_SetSpritePosition (int nsprite, int x, int y);
func (sprite Sprite) SetSpritePosition(x, y int) bool {
	return convertBool(C.TLN_SetSpritePosition(CInt(sprite), CInt(x), CInt(y)))
}

// TLNAPI bool TLN_SetSpritePicture (int nsprite, int entry);
func (sprite Sprite) SetSpritePicture(entry int) bool {
	return convertBool(C.TLN_SetSpritePicture(CInt(sprite), CInt(entry)))
}

// TLNAPI bool TLN_SetSpritePalette (int nsprite, TLN_Palette palette);
func (sprite Sprite) SetSpritePalette(palette Palette) bool {
	return convertBool(C.TLN_SetSpritePalette(CInt(sprite), palette.data))
}

// TLNAPI bool TLN_SetSpriteBlendMode (int nsprite, TLN_Blend mode, uint8_t factor);
func (sprite Sprite) SetSpriteBlendMode(mode Blend, factor uint8) bool {
	return convertBool(C.TLN_SetSpriteBlendMode(CInt(sprite), mode, CUint8(factor)))
}

// TLNAPI bool TLN_SetSpriteScaling (int nsprite, float sx, float sy);
func (sprite Sprite) SetSpriteScaling(sx, sy float32) bool {
	return convertBool(C.TLN_SetSpriteScaling(CInt(sprite), CFloat(sx), CFloat(sy)))
}

// TLNAPI bool TLN_ResetSpriteScaling (int nsprite);
func ResetSpriteScaling(nsprite int) bool {
	return convertBool(C.TLN_ResetSpriteScaling(CInt(nsprite)))
}

// TLNAPI bool TLN_SetSpriteRotation (int nsprite, float angle);
// func TLN_SetSpriteRotation (nsprite CInt, angle CFloat) CBool {
//	return C.TLN_SetSpriteRotation ( nsprite, angle)
// }

// TLNAPI bool TLN_ResetSpriteRotation (int nsprite);
// func TLN_ResetSpriteRotation(nsprite CInt) CBool {
//	return C.TLN_ResetSpriteRotation(nsprite)
// }

// TLNAPI int  TLN_GetSpritePicture (int nsprite);
func (sprite Sprite) GetSpritePicture() int {
	return int(C.TLN_GetSpritePicture(CInt(sprite)))
}

// TLNAPI int TLN_GetSpriteX(int nsprite);
func (sprite Sprite) GetSpriteX() int {
	return int(C.TLN_GetSpriteX(CInt(sprite)))
}

// TLNAPI int TLN_GetSpriteY(int nsprite);
func (sprite Sprite) GetSpriteY() int {
	return int(C.TLN_GetSpriteY(CInt(sprite)))
}

// TLNAPI int  TLN_GetAvailableSprite (void);
func GetAvailableSprite() Sprite {
	return Sprite(C.TLN_GetAvailableSprite())
}

// TLNAPI bool TLN_EnableSpriteCollision (int nsprite, bool enable);
func (sprite Sprite) EnableSpriteCollision(enable bool) bool {
	// e := 0
	// if enable {
	// 	e = 1
	// }
	return convertBool(C.TLN_EnableSpriteCollision(CInt(sprite), CBool(boolToUint8(enable))))
}

// TLNAPI bool TLN_GetSpriteCollision (int nsprite);
func (sprite Sprite) GetSpriteCollision() bool {
	return convertBool(C.TLN_GetSpriteCollision(CInt(sprite)))
}

// TLNAPI bool TLN_GetSpriteState(int nsprite, TLN_SpriteState* state);
func getSpriteState(nsprite int, state *C.TLN_SpriteState) bool {
	return convertBool(C.TLN_GetSpriteState(CInt(nsprite), state))
}

func (sprite Sprite) GetSpriteState(state *SpriteState) bool {
	return convertBool(C.TLN_GetSpriteState(CInt(sprite), (*C.TLN_SpriteState)(unsafe.Pointer(state))))
}

// TLNAPI bool TLN_SetFirstSprite(int nsprite);
func (sprite Sprite) SetFirstSprite() bool {
	return convertBool(C.TLN_SetFirstSprite(CInt(sprite)))
}

// TLNAPI bool TLN_SetNextSprite(int nsprite, int next);
func (sprite Sprite) SetNextSprite(next Sprite) bool {
	return convertBool(C.TLN_SetNextSprite(CInt(sprite), CInt(next)))
}

// TLNAPI bool TLN_EnableSpriteMasking(int nsprite, bool enable);
func (sprite Sprite) EnableSpriteMasking(enable bool) bool {
	// e := 0
	// if enable {
	// 	e = 1
	// }
	return convertBool(C.TLN_EnableSpriteMasking(CInt(sprite), CBool(boolToUint8(enable))))
}

// TLNAPI void TLN_SetSpritesMaskRegion(int top_line, int bottom_line);
func SetSpritesMaskRegion(top_line int, bottom_line int) {
	C.TLN_SetSpritesMaskRegion(CInt(top_line), CInt(bottom_line))
}

// TLNAPI bool TLN_SetSpriteAnimation (int nsprite, TLN_Sequence sequence, int loop);
func (sprite Sprite) SetSpriteAnimation(sequence Sequence, loop int) bool {
	return convertBool(C.TLN_SetSpriteAnimation(CInt(sprite), sequence.data, CInt(loop)))
}

// TLNAPI bool TLN_DisableSpriteAnimation(int nsprite);
func (sprite Sprite) DisableSpriteAnimation() bool {
	return convertBool(C.TLN_DisableSpriteAnimation(CInt(sprite)))
}

// TLNAPI bool TLN_PauseSpriteAnimation(int index);
func (sprite Sprite) PauseSpriteAnimation() bool {
	return convertBool(C.TLN_PauseSpriteAnimation(CInt(sprite)))
}

// TLNAPI bool TLN_ResumeSpriteAnimation(int index);
func (sprite Sprite) ResumeSpriteAnimation() bool {
	return convertBool(C.TLN_ResumeSpriteAnimation(CInt(sprite)))
}

// TLNAPI bool TLN_DisableAnimation(int index);
func (sprite Sprite) DisableAnimation() bool {
	return convertBool(C.TLN_DisableAnimation(CInt(sprite)))
}

// TLNAPI bool TLN_DisableSprite (int nsprite);
func (sprite Sprite) DisableSprite() bool {
	return convertBool(C.TLN_DisableSprite(CInt(sprite)))
}

// TLNAPI TLN_Palette TLN_GetSpritePalette (int nsprite);
func (sprite Sprite) GetSpritePalette() Palette {
	return Palette{data: C.TLN_GetSpritePalette(CInt(sprite))}
}

// }

// {
// TLNAPI TLN_Sequence TLN_CreateSequence (const char* name, int target, int num_frames, TLN_SequenceFrame* frames);
func createSequence(name string, target, num_frames int, frames *C.TLN_SequenceFrame) Sequence {
	return Sequence{data: C.TLN_CreateSequence(C.CString(name), CInt(target), CInt(num_frames), frames)}
}

func CreateSequence(name string, target, num_frames int, frames *SequenceFrame) Sequence {
	return createSequence(name, target, num_frames, (*C.TLN_SequenceFrame)(unsafe.Pointer(frames)))
}

// TLNAPI TLN_Sequence TLN_CreateCycle (const char* name, int num_strips, TLN_ColorStrip* strips);
func createCycle(name string, num_strips int, strips *C.TLN_ColorStrip) Sequence {
	return Sequence{data: C.TLN_CreateCycle(C.CString(name), CInt(num_strips), strips)}
}

func CreateCycle(name string, num_strips int, strips *ColorStrip) Sequence {
	return createCycle(name, num_strips, (*C.TLN_ColorStrip)(unsafe.Pointer(strips)))
}

// TLNAPI TLN_Sequence TLN_CreateSpriteSequence(const char* name, TLN_Spriteset spriteset, const char* basename, int delay);
func CreateSpriteSequence(name string, spriteset Spriteset, basename string, delay int) Sequence {
	return Sequence{data: C.TLN_CreateSpriteSequence(C.CString(name), spriteset.data, C.CString(basename), CInt(delay))}
}

// TLNAPI TLN_Sequence TLN_CloneSequence (TLN_Sequence src);
func (src Sequence) CloneSequence() Sequence {
	return Sequence{data: C.TLN_CloneSequence(src.data)}
}

// TLNAPI bool TLN_GetSequenceInfo (TLN_Sequence sequence, TLN_SequenceInfo* info);
func (sequence Sequence) GetSequenceInfo(info *SequenceInfo) bool {
	return convertBool(C.TLN_GetSequenceInfo(sequence.data, (*C.TLN_SequenceInfo)(unsafe.Pointer(info))))
	// return convertBool(C.TLN_GetSequenceInfo(sequence, (*C.TLN_SequenceInfo)(unsafe.Pointer(info))))
}

// func (sequence Sequence) getSequenceInfo(info *C.TLN_SequenceInfo) bool {
// 	return convertBool(C.TLN_GetSequenceInfo(sequence.data, info))
// }

// TLNAPI bool TLN_DeleteSequence (TLN_Sequence sequence);
func (sequence *Sequence) DeleteSequence() bool {
	a := convertBool(C.TLN_DeleteSequence(sequence.data))
	if a {
		sequence.data = nil
	}
	return a
}

// }

// {
// TLNAPI TLN_SequencePack TLN_CreateSequencePack (void);
func CreateSequencePack() SequencePack {
	return SequencePack{data: C.TLN_CreateSequencePack()}
}

// TLNAPI TLN_SequencePack TLN_LoadSequencePack (const char* filename);
func LoadSequencePack(filename string) SequencePack {
	return SequencePack{data: C.TLN_LoadSequencePack(C.CString(filename))}
}

// TLNAPI TLN_Sequence TLN_GetSequence (TLN_SequencePack sp, int index);
func (sp SequencePack) GetSequence(index int) Sequence {
	return Sequence{data: C.TLN_GetSequence(sp.data, CInt(index))}
}

// TLNAPI TLN_Sequence TLN_FindSequence (TLN_SequencePack sp, const char* name);
func (sp SequencePack) FindSequence(name string) Sequence {
	return Sequence{data: C.TLN_FindSequence(sp.data, C.CString(name))}
}

// TLNAPI int TLN_GetSequencePackCount (TLN_SequencePack sp);
func (sp SequencePack) GetSequencePackCount() int {
	return int(C.TLN_GetSequencePackCount(sp.data))
}

// TLNAPI bool TLN_AddSequenceToPack (TLN_SequencePack sp, TLN_Sequence sequence);
func (sp SequencePack) AddSequenceToPack(sequence Sequence) bool {
	return convertBool(C.TLN_AddSequenceToPack(sp.data, sequence.data))
}

// TLNAPI bool TLN_DeleteSequencePack (TLN_SequencePack sp);
func (sp *SequencePack) DeleteSequencePack() bool {
	a := convertBool(C.TLN_DeleteSequencePack(sp.data))
	if a {
		sp.data = nil
	}
	return a
}

// }

// {
// TLNAPI bool TLN_SetPaletteAnimation (int index, TLN_Palette palette, TLN_Sequence sequence, bool blend);
func SetPaletteAnimation(index int, palette Palette, sequence Sequence, blend bool) bool {
	// b := 0
	// if blend {
	// 	b = 1
	// }
	return convertBool(C.TLN_SetPaletteAnimation(CInt(index), palette.data, sequence.data, CBool(boolToUint8(blend))))
}

// TLNAPI bool TLN_SetPaletteAnimationSource (int index, TLN_Palette);
func SetPaletteAnimationSource(index int, palette Palette) bool {
	return convertBool(C.TLN_SetPaletteAnimationSource(CInt(index), palette.data))
}

// TLNAPI bool TLN_GetAnimationState (int index);
func GetAnimationState(index int) bool {
	return convertBool(C.TLN_GetAnimationState(CInt(index)))
}

// TLNAPI bool TLN_SetAnimationDelay (int index, int frame, int delay);
func SetAnimationDelay(index, frame, delay int) bool {
	return convertBool(C.TLN_SetAnimationDelay(CInt(index), CInt(frame), CInt(delay)))
}

// TLNAPI int  TLN_GetAvailableAnimation (void);
func GetAvailableAnimation() int {
	return int(C.TLN_GetAvailableAnimation())
}

// TLNAPI bool TLN_DisablePaletteAnimation(int index);
func DisablePaletteAnimation(index int) bool {
	return convertBool(C.TLN_DisablePaletteAnimation(CInt(index)))
}

// }

// {
// TLNAPI bool TLN_LoadWorld(const char* tmxfile, int first_layer);
func LoadWorld(tmxfile string, first_layer int) bool {
	return convertBool(C.TLN_LoadWorld(C.CString(tmxfile), CInt(first_layer)))
}

// TLNAPI void TLN_SetWorldPosition(int x, int y);
func SetWorldPosition(x, y int) {
	C.TLN_SetWorldPosition(CInt(x), CInt(y))
}

// TLNAPI bool TLN_SetLayerParallaxFactor(int nlayer, float x, float y);
func SetLayerParallaxFactor(nlayer int, x, y float32) bool {
	return convertBool(C.TLN_SetLayerParallaxFactor(CInt(nlayer), CFloat(x), CFloat(y)))
}

// TLNAPI bool TLN_SetSpriteWorldPosition(int nsprite, int x, int y);
func (sprite Sprite) SetSpriteWorldPosition(x, y int) bool {
	return convertBool(C.TLN_SetSpriteWorldPosition(CInt(sprite), CInt(x), CInt(y)))
}

// TLNAPI void TLN_ReleaseWorld(void);
func ReleaseWorld() {
	C.TLN_ReleaseWorld()
}

// }

// func TestPointer(ptr unsafe.Pointer) int {
// 	return 0
// }

// func ConvertPointer[T any](ptr unsafe.Pointer) *T { return (*T)(ptr) }
