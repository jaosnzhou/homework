package main

import "fmt"

//var animalList = []string{"fly", "spider", "bird", "cat", "dog", "cow", "horse"}

type animal struct {
	Name     string
	position string
}

var animalList = []animal{
	animal{Name: "spider", position: "first"},
	animal{Name: "bird", position: "middle"},
	animal{Name: "cat", position: "middle"},
	animal{Name: "cow", position: "last"},
}

func main() {
	generateSong(animalList)
}
func generateSong(aList []animal) {
	for i := 0; i < len(aList); i++ {
		openingBar(aList[i])
		if i > 0 && i < len(aList)-1 {
			animalFeature(aList[i])
		}
		for j := i; j >= 1 && j < len(aList)-1; j-- {
			mainLyric(aList[j], aList[j-1])
		}
		if aList[i].position == "last" {
			endingLyric()
			break
		}
		endingBar(aList[0])
	}
}
func openingBar(a animal) {
	if a.position == "first" {
		fmt.Println("There was an old lady who swallowed a " + a.Name + ".")
	} else if a.position == "last" {
		fmt.Println("There was an old lady who swallowed a " + a.Name + "...")
	} else {
		fmt.Println("There was an old lady who swallowed a " + a.Name + ";")
	}

}
func animalFeature(a animal) {
	switch a.Name {
	case "spider":
		fmt.Println("That wriggled and wiggled and tickled inside her.")
	case "bird":
		fmt.Println("How absurd to swallow a bird.")
	case "cat":
		fmt.Println("Fancy that to swallow a cat!")
	case "dog":
		fmt.Println("What a hog, to swallow a dog!")
	case "cow":
		fmt.Println("I don't know how she swallowed a cow!")
	}
}
func mainLyric(a1 animal, a2 animal) {
	if a2.position == "first" {
		fmt.Println("She swallowed the " + a1.Name + " to catch the " + a2.Name + ";")
	} else {
		fmt.Println("She swallowed the " + a1.Name + " to catch the " + a2.Name + ",")
	}
}
func endingLyric() {
	fmt.Println("...She's dead, of course!")
}
func endingBar(a animal) {
	fmt.Println("I don't know why she swallowed a " + a.Name + " - perhaps she'll die!")
}

//
// There was an old lady who swallowed a fly.                    -----openingBar
// I don't know why she swallowed a fly - perhaps she'll die!    -----endingBar
//
// There was an old lady who swallowed a spider;                 -----openingBar
// That wriggled and wiggled and tickled inside her.             -----animalFeature
// She swallowed the spider to catch the fly;                    -----mainLyric
// I don't know why she swallowed a fly - perhaps she'll die!    -----endingBar
//
// There was an old lady who swallowed a horse...                -----openingBar
// ...She's dead, of course!                                     -----endingLyric
