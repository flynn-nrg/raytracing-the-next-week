// Package camera implements a set of functions to work with cameras.
package camera

import (
	"math"
	"math/rand"

	"github.com/flynn-nrg/raytracing-the-next-week/pkg/ray"
	"github.com/flynn-nrg/raytracing-the-next-week/pkg/vec3"
)

// Camera represents a camera in the world.
type Camera struct {
	lensRadius      float64
	u               *vec3.Vec3Impl
	v               *vec3.Vec3Impl
	origin          *vec3.Vec3Impl
	lowerLeftCorner *vec3.Vec3Impl
	horizontal      *vec3.Vec3Impl
	vertical        *vec3.Vec3Impl
}

// New returns an instance of a camera.
func New(lookFrom *vec3.Vec3Impl, lookAt *vec3.Vec3Impl, vup *vec3.Vec3Impl,
	vfov float64, aspect float64, aperture float64, focusDist float64) *Camera {

	lensRadius := aperture / 2.0
	theta := vfov * math.Pi / 180
	halfHeight := math.Tan(theta / 2.0)
	halfWidth := aspect * halfHeight
	w := vec3.UnitVector(vec3.Sub(lookFrom, lookAt))
	u := vec3.UnitVector(vec3.Cross(vup, w))
	v := vec3.Cross(w, u)

	// origin - halfWidth*focusDist*u - halfHeight*focusDist*v - focusDist*w
	lowerLeftCorner := vec3.Sub(lookFrom, vec3.ScalarMul(u, halfWidth*focusDist), vec3.ScalarMul(v, halfHeight*focusDist), vec3.ScalarMul(w, focusDist))
	horizontal := vec3.ScalarMul(u, 2.0*halfWidth*focusDist)
	vertical := vec3.ScalarMul(v, 2.0*halfHeight*focusDist)
	origin := lookFrom

	return &Camera{
		lensRadius:      lensRadius,
		u:               u,
		v:               v,
		lowerLeftCorner: lowerLeftCorner,
		horizontal:      horizontal,
		vertical:        vertical,
		origin:          origin,
	}
}

// GetRay returns the ray associated for the supplied u and v.
func (c *Camera) GetRay(s float64, t float64) *ray.RayImpl {
	rd := vec3.ScalarMul(randomInUnitDisc(), c.lensRadius)
	offset := vec3.Add(vec3.ScalarMul(c.u, rd.X), vec3.ScalarMul(c.v, rd.Y))
	return ray.New(vec3.Add(c.origin, offset),
		// lowerLeftCorner + s*horizontal + t*vertical - origin - offset
		vec3.Sub(vec3.Add(c.lowerLeftCorner, vec3.ScalarMul(c.horizontal, s),
			vec3.ScalarMul(c.vertical, t)), c.origin, offset))
}

func randomInUnitDisc() *vec3.Vec3Impl {
	for {
		p := vec3.Sub(vec3.ScalarMul(&vec3.Vec3Impl{X: rand.Float64(), Y: rand.Float64()}, 2.0), &vec3.Vec3Impl{X: 1.0, Y: 1.0})
		if vec3.Dot(p, p) < 1.0 {
			return p
		}
	}
}
