package rt

import (
	"github.com/go-gl/mathgl/mgl64"
)

type Ray struct {
	Origin, Direction mgl64.Vec3
}

type Hit struct {
	Normal, Intersection mgl64.Vec3
	T float64
}

func (ray *Ray) PointAt(t float64) mgl64.Vec3 {
	return ray.Origin.Add(ray.Direction.Mul(t))
}

type Object interface {
	GetColor() mgl64.Vec3
	Intersect(ray *Ray, tMin, tMax float64) *Hit
}