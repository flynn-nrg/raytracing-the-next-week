package hitable

import (
	"github.com/flynn-nrg/raytracing-the-next-week/pkg/hitrecord"
	"github.com/flynn-nrg/raytracing-the-next-week/pkg/material"
	"github.com/flynn-nrg/raytracing-the-next-week/pkg/ray"
)

// Ensure interface compliance.
var _ Hitable = (*HitableSlice)(nil)

// HitableSlice represents a list of hitable entities.
type HitableSlice struct {
	hitables []Hitable
}

// NewSlice returns an instance of HitableSlice.
func NewSlice(hitables []Hitable) *HitableSlice {
	return &HitableSlice{
		hitables: hitables,
	}
}

// Hit computes whether a ray intersects with any of the elements in the slice.
func (hs *HitableSlice) Hit(r ray.Ray, tMin float64, tMax float64) (*hitrecord.HitRecord, material.Material, bool) {
	var rec *hitrecord.HitRecord
	var mat material.Material
	var hitAnything bool
	closestSoFar := tMax

	for _, h := range hs.hitables {
		if tempRec, tempMat, ok := h.Hit(r, tMin, closestSoFar); ok {
			rec = tempRec
			mat = tempMat
			hitAnything = ok
			closestSoFar = rec.T()
		}
	}

	return rec, mat, hitAnything
}
