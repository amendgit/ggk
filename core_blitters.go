package ggk

/** Returns nullptr if no SkRasterPipeline blitter can be constructed for this paint. */
func CreateRasterPipelineBlitter(dst *Pixmap, paint *Paint) Blitter {
	return NewRasterPipelineBlitter(dst, paint)
}
