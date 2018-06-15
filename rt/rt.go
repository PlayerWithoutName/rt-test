package rt

import (
	"github.com/go-gl/mathgl/mgl64"
	"image/color"
	"math"
	"math/rand"
)

var sky1 mgl64.Vec3 = [3]float64{1.0, 1.0, 1.0}
var sky2 mgl64.Vec3 = [3]float64{0.5, 0.7, 1.0}

func randomUnitVector() mgl64.Vec3 {
	vec := mgl64.Vec3{1.0, 1.0, 1.0}
	for ok := true; ok; ok = vec.Dot(vec) > 1.0 {
		vec = mgl64.Vec3{rand.Float64(), rand.Float64(), rand.Float64()}.Mul(2.0).Sub(mgl64.Vec3{1.0, 1.0, 1.0})
	}
	return vec
}

func sky(t float64) mgl64.Vec3 {
	c1 := sky1.Mul(1-t)
	c2 := sky2.Mul(t)
	return c1.Add(c2)
}

func col(ray *Ray, tMin, tMax float64, objects *[]Object) mgl64.Vec3 {
	for _, obj := range *objects {
		if hit := obj.Intersect(ray, tMin, tMax); hit != nil {
			target := hit.Intersection.Add(hit.Normal).Add(randomUnitVector())
			return col(&Ray{hit.Intersection, target.Sub(hit.Intersection)}, tMin, tMax, objects).Mul(0.5)
		}
	}
	dir := ray.Direction.Normalize()
	t := 0.5*(dir.Y() + 1.0)
	return sky(t)
}

func Trace() {
	res := &Result{
		W: 200,
		H: 100,
		Path: "test.png",
	}

	res.Init()

	ar := float64(res.W)/float64(res.H)

	samples := 100

	lowerLeft := mgl64.Vec3{-ar, -1.0, -1.0}

	horizontal := mgl64.Vec3{2*ar, 0.0, 0.0}
	vertical := mgl64.Vec3{0.0, 2.0, 0.0}

	origin := mgl64.Vec3{0.0, 0.0, 0.0}

	objects := []Object{
		&Sphere{
			Position:mgl64.Vec3{0.0, 0.0, -1.0},
			Radius:0.5,
			Color:mgl64.Vec3{1.0, 0.0, 0.0},
		},
		&Sphere{
			Position:mgl64.Vec3{0.0, -100.5, -1.0},
			Radius:100,
			Color:mgl64.Vec3{1.0, 0.0, 0.0},
		},
	}

	for x := 0; x < res.W; x++ {
		for y := 1; y <= res.H; y++ {
			outCol := mgl64.Vec3{0.0, 0.0, 0.0}

			for i := 0; i < samples; i++ {
				u := (float64(x) + rand.Float64()) / float64(res.W)
				v := (float64(y) + rand.Float64()) / float64(res.H)

				r := &Ray{origin, lowerLeft.Add(horizontal.Mul(u)).Add(vertical.Mul(v))}
				outCol = outCol.Add(col(r, 0.0, math.MaxFloat64, &objects))
			}

			outCol = outCol.Mul(1.0/float64(samples))
			outCol = mgl64.Vec3{math.Sqrt(outCol.X()), math.Sqrt(outCol.Y()), math.Sqrt(outCol.Z())}

			res.SetPixel(x, res.H-y, &color.NRGBA64{uint16(outCol.X()*65535), uint16(outCol.Y()*65535), uint16(outCol.Z()*65535), 65535})
		}
	}

	res.Save()
}