package point

import (
	"Matrix/inter"
)

type Point2D[T inter.I_F] struct {
	X, Y, W T
}
type Point3D[T inter.I_F] struct {
	X, Y, Z, W T
}

func NewPoint2D[T inter.I_F](x, y T) *Point2D[T] {
	return &Point2D[T]{x, y, 1}
}
func NewPoint3D[T inter.I_F](x, y, z T) *Point3D[T] {
	return &Point3D[T]{x, y, z, 1}
}
func (p1 *Point2D[T]) Set(x, y, w T) {
	p1.X = x
	p1.Y = y
	p1.W = w
}
func (p1 *Point3D[T]) Set(x, y, z, w T) {
	p1.X = x
	p1.Y = y
	p1.Z = z
	p1.W = w
}

func (p1 *Point2D[T]) Subto(p2 *Point2D[T]) *Point2D[T] {
	p1.X = p1.X - p2.X
	p1.Y = p1.Y - p2.Y
	p1.W = p1.W - p2.W
	return p1
}
func (p1 *Point3D[T]) Subto(p2 *Point3D[T]) *Point3D[T] {
	p1.X = p1.X - p2.X
	p1.Y = p1.Y - p2.Y
	p1.Z = p1.Z - p2.Z
	p1.W = p1.W - p2.W
	return p1
}
