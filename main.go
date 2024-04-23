package main

import (
	"errors"
	"fmt"
	"strconv"
)

func getPath() string {
	return "dsadasdas"
}

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Dividing by zero")
	}

	return a / b, nil
}

// func (структура, которую расширяем) имя(аргументы) (возвращаемый тип данных)
func (u User) getData() string {
	return u.name + " " + strconv.Itoa(u.id)
}

func (u *User) setName(name string) {
	u.name = name
}

func printUserData(us User) {
	us.id = 5
	fmt.Println(us)
}

func printUserData2(us *User) {
	us.id = 25
	fmt.Println(us)
	fmt.Println(*us)
}

type Pizdabol interface {
	pizdet() string
}

type Sobaka struct {
	poroda string
	weight int
}

type User struct {
	Sobaka
	id   int
	name string
	age  int
}

func (u User) pizdet() string {
	return "я люблю пиздеть"
}

func spizdanut(p Pizdabol) {
	fmt.Println(p.pizdet())
}

func main() {
	/*	var a int32
		a = 5

		b := 5

		c := getPath()

		fmt.Println("Hello world for a", a)
		fmt.Println("Hello world for b", b)

		fmt.Println("Hello world for c", c)

		if d := getPath(); d != "" {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}

		for i := 1; i < 10; i++ {
			fmt.Print(i)
		}
		fmt.Println()

		j := 1
		for j < 10 {
			fmt.Print(j, " ")
			j++
		}
		fmt.Println()
	*/
	/*var array [3]int
	array[0] = 1
	array[1] = 2
	array[2] = 3

	fmt.Println(array)
	fmt.Println(len(array))
	fmt.Println(cap(array))

	fmt.Println()

		type, len, cap

	slice := make([]int, 2, 5)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice = append(slice, 3, 4, 5)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice = append(slice, 6)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	for key, value := range slice {
		fmt.Println(key, value)
	}

	a := 14
	b := 0

	if dividingResult, err := divide(a, b); err == nil {
		fmt.Println(dividingResult)
	} else {
		fmt.Println(err)
	}*/

	/*fmt.Printf("%+v", user)
	fmt.Printf("%s %d", user.name, user.id)*/

	/*u := User{
		id:   1,
		name: "dsadas",
		age:  15,
	}

	printUserData(u)
	fmt.Println(u)

	printUserData2(&u)

	fmt.Println(u)
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()

	fmt.Println(u.getData())
	u.setName("pups")
	fmt.Println(u.getData())

	fmt.Println(
		User.getData(u),
	)

	User.setName(u, "dsadsa")
	*/
	posts := []string{
		"post 1",
		"post 2",
		"post 3",
		"post 4",
	}

	fmt.Println(posts[2:])
	b := posts[1:3]
	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Println(cap(b))

	fmt.Println(posts[:3])
	fmt.Println(posts[:3])

	newPosts := posts[1:3:4]

	fmt.Println(newPosts)
	fmt.Println(len(newPosts))
	fmt.Println(cap(newPosts))

	strings := append(newPosts, "aa", "bb")
	fmt.Println(strings)
	fmt.Println(len(strings))
	fmt.Println(cap(strings))

	u := User{
		id:   0,
		name: "dsadas",
		age:  0,
		Sobaka: Sobaka{
			poroda: "poroda",
			weight: 100,
		},
	}

	spizdanut(u)

	var hashmp map[string]int

	hashmp = make(map[string]int)

	hashmp["dd"] = 5

	value, ok := hashmp["puk"]

	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("такого ключа нет")
	}

}
