package ggk

/** Blitter
blitter and its subclasses are responsible for actually writing pixels
into memory. Besides efficiency, they handle clipping and antialiasing.
A SkBlitter subclass contains all the context needed to generate pixels
for the destination and how src/generated pixels map to the destination.
The coordinates passed to the blitX calls are in destination pixel space. */
type Blitter interface {
	/** Blit a horizontal run of one or more pixels. */
	BlitH(x, y, width int)

	/** BlitAntiH
	Blit a horizontal run of antialiased pixels; runs[] is a *sparse*
	zero-terminated run-length encoding of spans of constant alpha values.
	The runs[] and antialias[] work together to represent long runs of pixels with the same
	alphas. The runs[] contains the number of pixels with the same alpha, and antialias[]
	contain the coverage value for that number of pixels. The runs[] (and antialias[]) are
	encoded in a clever way. The runs array is zero terminated, and has enough entries for
	each pixel plus one, in most cases some of the entries will not contain valid data. An entry
	in the runs array contains the number of pixels (np) that have the same alpha value. The
	next np value is found np entries away. For example, if runs[0] = 7, then the next valid
	entry will by at runs[7]. The runs array and antialias[] are coupled by index. So, if the
	np entry is at runs[45] = 12 then the alpha value can be found at antialias[45] = 0x88.
	This would mean to use an alpha value of 0x88 for the next 12 pixels starting at pixel 45. */
	BlitAntiH(x, y int, antialias []Alpha, runs []int16)

	/** BlitV blit a vertical run of pixels with a constant alpha value. */
	BlitV(x, y, height int, alpha Alpha)

	/** BlitRect blit a solid rectangle one or more pixels wide. */
	BlitRect(x, y, width, height int)

	/** BlitAntiRect
	blit a rectangle with one alpha-blended column on the left,
	width (zero or more) opaque pixels, and one alpha-blended column
	on the right.
	The result will always be at least two pixels wide. */
	BlitAntiRect(x, y, width, height, leftAlpha Alpha, rightAlpha Alpha)

	/** BlitMask
	Blit a pattern of pixels defined by a rectangle-clipped mask;
	typically used for text. */
	BlitMask(mask *Mask, clip Rect)

	/** JustAnOpaqueColor
	If the blitter just sets a single value for each pixel, return the
	bitmap it draws into, and assign value. If not, return nullptr and ignore
	the value parameter. */
	JustAnOpaqueColor(value uint32) *Pixmap

	/** (x, y), (x + 1, y) */
	BlitAntiH2(x, y int, a0, a1 uint8)

	/** (x, y), (x, y + 1) */
	BlitAntiV2(x, y int, a0, a1 uint8)

	/** IsNullBlitter
	Special method just to identify the null blitter, which is returned
	from Choose() if the request cannot be fulfilled. Default impl
	returns false. */
	IsNullBlitter() bool

	/** ResetShaderContext
	Special methods for SkShaderBlitter. On all other classes this is a no-op. */
	ResetShaderContext(context *tShaderContextRec) bool

	GetShaderContext() *ShaderContext

	/** RequestRowsPreserved
	Special methods for blitters that can blit more than one row at a time.
	This function returns the number of rows that this blitter could optimally
	process at a time. It is still required to support blitting one scanline
	at a time. */
	RequestRowsPreserved() int

	/** AllocBlitMemory
	This function allocates memory for the blitter that the blitter then owns.
	The memory can be used by the calling function at will, but it will be
	released when the blitter's destructor is called. This function returns
	nullptr if no persistent memory is needed by the blitter. */
	AllocBlitMemory(sz int)
}

/** BlitterClipper
Factory to set up the appropriate most-efficient wrapper blitter
to apply a clip. Returns a pointer to a member, so lifetime must
be managed carefully. */
type BlitterClipper struct {
	// toimpl
}

func (*BlitterClipper)apply(blitter Blitter, clip *Region, bounds Rect) {
	toimpl()
}

func BlitterChoose(device *Pixmap, matrix *Matrix, origPaint *Paint, drawCoverage bool) Blitter {
	toimpl()
	return nil
}