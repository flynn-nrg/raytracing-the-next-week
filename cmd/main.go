package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/flynn-nrg/raytracing-the-next-week/pkg/camera"
	"github.com/flynn-nrg/raytracing-the-next-week/pkg/hitable"
	"github.com/flynn-nrg/raytracing-the-next-week/pkg/ray"
	"github.com/flynn-nrg/raytracing-the-next-week/pkg/scenes"
	"github.com/flynn-nrg/raytracing-the-next-week/pkg/vec3"
)

func color(r ray.Ray, world *hitable.HitableSlice, depth int) *vec3.Vec3Impl {
	if rec, mat, ok := world.Hit(r, 0.001, math.MaxFloat64); ok {
		scattered, attenuation, ok := mat.Scatter(r, rec)
		if depth < 50 && ok {
			return vec3.Mul(attenuation, color(scattered, world, depth+1))
		}

	}
	unitDirection := vec3.UnitVector(r.Direction())
	t := 0.5*unitDirection.Y + 1.0
	return vec3.Add(vec3.ScalarMul(&vec3.Vec3Impl{X: 1.0, Y: 1.0, Z: 1.0}, (1.0-t)),
		vec3.ScalarMul(&vec3.Vec3Impl{X: 0.5, Y: 0.7, Z: 1.0}, t))
}

func main() {
	nx := 200
	ny := 100
	ns := 100

	rand.Seed(time.Now().UnixNano())

	fmt.Printf("P3\n%v %v\n255\n", nx, ny)

	world := scenes.TextureMappedSphere()
	lookFrom := &vec3.Vec3Impl{X: 13.0, Y: 2.0, Z: 3.0}
	lookAt := &vec3.Vec3Impl{}
	vup := &vec3.Vec3Impl{Y: 1}
	distToFocus := 10.0
	aperture := 0.0
	aspect := float64(nx) / float64(ny)
	vfov := float64(20.0)
	time0 := 0.0
	time1 := 1.0
	cam := camera.New(lookFrom, lookAt, vup, vfov, aspect, aperture, distToFocus, time0, time1)

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			col := &vec3.Vec3Impl{}
			for s := 0; s < ns; s++ {
				u := (float64(i) + rand.Float64()) / float64(nx)
				v := (float64(j) + rand.Float64()) / float64(ny)
				r := cam.GetRay(u, v)
				col = vec3.Add(col, color(r, world, 0))
			}

			col = vec3.ScalarDiv(col, float64(ns))
			// gamma 2
			col = &vec3.Vec3Impl{X: math.Sqrt(col.X), Y: math.Sqrt(col.Y), Z: math.Sqrt(col.Z)}
			ir := int(255.99 * col.X)
			ig := int(255.99 * col.Y)
			ib := int(255.99 * col.Z)

			fmt.Printf("%v %v %v\n", ir, ig, ib)
		}
	}
}
