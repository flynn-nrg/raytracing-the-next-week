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
	center0  *vec3.Vec3Impl
	center1  *vec3.Vec3Impl
	time0    float64
	time1    float64
	radius   float64
	material material.Material
}

// NewSphere returns a new instance of Sphere.
func NewSphere(center0 *vec3.Vec3Impl, center1 *vec3.Vec3Impl, time0 float64, time1 float64, radius float64, material material.Material) *Sphere {
	return &Sphere{
		center0:  center0,
		center1:  center1,
		time0:    time0,
		time1:    time1,
		radius:   radius,
		material: material,
	}
}

// Hit computes whether a ray intersects with the defined sphere.
func (s *Sphere) Hit(r ray.Ray, tMin float64, tMax float64) (*hitrecord.HitRecord, material.Material, bool) {
	oc := vec3.Sub(r.Origin(), s.center(r.Time()))
	a := vec3.Dot(r.Direction(), r.Direction())
	b := vec3.Dot(oc, r.Direction())
	c := vec3.Dot(oc, oc) - (s.radius * s.radius)

	discriminant := (b * b) - (a * c)
	if discriminant > 0 {
		temp := (-b - math.Sqrt(b*b-a*c)) / a
		if temp < tMax && temp > tMin {
			return hitrecord.New(temp, r.PointAtParameter(temp),
				vec3.ScalarDiv(vec3.Sub(r.PointAtParameter(temp), s.center(r.Time())),
					s.radius)), s.material, true
		}

		temp = (-b + math.Sqrt(b*b-a*c)) / a
		if temp < tMax && temp > tMin {
			return hitrecord.New(temp,
				r.PointAtParameter(temp),
				vec3.ScalarDiv(vec3.Sub(r.PointAtParameter(temp), s.center(r.Time())), s.radius)), s.material, true
		}
	}

	return nil, nil, false
}

func (s *Sphere) center(time float64) *vec3.Vec3Impl {
	return vec3.Add(s.center0, vec3.ScalarMul(vec3.Sub(s.center1, s.center0), ((time-s.time0)/(s.time1-s.time0))))
}