// Package material implements the different materials and their properties.
package material

import (
	"github.com/flynn-nrg/raytracing-the-next-week/pkg/hitrecord"
	"github.com/flynn-nrg/raytracing-the-next-week/pkg/ray"
	"github.com/flynn-nrg/raytracing-the-next-week/pkg/vec3"
)

// Material defines the methods to handle materials.
type Material interface {
	Scatter(r ray.Ray, hr *hitrecord.HitRecord) (*ray.RayImpl, *vec3.Vec3Impl, bool)
	Emitted(u float64, v float64, p *vec3.Vec3Impl) *vec3.Vec3Impl
}
