package main

import "fmt"

func main() {
	sub1 := &Subscriber{"Oliver1999"}
	sub2 := &Subscriber{"Jorah0_0"}
	sub3 := &Subscriber{"Isabella000"}
	sub4 := &Subscriber{"Lily_Oscar"}
	sub5 := &Subscriber{"HarryPotter"}
	blog1 := BloggerX{
		Name:   "Traveler's Diary",
		People: map[string]Observer{},
	}
	blog2 := BloggerX{
		"Reader's Club",
		map[string]Observer{},
	}
	blog3 := BloggerX{
		"The Gamer's Life",
		map[string]Observer{},
	}
	blog1.Subscribe(sub3)
	blog1.Subscribe(sub4)
	blog1.Notify()
	blog2.Subscribe(sub2)
	blog2.Notify()
	blog3.Subscribe(sub1)
	blog3.Subscribe(sub5)
	blog3.Notify()
	fmt.Println("User HarryPotter has unsubscribed ")
	blog3.UnSubscribe(sub5)
	blog3.Notify()
}

type Subject interface {
	Subscribe(Observer)
	UnSubscribe(Observer)
	Notify()
}
type Observer interface {
	Update(blogName string)
	GetName() string
}
type BloggerX struct {
	Name   string
	People map[string]Observer
}

func (blo *BloggerX) Subscribe(people Observer) {
	blo.People[people.GetName()] = people

}

func (blo *BloggerX) UnSubscribe(people Observer) {
	delete(blo.People, people.GetName())

}
func (blo *BloggerX) Notify() {
	for _, people := range blo.People {
		people.Update(blo.Name)
	}

}

type Subscriber struct {
	Name string
}

func (obs *Subscriber) Update(blogName string) {
	fmt.Printf("%s you are subscribed to the channel %s"+"\n", obs.GetName(), blogName)

}

func (obs *Subscriber) GetName() string {
	return obs.Name
}
