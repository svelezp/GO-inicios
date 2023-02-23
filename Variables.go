package main
import "fmt"
//el comando "var" se usa oara declarar variables, se pueden declarar varias del mismo tipo en una sola linea, o todas por separado


var (
    name, location string
    age int
)

var name2 string
var age2 int
var location2 string

//Las declaracion de variable puede tener un inicializador
 var (
    name3 string = "NeoStarlight"
    age3 int = 21
    location3 string = "NeoMoona"
 )

//Si hay un inicializador, se puede saltar la definición del tipo de variable

 var (
     name4 = "NeoStarlight"
     age4 = 21
     location4 = "NeoMoona"
 )

 //De igual forma, se pueden inicializar varias variables al tiempo

 var (
    name5, age5, location5 = "NeoStarlight", 21, "NeoMoona"
 )

#//entro de una función, se puede remplazar "var" por :=

func main(){
    name6, location6 := "NeoStarlight",  "NeoMoona"
    age := 32
    fmt.Printf ("%s age %d from %s ", name6, age6, location6)
}