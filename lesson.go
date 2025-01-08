// Go:build main
/*
python

class Vertex(object):
	def __init__(self, x, y):
		self.x = x
		self.y = y

	def area(self):
		return self.x * self.y

	def scale(self, i):
		self.x = self.x * i
		self.y = self.y * i


class Vertex3D(Vertex):
	def __init__(self, x, y, z):
		super.__init__(x, y)
		self.z = z

	def area_3d(self):
		return self.x * self.y * self.z

	def scale_3d(self, i):
		self.x = self.x * i
		self.y = self.y * i
		self.z = self.z * i


v = Vertex(3, 4, 5)
v.scale(10)
print(v.area())
print(v.area_3d())
*/

package main

import "fmt"

type Vertex struct {
	x, y int
}

func (v Vertex) Area() int {
	return v.x * v.y
}

func (v *Vertex) Scale(i int) {
	v.x = v.x * i
	v.y = v.y * i
}

type Vertex3D struct {
	Vertex
	z int
}

func (v Vertex3D) Area3D() int {
	return v.x * v.y * v.z
}

func (v *Vertex3D) Scale3D(i int) {
	v.x = v.x * i
	v.y = v.y * i
	v.z = v.z * i
}

func New(x, y, z int) *Vertex3D {
	return &Vertex3D{Vertex{x, y}, z}
}

func main() {

	v := New(3, 4, 5)
	v.Scale3D(10)

	// fmt.Println(v.Area())
	fmt.Println(v.Area3D())

}
