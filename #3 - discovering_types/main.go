package main

import "fmt"


func main()  {

	var (
		b bool
		s string
		i int
		u uint // c'est tous simplement pour les nombre positive
		u8 uint8 // c'est pour les nombres positif entre 0 à 255
		i8 int8  // c'est pour les nombres negatif entre -128 à -127
		i16 int16 
		u16 uint16 
		f float32
		)	

	b = true
	s = "Tahina"
	i = -15
	u = 15
	u8 = 25
	i8 = 127
	i16 = 21500
	u16 = 40000
	f = 3.14

	fmt.Println(b, s, i, u, u8, i8, i16, u16, f)
}



/*


string => pour les textes 
bool => pour booléan
int => pour les nombre mais il y a beaucoup de different type dans le int 
int8, uint8, int16, uint16, int32, uint32, int64, uint64, int, uint, uintptr => pour les entiers
float32, float64 => pour les Décimaux
complex64, complex128 => pour complexes
Byte => alias de uint8
Rune => alias de int32 

*/