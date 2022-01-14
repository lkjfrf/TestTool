package transform

import (
	"math"
) // 3D vector: (x, y, z)
type Vector3 struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

func (this *Vector3) Equal(v Vector3) bool {
	return this.X == v.X && this.Y == v.Y && this.Z == v.Z
}

//  3D vector: set value
func (this *Vector3) Set(x, y, z float64) {
	this.X = x
	this.Y = y
	this.Z = z
}

//  3D vector: copy
func (this *Vector3) Clone() Vector3 {
	return NewVector3(this.X, this.Y, this.Z)
}

//  3D vector: plus
// this = this + v
func (this *Vector3) Add(v Vector3) {
	this.X += v.X
	this.Y += v.Y
	this.Z += v.Z
}

//  3D vector: minus
// this = this - v
func (this *Vector3) Sub(v Vector3) {
	this.X -= v.X
	this.Y -= v.Y
	this.Z -= v.Z
}

//  3D vector: number multiplication
func (this *Vector3) Multiply(scalar float64) {
	this.X *= scalar
	this.Y *= scalar
	this.Z *= scalar
}

func (this *Vector3) Divide(scalar float64) {
	if scalar == 0 {
		panic("The denominator cannot be zero!")
	}
	this.Multiply(1 / scalar)
}

//  3D vector: points
func (this *Vector3) Dot(v Vector3) float64 {
	return this.X*v.X + this.Y*v.Y + this.Z*v.Z
}

//  3D vector: fork
func (this *Vector3) Cross(v Vector3) {
	x, y, z := this.X, this.Y, this.Z
	this.X = y*v.Z - z*v.Y
	this.Y = z*v.X - x*v.Z
	this.Z = x*v.Y - y*v.X
}

//  3D vector: length
func (this *Vector3) Length() float64 {
	return math.Sqrt(this.X*this.X + this.Y*this.Y + this.Z*this.Z)
}

//  3D vector: length square
func (this *Vector3) LengthSq() float64 {
	return this.X*this.X + this.Y*this.Y + this.Z*this.Z
}

//  3D vector: unitization
func (this *Vector3) Normalize() {
	this.Divide(this.Length())
}

//  Returns: New Vector
func NewVector3(x, y, z float64) Vector3 {
	return Vector3{X: x, Y: y, Z: z}
}

//  Return: Zero Vector (0,0,0)
func Zero3() Vector3 {
	return Vector3{X: 0, Y: 0, Z: 0}
}

//  X-axis unit vector
func XAxis3() Vector3 {
	return Vector3{X: 1, Y: 0, Z: 0}
}

//  Y-axis unit vector
func YAxis3() Vector3 {
	return Vector3{X: 0, Y: 1, Z: 0}
}

//  Z-axis unit vector
func ZAxis3() Vector3 {
	return Vector3{X: 0, Y: 0, Z: 1}
}
func XYAxis3() Vector3 {
	return Vector3{X: 1, Y: 1, Z: 0}
}
func XZAxis3() Vector3 {
	return Vector3{X: 1, Y: 0, Z: 1}
}
func YZAxis3() Vector3 {
	return Vector3{X: 0, Y: 1, Z: 1}
}
func XYZAxis3() Vector3 {
	return Vector3{X: 1, Y: 1, Z: 1}
}

//  Returns: a + b vector
func Add3(a, b Vector3) Vector3 {
	return Vector3{X: a.X + b.X, Y: a.Y + b.Y, Z: a.Z + b.Z}
}

//  Returns: a - b vector
func Sub3(a, b Vector3) Vector3 {
	return Vector3{X: a.X - b.X, Y: a.Y - b.Y, Z: a.Z - b.Z}
}

//  Return: a x b vector (X fork)
func Cross3(a, b Vector3) Vector3 {
	return Vector3{X: a.Y*b.Z - a.Z*b.Y, Y: a.Z*b.X - a.X*b.Z, Z: a.X*b.Y - a.Y*b.X}
}

func AddArray3(vs []Vector3, dv Vector3) []Vector3 {
	for i, _ := range vs {
		vs[i].Add(dv)
	}
	return vs
}

func Multiply3(v Vector3, scalars []float64) []Vector3 {
	vs := []Vector3{}
	for _, value := range scalars {
		vector := v.Clone()
		vector.Multiply(value)
		vs = append(vs, vector)
	}
	return vs
}

//  Returns: unitization vector
func Normalize3(a Vector3) Vector3 {
	b := a.Clone()
	b.Normalize()
	return b
} //Seeking two points
func GetDistance(a Vector3, b Vector3) float64 {
	return math.Sqrt(math.Pow(a.X-b.X, 2) + math.Pow(a.Y-b.Y, 2) + math.Pow(a.Z-b.Z, 2))
}
