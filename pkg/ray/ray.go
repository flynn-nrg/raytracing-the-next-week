package ray

import "github.com/flynn-nrg/raytracing-the-next-week/pkg/vec3"

// Ensure interface compliance.
var _ Ray = (*RayImpl)(nil)

// RayImpl implements the Ray interface.
type RayImpl struct {
	origin    *vec3.Vec3Impl
	direction *vec3.Vec3Impl
}

// New returns a new ray with the supplied origin and direction vectors.
func New(origin *vec3.Vec3Impl, direction *vec3.Vec3Impl) *RayImpl {
	return &RayImpl{
		origin:    origin,
		direction: direction,
	}
}

// Origin returns the origin vector of this ray.
func (r *RayImpl) Origin() *vec3.Vec3Impl {
	return r.origin
}

// Direction returns the direction vector of this ray.
func (r *RayImpl) Direction() *vec3.Vec3Impl {
	return r.direction
}

// PointAtParameter is used to traverse the ray.
func (r *RayImpl) PointAtParameter(t float64) *vec3.Vec3Impl {
	return vec3.Add(r.origin, vec3.ScalarMul(r.direction, t))
}
