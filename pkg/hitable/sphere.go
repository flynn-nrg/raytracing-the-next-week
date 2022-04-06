package hitable

import (
	"math"

	"github.com/flynn-nrg/raytracing-the-next-week/pkg/hitrecord"
	"github.com/flynn-nrg/raytracing-the-next-week/pkg/material"
	"github.com/flynn-nrg/raytracing-the-next-week/pkg/ray"
	"github.com/flynn-nrg/raytracing-the-next-week/pkg/vec3"
)

// Ensure interface compliance.
var _ Hitable = (*Sphere)(nil)

// Sphere represents a sphere in the 3d world.
type Sphere struct {
	center   *vec3.Vec3Impl
	radius   float64
	material material.Material
}

// NewSphere returns a new instance of Sphere.
func NewSphere(center *vec3.Vec3Impl, radius float64, material material.Material) *Sphere {
	return &Sphere{
		center:   center,
		radius:   radius,
		material: material,
	}
}

// Hit computes whether a ray intersects with the defined sphere.
func (s *Sphere) Hit(r ray.Ray, tMin float64, tMax float64) (*hitrecord.HitRecord, material.Material, bool) {
	oc := vec3.Sub(r.Origin(), s.center)
	a := vec3.Dot(r.Direction(), r.Direction())
	b := vec3.Dot(oc, r.Direction())
	c := vec3.Dot(oc, oc) - (s.radius * s.radius)

	discriminant := (b * b) - (a * c)
	if discriminant > 0 {
		temp := (-b - math.Sqrt(b*b-a*c)) / a
		if temp < tMax && temp > tMin {
			return hitrecord.New(temp, r.PointAtParameter(temp),
				vec3.ScalarDiv(vec3.Sub(r.PointAtParameter(temp), s.center),
					s.radius)), s.material, true
		}

		temp = (-b + math.Sqrt(b*b-a*c)) / a
		if temp < tMax && temp > tMin {
			return hitrecord.New(temp,
				r.PointAtParameter(temp),
				vec3.ScalarDiv(vec3.Sub(r.PointAtParameter(temp), s.center), s.radius)), s.material, true
		}
	}

	return nil, nil, false
}
