package hitable

import (
	"math"
	"math/rand"

	"github.com/flynn-nrg/raytracing-the-next-week/pkg/aabb"
	"github.com/flynn-nrg/raytracing-the-next-week/pkg/hitrecord"
	"github.com/flynn-nrg/raytracing-the-next-week/pkg/material"
	"github.com/flynn-nrg/raytracing-the-next-week/pkg/ray"
	"github.com/flynn-nrg/raytracing-the-next-week/pkg/texture"
	"github.com/flynn-nrg/raytracing-the-next-week/pkg/vec3"
)

// Ensure interface compliance.
var _ Hitable = (*ConstantMedium)(nil)

// ConstantMedium represents a medium with constant density.
type ConstantMedium struct {
	hitable       Hitable
	density       float64
	phaseFunction material.Material
}

// NewConstantMedium returns a new instance of the constant medium hitable.
func NewConstantMedium(hitable Hitable, density float64, a texture.Texture) *ConstantMedium {
	return &ConstantMedium{
		hitable:       hitable,
		density:       density,
		phaseFunction: material.NewIsotropic(a),
	}
}

func (cm *ConstantMedium) Hit(r ray.Ray, tMin float64, tMax float64) (*hitrecord.HitRecord, material.Material, bool) {
	if rec1, _, ok := cm.hitable.Hit(r, -math.MaxFloat64, math.MaxFloat64); ok {
		if rec2, _, ok := cm.hitable.Hit(r, rec1.T()+0.0001, math.MaxFloat64); ok {
			rec1t := rec1.T()
			rec2t := rec2.T()
			if rec1t < tMin {
				rec1t = tMin
			}
			if rec2t < tMax {
				rec2t = tMax
			}
			if rec1t >= rec2t {
				return nil, nil, false
			}
			if rec1t < 0 {
				rec1t = 0
			}

			distanceInsideBoundary := (rec2t - rec1t) * r.Direction().Length()
			hitDistance := -(1 / cm.density) * math.Log(rand.Float64())
			if hitDistance < distanceInsideBoundary {
				t := rec1t + hitDistance/r.Direction().Length()
				// arbitrary
				normal := &vec3.Vec3Impl{X: 1}
				hr := hitrecord.New(t, 0, 0, r.PointAtParameter(t), normal)
				return hr, cm.phaseFunction, true
			}
		}
	}

	return nil, nil, false
}

func (cm *ConstantMedium) BoundingBox(time0 float64, time1 float64) (*aabb.AABB, bool) {
	return cm.hitable.BoundingBox(time0, time1)
}
