package matrix

import (
	"Matrix/inter"
	"Matrix/vector"
	"math"
)

type Matrix4 [16]float64

func NewMatrix4() *Matrix4 {
	m := &Matrix4{}
	m.Set(
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	)
	return m
}
func (m *Matrix4) Set(n11, n12, n13, n14, n21, n22, n23, n24, n31, n32, n33, n34, n41, n42, n43, n44 float64) {
	m[0] = n11
	m[4] = n21
	m[8] = n31
	m[12] = n41
	m[1] = n12
	m[5] = n22
	m[9] = n32
	m[13] = n42
	m[2] = n13
	m[6] = n23
	m[10] = n33
	m[14] = n43
	m[3] = n14
	m[7] = n24
	m[11] = n34
	m[15] = n44
}
func (a *Matrix4) Add(b *Matrix4) Matrix4 {
	m := NewMatrix4()
	m[0] = a[0] + b[0]
	m[4] = a[4] + b[4]
	m[8] = a[8] + b[8]
	m[12] = a[12] + b[12]
	m[1] = a[1] + b[1]
	m[5] = a[5] + b[5]
	m[9] = a[9] + b[9]
	m[13] = a[13] + b[13]
	m[2] = a[2] + b[2]
	m[6] = a[6] + b[6]
	m[10] = a[10] + b[10]
	m[14] = a[14] + b[14]
	m[3] = a[3] + b[3]
	m[7] = a[7] + b[7]
	m[11] = a[11] + b[11]
	m[15] = a[15] + b[15]
	return *m
}
func (a *Matrix4) Subto(b *Matrix4) Matrix4 {
	m := NewMatrix4()
	m[0] = a[0] - b[0]
	m[4] = a[4] - b[4]
	m[8] = a[8] - b[8]
	m[12] = a[12] - b[12]
	m[1] = a[1] - b[1]
	m[5] = a[5] - b[5]
	m[9] = a[9] - b[9]
	m[13] = a[13] - b[13]
	m[2] = a[2] - b[2]
	m[6] = a[6] - b[6]
	m[10] = a[10] - b[10]
	m[14] = a[14] - b[14]
	m[3] = a[3] - b[3]
	m[7] = a[7] - b[7]
	m[11] = a[11] - b[11]
	m[15] = a[15] - b[15]
	return *m
}
func (a *Matrix4) Mul_lamda(k float64) Matrix4 {
	m := *a
	m[0] *= k
	m[4] *= k
	m[8] *= k
	m[12] *= k
	m[1] *= k
	m[5] *= k
	m[9] *= k
	m[13] *= k
	m[2] *= k
	m[6] *= k
	m[10] *= k
	m[14] *= k
	m[3] *= k
	m[7] *= k
	m[11] *= k
	m[15] *= k
	return m
}
func (a *Matrix4) Mul(b *Matrix4) *Matrix4 {

	ans := &Matrix4{}
	ans[0] = a[0]*b[0] + a[1]*b[4] + a[2]*b[8] + a[3]*b[12]
	ans[1] = a[0]*b[1] + a[1]*b[5] + a[2]*b[9] + a[3]*b[13]
	ans[2] = a[0]*b[2] + a[1]*b[6] + a[2]*b[10] + a[3]*b[14]
	ans[3] = a[0]*b[3] + a[1]*b[7] + a[2]*b[11] + a[3]*b[15]

	ans[4] = a[4]*b[0] + a[5]*b[4] + a[6]*b[8] + a[7]*b[12]
	ans[5] = a[4]*b[1] + a[5]*b[5] + a[6]*b[9] + a[7]*b[13]
	ans[6] = a[4]*b[2] + a[5]*b[6] + a[6]*b[10] + a[7]*b[14]
	ans[7] = a[4]*b[3] + a[5]*b[7] + a[6]*b[11] + a[7]*b[15]

	ans[8] = a[8]*b[0] + a[9]*b[4] + a[10]*b[8] + a[11]*b[12]
	ans[9] = a[8]*b[1] + a[9]*b[5] + a[10]*b[9] + a[11]*b[13]
	ans[10] = a[8]*b[2] + a[9]*b[6] + a[10]*b[10] + a[11]*b[14]
	ans[11] = a[8]*b[3] + a[9]*b[7] + a[10]*b[11] + a[11]*b[15]

	ans[12] = a[12]*b[0] + a[13]*b[4] + a[14]*b[8] + a[15]*b[12]
	ans[13] = a[12]*b[1] + a[13]*b[5] + a[14]*b[9] + a[15]*b[13]
	ans[14] = a[12]*b[2] + a[13]*b[6] + a[14]*b[10] + a[15]*b[14]
	ans[15] = a[12]*b[3] + a[13]*b[7] + a[14]*b[11] + a[15]*b[15]
	return ans
}
func (a *Matrix4) SetScale(x, y, z float64) *Matrix4 {
	a.Set(
		x, 0, 0, 0,
		0, y, 0, 0,
		0, 0, z, 0,
		0, 0, 0, 1,
	)
	return a
}
func (a *Matrix4) MulScale(x, y, z float64) *Matrix4 {
	m := NewMatrix4()

	m.Set(
		x, 0, 0, 0,
		0, y, 0, 0,
		0, 0, z, 0,
		0, 0, 0, 1,
	)
	a = a.Mul(m)
	return a
}
func (a *Matrix4) SetTranslation(x, y, z float64) *Matrix4 {
	a.Set(
		1, 0, 0, x,
		0, 1, 0, y,
		0, 0, 1, z,
		0, 0, 0, 1,
	)
	return a

}
func (a *Matrix4) MulTranslation(x, y, z float64) *Matrix4 {
	m := NewMatrix4()
	m.Set(
		1, 0, 0, x,
		0, 1, 0, y,
		0, 0, 1, z,
		0, 0, 0, 1,
	)
	return a.Mul(m)
}
func (a *Matrix4) MulRoTationX(x float64) *Matrix4 {
	m := NewMatrix4()
	var (
		c = float64(math.Cos(inter.UnitAngle * float64(x)))
		s = float64(math.Sin(inter.UnitAngle * float64(x)))
	)
	m.Set(
		1, 0, 0, 0,
		0, c, -s, 0,
		0, s, c, 0,
		0, 0, 0, 1,
	)
	return a.Mul(m)
}
func (a *Matrix4) MulRoTationY(x float64) *Matrix4 {
	m := NewMatrix4()
	var (
		c = float64(math.Cos(inter.UnitAngle * float64(x)))
		s = float64(math.Sin(inter.UnitAngle * float64(x)))
	)
	m.Set(
		c, 0, s, 0,
		0, 1, 0, 0,
		-s, 0, c, 0,
		0, 0, 0, 1,
	)
	return a.Mul(m)
}
func (a *Matrix4) MulRoTationZ(x float64) *Matrix4 {
	m := NewMatrix4()
	var (
		c = float64(math.Cos(inter.UnitAngle * float64(x)))
		s = float64(math.Sin(inter.UnitAngle * float64(x)))
	)
	m.Set(
		c, -s, 0, 0,
		s, c, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	)
	return m.Mul(a)
}
func (a *Matrix4) Transposition() {
	a[3], a[12] = a[12], a[3]
	a[2], a[7], a[8], a[13] = a[8], a[13], a[2], a[2]
	a[1], a[6], a[11], a[4], a[9], a[14] = a[4], a[9], a[14], a[1], a[6], a[11]
}
func (a *Matrix4) MulVector3D(v *vector.Vector3D) *vector.Vector3D {
	vd := &vector.Vector3D{}
	vd.X = a[0]*v.X + a[1]*v.Y + a[2]*v.Z + a[3]*v.W
	vd.Y = a[4]*v.X + a[5]*v.Y + a[6]*v.Z + a[7]*v.W
	vd.Z = a[8]*v.X + a[9]*v.Y + a[10]*v.Z + a[11]*v.W
	vd.W = a[12]*v.X + a[13]*v.Y + a[14]*v.Z + a[15]*v.W
	return vd
}
func (a *Matrix4) LookAt(eye, target, up *vector.Vector3D) *Matrix4 {
	var (
		z *vector.Vector3D
		x *vector.Vector3D
		y *vector.Vector3D
		T *Matrix4 = NewMatrix4()
		R *Matrix4 = NewMatrix4()
	)
	T.Set(
		1, 0, 0, -eye.X,
		0, 1, 0, -eye.Y,
		0, 0, 1, -eye.Z,
		0, 0, 0, 1,
	)
	z = eye.Subto(target).Normal() //.Normal()
	temp := *z
	// temp.X, temp.Y, temp.Z = -z.X, -z.Y, -z.Z
	x = temp.Corss(up).Normal() //.Normal()
	y = up                      //.Normal()
	// math32.NewMatrix4().LookAt()
	R.Set(
		x.X, x.Y, x.Z, 0,
		y.X, y.Y, y.Z, 0,
		z.X, z.Y, z.Z, 0,
		0, 0, 0, 1,
	)
	a = R.Mul(T)
	return a
}
func (a *Matrix4) MakePerspective(fov, aspect, near, far float64) *Matrix4 {
	var (
		ymax float64 = float64(math.Tan(float64(fov*0.5)*inter.UnitAngle)) * near
		ymin float64 = -ymax
		xmax float64 = ymax * aspect
		xmin float64 = ymin * aspect
	)
	a.MakeFrustum(near, far, ymin, ymax, xmax, xmin)
	return a
}
func (a *Matrix4) MakeFrustum(n, f, t, b, l, r float64) *Matrix4 {
	var (
		dx float64 = r - l
		dy float64 = t - b
		dz float64 = n - f
	)
	a.Set(
		(2*n)/dx, 0, -(l+r)/(dx), 0,
		0, (2*n)/dy, -(b+t)/(dy), 0,
		0, 0, (n+f)/dz, -(2*n*f)/dz,
		0, 0, 1, 0,
	)
	return a
}
