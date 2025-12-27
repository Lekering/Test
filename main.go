package main

type Greeter interface {
	LanguageName() string
	Greet(string) string
}

func SayHallo(name string, interf Greeter) {

}

type Ger struct {
}

func (g Ger) LanguageName() string {
	return "German"
}

func (g Ger) Greet(s string) string {
	return "I can speak " + g.LanguageName() + ": Hallo " + s
}

func main() {
	G := Ger{}
	SayHallo("Dima", G)
}
