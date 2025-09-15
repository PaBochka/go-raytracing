package vector3

import (
	"fmt"
	"math"
)

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

type Vector3 struct {
	X, Y, Z float64
}

func Add(v1 Vector3, v2 Vector3) Vector3 {
	return Vector3{
		v1.X + v2.X,
		v1.Y + v2.Y,
		v1.Z + v2.Z,
	}
}

func AddScalar(v Vector3, scalar float64) Vector3 {
	return Vector3{
		v.X + scalar,
		v.Y + scalar,
		v.Z + scalar,
	}
}

func AddScalars(v Vector3, x float64, y float64, z float64) Vector3 {
	return Vector3{
		v.X + x,
		v.Y + y,
		v.Z + z,
	}
}

func Sub(v1 Vector3, v2 Vector3) Vector3 {
	return Vector3{
		v1.X - v2.X,
		v1.Y - v2.Y,
		v1.Z - v2.Z,
	}
}

func SubScalar(v Vector3, scalar float64) Vector3 {
	return Vector3{
		v.X - scalar,
		v.Y - scalar,
		v.Z - scalar,
	}
}

func SubScalars(v Vector3, x float64, y float64, z float64) Vector3 {
	return Vector3{
		v.X - x,
		v.Y - y,
		v.Z - z,
	}
}

func Mul(v1 Vector3, v2 Vector3) Vector3 {
	return Vector3{
		v1.X * v2.X,
		v1.Y * v2.Y,
		v1.Z * v2.Z,
	}
}

func MulScalar(v Vector3, scalar float64) Vector3 {
	return Vector3{
		v.X * scalar,
		v.Y * scalar,
		v.Z * scalar,
	}
}

func MulScalars(v Vector3, x float64, y float64, z float64) Vector3 {
	return Vector3{
		v.X * x,
		v.Y * y,
		v.Z * z,
	}
}

func Div(v1 Vector3, v2 Vector3) Vector3 {
	return Vector3{
		v1.X / v2.X,
		v1.Y / v2.Y,
		v1.Z / v2.Z,
	}
}

func DivScalar(v Vector3, scalar float64) Vector3 {
	return Vector3{
		v.X / scalar,
		v.Y / scalar,
		v.Z / scalar,
	}
}

func DivScalars(v Vector3, x float64, y float64, z float64) Vector3 {
	return Vector3{
		v.X / x,
		v.Y / y,
		v.Z / z,
	}
}

func Dot(v1 Vector3, v2 Vector3) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

func Cross(v1 Vector3, v2 Vector3) Vector3 {
	return Vector3{
		v1.Y*v2.Z - v1.Z*v2.Y,
		v1.Z*v2.X - v1.X*v2.Z,
		v1.X*v2.Y - v1.Y*v2.X,
	}
}

func Lerp(a Vector3, b Vector3, t float64) Vector3 {
	return Vector3{
		a.X + (b.X-a.X)*t,
		a.Y + (b.Y-a.Y)*t,
		a.Z + (b.Z-a.Z)*t,
	}
}

func Distance(a Vector3, b Vector3) float64 {
	dx := a.X - b.X
	dy := a.Y - b.Y
	dz := a.Z - b.Z
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func Reflect(v1 Vector3, v2 Vector3) Vector3 {
	factor := -2.0 * Dot(v1, v2)
	return Vector3{
		factor*v1.X + v2.X,
		factor*v1.Y + v2.Y,
		factor*v1.Z + v2.Z,
	}
}

func (v *Vector3) Copy() Vector3 {
	return Vector3{v.X, v.Y, v.Z}
}

func (v *Vector3) Set(x float64, y float64, z float64) {
	v.X = x
	v.Y = y
	v.Z = z
}

func (v *Vector3) Add(other Vector3) Vector3 {
	return Vector3{
		v.X + other.X,
		v.Y + other.Y,
		v.Z + other.Z,
	}
}

func (v *Vector3) AddScalar(scalar float64) Vector3 {
	return Vector3{
		v.X + scalar,
		v.Y + scalar,
		v.Z + scalar,
	}
}

func (v *Vector3) AddScalars(x float64, y float64, z float64) Vector3 {
	return Vector3{
		v.X + x,
		v.Y + y,
		v.Z + z,
	}
}

func (v *Vector3) Sub(other Vector3) Vector3 {
	return Vector3{
		v.X - other.X,
		v.Y - other.Y,
		v.Z - other.Z,
	}
}

func (v *Vector3) SubScalar(scalar float64) Vector3 {
	return Vector3{
		v.X - scalar,
		v.Y - scalar,
		v.Z - scalar,
	}
}

func (v *Vector3) SubScalars(x float64, y float64, z float64) Vector3 {
	return Vector3{
		v.X - x,
		v.Y - y,
		v.Z - z,
	}
}

func (v *Vector3) Mul(other Vector3) Vector3 {
	return Vector3{
		v.X * other.X,
		v.Y * other.Y,
		v.Z * other.Z,
	}
}

func (v *Vector3) MulScalar(scalar float64) Vector3 {
	return Vector3{
		v.X * scalar,
		v.Y * scalar,
		v.Z * scalar,
	}
}

func (v *Vector3) MulScalars(x float64, y float64, z float64) Vector3 {
	return Vector3{
		v.X * x,
		v.Y * y,
		v.Z * z,
	}
}

func (v *Vector3) Div(other Vector3) Vector3 {
	return Vector3{
		v.X / other.X,
		v.Y / other.Y,
		v.Z / other.Z,
	}
}

func (v *Vector3) DivScalar(scalar float64) Vector3 {
	return Vector3{
		v.X / scalar,
		v.Y / scalar,
		v.Z / scalar,
	}
}

func (v *Vector3) DivScalars(x float64, y float64, z float64) Vector3 {
	return Vector3{
		v.X / x,
		v.Y / y,
		v.Z / z,
	}
}

func (v *Vector3) Negate() Vector3 {
	return Vector3{
		-v.X,
		-v.Y,
		-v.Z,
	}
}

func (v *Vector3) Distance(other Vector3) float64 {
	dx := v.X - other.X
	dy := v.Y - other.Y
	dz := v.Z - other.Z
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func (v *Vector3) Dot(other Vector3) float64 {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z
}

func (v *Vector3) Cross(other Vector3) Vector3 {
	return Vector3{
		v.Y*other.Z - v.Z*other.Y,
		v.Z*other.X - v.X*other.Z,
		v.X*other.Y - v.Y*other.X,
	}
}

func (v *Vector3) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v *Vector3) Normalize() Vector3 {
	m := v.Length()

	if m > 0.0 {
		return v.DivScalar(m)
	} else {
		return v.Copy()
	}
}

func (v *Vector3) Reflect(other Vector3) Vector3 {
	factor := -2.0 * v.Dot(other)
	return Vector3{
		factor*v.X + other.X,
		factor*v.Y + other.Y,
		factor*v.Z + other.Z,
	}
}

func (v *Vector3) Equals(other Vector3) bool {
	return v.X == other.X && v.Y == other.Y && v.Z == other.Z
}

func (v *Vector3) String() string {
	return fmt.Sprintf("Vector3(%f, %f, %f)", v.X, v.Y, v.Z)
}
