package ggk

/** \class Shader
 *
 *  Shaders specify the source color(s) for what is being drawn. If a paint
 *  has no shader, then the paint's color is used. If the paint has a
 *  shader, then the shader's color(s) are use instead, but they are
 *  modulated by the paint's alpha. This makes it easy to create a shader
 *  once (e.g. bitmap tiling or gradient) and then change its transparency
 *  w/o having to modify the original shader... only the paint's alpha needs
 *  to be modified.
 */
type Shader struct {
}

func NewShader_Color(color Color) *Shader {
	toimpl()
	return nil
}

func (shader *Shader) MakeWithColorFilter(filter *ColorFilter) *Shader {
	toimpl()
	return nil
}

func (shader *Shader) ContextSize(rec *ShaderContextRec) int {
	toimpl()
	return 0
}

func (shader *Shader) CreateContext(rec *ShaderContextRec) *ShaderContext {
	toimpl()
	return nil
}


type ShaderContextRec struct {
}

type ShaderDstType int

const (
	KShaderDstTypePMColor = ShaderDstType(iota) //< clients prefer shading into PMColor dest.
	KSahderDstTypePM4f                          //< clients prefer shading into PM4f dest.
)

func NewShaderContextRec(paint *Paint, matrix *Matrix, localM *Matrix, dstType ShaderDstType) *ShaderContextRec {
	toimpl()
	return nil
}

func BlitterPreferredShaderDest(dstInfo *ImageInfo) ShaderDstType {
	toimpl()
	return KShaderDstTypePMColor
}

type ShaderContext struct {
}
