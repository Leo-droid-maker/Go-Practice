package main

import (
	"fmt"
)

/* PROBLEMS without interfaces

type EnglishBot struct{}
type SpanishBot struct{}

func main() {
	eb := EnglishBot{}
	sp := SpanishBot{}  <- Its not work

	printGreeting(eb)
	printGreeting(sp)
}

func printGreeting(eb EnglishBot) {  <- Same name of function doesn't work
	fmt.Println(eb.getGreeting())
}

func printGreeting(sp SpanishBot) {
	fmt.Println(sp.getGreeting())
}

func (EnglishBot) getGreeting() string {  < - Same name
	// VERY cuntom logic for generating an english greeting
	return "Hi there"
}

func (SpanishBot) getGreeting() string {
	return "Hola"
}

--------------------------------------------------------------
Another example:

type user struct{
	name string
}

type bot interface {
	getGreeting(string, int) (string, error)
	getBotVersion() float
	respondToUser(user) string
}
--------------------------------------------------------------
*/

type bot interface {
	getGreeting() string // Solution
}

type PowerDrawer interface {
	Draw() int
}

type EnglishBot struct{}
type SpanishBot struct{}

type blender struct{}
type kettle struct{}
type mixer struct{}

func main() {
	eb := EnglishBot{}
	sp := SpanishBot{}
	bl, kt, mx := blender{}, kettle{}, mixer{}

	printGreeting(eb)
	printGreeting(sp)

	plug(bl)
	plug(kt)
	plug(mx)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting()) // Solution
}

func plug(device PowerDrawer) {
	fmt.Println(device.Draw())
}

func (blender) Draw() int {
	return 10
}

func (kettle) Draw() int {
	return 15
}

func (mixer) Draw() int {
	return 9
}

func (EnglishBot) getGreeting() string {
	// VERY cuntom logic for generating an english greeting
	return "Hi there"
}

func (SpanishBot) getGreeting() string {
	return "Hola"
}
