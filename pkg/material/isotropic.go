package material

import (
	"github.com/flynn-nrg/raytracing-the-next-week/pkg/hitrecord"
	"github.com/flynn-nrg/raytracing-the-next-week/pkg/ray"
	"github.com/flynn-nrg/raytracing-the-next-week/pkg/texture"
	"github.com/flynn-nrg/raytracing-the-next-week/pkg/vec3"
)

// Ensure interface compliance.
var _ Material = (*Isotropic)(nil)

// Isotropic represents an isotropic material.
type Isotropic struct {
	albedo texture.Texture
}

// NewIsotropic returns a new instances of the isotropic material.
func NewIsotropic(albedo texture.Texture) *Isotropic {
	return &Isotropic{
		albedo: albedo,
	}
}

// Scatter computes how the ray bounces off the surface of a diffuse material.
func (i *Isotropic) Scatter(r ray.Ray, hr *hitrecord.HitRecord) (*ray.RayImpl, *vec3.Vec3Impl, bool) {
	scattered := ray.New(hr.P(), randomInUnitSphere(), r.Time())
	attenuation := i.albedo.Value(hr.U(), hr.V(), hr.P())
	return scattered, attenuation, true
}

// Emitted returns black for isotropics materials.
func (i *Isotropic) Emitted(_ float64, _ float64, _ *vec3.Vec3Impl) *vec3.Vec3Impl {
	return &vec3.Vec3Impl{}
}
