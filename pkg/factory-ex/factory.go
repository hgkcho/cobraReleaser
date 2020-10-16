package main

import (
	"fmt"

	"github.com/bluele/factory-go/factory"
)

type User struct {
	ID       int
	Name     string
	Location string
}

var locations = map[int]string{1: "Kyoto"}

// 'Location: "Tokyo"' is default value.
var UserFactory = factory.NewFactory(
	&User{Location: "Tokyo"},
).SeqInt("ID", func(n int) (interface{}, error) {
	return n, nil
}).Attr("Name", func(args factory.Args) (interface{}, error) {
	user := args.Instance().(*User)
	return fmt.Sprintf("user-%d", user.ID), nil
}).Attr("Location", func(args factory.Args) (interface{}, error) {
	user := args.Instance().(*User)
	location, _ := locations[user.ID]
	if location != "" {
		return location, nil
	}
	return user.Location, nil
})

func main() {
	for i := 0; i < 3; i++ {
		user := UserFactory.MustCreate().(*User)
		fmt.Println("ID:", user.ID, " Name:", user.Name, " Location:", user.Location)
	}
}
