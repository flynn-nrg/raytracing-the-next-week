package hitrecord

import "github.com/flynn-nrg/raytracing-the-next-week/pkg/vec3"

// HitRecord contains data related to an intersection between a ray and an object.
type HitRecord struct {
	t      float64
	p      *vec3.Vec3Impl
	normal *vec3.Vec3Impl
}

func New(t float64, p *vec3.Vec3Impl, normal *vec3.Vec3Impl) *HitRecord {
	return &HitRecord{
		t:      t,
		p:      p,
		normal: normal,
	}
}

// Normal returns the normal vector at the intersection point.
func (hr *HitRecord) Normal() *vec3.Vec3Impl {
	return hr.normal
}

// P returns the intersection point.
func (hr *HitRecord) P() *vec3.Vec3Impl {
	return hr.p
}

// T returns the t value.
func (hr *HitRecord) T() float64 {
	return hr.t
}
