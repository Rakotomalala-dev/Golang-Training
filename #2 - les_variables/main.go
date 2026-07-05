package main

import "fmt"

func main() {
	//1 exemple
	var x int
	x = 15

	//2 exemple
	y := 17

	//3 exemple
	var (
		age  int
		name string
		z    bool
	)
	age = 21
	name = "Tahina"
	z = true

	fmt.Println("exemple 1 est : ", x)
	fmt.Println("exemple 2 est : ", y)
	fmt.Printf("Mon nom est %v, j'ai %v ans, j'ai plus de 18 ans : %v\n", name, age, z)

}

/*

var x int => ici le variable est de type APRES l'identifiant
x = 16 => ici le variable assignée après la déclaraton
y:= 17 => ici le variable déclarée Et assignée

%v	la valeur dans un format par défaut
	lors de l'affichage de structures, le drapeau « plus » (%+v) ajoute les noms des champs
%#v	une représentation de la valeur selon la syntaxe Go
	(les valeurs infinies et NaN en virgule flottante s'affichent sous la forme ±Inf et NaN)
%T	une représentation du type de la valeur selon la syntaxe Go
%%	un signe pourcentage littéral ; ne consomme aucune valeur

*/
