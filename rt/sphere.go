package rt

import (
	"github.com/go-gl/mathgl/mgl64"
	"math"
)

type Sphere struct {
	Object
	Position mgl64.Vec3
	Radius float64
	Color mgl64.Vec3
}

func (sphere *Sphere) Intersect(ray *Ray, tMin, tMax float64) *Hit {
	oc := ray.Origin.Sub(sphere.Position)
	a := ray.Direction.Dot(ray.Direction)
	b := oc.Dot(ray.Direction)
	c := oc.Dot(oc) - sphere.Radius * sphere.Radius
	discriminant := b*b - a*c
	if discriminant > 0 {
		sqrt := math.Sqrt(discriminant)

		hit := &Hit{}

		temp := (-b - sqrt) / a
		if temp > tMin && temp < tMax {
			hit.T = temp
			hit.Intersection = ray.PointAt(hit.T)
			hit.Normal = hit.Intersection.Sub(sphere.Position).Normalize()
			return hit
		}

		temp = (-b + sqrt) / a
		if temp > tMin && temp < tMax {
			hit.T = temp
			hit.Intersection = ray.PointAt(hit.T)
			hit.Normal = hit.Intersection.Sub(sphere.Position).Normalize()
			return hit
		}
	}
	return nil
}

func (sphere *Sphere) GetColor() mgl64.Vec3 {
	return sphere.Color
}