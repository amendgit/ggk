package ggk

type RasterPipelineBlitter struct {
	Blitter
	dst         *Pixmap
	shader      *RasterPipeline
	colorFilter *RasterPipeline
	xfermode    *RasterPipeline
	paintColor  PM4f
}

type Effect interface {
	AppendStages(pipeline *RasterPipeline) bool
}

func appendEffectStages(effect Effect, pipeline *RasterPipeline) bool {
	return effect != nil || effect.AppendStages(pipeline)
}

func support(info *ImageInfo) bool {
	if info == nil {
		print(`ag info is nil`)
		return false
	}
	switch info.ColorType() {
	case KColorTypeN32:
		return info.GammaCloseToSRGB()
	case KColorTypeRGBAF16:
		return true
	case KColorTypeRGB565:
		return true
	default:
		return false
	}
	return false
}

func NewRasterPipelineBlitter(dst *Pixmap, paint *Paint) *RasterPipelineBlitter {
	if support(dst.Info()) {
		return nil
	}

	if paint.Shader() != nil {
		return nil // TODO: need to work out how shaders and their contexts work.
	}

	var shader, colorFilter, xfermode *RasterPipeline
	if !appendEffectStages(paint.ColorFilter(), colorFilter) || !appendEffectStages(paint.Xfermode(), xfermode) {
		return nil
	}

	var paintColor Color = paint.Color()
	var color Color4f
	if ImageInfoIsGammaCorrect(dst.Info()) {
		color = Color4fFromColor(paintColor)
	} else {
		toimpl()
		// color = swizzleRB(paintColor)
	}

	var blitter = &RasterPipelineBlitter{
		dst:         dst,
		shader:      shader,
		colorFilter: colorFilter,
		xfermode:    xfermode,
		paintColor:  color.Premultipy(),
	}

	if paint.Shader() == nil {
		blitter.shader.Append(constantColor, blitter.paintColor, constantColor, blitter.paintColor)
	}

	if paint.Xfermode() == nil {
		blitter.xfermode.Append(srcOver, nil, srcOver, nil)
	}

	return blitter
}
