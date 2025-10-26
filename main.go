// Material for ray tracing got from https://gabrielgambetta.com/computer-graphics-from-scratch/
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"math"
	"os"
	"runtime"
	"sync"

	"raytracing/vector3"
)

type Vec3 = vector3.Vector3
type Color = color.RGBA

type LightType uint32

const (
	Point   LightType = 0
	Ambient LightType = 1
)

var Epsilon float64 = 0.001
var MaxIntensity float64 = 1.

type Light struct {
	lightType LightType
	position  Vec3
	intensity float64
}

func (light *Light) ComputeLighting(point Vec3, normal Vec3, inverseDir Vec3, specular float64, spheres []Sphere) float64 {
	resIntensity := 0.
	lightDir := vector3.Vector3{}
	tMax := math.MaxFloat64
	switch light.lightType {
	case Ambient:
		return light.intensity
	case Point:
		lightDir = vector3.Sub(light.position, point)
		tMax = 1.
	}
	tMin := Epsilon
	closestSphere, _ := FindClosest(point, lightDir, spheres, tMin, tMax)

	if closestSphere.IsNull() {
		lightValue := math.Max(0., vector3.Dot(lightDir, normal))
		resIntensity += light.intensity * lightValue / (point.Length() * normal.Length())
		if specular > -1 {
			reflectDir := ReflectRay(lightDir, normal)
			specularValue := reflectDir.Dot(inverseDir)
			reflectDirLenght := reflectDir.Length()
			inverseDirLenght := inverseDir.Length()
			if reflectDirLenght == 0.0 || inverseDirLenght == 0.0 {
				panic("ComputeLighting: Division by zero")
			}
			resIntensity += light.intensity * math.Pow((math.Max(0., specularValue)/(reflectDir.Length()*inverseDir.Length())), specular)
		}
	}
	return math.Max(0., resIntensity)
}

type Sphere struct {
	radius     float64
	center     Vec3
	color      Color
	specular   float64
	reflective float64
}

func (s *Sphere) IsNull() bool {
	return (math.Abs(s.radius-0.) <= Epsilon) && (math.Abs(s.center.Length()-0.) <= Epsilon)
}

func (s *Sphere) ComputeIntersection(startPoint Vec3, direction Vec3) (float64, float64) {
	oc := startPoint.Sub(s.center)
	a := vector3.Dot(direction, direction)
	if a == 0.0 {
		panic("ComputeIntersection: Division by zero")
	}
	b := 2 * vector3.Dot(oc, direction)
	c := vector3.Dot(oc, oc) - s.radius*s.radius

	//intersection equantion of a^2 + 2b + c
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return -1., -1.
	}
	t1 := (-b + math.Sqrt(discriminant)) / (2 * a)
	t2 := (-b - math.Sqrt(discriminant)) / (2 * a)
	return t1, t2
}

func FindClosest(startPoint Vec3, direction Vec3, spheres []Sphere, tMin float64, tMax float64) (Sphere, float64) {
	closestT := math.MaxFloat64
	closestSphere := Sphere{radius: 0, center: Vec3{X: 0, Y: 0, Z: 0}, color: Color{R: 0, G: 0, B: 0, A: 255}}

	for _, sphere := range spheres {
		t1, t2 := sphere.ComputeIntersection(startPoint, direction)
		if t1 >= tMin && t1 <= tMax && t1 < closestT {
			closestSphere = sphere
			closestT = t1
		}
		if t2 >= tMin && t2 <= tMax && t2 < closestT {
			closestSphere = sphere
			closestT = t2
		}
	}
	return closestSphere, closestT
}

func ReflectRay(ray Vec3, normal Vec3) Vec3 {
	//in physics reflect = l - 2*n*dot(n,l)
	//due to negate ligth vector
	reflect := normal.Reflect(ray.Negate())
	return reflect
}

func TraceRay(startPoint Vec3, direction Vec3, spheres []Sphere, lights []Light, recursionDepth int8, tMin float64, tMax float64) Color {
	if direction.Length() == 0.0 {
		fmt.Println("Warning: ray direction is zero")
	}

	closestSphere, closestT := FindClosest(startPoint, direction, spheres, tMin, tMax)
	if closestSphere.IsNull() {
		return color.RGBA{R: 125, G: 125, B: 125, A: 255}
	}
	// P = O + tD
	pointIntersect := vector3.Add(startPoint, direction.MulScalar(closestT))
	// N = P - C
	normal := vector3.Sub(pointIntersect, closestSphere.center)
	normal = normal.Normalize()
	lightVal := 0.
	for _, light := range lights {
		lightVal += light.ComputeLighting(pointIntersect, normal, direction.Negate(), closestSphere.specular, spheres)
	}
	lightVal = math.Min(MaxIntensity, lightVal)
	closestSphere.color.R = uint8(float64(closestSphere.color.R) * lightVal)
	closestSphere.color.G = uint8(float64(closestSphere.color.G) * lightVal)
	closestSphere.color.B = uint8(float64(closestSphere.color.B) * lightVal)

	localColor := closestSphere.color
	if closestSphere.reflective <= 0 || recursionDepth <= 0 {
		return localColor
	}

	reflectedRay := ReflectRay(direction.Negate(), normal)
	tMin = Epsilon //Necessary offset for avoid intersection with itself
	reflectedColor := TraceRay(pointIntersect, reflectedRay, spheres, lights, recursionDepth-1, tMin, tMax)

	localColor.R = uint8(float64(localColor.R) * (1 - closestSphere.reflective))
	localColor.G = uint8(float64(localColor.G) * (1 - closestSphere.reflective))
	localColor.B = uint8(float64(localColor.B) * (1 - closestSphere.reflective))

	reflectedColor.R = uint8(float64(reflectedColor.R)*closestSphere.reflective) + localColor.R
	reflectedColor.G = uint8(float64(reflectedColor.G)*closestSphere.reflective) + localColor.G
	reflectedColor.B = uint8(float64(reflectedColor.B)*closestSphere.reflective) + localColor.B

	return reflectedColor
}

func main() {

	spheres := [4]Sphere{{radius: 1, center: Vec3{X: 0, Y: -1, Z: 3}, color: Color{R: 255, G: 0, B: 0, A: 255}, specular: 100, reflective: 0.01},
		{radius: 1, center: Vec3{X: -2, Y: 0, Z: 3}, color: Color{R: 0, G: 255, B: 0, A: 255}, specular: 25, reflective: 0.5},
		{radius: 1, center: Vec3{X: 2, Y: 0, Z: 3}, color: Color{R: 0, G: 0, B: 255, A: 255}, specular: 15, reflective: 0.1},
		{radius: 2000, center: Vec3{X: 0, Y: -2001, Z: 5}, color: Color{R: 255, G: 255, B: 0, A: 255}, specular: 1000, reflective: 0.}}

	lights := [3]Light{{lightType: Point, position: Vec3{X: -4, Y: 5, Z: 2}, intensity: 0.2},
		{lightType: Point, position: Vec3{X: 2, Y: 1, Z: 0}, intensity: 0.2},
		{lightType: Ambient, intensity: 0.3}}
	rayDistance := 1.

	start := Vec3{X: 0, Y: 0, Z: 0}
	w, h := 2048, 2048
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	if img == nil {
		panic("Can't create RGBA Image")
	}

	cpus := runtime.NumCPU()
	var wg sync.WaitGroup

	for i := 0; i < cpus; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for row := i; row < h; row += cpus {
				for col := 0; col < w; col++ {
					// Normalized pixed coordinates to [-1, 1]
					x := ((float64(row)+0.5)*2/float64(w) - 1)
					y := 1 - ((float64(col) + 0.5) * 2 / float64(h))
					rayDirection := Vec3{X: float64(x), Y: float64(y), Z: rayDistance}
					tMin := 1.
					tMax := math.MaxFloat64
					recursionDepth := 3
					clr := TraceRay(start, rayDirection, spheres[:], lights[:], int8(recursionDepth), tMin, tMax)
					img.Set(row, col, clr)
				}
			}
		}(i)
	}
	wg.Wait()
	f, err := os.Create("img.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if err = jpeg.Encode(f, img, nil); err != nil {
		fmt.Printf("failed to encode: %v", err)
	}
}
