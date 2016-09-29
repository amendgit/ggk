package ggk

import (
	"errors"
	"sync/atomic"
)

type Bitmap struct {
	rowBytes int
	flags    uint8

	info       *ImageInfo
	colorTable *ColorTable

	pixels         *Pixels
	pixelOrigin    Point
	pixelLockCount int32
}

// Swap the fields of the two bitmaps. This routine is guaranteed to never fail or throw.
func (bmp *Bitmap) Swap(otr *Bitmap) {
	*bmp, *otr = *otr, *bmp
}

func (bmp *Bitmap) Info() *ImageInfo {
	return bmp.info
}

func (bmp *Bitmap) Width() Scalar {
	return bmp.info.Width()
}

func (bmp *Bitmap) Height() Scalar {
	return bmp.info.Height()
}

func (bmp *Bitmap) ColorType() ColorType {
	return bmp.info.colorType
}

func (bmp *Bitmap) AlphaType() AlphaType {
	return bmp.info.alphaType
}

func (bmp *Bitmap) ProfileType() ColorProfileType {
	return bmp.info.profileType
}

// Return the number of bytes per pixel based on the colortype. If the colortype is
// KColorTypeUnknown, then 0 is returend.
func (bmp *Bitmap) BytesPerPixel() int {
	return bmp.info.BytesPerPixel()
}

// Return the rowBytes expressed as a number of pixels (like width and height).
// If the colortype is KColorTypeUnknown, then 0 is returend.
func (bmp *Bitmap) RowBytesAsPixels() int {
	return bmp.rowBytes >> uint(bmp.ShiftPerPixel())
}

// Return the shift amount per pixel (i.e. 0 for 1-byte per pixel, 1 for 2-bytes per pixel
// colortypes. 2 for 4-bytes per pixel colortypes). Returns 0 for ColorType_Unknown.
func (bmp *Bitmap) ShiftPerPixel() int {
	return bmp.BytesPerPixel() >> 1
}

// IsEmpty returns true iff the bitmap has empty dimensions.
// Hey!  Before you use this, see if you really want to know DrawNothing() intead.
func (bmp *Bitmap) IsEmpty() bool {
	return bmp.info.IsEmpty()
}

// Return true iff the bitmap has no pixelref. Note: this can return true even if the
// dimensions of the bitmap are > 0 (see IsEmpty()).
// Hey! Before you use this, see if you really want to know DrawNothing() intead.
func (bmp *Bitmap) IsNull() bool {
	return bmp.pixels == nil
}

var ErrBitmapIsNotValid = errors.New(`error: bitmap is not valid.`)

// IsValid return true iff the bitmap has valid imageInfo, pixels and colorTable
func (bmp *Bitmap) IsValid() bool {
	if !bmp.info.IsValid() {
		return false
	}
	if !bmp.info.ValidRowBytes(bmp.rowBytes) {
		return false
	}
	if bmp.info.ColorType() == KColorTypeRGB565 &&
		bmp.info.AlphaType() != KAlphaTypeOpaque {
		return false
	}
	// TOIMPL
	if bmp.pixels != nil {
		if bmp.pixelLockCount <= 0 &&
			//    !bmp.pixels.IsLock() &&
			bmp.rowBytes < bmp.info.MinRowBytes() &&
			bmp.pixelOrigin.X < 0 &&
			bmp.pixelOrigin.Y < 0 &&
			bmp.info.Width() < bmp.Width()+bmp.pixelOrigin.X &&
			bmp.info.Height() < bmp.Height()+bmp.pixelOrigin.Y {
			return false
		}
	} else {
		if bmp.colorTable != nil {
			return false
		}
	}
	return true
}

// Return true iff drawing the bitmap has no effect.
func (bmp *Bitmap) DrawNothing() bool {
	return bmp.IsEmpty() || bmp.IsNil()
}

// Return the number of bytes between subsequent rows of the bitmap.
func (bmp *Bitmap) RowBytes() int {
	return bmp.rowBytes
}

// Set the bitmap's alphaType, returing true on success. If false is
// returned, then the specified new alphaType is incompatible with the
// colortype, and the current alphaType is unchanged.
//
// Note: this changes the alpahType for the underlying types, which means
// that all bitmaps that might be sharing (subsets of) the pixels will
// be affected.
func (bmp *Bitmap) SetAlphaType(alphaType AlphaType) bool {
	alphaType, err := bmp.info.colorType.ValidateAlphaType(alphaType)
	if err != nil {
		return false
	}
	if bmp.info.alphaType != alphaType {
		bmp.info.SetAlphaType(alphaType)
	}
	return true
}

// Return the address of the pixels for this Bitmap.
func (bmp *Bitmap) Pixels() *Pixels {
	return bmp.pixels
}

// Return the bytes of the pixels for this bitmap.
func (bmp *Bitmap) PixelsBytes() []byte {
	bmp.pixels.LockPixels()
	var bytes = bmp.pixels.Bytes()
	return bytes
}

// Return the byte size of the pixels, based on the height and rowBytes.
// Note this truncates the result to 32bits. Call Size64() to detect 
// if the real size exceeds 32bit.
func (bmp *Bitmap) Size() int32 {
	toimpl()
	return 0
}

// Return the number of bytes from the pointer returned by getPixels()
// to the end of the allocated space in the buffer. Required in
// cases where extractSubset has been called.
func (bmp *Bitmap) SafeSize() int32 {
	toimpl()
	return 0
}

// Return the full size of the bitmap, in bytes.
func (bmp *Bitmap) ComputeSize64() int64 {
	toimpl()
	return 0
}

// Return the number of bytes from the pointer returned by getPixels()
// to the end of the allocated space in the buffer. This may be smaller
// than computeSize64() if there is any rowbytes padding beyond the width.
func (bmp *Bitmap) ComputeSafeSize64() int64 {
	toimpl()
	return 0
}

// Returns true if this bitmap is marked as immutable, meaning that the
// contents of its pixels will not change for the lifetime of the bitmap.
func (bmp *Bitmap) IsImmutable() bool {
	toimpl()
	return false
}

// Marks this bitmap as immutable, meaning that the contents of its
// pixels will not change for the lifetime of the bitmap and of the
// underlying pixelref. This state can be set, but it cannot be
// cleared once it is set. This state propagates to all other bitmaps
// that share the same pixelref.
func (bmp *Bitmap) SetIsImmutable() {
	toimpl()
}

// Return true if the bitmap is opaque (has no translucent/transparent pixels)
func (bmp *Bitmap) IsOpaque() bool {
	toimpl()
	return false
}

// Returns true if the bitmap is volatile (i.e. should not be cached by devices.)
func (bmp *Bitmap) IsVolatile() bool {
	toimpl()
	return false
}

// Specify whether this bitmap is volatile. Bitmaps are not volatile by
// default. Temporary bitmaps that are discarded after use should be
// marked as volatile. This provides a hint to the device that the bitmap
// should not be cached. Providing this hint when appropriate can
// improve performance by avoiding unnecessary overhead and resource
// consumption on the device.
func (bmp *Bitmap) SetIsVolatile(isVolatile bool) {
	toimpl()
	return false
}

// Reset the bitmap to its initial state (see default constructor). If we are a (shared)
// owner of the pixels, that ownership is decremented.
func (bmp *Bitmap) Reset() {
	bmp.freePixels()
	var zero Bitmap
	*bmp = zero
}

// This will brute-force return true if all of the pixels in the bitmap
// are opaque. If it fails to read the pixels, or encounters an error,
// it will return false.
//
// Since this can be an expensive operation, the bitmap stores a flag for
// this (isOpaque). Only call this if you need to compute this value from
// "unknown" pixels.
func (bmp *Bitmap) ComputeIsOpaque() bool {
	toimpl()
	return false
}

// Return the bitmap's bounds [0, 0, width, height] as an Rect.
func (bmp *Bitmap) Bounds() Rect {
	var (
		x      = bmp.pixelOrigin.X
		y      = bmp.pixelOrigin.Y
		width  = bmp.info.Width()
		height = bmp.info.Height()
	)
	return MakeRect(x, y, width, height)
}

func (bmp *Bitmap) Dimensions() Size {
	toimpl()
	return SizeZero
}

// Returns the bounds of this bitmap, offset by its pixelref origin.
func (bmp *Bitmap) Subset() Rect {
	toimpl()
	return RectZero
}

func (bmp *Bitmap) SetInfo(imageInfo ImageInfo, rowBytes int) bool {
	alphaType, err := imageInfo.ColorType().ValidateAlphaType(imageInfo.AlphaType())
	if err != nil {
		bmp.Reset()
		return false
	}
	// alphaType is the real value.
	var minRowBytes int64 = imageInfo.MinRowBytes64()
	if int64(int32(minRowBytes)) != minRowBytes {
		bmp.Reset()
		return false
	}
	if imageInfo.Width() < 0 || imageInfo.Height() < 0 {
		bmp.Reset()
		return false
	}
	if imageInfo.ColorType() == KColorTypeUnknown {
		rowBytes = 0
	} else if rowBytes == 0 {
		rowBytes = int(minRowBytes)
	} else if !imageInfo.ValidRowBytes(rowBytes) {
		bmp.Reset()
		return false
	}
	bmp.freePixels()
	bmp.info = imageInfo.MakeAlphaType(alphaType)
	bmp.rowBytes = rowBytes
	return true
}

 // Allocate the bitmap's pixels to match the requested image info. If the Factory
 // is non-null, call it to allcoate the pixelref. If the ImageInfo requires
 // a colortable, then ColorTable must be non-null, and will be ref'd.
 // On failure, the bitmap will be set to empty and return false.
func (bmp *Bitmap) TryAllocPixels(info *ImageInfo, factory *PixelsRefFactory, ct *ColorTable) bool {
	toimpl()
	return false
}

var ErrAllocPixels = errors.New(`ERROR: bad imageInfo, rowBytes. or allocate failed`)

func (bmp *Bitmap) AllocPixels(requestedInfo ImageInfo, rowBytes int) error {
	if requestedInfo.ColorType() == KColorTypeIndex8 {
		bmp.Reset()
		return ErrAllocPixels
	}
	if !bmp.SetInfo(requestedInfo, rowBytes) {
		bmp.Reset()
		return ErrAllocPixels
	}
	// SetInfo may have corrected info (e.g. 565 is always opaque).
	var correctedInfo = bmp.Info()
	// SetInfo may have computed a valid rowBytes if 0 were passed in
	rowBytes = bmp.RowBytes()
	// Allocate memories.
	var pixels = NewMemoryPixelsAlloc(correctedInfo, rowBytes)
	if pixels == nil {
		bmp.Reset()
		return ErrAllocPixels
	}
	bmp.pixels = pixels.Pixels
	if bmp.LockPixels() != nil {
		bmp.Reset()
		return ErrAllocPixels
	}
	return ErrAllocPixels
}

// Install a pixelref that wraps the specified pixels and rowBytes, and
// optional ReleaseProc and context. When the pixels are no longer
// referenced, if releaseProc is not null, it will be called with the
// pixels and context as parameters.
// On failure, the bitmap will be set to empty and return false.
//
// If specified, the releaseProc will always be called, even on failure. It is also possible
// for success but the releaseProc is immediately called (e.g. valid Info but NULL pixels).
// func (bmp *Bitmap) InstallPixels(requestedInfo ImageInfo, pixelsBytes []byte, rowbytes int, ct *ColorTable, 
// 	releaseProc ReleaseProc, context *interface{}) bool {
// 	toimpl()
// 	return false
// }

// Call installPixels with no ReleaseProc specified. This means that the
// caller must ensure that the specified pixels are valid for the lifetime
// of the created bitmap (and its pixelRef).
func (bmp *Bitmap) InstallPixels(requestedInfo ImageInfo, pixelsBytes []byte, rowBytes int, ct *ColorTable) bool {
	if !bmp.SetInfo(requestedInfo, rowBytes) {
		// release pixels
		bmp.Reset()
		return false
	}
	if pixelsBytes == nil {
		// release pixels
		return true // we behaved as if they called setInfo()
	}
	var pixels = NewMemoryPixelsDirect(pixelsBytes)
	if pixels == nil {
		bmp.Reset()
		return false
	}
	bmp.pixels = pixels.Pixels
	// since we're already allocated, we LockPixels right away.
	bmp.LockPixels()
	if !bmp.IsValid() {
		// 	log.Printf(`xyz`)
	}
	return true
}

// Assign a pixels and origin to the bitmap. Pixels are reference.
// so the existing one (if any) will be unref'd and the new one will be
// ref'd. (x,y) specify the offset within the pixelRef's pixels for the
// top/left corner of the bitmap. For a bitmap that encompass the entire
// pixels of the pixel ref, these will be (0,0).
func (bmp *Bitmap) SetPixels(pixels *Pixels, origin Point) {
	toimpl()
}

// Return the current pixelref object or NULL if there is none. This does
// not affect the refcount of the pixelref.
func (bmp *Bitmap) PixelRef() *PixelRef {
	toimpl()
	return nil
}

// A bitmap can reference a subset of a pixelref's pixels. That means the
// bitmap's width/height can be <= the dimensions of the pixelref. The
// pixelref origin is the x,y location within the pixelref's pixels for
// the bitmap's top/left corner. To be valid the following must be true:
//
// origin_x + bitmap_width  <= pixelref_width
// origin_y + bitmap_height <= pixelref_height
//
// PixelRefOrigin() returns this origin, or (0,0) if there is no pixelRef.
func (bmp *Bitmap) PixelRefOrigin() Point {
	toimpl()
	return PointZero
}

// Assign a pixelref and origin to the bitmap. Pixelrefs are reference,
// so the existing one (if any) will be unref'd and the new one will be
// ref'd. (x,y) specify the offset within the pixelref's pixels for the
// top/left corner of the bitmap. For a bitmap that encompases the entire
// pixels of the pixelref, these will be (0,0).
func (bmp *Bitmap) SetPixelRef(pr *PixelRef, dx, dy int) *PixelRef {
	toimpl()
	return nil
}

// Call this to ensure that the bitmap points to the current pixel address
// in the pixels. Balance it with a call to UnlockPixels(). These calls
// are harmless if there is no pixelRef.
func (bmp *Bitmap) LockPixels() error {
	if bmp.pixels != nil && atomic.AddInt32(&bmp.pixelLockCount, 1) == 1 {
		bmp.pixels.LockPixels()
	}
	return nil
}

// When you are finished access the pixel memory, call this to balance a
// previous call to LockPixels(). This allows pixelRefs that implement
// cached/deferred image decoding to know when there are active clients of
// a given image.
func (bmp *Bitmap) UnlockPixels() error {
	if bmp.pixels != nil && atomic.AddInt32(&bmp.pixelLockCount, -1) == 0 {
		bmp.pixels.UnlockPixels()
	}
	return nil
}

func (bmp *Bitmap) requestLock(result *AutoPixmapLock) bool {
	toimpl()
	return false
}

// Call this to be sure that the bitmap is valid enough to be drawn (i.e.
// it has non-null pixels, and if required by its colortype, it has a
// non-null colortable. Returns true if all of the above are met.
func (bmp *Bitmap) ReadyToDraw() {
	toimpl()
}

// Unreference any pixels or colorTables.
func (bmp *Bitmap) FreePixels() {
	if bmp.pixels != nil {
		if bmp.pixelLockCount > 0 {
			bmp.UnlockPixels()
		}
		bmp.pixels = nil
		bmp.pixelOrigin = PointZero
	}
	bmp.pixelLockCount = 0
	bmp.pixels = nil
	bmp.colorTable = nil
}

// Return the bitmap's colortable, if it uses one (i.e. colorType is
// Index_8) and the pixels are locked.
// Otherwise returns NULL. Does not affect the colortable's
// reference count.
func (bmp *Bitmap) ColorTable() *ColorTable {
	toimpl()
	return nil
}

// Returns a non-zero, unique value corresponding to the pixels in our
// pixelref. Each time the pixels are changed (and notifyPixelsChanged
// is called), a different generation ID will be returned. Finally, if
// there is no pixelRef then zero is returned.
func (bmp *Bitmap) GenerationID() uint32 {
	toimpl()
	return 0
}

// Call this if you have changed the contents of the pixels. This will in-
// turn cause a different generation ID value to be returned from
// getGenerationID().
func (bmp *Bitmap) NotifyPixelsChanged() {
	toimpl()
}

// Fill the entire bitmap with the specified color.
// If the bitmap's colortype does not support alpha (e.g. 565) then the alpha
// of the color is ignored (treated as opaque). If the colortype only supports
// alpha (e.g. A1 or A8) then the color's r,g,b components are ignored.
func (bmp *Bitmap) EraseColor(c Color) {
	toimpl()
}

// Fill the entire bitmap with the specified color.
// If the bitmap's colortype does not support alpha (e.g. 565) then the alpha
// of the color is ignored (treated as opaque). If the colortype only supports
// alpha (e.g. A1 or A8) then the color's r,g,b components are ignored.
func (bmp *Bitmap) EraseARGB(a, r, g, b uint8) {
	toimpl()
}

// Fill the specified area of this bitmap with the specified color.
// If the bitmap's colortype does not support alpha (e.g. 565) then the alpha
// of the color is ignored (treated as opaque). If the colortype only supports
// alpha (e.g. A1 or A8) then the color's r,g,b components are ignored.
func (bmp *Bitmap) EraseRect(c Color, area Rect) {
	toimpl()
}

// Return the SkColor of the specified pixel.  In most cases this will
// require un-premultiplying the color.  Alpha only colortypes (e.g. kAlpha_8_SkColorType)
// return black with the appropriate alpha set.  The value is undefined
// for kUnknown_SkColorType or if x or y are out of bounds, or if the bitmap
// does not have any pixels (or has not be locked with lockPixels()).
func (bmp *Bitmap) ColorAt(x, y) Color {
	return Color(0, 0, 0, 0)
}