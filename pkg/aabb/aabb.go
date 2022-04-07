// Package aabb implements functions to work with axis-aligned bounding boxes.
package aabb

import (
	"math"

	"github.com/flynn-nrg/raytracing-the-next-week/pkg/ray"
	"github.com/flynn-nrg/raytracing-the-next-week/pkg/vec3"
)

// AABB represents an axis-aligned bounding box.
type AABB struct {
	min *vec3.Vec3Impl
	max *vec3.Vec3Impl
}

// New returns a new axis-aligned bounding box.
func New(min *vec3.Vec3Impl, max *vec3.Vec3Impl) *AABB {
	return &AABB{
		min: min,
		max: max,
	}
}

// SurroundingBox computes the box that encloses the two supplied boxes.
func SurroundingBox(box0 *AABB, box1 *AABB) *AABB {
	small := &vec3.Vec3Impl{
		X: math.Min(box0.min.X, box1.min.X),
		Y: math.Min(box0.min.Y, box1.min.Y),
		Z: math.Min(box0.min.Z, box1.min.Z),
	}
	big := &vec3.Vec3Impl{
		X: math.Max(box0.max.X, box1.max.X),
		Y: math.Max(box0.max.Y, box1.max.Y),
		Z: math.Max(box0.max.Z, box1.max.Z),
	}

	return New(small, big)
}

// BoxLessX sorts two boxes by their min X value.
func BoxLessX(box0 *AABB, box1 *AABB) bool {
	return box0.min.X < box1.min.X
}

// BoxLessY sorts two boxes by their min Y value.
func BoxLessY(box0 *AABB, box1 *AABB) bool {
	return box0.min.Y < box1.min.Y
}

// BoxLessZ sorts two boxes by their min Z value.
func BoxLessZ(box0 *AABB, box1 *AABB) bool {
	return box0.min.Z < box1.min.Z
}

// Hit returns true if a ray intersects with the bounding box.
func (a *AABB) Hit(r ray.Ray, tMin float64, tMax float64) bool {

	mins := []float64{a.min.X, a.min.Y, a.min.Z}
	maxs := []float64{a.max.X, a.max.Y, a.max.Z}
	origs := []float64{r.Origin().X, r.Origin().Y, r.Origin().Z}
	dirs := []float64{r.Direction().X, r.Direction().Y, r.Direction().Z}

	for i := range mins {
		invD := 1.0 / dirs[i]
		t0 := (mins[i] - origs[i]) * invD
		t1 := (maxs[i] - origs[i]) * invD
		if invD < 0.0 {
			t := t0
			t0 = t1
			t1 = t
		}

		tMin = math.Max(t0, tMin)
		tMax = math.Min(t1, tMax)
		if tMax <= tMin {
			return false
		}
	}

	return true
}
