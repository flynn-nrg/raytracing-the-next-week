package hitable

import (
	"fmt"
	"math/rand"
	"sort"

	"github.com/flynn-nrg/raytracing-the-next-week/pkg/aabb"
	"github.com/flynn-nrg/raytracing-the-next-week/pkg/hitrecord"
	"github.com/flynn-nrg/raytracing-the-next-week/pkg/material"
	"github.com/flynn-nrg/raytracing-the-next-week/pkg/ray"
)

// Ensure interface compliance.
var _ Hitable = (*BVHNode)(nil)

// BVHNode represents a bounding volume hierarchy node.
type BVHNode struct {
	left  Hitable
	right Hitable
	time0 float64
	time1 float64
	box   *aabb.AABB
}

func NewBVH(hitables []Hitable, time0 float64, time1 float64) *BVHNode {
	bn := &BVHNode{
		time0: time0,
		time1: time1,
	}

	axis := int(3 * rand.Float64())
	switch axis {
	case 0:
		sort.Slice(hitables, func(i, j int) bool {
			var box0, box1 *aabb.AABB
			var ok bool
			if box0, ok = hitables[i].BoundingBox(0, 0); !ok {
				fmt.Printf("no bounding box in BVH node\n")
				return false
			}
			if box1, ok = hitables[j].BoundingBox(0, 0); !ok {
				fmt.Printf("no bounding box in BVH node\n")
				return false
			}
			return aabb.BoxLessX(box0, box1)
		})

	case 1:
		sort.Slice(hitables, func(i, j int) bool {
			var box0, box1 *aabb.AABB
			var ok bool
			if box0, ok = hitables[i].BoundingBox(0, 0); !ok {
				fmt.Errorf("no bounding box in BVH node\n")
				return false
			}
			if box1, ok = hitables[j].BoundingBox(0, 0); !ok {
				fmt.Errorf("no bounding box in BVH node\n")
				return false
			}
			return aabb.BoxLessY(box0, box1)
		})

	case 2:
		sort.Slice(hitables, func(i, j int) bool {
			var box0, box1 *aabb.AABB
			var ok bool
			if box0, ok = hitables[i].BoundingBox(0, 0); !ok {
				return false
			}
			if box1, ok = hitables[j].BoundingBox(0, 0); !ok {
				return false
			}
			return aabb.BoxLessZ(box0, box1)
		})

	}

	if len(hitables) == 1 {
		bn.left = hitables[0]
		bn.right = bn.left
	} else if len(hitables) == 2 {
		bn.left = hitables[0]
		bn.right = hitables[1]
	} else {
		bn.left = NewBVH(hitables[:len(hitables)/2], time0, time1)
		bn.right = NewBVH(hitables[len(hitables)/2:], time0, time1)
	}

	var leftBox, rightBox *aabb.AABB
	var ok bool
	if leftBox, ok = bn.left.BoundingBox(time0, time1); !ok {
		fmt.Printf("no bounding box in BVH node\n")
	}
	if rightBox, ok = bn.right.BoundingBox(time0, time1); !ok {
		fmt.Printf("no bounding box in BVH node\n")
	}

	bn.box = aabb.SurroundingBox(leftBox, rightBox)

	return bn
}

func (bn *BVHNode) Hit(r ray.Ray, tMin float64, tMax float64) (*hitrecord.HitRecord, material.Material, bool) {
	if bn.box.Hit(r, tMin, tMax) {
		leftRec, leftMat, hitLeft := bn.left.Hit(r, tMin, tMax)
		rightRec, rightMat, hitRight := bn.right.Hit(r, tMin, tMax)

		if hitLeft && hitRight {
			if leftRec.T() < rightRec.T() {
				return leftRec, leftMat, true
			}
			return rightRec, rightMat, true
		}

		if hitLeft {
			return leftRec, leftMat, true
		}

		if hitRight {
			return rightRec, rightMat, true
		}
	}

	return nil, nil, false
}

func (bn *BVHNode) BoundingBox(time0 float64, time1 float64) (*aabb.AABB, bool) {
	return bn.box, true
}
