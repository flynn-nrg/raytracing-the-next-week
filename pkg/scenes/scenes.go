// Package scenes implements some sample scenes.
package scenes

import (
	"log"
	"math/rand"
	"os"

	"github.com/flynn-nrg/raytracing-the-next-week/pkg/hitable"
	"github.com/flynn-nrg/raytracing-the-next-week/pkg/material"
	"github.com/flynn-nrg/raytracing-the-next-week/pkg/texture"
	"github.com/flynn-nrg/raytracing-the-next-week/pkg/vec3"
)

// RandomScene returns a random scene.
func RandomScene() *hitable.HitableSlice {
	checker := texture.NewChecker(texture.NewConstant(&vec3.Vec3Impl{X: 0.2, Y: 0.3, Z: 0.1}),
		texture.NewConstant(&vec3.Vec3Impl{X: 0.9, Y: 0.9, Z: 0.9}))
	spheres := []hitable.Hitable{hitable.NewSphere(&vec3.Vec3Impl{X: 0, Y: -1000, Z: 0}, &vec3.Vec3Impl{X: 0, Y: -1000, Z: 0}, 0, 1, 1000, material.NewLambertian(checker))}
	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMat := rand.Float64()
			center := &vec3.Vec3Impl{X: float64(a) + 0.9*rand.Float64(), Y: 0.2, Z: float64(b) + 0.9*rand.Float64()}
			if vec3.Sub(center, &vec3.Vec3Impl{X: 4, Y: 0.2, Z: 0}).Length() > 0.9 {
				if chooseMat < 0.8 {
					// diffuse
					spheres = append(spheres, hitable.NewSphere(center,
						vec3.Add(center, &vec3.Vec3Impl{Y: 0.5 * rand.Float64()}), 0.0, 1.0, 0.2,
						material.NewLambertian(texture.NewConstant(&vec3.Vec3Impl{
							X: rand.Float64() * rand.Float64(),
							Y: rand.Float64() * rand.Float64(),
							Z: rand.Float64() * rand.Float64(),
						}))))
				} else if chooseMat < 0.95 {
					// metal
					spheres = append(spheres, hitable.NewSphere(center, center, 0.0, 1.0, 0.2,
						material.NewMetal(&vec3.Vec3Impl{
							X: 0.5 * (1.0 - rand.Float64()),
							Y: 0.5 * (1.0 - rand.Float64()),
							Z: 0.5 * (1.0 - rand.Float64()),
						}, 0.2*rand.Float64())))
				} else {
					// glass
					spheres = append(spheres, hitable.NewSphere(center, center, 0.0, 1.0, 0.2, material.NewDielectric(1.5)))
				}
			}
		}
	}

	spheres = append(spheres, hitable.NewSphere(&vec3.Vec3Impl{Y: 1.0}, &vec3.Vec3Impl{Y: 1.0}, 0.0, 1.0, 1.0, material.NewDielectric(1.5)))
	spheres = append(spheres, hitable.NewSphere(&vec3.Vec3Impl{X: -4.0, Y: 1.0}, &vec3.Vec3Impl{X: -4.0, Y: 1.0}, 0.0, 1.0, 1.0, material.NewLambertian(texture.NewConstant(&vec3.Vec3Impl{X: 0.4, Y: 0.2, Z: 0.1}))))
	spheres = append(spheres, hitable.NewSphere(&vec3.Vec3Impl{X: 4.0, Y: 1.0}, &vec3.Vec3Impl{X: 4.0, Y: 1.0}, 0.0, 1.0, 1.0, material.NewMetal(&vec3.Vec3Impl{X: 0.7, Y: 0.6, Z: 0.5}, 0.0)))

	return hitable.NewSlice(spheres)
}

// TwoSpheres returns a scene containing two spheres.
func TwoSpheres() *hitable.HitableSlice {
	checker := texture.NewChecker(texture.NewConstant(&vec3.Vec3Impl{X: 0.2, Y: 0.3, Z: 0.1}),
		texture.NewConstant(&vec3.Vec3Impl{X: 0.9, Y: 0.9, Z: 0.9}))
	spheres := []hitable.Hitable{
		hitable.NewSphere(&vec3.Vec3Impl{X: 0, Y: -10, Z: 0}, &vec3.Vec3Impl{X: 0, Y: -10, Z: 0}, 0, 1, 10, material.NewLambertian(checker)),
		hitable.NewSphere(&vec3.Vec3Impl{X: 0, Y: 10, Z: 0}, &vec3.Vec3Impl{X: 0, Y: 10, Z: 0}, 0, 1, 10, material.NewLambertian(checker)),
	}

	return hitable.NewSlice(spheres)
}

// TwoPerlinSpheres returns a scene containing two spheres with Perlin noise.
func TwoPerlinSpheres() *hitable.HitableSlice {
	perText := texture.NewNoise(4.0)
	spheres := []hitable.Hitable{
		hitable.NewSphere(&vec3.Vec3Impl{X: 0, Y: -1000, Z: 0}, &vec3.Vec3Impl{X: 0, Y: -1000, Z: 0}, 0, 1, 1000, material.NewLambertian(perText)),
		hitable.NewSphere(&vec3.Vec3Impl{X: 0, Y: 2, Z: 0}, &vec3.Vec3Impl{X: 0, Y: 2, Z: 0}, 0, 1, 2, material.NewLambertian(perText)),
	}

	return hitable.NewSlice(spheres)
}

// TwoPerlinSpheres returns a scene containing two spheres with Perlin noise.
func TextureMappedSphere() *hitable.HitableSlice {
	file, err := os.Open("../images/earth.png")
	if err != nil {
		log.Fatalf("could not read texture file; %v", err)
	}
	imgText, err := texture.NewFromPNG(file)
	if err != nil {
		log.Fatalf("failed to decode image; %v", err)
	}
	spheres := []hitable.Hitable{
		hitable.NewSphere(&vec3.Vec3Impl{X: 0, Y: 0, Z: 0}, &vec3.Vec3Impl{X: 0, Y: 0, Z: 0}, 0, 1, 1, material.NewLambertian(imgText)),
	}

	return hitable.NewSlice(spheres)
}
