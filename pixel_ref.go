package ggk

type PixelRefImpl interface {
	/**
     *  On success, returns true and fills out the LockRec for the pixels. On
     *  failure returns false and ignores the LockRec parameter.
     *
     *  The caller will have already acquired a mutex for thread safety, so this
     *  method need not do that.
     */
	//OnNewLockPixels(rec *LockRec) bool

	/**
     *  Balancing the previous successful call to onNewLockPixels. The locked
     *  pixel address will no longer be referenced, so the subclass is free to
     *  move or discard that memory.
     *
     *  The caller will have already acquired a mutex for thread safety, so this
     *  method need not do that.
     */
	OnUnlockPixels();

	/** Default impl returns true */
	OnLockPixelsAreWritable() bool;

	/**
	 *  For pixelrefs that don't have access to their raw pixels, they may be
	 *  able to make a copy of them (e.g. if the pixels are on the GPU).
	 *
	 *  The base class implementation returns false;
	 */
	OnReadPixels(dst *Bitmap, colorType ColorType, subsetOrNull Rect) bool

	// default impl returns NULL.
	//OnRefEncodedData() *Data

	// default impl does nothing.
	OnNotifyPixelsChanged()

	//OnQueryYUV8(SkYUVSizeInfo*, SkYUVColorSpace*) bool
	//OnGetYUV8Planes(const SkYUVSizeInfo&, void*[3] /*planes*/) bool

	/**
	 *  Returns the size (in bytes) of the internally allocated memory.
	 *  This should be implemented in all serializable SkPixelRef derived classes.
	 *  SkBitmap::fPixelRefOffset + SkBitmap::getSafeSize() should never overflow this value,
	 *  otherwise the rendering code may attempt to read memory out of bounds.
	 *
	 *  @return default impl returns 0.
	 */
	GetAllocatedSizeInBytes() int;

	//OnRequestLock(LockRequest, *LockResult) bool

	OnIsLazyGenerated() bool
}

/** \class PixelRef

    This class is the smart container for pixel memory, and is used with
    SkBitmap. A pixelref is installed into a bitmap, and then the bitmap can
    access the actual pixel memory by calling lockPixels/unlockPixels.

    This class can be shared/accessed between multiple threads.
*/
type PixelRef struct {
	impl PixelRefImpl
}

type PixelRefFactory interface {
	/**
     *  Allocate a new pixelref matching the specified ImageInfo, allocating
     *  the memory for the pixels. If the ImageInfo requires a ColorTable,
     *  the pixelref will ref() the colortable.
     *  On failure return NULL.
     */
	Create(imageInfo *ImageInfo, rowBytes int, colorTable *ColorTable) PixelRef
}