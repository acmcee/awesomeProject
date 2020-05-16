package main

type Behavior interface {
	Run() string
	Eat() string
}

type Dog struct {
	ID int
	Name string
	Age int
}

func (d *Dog)Eat() string {  // 已经隐式实现了接口
	return "eating"
}

func (d *Dog)Run()  string { 	// 已经隐式实现了接口
	return "running"
}


func main() {

	dog := new(Dog)
	dog.Name = "旺财"
	dog.ID = 1
	dog.Age = 5
	dog.Eat()
	dog.Run()

}
