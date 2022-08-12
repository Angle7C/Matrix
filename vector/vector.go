package vector

import (
	"math"
)

type Vector2D struct {
	X, Y, W float64
}

func NewVector2D(x, y float64) *Vector2D {
	v := &Vector2D{x, y, 0}
	return v
}
func (v *Vector2D) Subto(t *Vector2D) *Vector2D {
	v.X -= t.X
	v.Y -= t.Y
	v.W -= t.W
	return v
}
func (v *Vector2D) Add(t *Vector2D) *Vector2D {
	v.X += t.X
	v.Y += t.Y
	v.W += t.W
	return v
}
func (v *Vector2D) Mul_lamda(k float64) *Vector2D {
	v.X *= k
	v.Y *= k
	v.W *= k
	return v
}
func (v *Vector2D) Dot(t *Vector2D) float64 {
	return v.X*t.X + v.Y*t.Y
}
func (v *Vector2D) Corss(t *Vector2D) float64 {
	return v.X*t.Y - t.X*v.Y
}
func (v *Vector2D) Normality() float64 {
	return float64((math.Sqrt(float64(v.Normality2()))))
}
func (v *Vector2D) Normality2() float64 {
	return v.X*v.X + v.Y + v.Y
}
func (v *Vector2D) Normal() *Vector2D {
	nomral := v.Normality()
	v.X /= nomral
	v.Y /= nomral
	return v
}

type Vector3D struct {
	X, Y, Z, W float64
}

func NewVector3D(x, y, z float64) *Vector3D {
	v := &Vector3D{}
	v.X = x
	v.Y = y
	v.Z = z
	v.W = 0
	return v
}
func NewVector3DFromVector(v *Vector3D) *Vector3D {
	return NewVector3D(v.X, v.Y, v.Z)
}
func NewVector3DFromArrayPoint(arr []float64, start int) *Vector3D {
	vd := NewVector3D(arr[start], arr[start+1], arr[start+2])
	vd.W = 1.0
	return vd
}
func NewVector3DFromArrayVector(arr []float64, start int) *Vector3D {
	return NewVector3D(arr[start], arr[start+1], arr[start+2])
}
func (v *Vector3D) Subto(t *Vector3D) *Vector3D {
	temp := Vector3D{}
	temp.X = v.X - t.X
	temp.Y = v.Y - t.Y
	temp.Z = v.Z - t.Z
	temp.W = v.W - t.W
	return &temp
}
func (v *Vector3D) Add(t *Vector3D) *Vector3D {
	temp := Vector3D{}
	temp.X = v.X + t.X
	temp.Y = v.Y + t.Y
	temp.Z = v.Z + t.Z
	temp.W = v.W + t.W
	return &temp
}
func (v *Vector3D) Mul_lamda(k float64) *Vector3D {
	temp := Vector3D{}
	temp.X = v.X * k
	temp.Y = v.Y * k
	temp.Z = v.Z * k
	temp.W = v.W * k
	return &temp
}
func (v *Vector3D) Dot(t *Vector3D) float64 {
	return v.X*t.X + v.Y*t.Y + v.Z*t.Z
}
func (a *Vector3D) Corss(b *Vector3D) *Vector3D {
	return &Vector3D{a.Y*b.Z - a.Z*b.Y, a.Z*b.X - a.X*b.Z, a.X*b.Y - a.Y*b.X, 0}
}
func (v *Vector3D) Normality() float64 {
	return math.Sqrt(v.Normality2())
}
func (v *Vector3D) Normality2() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}
func (v *Vector3D) Normal() *Vector3D {
	nomral := v.Normality()
	v.X /= nomral
	v.Y /= nomral
	v.Z /= nomral
	// v.W /= nomral
	return v
}
func (v *Vector3D) ToArray(arr []float64, start int) {
	arr[start] = v.X
	arr[start+1] = v.Y
	arr[start+2] = v.Z
}
