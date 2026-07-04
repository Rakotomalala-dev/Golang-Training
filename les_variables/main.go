package main

import "fmt"

func main() {
	var x int

	x = 15

	fmt.Printf("mon nombre est : %v", x)

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
