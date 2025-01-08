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

type Human interface {
	Say() string
}

type Person struct {
	Name string
}

type Dog struct {
	Name string
}

func (p *Person) Say() string {
	p.Name = "Mr." + p.Name
	fmt.Println(p.Name)
	return p.Name
}

func DriveCar(human Human) {
	if human.Say() == "Mr.Mike" {
		fmt.Println("Drive")
	} else {
		fmt.Println("Get out")
	}
}

func main() {
	var mike Human = &Person{"Mike"}
	var x Human = &Person{"X"}
	var dog Dog = Dog{"dog"}
	DriveCar(mike)
	DriveCar(x)
	DriveCar(dog)

}
