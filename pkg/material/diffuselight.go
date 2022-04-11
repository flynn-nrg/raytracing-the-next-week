package material

import (
	"github.com/flynn-nrg/raytracing-the-next-week/pkg/hitrecord"
	"github.com/flynn-nrg/raytracing-the-next-week/pkg/ray"
	"github.com/flynn-nrg/raytracing-the-next-week/pkg/texture"
	"github.com/flynn-nrg/raytracing-the-next-week/pkg/vec3"
)

// Ensure interface compliance.
var _ Material = (*DiffuseLight)(nil)

// DiffuseLight represents a diffuse light material.
type DiffuseLight struct {
	emit texture.Texture
}

// NewDiffuseLight returns an instance of a diffuse light.
func NewDiffuseLight(emit texture.Texture) *DiffuseLight {
	return &DiffuseLight{
		emit: emit,
	}
}

// Scatter returns false for diffuse light materials.
func (dl *DiffuseLight) Scatter(_ ray.Ray, _ *hitrecord.HitRecord) (*ray.RayImpl, *vec3.Vec3Impl, bool) {
	return nil, nil, false
}

// Emitted returns the texture value at that point.
func (dl *DiffuseLight) Emitted(u float64, v float64, p *vec3.Vec3Impl) *vec3.Vec3Impl {
	return dl.emit.Value(u, v, p)
}
