package gotilengine

/*
#cgo CFLAGS: -I${SRCDIR}
#cgo LDFLAGS: -lTilengine
#include "Tilengine.h"
*/
import "C"
import "unsafe"

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
	TILENGINE_VER_MAJ = 2
	TILENGINE_VER_MIN = 13
	TILENGINE_VER_REV = 2
	TILENGINE_HEADER_VERSION = ((TILENGINE_VER_MAJ << 16) | (TILENGINE_VER_MIN << 8) | TILENGINE_VER_REV)
)

type TLN_TileFlags = C.TLN_TileFlags

const (
	FLAG_NONE TLN_TileFlags = 0	
	FLAG_FLIPX		= 1 << 15
	FLAG_FLIPY		= 1 << 14
	FLAG_ROTATE		= 1 << 13
	FLAG_PRIORITY	        = 1 << 12
	FLAG_MASKED		= 1 << 11
	FLAG_TILESET            = 7 << 8
	FLAG_PALETTE	        = 7 << 5
)

type TLN_Blend = C.TLN_Blend

const (
	BLEND_NONE TLN_Blend = iota		
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

type TLN_LayerType = C.TLN_LayerType

const (
	LAYER_NONE TLN_LayerType = iota		
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

type TLN_Affine = C.TLN_Affine


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

type Tile = C.Tile 


// typedef struct
// {
// 	int index;		
// 	int delay;		
// }
// TLN_SequenceFrame;

type TLN_SequenceFrame = C.TLN_SequenceFrame


// typedef struct
// {
// 	int delay;		
// 	uint8_t first;	
// 	uint8_t count;	
// 	uint8_t dir;	
// }
// TLN_ColorStrip;

type TLN_ColorStrip = C.TLN_ColorStrip


// typedef struct
// {
// 	char name[32];	
// 	int num_frames;	
// }
// TLN_SequenceInfo;

type TLN_SequenceInfo = C.TLN_SequenceInfo


// typedef struct
// {
// 	char name[64];	
// 	int x;			
// 	int y;			
// 	int w;			
// 	int h;			
// }
// TLN_SpriteData;

type TLN_SpriteData = C.TLN_SpriteData


// typedef struct
// {
// 	int w;			
// 	int h;			
// }
// TLN_SpriteInfo;

type TLN_SpriteInfo = C.TLN_SpriteInfo


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

type TLN_TileInfo = C.TLN_TileInfo


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

type TLN_ObjectInfo = C.TLN_ObjectInfo


// typedef struct
// {
// 	uint8_t	type;		
// 	bool	priority;	
// }
// TLN_TileAttributes;

type TLN_TileAttributes = C.TLN_TileAttributes


const (
	TLN_OVERLAY_NONE = 0
	TLN_OVERLAY_SHADOWMASK = 0
	TLN_OVERLAY_APERTURE = 0
	TLN_OVERLAY_SCANLINES = 0
	TLN_OVERLAY_CUSTOM = 0
)

type TLN_CRT = C.TLN_CRT


const (
	TLN_CRT_SLOT TLN_CRT = iota
	TLN_CRT_APERTURE	
	TLN_CRT_SHADOW 	
)


// typedef struct
// {
// 	int16_t dx;		
// 	int16_t dy;		
// }
// TLN_PixelMap;

type TLN_PixelMap = C.TLN_PixelMap

// typedef struct Engine*		 TLN_Engine;			
type TLN_Engine = C.TLN_Engine

// typedef union  Tile*		 TLN_Tile; 
type TLN_Tile = C.TLN_Tile

// typedef struct Tileset*		 TLN_Tileset; 
type TLN_Tileset = C.TLN_Tileset

// typedef struct Tilemap*		 TLN_Tilemap;			
type TLN_Tilemap = C.TLN_Tilemap

// typedef struct Palette*		 TLN_Palette;			
type TLN_Palette = C.TLN_Palette

// typedef struct Spriteset*	 TLN_Spriteset;			
type TLN_Spriteset = C.TLN_Spriteset

// typedef struct Sequence*	 TLN_Sequence;			
type TLN_Sequence = C.TLN_Sequence

// typedef struct SequencePack* TLN_SequencePack;		
type TLN_SequencePack = C.TLN_SequencePack

// typedef struct Bitmap*		 TLN_Bitmap;			
type TLN_Bitmap = C.TLN_Bitmap

// typedef struct ObjectList*	 TLN_ObjectList;		
type TLN_ObjectList = C.TLN_ObjectList


// typedef struct
// {
// 	TLN_Bitmap bitmap;
// 	uint16_t id;
// 	uint8_t	type;
// }
// TLN_TileImage;

type TLN_TileImage = C.TLN_TileImage


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

type TLN_SpriteState = C.TLN_SpriteState


// typedef union SDL_Event SDL_Event;
type SDL_Event = C.SDL_Event

// typedef void(*TLN_VideoCallback)(int scanline);
type TLN_VideoCallback = C.TLN_VideoCallback

// typedef uint8_t(*TLN_BlendFunction)(uint8_t src, uint8_t dst);
type TLN_BlendFunction = C.TLN_BlendFunction

// typedef void(*TLN_SDLCallback)(SDL_Event*);
type TLN_SDLCallback = C.TLN_SDLCallback

type TLN_Player = C.TLN_Player

const (
	PLAYER1 TLN_Player = iota	
	PLAYER2	
	PLAYER3	
	PLAYER4	
)

type TLN_Input = C.TLN_Input


const (
	INPUT_NONE TLN_Input = iota		
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
	CWF_FULLSCREEN	        = (1 << 0)	
	CWF_VSYNC		= (1 << 1)	
	CWF_S1			= (1 << 2)	
	CWF_S2			= (2 << 2)	
	CWF_S3			= (3 << 2)	
	CWF_S4			= (4 << 2)	
	CWF_S5			= (5 << 2)	
	CWF_NEAREST		= (1 << 6)
)

type TLN_Error = C.TLN_Error


const (
	TLN_ERR_OK TLN_Error = iota				
	TLN_ERR_OUT_OF_MEMORY	
	TLN_ERR_IDX_LAYER		
	TLN_ERR_IDX_SPRITE		
	TLN_ERR_IDX_ANIMATION	
	TLN_ERR_IDX_PICTURE	
	TLN_ERR_REF_TILESET	
	TLN_ERR_REF_TILEMAP	
	TLN_ERR_REF_SPRITESET	
	TLN_ERR_REF_PALETTE	
	TLN_ERR_REF_SEQUENCE	
	TLN_ERR_REF_SEQPACK	
	TLN_ERR_REF_BITMAP		
	TLN_ERR_NULL_POINTER	
	TLN_ERR_FILE_NOT_FOUND	
	TLN_ERR_WRONG_FORMAT	
	TLN_ERR_WRONG_SIZE		
	TLN_ERR_UNSUPPORTED	
	TLN_ERR_REF_LIST		
	TLN_ERR_IDX_PALETTE	
	TLN_MAX_ERR
)

type TLN_LogLevel = C.TLN_LogLevel


const (
	TLN_LOG_NONE TLN_LogLevel = iota		
	TLN_LOG_ERRORS		
	TLN_LOG_VERBOSE	
)


// {
// TLNAPI TLN_Engine TLN_Init (int hres, int vres, int numlayers, int numsprites, int numanimations);
func TLN_Init(hres CInt, vres CInt, numlayers CInt, numsprites CInt, numanimations CInt) TLN_Engine {
	return C.TLN_Init(hres, vres, numlayers, numsprites, numanimations)
}

// TLNAPI void TLN_Deinit (void);
func TLN_Deinit() {
	C.TLN_Deinit()
}

// TLNAPI bool TLN_DeleteContext (TLN_Engine context);
func TLN_DeleteContext(context TLN_Engine) CBool {
	return C.TLN_DeleteContext(context)
}

// TLNAPI bool TLN_SetContext(TLN_Engine context);
func TLN_SetContext(context TLN_Engine) CBool {
	return C.TLN_SetContext(context)
}

// TLNAPI TLN_Engine TLN_GetContext(void);
func TLN_GetContext() TLN_Engine {
	return C.TLN_GetContext()
}

// TLNAPI int TLN_GetWidth (void);
func TLN_GetWidth() CInt {
	return C.TLN_GetWidth()
}

// TLNAPI int TLN_GetHeight (void);
func TLN_GetHeight() CInt {
	return C.TLN_GetHeight()
}

// TLNAPI uint32_t TLN_GetNumObjects (void);
func TLN_GetNumObjects() CUint32 {
	return C.TLN_GetNumObjects()
}

// TLNAPI uint32_t TLN_GetUsedMemory (void);
func TLN_GetUsedMemory() CUint32 {
	return C.TLN_GetUsedMemory()
}

// TLNAPI uint32_t TLN_GetVersion (void);
func TLN_GetVersion() CUint32 {
	return C.TLN_GetVersion()
}

// TLNAPI int TLN_GetNumLayers (void);
func TLN_GetNumLayers() CInt {
	return C.TLN_GetNumLayers()
}

// TLNAPI int TLN_GetNumSprites (void);
func TLN_GetNumSprites() CInt {
	return C.TLN_GetNumSprites()
}

// TLNAPI void TLN_SetBGColor (uint8_t r, uint8_t g, uint8_t b);
func TLN_SetBGColor(r CUint8, g CUint8, b CUint8) {
	C.TLN_SetBGColor (r, g, b)
}

// TLNAPI bool TLN_SetBGColorFromTilemap (TLN_Tilemap tilemap);
func TLN_SetBGColorFromTilemap(tilemap TLN_Tilemap) CBool {
	return C.TLN_SetBGColorFromTilemap(tilemap)
}

// TLNAPI void TLN_DisableBGColor (void);
func TLN_DisableBGColor() {
	C.TLN_DisableBGColor()
}

// TLNAPI bool TLN_SetBGBitmap (TLN_Bitmap bitmap);
func TLN_SetBGBitmap(bitmap TLN_Bitmap) CBool {
	return C.TLN_SetBGBitmap(bitmap)
}

// TLNAPI bool TLN_SetBGPalette (TLN_Palette palette);
func TLN_SetBGPalette (palette TLN_Palette) CBool {
	return C.TLN_SetBGPalette(palette)
}

// TLNAPI bool TLN_SetGlobalPalette(int index, TLN_Palette palette);
func TLN_SetGlobalPalette(index CInt, palette TLN_Palette) CBool {
	return C.TLN_SetGlobalPalette(index, palette)
}

// TLNAPI void TLN_SetRasterCallback (TLN_VideoCallback);
func TLN_SetRasterCallback(cb TLN_VideoCallback) {
	C.TLN_SetRasterCallback(cb)
}

// TLNAPI void TLN_SetFrameCallback (TLN_VideoCallback);
func TLN_SetFrameCallback(cb TLN_VideoCallback) {
	C.TLN_SetFrameCallback(cb)
}

// TLNAPI void TLN_SetRenderTarget (uint8_t* data, int pitch);
func TLN_SetRenderTarget(data *CUint8, pitch CInt) {
	C.TLN_SetRenderTarget(data, pitch)
}

// TLNAPI void TLN_UpdateFrame (int frame);
func TLN_UpdateFrame(frame CInt) {
	C.TLN_UpdateFrame(frame)
}

// TLNAPI void TLN_SetLoadPath (const char* path);
func TLN_SetLoadPath(path CString) {
	C.TLN_SetLoadPath(path)
}

// TLNAPI void TLN_SetCustomBlendFunction (TLN_BlendFunction);
func TLN_SetCustomBlendFunction(cb TLN_BlendFunction) {
	C.TLN_SetCustomBlendFunction(cb)
}

// TLNAPI void TLN_SetLogLevel(TLN_LogLevel log_level);
func TLN_SetLogLevel(log_level TLN_LogLevel) {
	C.TLN_SetLogLevel(log_level)
}

// TLNAPI bool TLN_OpenResourcePack(const char* filename, const char* key);
func TLN_OpenResourcePack(filename CString, key CString) CBool {
	return C.TLN_OpenResourcePack(filename, key)
}

// TLNAPI void TLN_CloseResourcePack(void);
func TLN_CloseResourcePack() {
	C.TLN_CloseResourcePack()
}

// TLNAPI TLN_Palette TLN_GetGlobalPalette(int index);
func TLN_GetGlobalPalette(index CInt) TLN_Palette {
	return C.TLN_GetGlobalPalette(index)
}
// }


// {
// TLNAPI void TLN_SetLastError (TLN_Error error);
func TLN_SetLastError(error TLN_Error) {
	C.TLN_SetLastError(error)
}

// TLNAPI TLN_Error TLN_GetLastError (void);
func TLN_GetLastError() TLN_Error {
	return C.TLN_GetLastError()
}

// TLNAPI const char *TLN_GetErrorString (TLN_Error error);
func TLN_GetErrorString(error TLN_Error) CString {
	return C.TLN_GetErrorString(error)
}
// }


// {
// TLNAPI bool TLN_CreateWindow (const char* overlay, int flags);
func TLN_CreateWindow(overlay CString, flags CInt) CBool {
	return C.TLN_CreateWindow(overlay, flags)
}

// TLNAPI bool TLN_CreateWindowThread (const char* overlay, int flags);
func TLN_CreateWindowThread (overlay CString, flags CInt) CBool {
	return C.TLN_CreateWindowThread (overlay, flags)
}

// TLNAPI void TLN_SetWindowTitle (const char* title);
func TLN_SetWindowTitle (title CString) {
	C.TLN_SetWindowTitle (title)
}

// TLNAPI bool TLN_ProcessWindow (void);
func TLN_ProcessWindow() CBool {
	return C.TLN_ProcessWindow()
}

// TLNAPI bool TLN_IsWindowActive (void);
func TLN_IsWindowActive() CBool {
	return C.TLN_IsWindowActive()
}

// TLNAPI bool TLN_GetInput (TLN_Input id);
func TLN_GetInput (id TLN_Input) CBool {
	return C.TLN_GetInput (id)
}

// TLNAPI void TLN_EnableInput (TLN_Player player, bool enable);
func TLN_EnableInput (player TLN_Player, enable CBool) {
	C.TLN_EnableInput (player, enable)
}

// TLNAPI void TLN_AssignInputJoystick (TLN_Player player, int index);
func TLN_AssignInputJoystick (player TLN_Player, index CInt) {
	C.TLN_AssignInputJoystick (player, index)
}

// TLNAPI void TLN_DefineInputKey (TLN_Player player, TLN_Input input, uint32_t keycode);
func TLN_DefineInputKey (player TLN_Player, input TLN_Input, keycode CUint32) {
	C.TLN_DefineInputKey (player, input, keycode)
}

// TLNAPI void TLN_DefineInputButton (TLN_Player player, TLN_Input input, uint8_t joybutton);
func TLN_DefineInputButton (player TLN_Player, input TLN_Input, joybutton CUint8) {
	C.TLN_DefineInputButton (player, input, joybutton)
}

// TLNAPI void TLN_DrawFrame (int frame);
func TLN_DrawFrame (frame CInt) {
	C.TLN_DrawFrame (frame)
}

// TLNAPI void TLN_WaitRedraw (void);
func TLN_WaitRedraw() {
	C.TLN_WaitRedraw()
}

// TLNAPI void TLN_DeleteWindow (void);
func TLN_DeleteWindow() {
	C.TLN_DeleteWindow()
}

// TLNAPI void TLN_EnableBlur (bool mode);
func TLN_EnableBlur (mode CBool) {
	C.TLN_EnableBlur (mode)
}

// TLNAPI void TLN_ConfigCRTEffect(TLN_CRT type, bool blur);
func TLN_ConfigCRTEffect(_type TLN_CRT, blur CBool) {
	C.TLN_ConfigCRTEffect(_type, blur)
}

// TLNAPI void TLN_EnableCRTEffect (int overlay, uint8_t overlay_factor, uint8_t threshold, uint8_t v0, uint8_t v1, uint8_t v2, uint8_t v3, bool blur, uint8_t glow_factor);
func TLN_EnableCRTEffect (overlay CInt, overlay_factor CUint8, threshold CUint8, v0 CUint8, v1 CUint8, v2 CUint8, v3 CUint8, blur CBool, glow_factor CUint8) {
	C.TLN_EnableCRTEffect (overlay, overlay_factor, threshold, v0, v1, v2, v3, blur, glow_factor)
}

// TLNAPI void TLN_DisableCRTEffect (void);
func TLN_DisableCRTEffect() {
	C.TLN_DisableCRTEffect()
}

// TLNAPI void TLN_SetSDLCallback(TLN_SDLCallback);
func TLN_SetSDLCallback(cb TLN_SDLCallback) {
	C.TLN_SetSDLCallback(cb)
}

// TLNAPI void TLN_Delay (uint32_t msecs);
func TLN_Delay (msecs CUint32) {
	C.TLN_Delay (msecs)
}

// TLNAPI uint32_t TLN_GetTicks (void);
func TLN_GetTicks() CUint32 {
	return C.TLN_GetTicks()
}

// TLNAPI int TLN_GetWindowWidth(void);
func TLN_GetWindowWidth() CInt {
	return C.TLN_GetWindowWidth()
}

// TLNAPI int TLN_GetWindowHeight(void);
func TLN_GetWindowHeight() CInt {
	return C.TLN_GetWindowHeight()
}
// }


// {
// TLNAPI TLN_Spriteset TLN_CreateSpriteset (TLN_Bitmap bitmap, TLN_SpriteData* data, int num_entries);
func TLN_CreateSpriteset(bitmap TLN_Bitmap, data *TLN_SpriteData, num_entries CInt) TLN_Spriteset {
	return C.TLN_CreateSpriteset(bitmap, data, num_entries)
}

// TLNAPI TLN_Spriteset TLN_LoadSpriteset (const char* name);
func TLN_LoadSpriteset(name CString) TLN_Spriteset {
	return C.TLN_LoadSpriteset(name)
}

// TLNAPI TLN_Spriteset TLN_CloneSpriteset (TLN_Spriteset src);
func TLN_CloneSpriteset(src TLN_Spriteset) TLN_Spriteset {
	return C.TLN_CloneSpriteset(src)
}

// TLNAPI bool TLN_GetSpriteInfo (TLN_Spriteset spriteset, int entry, TLN_SpriteInfo* info);
func TLN_GetSpriteInfo(spriteset TLN_Spriteset, entry CInt, info *TLN_SpriteInfo) CBool {
	return C.TLN_GetSpriteInfo(spriteset, entry, info)
}

// TLNAPI TLN_Palette TLN_GetSpritesetPalette (TLN_Spriteset spriteset);
func TLN_GetSpritesetPalette(spriteset TLN_Spriteset) TLN_Palette {
	return C.TLN_GetSpritesetPalette(spriteset)
}

// TLNAPI int TLN_FindSpritesetSprite (TLN_Spriteset spriteset, const char* name);
func TLN_FindSpritesetSprite(spriteset TLN_Spriteset, name CString) CInt {
	return C.TLN_FindSpritesetSprite(spriteset, name)
}

// TLNAPI bool TLN_SetSpritesetData (TLN_Spriteset spriteset, int entry, TLN_SpriteData* data, void* pixels, int pitch);
func TLN_SetSpritesetData(spriteset TLN_Spriteset, entry CInt, data *TLN_SpriteData, pixels unsafe.Pointer, pitch CInt) CBool {
	return C.TLN_SetSpritesetData(spriteset, entry, data, pixels, pitch)
}
// TLNAPI bool TLN_DeleteSpriteset (TLN_Spriteset Spriteset);
func TLN_DeleteSpriteset(Spriteset TLN_Spriteset) CBool {
	return C.TLN_DeleteSpriteset(Spriteset)
}
// }


// {
// TLNAPI TLN_Tileset TLN_CreateTileset (int numtiles, int width, int height, TLN_Palette palette, TLN_SequencePack sp, TLN_TileAttributes* attributes);
func TLN_CreateTileset(numtiles CInt, width CInt, height CInt, palette TLN_Palette, sp TLN_SequencePack, attributes *TLN_TileAttributes) TLN_Tileset {
	return C.TLN_CreateTileset(numtiles, width, height, palette, sp, attributes)
}

// TLNAPI TLN_Tileset TLN_CreateImageTileset(int numtiles, TLN_TileImage* images);
func TLN_CreateImageTileset(numtiles CInt, images *TLN_TileImage) TLN_Tileset {
	return C.TLN_CreateImageTileset(numtiles, images)
}

// TLNAPI TLN_Tileset TLN_LoadTileset (const char* filename);
func TLN_LoadTileset(filename CString) TLN_Tileset {
	return C.TLN_LoadTileset(filename)
}

// TLNAPI TLN_Tileset TLN_CloneTileset (TLN_Tileset src);
func TLN_CloneTileset(src TLN_Tileset) TLN_Tileset {
	return C.TLN_CloneTileset(src)
}

// TLNAPI bool TLN_SetTilesetPixels (TLN_Tileset tileset, int entry, uint8_t* srcdata, int srcpitch);
func TLN_SetTilesetPixels(tileset TLN_Tileset, entry CInt, srcdata *CUint8, srcpitch CInt) CBool {
	return C.TLN_SetTilesetPixels(tileset, entry, srcdata, srcpitch)
}

// TLNAPI int TLN_GetTileWidth (TLN_Tileset tileset);
func TLN_GetTileWidth(tileset TLN_Tileset) CInt {
	return C.TLN_GetTileWidth(tileset)
}

// TLNAPI int TLN_GetTileHeight (TLN_Tileset tileset);
func TLN_GetTileHeight(tileset TLN_Tileset) CInt {
	return C.TLN_GetTileHeight(tileset)
}

// TLNAPI int TLN_GetTilesetNumTiles(TLN_Tileset tileset);
func TLN_GetTilesetNumTiles(tileset TLN_Tileset) CInt {
	return C.TLN_GetTilesetNumTiles(tileset)
}

// TLNAPI TLN_Palette TLN_GetTilesetPalette (TLN_Tileset tileset);
func TLN_GetTilesetPalette(tileset TLN_Tileset) TLN_Palette {
	return C.TLN_GetTilesetPalette(tileset)
}

// TLNAPI TLN_SequencePack TLN_GetTilesetSequencePack (TLN_Tileset tileset);
func TLN_GetTilesetSequencePack(tileset TLN_Tileset) TLN_SequencePack {
	return C.TLN_GetTilesetSequencePack(tileset)
}

// TLNAPI bool TLN_DeleteTileset (TLN_Tileset tileset);
func TLN_DeleteTileset(tileset TLN_Tileset) CBool {
	return C.TLN_DeleteTileset(tileset)
}
// }


// {
// TLNAPI TLN_Tilemap TLN_CreateTilemap (int rows, int cols, TLN_Tile tiles, uint32_t bgcolor, TLN_Tileset tileset);
func TLN_CreateTilemap(rows CInt, cols CInt, tiles TLN_Tile, bgcolor CUint32, tileset TLN_Tileset) TLN_Tilemap {
	return C.TLN_CreateTilemap(rows, cols, tiles, bgcolor, tileset)
}

// TLNAPI TLN_Tilemap TLN_LoadTilemap (const char* filename, const char* layername);
func TLN_LoadTilemap(filename CString, layername CString) TLN_Tilemap {
	return C.TLN_LoadTilemap(filename, layername)
}

// TLNAPI TLN_Tilemap TLN_CloneTilemap (TLN_Tilemap src);
func TLN_CloneTilemap(src TLN_Tilemap) TLN_Tilemap {
	return C.TLN_CloneTilemap(src)
}

// TLNAPI int TLN_GetTilemapRows (TLN_Tilemap tilemap);
func TLN_GetTilemapRows(tilemap TLN_Tilemap) CInt {
	return C.TLN_GetTilemapRows(tilemap)
}

// TLNAPI int TLN_GetTilemapCols (TLN_Tilemap tilemap);
func TLN_GetTilemapCols(tilemap TLN_Tilemap) CInt {
	return C.TLN_GetTilemapCols(tilemap)
}

// TLNAPI bool TLN_SetTilemapTileset(TLN_Tilemap tilemap, TLN_Tileset tileset);
func TLN_SetTilemapTileset(tilemap TLN_Tilemap, tileset TLN_Tileset) CBool {
	return C.TLN_SetTilemapTileset(tilemap, tileset)
}

// TLNAPI TLN_Tileset TLN_GetTilemapTileset (TLN_Tilemap tilemap);
func TLN_GetTilemapTileset(tilemap TLN_Tilemap) TLN_Tileset {
	return C.TLN_GetTilemapTileset(tilemap)
}

// TLNAPI bool TLN_SetTilemapTileset2(TLN_Tilemap tilemap, TLN_Tileset tileset, int index);
func TLN_SetTilemapTileset2(tilemap TLN_Tilemap, tileset TLN_Tileset, index CInt) CBool {
	return C.TLN_SetTilemapTileset2(tilemap, tileset, index)
}

// TLNAPI TLN_Tileset TLN_GetTilemapTileset2(TLN_Tilemap tilemap, int index);
func TLN_GetTilemapTileset2(tilemap TLN_Tilemap, index CInt) TLN_Tileset {
	return C.TLN_GetTilemapTileset2(tilemap, index)
}

// TLNAPI bool TLN_GetTilemapTile (TLN_Tilemap tilemap, int row, int col, TLN_Tile tile);
func TLN_GetTilemapTile(tilemap TLN_Tilemap, row CInt, col CInt, tile TLN_Tile) CBool {
	return C.TLN_GetTilemapTile(tilemap, row, col, tile)
}

// TLNAPI bool TLN_SetTilemapTile (TLN_Tilemap tilemap, int row, int col, TLN_Tile tile);
func TLN_SetTilemapTile(tilemap TLN_Tilemap, row CInt, col CInt, tile TLN_Tile) CBool {
	return C.TLN_SetTilemapTile(tilemap, row, col, tile)
}

// TLNAPI bool TLN_CopyTiles (TLN_Tilemap src, int srcrow, int srccol, int rows, int cols, TLN_Tilemap dst, int dstrow, int dstcol);
func TLN_CopyTiles(src TLN_Tilemap, srcrow CInt, srccol CInt, rows CInt, cols CInt, dst TLN_Tilemap, dstrow CInt, dstcol CInt) CBool {
	return C.TLN_CopyTiles(src, srcrow, srccol, rows, cols, dst, dstrow, dstcol)
}

// TLNAPI TLN_Tile TLN_GetTilemapTiles(TLN_Tilemap tilemap, int row, int col);
func TLN_GetTilemapTiles(tilemap TLN_Tilemap, row CInt, col CInt) TLN_Tile {
	return C.TLN_GetTilemapTiles(tilemap, row, col)
}

// TLNAPI bool TLN_DeleteTilemap (TLN_Tilemap tilemap);
func TLN_DeleteTilemap(tilemap TLN_Tilemap) CBool {
	return C.TLN_DeleteTilemap(tilemap)
}
// }


// {
// TLNAPI TLN_Palette TLN_CreatePalette (int entries);
func TLN_CreatePalette(entries CInt) TLN_Palette {
	return C.TLN_CreatePalette(entries)
}

// TLNAPI TLN_Palette TLN_LoadPalette (const char* filename);
func TLN_LoadPalette(filename CString) TLN_Palette {
	return C.TLN_LoadPalette(filename)
}

// TLNAPI TLN_Palette TLN_ClonePalette (TLN_Palette src);
func TLN_ClonePalette(src TLN_Palette) TLN_Palette {
	return C.TLN_ClonePalette(src)
}

// TLNAPI bool TLN_SetPaletteColor (TLN_Palette palette, int color, uint8_t r, uint8_t g, uint8_t b);
func TLN_SetPaletteColor(palette TLN_Palette, color CInt, r CUint8, g CUint8, b CUint8) CBool {
	return C.TLN_SetPaletteColor(palette, color, r, g, b)
}

// TLNAPI bool TLN_MixPalettes (TLN_Palette src1, TLN_Palette src2, TLN_Palette dst, uint8_t factor);
func TLN_MixPalettes(src1 TLN_Palette, src2 TLN_Palette, dst TLN_Palette, factor CUint8) CBool {
	return C.TLN_MixPalettes(src1, src2, dst, factor)
}

// TLNAPI bool TLN_AddPaletteColor (TLN_Palette palette, uint8_t r, uint8_t g, uint8_t b, uint8_t start, uint8_t num);
func TLN_AddPaletteColor(palette TLN_Palette, r CUint8, g CUint8, b CUint8, start CUint8, num CUint8) CBool {
	return C.TLN_AddPaletteColor(palette, r, g, b, start, num)
}

// TLNAPI bool TLN_SubPaletteColor (TLN_Palette palette, uint8_t r, uint8_t g, uint8_t b, uint8_t start, uint8_t num);
func TLN_SubPaletteColor(palette TLN_Palette, r CUint8, g CUint8, b CUint8, start CUint8, num CUint8) CBool {
	return C.TLN_SubPaletteColor(palette, r, g, b, start, num)
}

// TLNAPI bool TLN_ModPaletteColor (TLN_Palette palette, uint8_t r, uint8_t g, uint8_t b, uint8_t start, uint8_t num);
func TLN_ModPaletteColor(palette TLN_Palette, r CUint8, g CUint8, b CUint8, start CUint8, num CUint8) CBool {
	return C.TLN_ModPaletteColor(palette, r, g, b, start, num)
}

// TLNAPI uint8_t* TLN_GetPaletteData (TLN_Palette palette, int index);
func TLN_GetPaletteData(palette TLN_Palette, index CInt) *CUint8 {
	return C.TLN_GetPaletteData(palette, index)
}

// TLNAPI bool TLN_DeletePalette (TLN_Palette palette);
func TLN_DeletePalette(palette TLN_Palette) CBool {
	return TLN_DeletePalette(palette)
}
// }


// {
// TLNAPI TLN_Bitmap TLN_CreateBitmap (int width, int height, int bpp);
func TLN_CreateBitmap(width CInt, height CInt, bpp CInt) TLN_Bitmap {
	return C.TLN_CreateBitmap(width, height, bpp)
}

// TLNAPI TLN_Bitmap TLN_LoadBitmap (const char* filename);
func TLN_LoadBitmap(filename CString) TLN_Bitmap {
	return C.TLN_LoadBitmap(filename)
}

// TLNAPI TLN_Bitmap TLN_CloneBitmap (TLN_Bitmap src);
func TLN_CloneBitmap(src TLN_Bitmap) TLN_Bitmap {
	return C.TLN_CloneBitmap(src)
}

// TLNAPI uint8_t* TLN_GetBitmapPtr (TLN_Bitmap bitmap, int x, int y);
func TLN_GetBitmapPtr(bitmap TLN_Bitmap, x CInt, y CInt) *CUint8 {
	return C.TLN_GetBitmapPtr(bitmap, x, y)
}

// TLNAPI int TLN_GetBitmapWidth (TLN_Bitmap bitmap);
func TLN_GetBitmapWidth(bitmap TLN_Bitmap) CInt {
	return C.TLN_GetBitmapWidth(bitmap)
}

// TLNAPI int TLN_GetBitmapHeight (TLN_Bitmap bitmap);
func TLN_GetBitmapHeight(bitmap TLN_Bitmap) CInt {
	return C.TLN_GetBitmapHeight(bitmap)
}

// TLNAPI int TLN_GetBitmapDepth (TLN_Bitmap bitmap);
func TLN_GetBitmapDepth(bitmap TLN_Bitmap) CInt {
	return C.TLN_GetBitmapDepth(bitmap)
}

// TLNAPI int TLN_GetBitmapPitch (TLN_Bitmap bitmap);
func TLN_GetBitmapPitch(bitmap TLN_Bitmap) CInt {
	return C.TLN_GetBitmapPitch(bitmap)
}

// TLNAPI TLN_Palette TLN_GetBitmapPalette (TLN_Bitmap bitmap);
func TLN_GetBitmapPalette(bitmap TLN_Bitmap) TLN_Palette {
	return C.TLN_GetBitmapPalette(bitmap)
}

// TLNAPI bool TLN_SetBitmapPalette (TLN_Bitmap bitmap, TLN_Palette palette);
func TLN_SetBitmapPalette(bitmap TLN_Bitmap, palette TLN_Palette) CBool {
	return C.TLN_SetBitmapPalette(bitmap, palette)
}

// TLNAPI bool TLN_DeleteBitmap (TLN_Bitmap bitmap);
func TLN_DeleteBitmap(bitmap TLN_Bitmap) CBool {
	return C.TLN_DeleteBitmap(bitmap)
}
// }


// {
// TLNAPI TLN_ObjectList TLN_CreateObjectList(void);
func TLN_CreateObjectList() TLN_ObjectList {
	return C.TLN_CreateObjectList()
}

// TLNAPI bool TLN_AddTileObjectToList(TLN_ObjectList list, uint16_t id, uint16_t gid, uint16_t flags, int x, int y);
func TLN_AddTileObjectToList(list TLN_ObjectList, id CUint16, gid CUint16, flags CUint16, x CInt, y CInt) CBool {
	return C.TLN_AddTileObjectToList(list, id, gid, flags, x, y)
}

// TLNAPI TLN_ObjectList TLN_LoadObjectList(const char* filename, const char* layername);
func TLN_LoadObjectList(filename CString, layername CString) TLN_ObjectList {
	return C.TLN_LoadObjectList(filename, layername)
}

// TLNAPI TLN_ObjectList TLN_CloneObjectList(TLN_ObjectList src);
func TLN_CloneObjectList(src TLN_ObjectList) TLN_ObjectList {
	return C.TLN_CloneObjectList(src)
}

// TLNAPI int TLN_GetListNumObjects(TLN_ObjectList list);
func TLN_GetListNumObjects(list TLN_ObjectList) CInt {
	return C.TLN_GetListNumObjects(list)
}

// TLNAPI bool TLN_GetListObject(TLN_ObjectList list, TLN_ObjectInfo* info);
func TLN_GetListObject(list TLN_ObjectList, info *TLN_ObjectInfo) CBool {
	return C.TLN_GetListObject(list, info)
}

// TLNAPI bool TLN_DeleteObjectList(TLN_ObjectList list);
func TLN_DeleteObjectList(list TLN_ObjectList) CBool {
	return C.TLN_DeleteObjectList(list)
}
// }


// {
// TLNAPI bool TLN_SetLayer (int nlayer, TLN_Tileset tileset, TLN_Tilemap tilemap);
func TLN_SetLayer (nlayer CInt, tileset TLN_Tileset, tilemap TLN_Tilemap) CBool {
	return C.TLN_SetLayer (nlayer, tileset, tilemap)
}

// TLNAPI bool TLN_SetLayerTilemap(int nlayer, TLN_Tilemap tilemap);
func TLN_SetLayerTilemap(nlayer CInt, tilemap TLN_Tilemap) CBool {
	return C.TLN_SetLayerTilemap(nlayer, tilemap)
}

// TLNAPI bool TLN_SetLayerBitmap(int nlayer, TLN_Bitmap bitmap);
func TLN_SetLayerBitmap(nlayer CInt, bitmap TLN_Bitmap) CBool {
	return C.TLN_SetLayerBitmap(nlayer, bitmap)
}

// TLNAPI bool TLN_SetLayerPalette (int nlayer, TLN_Palette palette);
func TLN_SetLayerPalette (nlayer CInt, palette TLN_Palette) CBool {
	return C.TLN_SetLayerPalette (nlayer, palette)
}

// TLNAPI bool TLN_SetLayerPosition (int nlayer, int hstart, int vstart);
func TLN_SetLayerPosition (nlayer CInt, hstart CInt, vstart CInt) CBool {
	return C.TLN_SetLayerPosition (nlayer, hstart, vstart)
}

// TLNAPI bool TLN_SetLayerScaling (int nlayer, float xfactor, float yfactor);
func TLN_SetLayerScaling (nlayer CInt, xfactor CFloat, yfactor CFloat) CBool {
	return C.TLN_SetLayerScaling (nlayer, xfactor, yfactor)
}

// TLNAPI bool TLN_SetLayerAffineTransform (int nlayer, TLN_Affine *affine);
func TLN_SetLayerAffineTransform (nlayer CInt, affine *TLN_Affine) CBool {
	return C.TLN_SetLayerAffineTransform (nlayer, affine)
}

// TLNAPI bool TLN_SetLayerTransform (int layer, float angle, float dx, float dy, float sx, float sy);
func TLN_SetLayerTransform (layer CInt, angle CFloat, dx CFloat, dy CFloat, sx CFloat, sy CFloat) CBool {
	return C.TLN_SetLayerTransform (layer, angle, dx, dy, sx, sy)
}

// TLNAPI bool TLN_SetLayerPixelMapping (int nlayer, TLN_PixelMap* table);
func TLN_SetLayerPixelMapping (nlayer CInt, table *TLN_PixelMap) CBool {
	return C.TLN_SetLayerPixelMapping (nlayer, table)
}

// TLNAPI bool TLN_SetLayerBlendMode (int nlayer, TLN_Blend mode, uint8_t factor);
func TLN_SetLayerBlendMode (nlayer CInt, mode TLN_Blend, factor CUint8) CBool {
	return C.TLN_SetLayerBlendMode (nlayer, mode, factor)
}

// TLNAPI bool TLN_SetLayerColumnOffset (int nlayer, int* offset);
func TLN_SetLayerColumnOffset (nlayer CInt, offset *CInt) CBool {
	return C.TLN_SetLayerColumnOffset (nlayer, offset)
}

// TLNAPI bool TLN_SetLayerClip (int nlayer, int x1, int y1, int x2, int y2);
func TLN_SetLayerClip (nlayer CInt, x1 CInt, y1 CInt, x2 CInt, y2 CInt) CBool {
	return C.TLN_SetLayerClip (nlayer, x1, y1, x2, y2)
}

// TLNAPI bool TLN_DisableLayerClip (int nlayer);
func TLN_DisableLayerClip (nlayer CInt) CBool {
	return C.TLN_DisableLayerClip (nlayer)
}

// TLNAPI bool TLN_SetLayerMosaic (int nlayer, int width, int height);
func TLN_SetLayerMosaic (nlayer CInt, width CInt, height CInt) CBool {
	return C.TLN_SetLayerMosaic (nlayer, width, height)
}

// TLNAPI bool TLN_DisableLayerMosaic (int nlayer);
func TLN_DisableLayerMosaic (nlayer CInt) CBool {
	return C.TLN_DisableLayerMosaic (nlayer)
}

// TLNAPI bool TLN_ResetLayerMode (int nlayer);
func TLN_ResetLayerMode (nlayer CInt) CBool {
	return C.TLN_ResetLayerMode (nlayer)
}

// TLNAPI bool TLN_SetLayerObjects(int nlayer, TLN_ObjectList objects, TLN_Tileset tileset);
func TLN_SetLayerObjects(nlayer CInt, objects TLN_ObjectList,  tileset TLN_Tileset) CBool {
	return C.TLN_SetLayerObjects(nlayer, objects, tileset)
}

// TLNAPI bool TLN_SetLayerPriority(int nlayer, bool enable);
func TLN_SetLayerPriority(nlayer CInt, enable CBool) CBool {
	return C.TLN_SetLayerPriority(nlayer, enable)
}

// TLNAPI bool TLN_SetLayerParent(int nlayer, int parent);
func TLN_SetLayerParent(nlayer CInt, parent CInt) CBool {
	return C.TLN_SetLayerParent(nlayer, parent)
}

// TLNAPI bool TLN_DisableLayerParent(int nlayer);
func TLN_DisableLayerParent(nlayer CInt) CBool {
	return C.TLN_DisableLayerParent(nlayer)
}

// TLNAPI bool TLN_DisableLayer (int nlayer);
func TLN_DisableLayer (nlayer CInt) CBool {
	return C.TLN_DisableLayer (nlayer)
}

// TLNAPI bool TLN_EnableLayer(int nlayer);
func TLN_EnableLayer(nlayer CInt) CBool {
	return C.TLN_EnableLayer(nlayer)
}

// TLNAPI TLN_LayerType TLN_GetLayerType(int nlayer);
func TLN_GetLayerType(nlayer CInt) TLN_LayerType {
	return C.TLN_GetLayerType(nlayer)
}

// TLNAPI TLN_Palette TLN_GetLayerPalette (int nlayer);
func TLN_GetLayerPalette (nlayer CInt) TLN_Palette {
	return C.TLN_GetLayerPalette (nlayer)
}

// TLNAPI TLN_Tileset TLN_GetLayerTileset(int nlayer);
func TLN_GetLayerTileset(nlayer CInt) TLN_Tileset {
	return C.TLN_GetLayerTileset(nlayer)
}

// TLNAPI TLN_Tilemap TLN_GetLayerTilemap(int nlayer);
func TLN_GetLayerTilemap(nlayer CInt) TLN_Tilemap {
	return C.TLN_GetLayerTilemap(nlayer)
}

// TLNAPI TLN_Bitmap TLN_GetLayerBitmap(int nlayer);
func TLN_GetLayerBitmap(nlayer CInt) TLN_Bitmap {
	return C.TLN_GetLayerBitmap(nlayer)
}

// TLNAPI TLN_ObjectList TLN_GetLayerObjects(int nlayer);
func TLN_GetLayerObjects(nlayer CInt) TLN_ObjectList {
	return C.TLN_GetLayerObjects(nlayer)
}

// TLNAPI bool TLN_GetLayerTile (int nlayer, int x, int y, TLN_TileInfo* info);
func TLN_GetLayerTile (nlayer CInt, x CInt, y CInt, info *TLN_TileInfo) CBool {
	return C.TLN_GetLayerTile (nlayer, x, y, info)
}

// TLNAPI int TLN_GetLayerWidth (int nlayer);
func TLN_GetLayerWidth (nlayer CInt) CInt {
	return C.TLN_GetLayerWidth (nlayer)
}

// TLNAPI int TLN_GetLayerHeight (int nlayer);
func TLN_GetLayerHeight (nlayer CInt) CInt {
	return C.TLN_GetLayerHeight (nlayer)
}

// TLNAPI int TLN_GetLayerX(int nlayer);
func TLN_GetLayerX(nlayer CInt) CInt {
	return C.TLN_GetLayerX(nlayer)
}

// TLNAPI int TLN_GetLayerY(int nlayer);
func TLN_GetLayerY(nlayer CInt) CInt {
	return C.TLN_GetLayerY(nlayer)
}
// }


// {
// TLNAPI bool TLN_ConfigSprite (int nsprite, TLN_Spriteset spriteset, uint32_t flags);
func TLN_ConfigSprite (nsprite CInt, spriteset TLN_Spriteset, flags CUint32) CBool {
	return C.TLN_ConfigSprite (nsprite, spriteset, flags)
}

// TLNAPI bool TLN_SetSpriteSet (int nsprite, TLN_Spriteset spriteset);
func TLN_SetSpriteSet (nsprite CInt, spriteset TLN_Spriteset) CBool {
	return C.TLN_SetSpriteSet (nsprite, spriteset)
}

// TLNAPI bool TLN_SetSpriteFlags (int nsprite, uint32_t flags);
func TLN_SetSpriteFlags (nsprite CInt, flags CUint32) CBool {
	return C.TLN_SetSpriteFlags (nsprite, flags)
}

// TLNAPI bool TLN_EnableSpriteFlag(int nsprite, uint32_t flag, bool enable);
func TLN_EnableSpriteFlag(nsprite CInt, flag CUint32, enable CBool) CBool {
	return C.TLN_EnableSpriteFlag(nsprite, flag, enable)
}

// TLNAPI bool TLN_SetSpritePivot(int nsprite, float px, float py);
func TLN_SetSpritePivot(nsprite CInt, px CFloat, py CFloat) CBool {
	return C.TLN_SetSpritePivot(nsprite, px, py)
}

// TLNAPI bool TLN_SetSpritePosition (int nsprite, int x, int y);
func TLN_SetSpritePosition (nsprite CInt, x CInt, y CInt) CBool {
	return C.TLN_SetSpritePosition (nsprite, x, y)
}

// TLNAPI bool TLN_SetSpritePicture (int nsprite, int entry);
func TLN_SetSpritePicture (nsprite CInt, entry CInt) CBool {
	return C.TLN_SetSpritePicture (nsprite, entry)
}

// TLNAPI bool TLN_SetSpritePalette (int nsprite, TLN_Palette palette);
func TLN_SetSpritePalette (nsprite CInt, palette TLN_Palette) CBool {
	return C.TLN_SetSpritePalette (nsprite, palette)
}

// TLNAPI bool TLN_SetSpriteBlendMode (int nsprite, TLN_Blend mode, uint8_t factor);
func TLN_SetSpriteBlendMode (nsprite CInt, mode TLN_Blend, factor CUint8) CBool {
	return C.TLN_SetSpriteBlendMode (nsprite, mode, factor)
}

// TLNAPI bool TLN_SetSpriteScaling (int nsprite, float sx, float sy);
func TLN_SetSpriteScaling (nsprite CInt, sx CFloat, sy CFloat) CBool {
	return C.TLN_SetSpriteScaling (nsprite, sx, sy)
}

// TLNAPI bool TLN_ResetSpriteScaling (int nsprite);
func TLN_ResetSpriteScaling (nsprite CInt) CBool {
	return C.TLN_ResetSpriteScaling (nsprite)
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
func TLN_GetSpritePicture (nsprite CInt) CInt {
	return C.TLN_GetSpritePicture (nsprite)
}

// TLNAPI int TLN_GetSpriteX(int nsprite);
func TLN_GetSpriteX(nsprite CInt) CInt {
	return C.TLN_GetSpriteX(nsprite)
}

// TLNAPI int TLN_GetSpriteY(int nsprite);
func TLN_GetSpriteY(nsprite CInt) CInt {
	return C.TLN_GetSpriteY(nsprite)
}

// TLNAPI int  TLN_GetAvailableSprite (void);
func TLN_GetAvailableSprite() CInt {
	return C.TLN_GetAvailableSprite();
}

// TLNAPI bool TLN_EnableSpriteCollision (int nsprite, bool enable);
func TLN_EnableSpriteCollision (nsprite CInt, enable CBool) CBool {
	return C.TLN_EnableSpriteCollision (nsprite, enable)
}

// TLNAPI bool TLN_GetSpriteCollision (int nsprite);
func TLN_GetSpriteCollision (nsprite CInt) CBool {
	return C.TLN_GetSpriteCollision (nsprite)
}

// TLNAPI bool TLN_GetSpriteState(int nsprite, TLN_SpriteState* state);
func TLN_GetSpriteState(nsprite CInt, state *TLN_SpriteState) CBool {
	return C.TLN_GetSpriteState(nsprite, state)
}

// TLNAPI bool TLN_SetFirstSprite(int nsprite);
func TLN_SetFirstSprite(nsprite CInt) CBool {
	return C.TLN_SetFirstSprite(nsprite)
}

// TLNAPI bool TLN_SetNextSprite(int nsprite, int next);
func TLN_SetNextSprite(nsprite CInt, next CInt) CBool {
	return C.TLN_SetNextSprite(nsprite, next)
}

// TLNAPI bool TLN_EnableSpriteMasking(int nsprite, bool enable);
func TLN_EnableSpriteMasking(nsprite CInt, enable CBool) CBool {
	return C.TLN_EnableSpriteMasking(nsprite, enable)
}

// TLNAPI void TLN_SetSpritesMaskRegion(int top_line, int bottom_line);
func TLN_SetSpritesMaskRegion(top_line CInt, bottom_line CInt) {
	C.TLN_SetSpritesMaskRegion(top_line, bottom_line)
}

// TLNAPI bool TLN_SetSpriteAnimation (int nsprite, TLN_Sequence sequence, int loop);
func TLN_SetSpriteAnimation (nsprite CInt, sequence TLN_Sequence, loop CInt) CBool {
	return C.TLN_SetSpriteAnimation (nsprite, sequence, loop)
}

// TLNAPI bool TLN_DisableSpriteAnimation(int nsprite);
func TLN_DisableSpriteAnimation(nsprite CInt) CBool {
	return C.TLN_DisableSpriteAnimation(nsprite)
}

// TLNAPI bool TLN_PauseSpriteAnimation(int index);
func TLN_PauseSpriteAnimation(index CInt) CBool {
	return C.TLN_PauseSpriteAnimation(index)
}

// TLNAPI bool TLN_ResumeSpriteAnimation(int index);
func TLN_ResumeSpriteAnimation(index CInt) CBool {
	return C.TLN_ResumeSpriteAnimation(index)
}

// TLNAPI bool TLN_DisableAnimation(int index);
func TLN_DisableAnimation(index CInt) CBool {
	return C.TLN_DisableAnimation(index)
}

// TLNAPI bool TLN_DisableSprite (int nsprite);
func TLN_DisableSprite (nsprite CInt) CBool {
	return C.TLN_DisableSprite (nsprite)
}

// TLNAPI TLN_Palette TLN_GetSpritePalette (int nsprite);
func TLN_GetSpritePalette (nsprite CInt) TLN_Palette {
	return C.TLN_GetSpritePalette (nsprite)
}
// }


// {
// TLNAPI TLN_Sequence TLN_CreateSequence (const char* name, int target, int num_frames, TLN_SequenceFrame* frames);
func TLN_CreateSequence(name CString, target CInt, num_frames CInt,  frames *TLN_SequenceFrame) TLN_Sequence {
	return C.TLN_CreateSequence(name, target, num_frames, frames)
}

// TLNAPI TLN_Sequence TLN_CreateCycle (const char* name, int num_strips, TLN_ColorStrip* strips);
func TLN_CreateCycle(name CString, num_strips CInt, strips *TLN_ColorStrip) TLN_Sequence {
	return C.TLN_CreateCycle(name, num_strips, strips)
}

// TLNAPI TLN_Sequence TLN_CreateSpriteSequence(const char* name, TLN_Spriteset spriteset, const char* basename, int delay);
func TLN_CreateSpriteSequence(name CString, spriteset TLN_Spriteset, basename CString, delay CInt) TLN_Sequence {
	return C.TLN_CreateSpriteSequence(name, spriteset, basename, delay)
}

// TLNAPI TLN_Sequence TLN_CloneSequence (TLN_Sequence src);
func TLN_CloneSequence(src TLN_Sequence) TLN_Sequence {
	return C.TLN_CloneSequence(src)
}

// TLNAPI bool TLN_GetSequenceInfo (TLN_Sequence sequence, TLN_SequenceInfo* info);
func TLN_GetSequenceInfo(sequence TLN_Sequence, info *TLN_SequenceInfo) CBool {
	return C.TLN_GetSequenceInfo(sequence, info)
}

// TLNAPI bool TLN_DeleteSequence (TLN_Sequence sequence);
func TLN_DeleteSequence(sequence TLN_Sequence) CBool {
	return C.TLN_DeleteSequence(sequence)
}
// }


// {
// TLNAPI TLN_SequencePack TLN_CreateSequencePack (void);
func TLN_CreateSequencePack() TLN_SequencePack {
	return C.TLN_CreateSequencePack()
}

// TLNAPI TLN_SequencePack TLN_LoadSequencePack (const char* filename);
func TLN_LoadSequencePack(filename CString) TLN_SequencePack {
	return C.TLN_LoadSequencePack(filename)
}

// TLNAPI TLN_Sequence TLN_GetSequence (TLN_SequencePack sp, int index);
func TLN_GetSequence(sp TLN_SequencePack, index CInt) TLN_Sequence {
	return C.TLN_GetSequence(sp, index)
}

// TLNAPI TLN_Sequence TLN_FindSequence (TLN_SequencePack sp, const char* name);
func TLN_FindSequence(sp TLN_SequencePack, name CString) TLN_Sequence {
	return C.TLN_FindSequence(sp, name)
}

// TLNAPI int TLN_GetSequencePackCount (TLN_SequencePack sp);
func TLN_GetSequencePackCount(sp TLN_SequencePack) CInt {
	return C.TLN_GetSequencePackCount(sp)
}

// TLNAPI bool TLN_AddSequenceToPack (TLN_SequencePack sp, TLN_Sequence sequence);
func TLN_AddSequenceToPack(sp TLN_SequencePack, sequence TLN_Sequence) CBool {
	return C.TLN_AddSequenceToPack(sp, sequence)
}

// TLNAPI bool TLN_DeleteSequencePack (TLN_SequencePack sp);
func TLN_DeleteSequencePack(sp TLN_SequencePack) CBool {
	return C.TLN_DeleteSequencePack(sp)
}
// }


// {
// TLNAPI bool TLN_SetPaletteAnimation (int index, TLN_Palette palette, TLN_Sequence sequence, bool blend);
func TLN_SetPaletteAnimation(index CInt, palette TLN_Palette, sequence TLN_Sequence, blend CBool) CBool {
	return C.TLN_SetPaletteAnimation(index, palette, sequence, blend)
}

// TLNAPI bool TLN_SetPaletteAnimationSource (int index, TLN_Palette);
func TLN_SetPaletteAnimationSource(index CInt, palette TLN_Palette) CBool {
	return C.TLN_SetPaletteAnimationSource(index, palette)
}

// TLNAPI bool TLN_GetAnimationState (int index);
func TLN_GetAnimationState(index CInt) CBool {
	return C.TLN_GetAnimationState(index)
}

// TLNAPI bool TLN_SetAnimationDelay (int index, int frame, int delay);
func TLN_SetAnimationDelay(index CInt, frame CInt, delay CInt) CBool {
	return C.TLN_SetAnimationDelay(index, frame, delay)
}

// TLNAPI int  TLN_GetAvailableAnimation (void);
func TLN_GetAvailableAnimation() CInt {
	return C.TLN_GetAvailableAnimation()
}

// TLNAPI bool TLN_DisablePaletteAnimation(int index);
func TLN_DisablePaletteAnimation(index CInt) CBool {
	return C.TLN_DisablePaletteAnimation(index)
}
// }


// {
// TLNAPI bool TLN_LoadWorld(const char* tmxfile, int first_layer);
func TLN_LoadWorld(tmxfile CString, first_layer CInt) CBool {
	return C.TLN_LoadWorld(tmxfile, first_layer)
}

// TLNAPI void TLN_SetWorldPosition(int x, int y);
func TLN_SetWorldPosition(x CInt, y CInt) {
	C.TLN_SetWorldPosition(x, y)
}

// TLNAPI bool TLN_SetLayerParallaxFactor(int nlayer, float x, float y);
func TLN_SetLayerParallaxFactor(nlayer CInt, x CFloat, y CFloat) CBool {
	return C.TLN_SetLayerParallaxFactor(nlayer, x, y)
}

// TLNAPI bool TLN_SetSpriteWorldPosition(int nsprite, int x, int y);
func TLN_SetSpriteWorldPosition(nsprite CInt, x CInt, y CInt) CBool {
	return C.TLN_SetSpriteWorldPosition(nsprite, x, y)
}

// TLNAPI void TLN_ReleaseWorld(void);
func TLN_ReleaseWorld() {
	C.TLN_ReleaseWorld()
}
// }
