package ggk

type RasterPipelineBlitter struct {
	Blitter
	dst         *Pixmap
	shader      *RasterPipeline
	colorFilter *RasterPipeline
	xfermode    *RasterPipeline
	paintColor  PM4f
}

func NewRasterPipelineBlitter(dst *Pixmap, paint *Paint) *RasterPipelineBlitter {
	toimpl()
	return &RasterPipelineBlitter{}
}