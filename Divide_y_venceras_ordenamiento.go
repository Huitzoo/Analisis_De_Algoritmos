/*Program that organize an array for the method of divide and conquer*/
package main
import (
  "fmt"
  "strings"
  "strconv"
  "math/rand"
  "math"
  "time"
)
var (
  cadena []int
)
func ordenar(arreglo []int)[]int{
  if len(arreglo) != 1 {
    for i := 0; i < len(arreglo)-1; i++ {
      for j := i+1; j < len(arreglo); j++ {
        if arreglo[j] < arreglo[i]{
          aux := arreglo[i]
          arreglo[i] = arreglo[j]
          arreglo[j] = aux
        }
      }
    }
  }
  return arreglo
}
func divide(div int,arreglo []int) []int{
  if len(arreglo) < int(div) {
    return ordenar(arreglo)
  }
  var (
    res []int
    i,j,k,r int = 0,1,0,0
  )
  var lista = make([][]int,int(div))
  tam := float64(len(arreglo))
  divi := float64(div)
  mod := len(arreglo)%div
  r = int(math.Floor(tam/divi))
  lista[0] = divide(div,arreglo[0:r+mod])
  for i,k,j=r+mod,1,1 ;j<div;j++{
    lista[k] = divide(div,arreglo[i:i+r])
    k++
    i=i+r
    fmt.Println(lista)
  }
  for p := 0; p < int(div); p++ {
    for q := 0;  q< len(lista[p]); q++ {
      res = append(res,int(lista[p][q]))
    }
  }
  return ordenar(res)
}
func main(){
var (
    datos string
    random = rand.New(rand.NewSource(time.Now().UnixNano()))
  )
  fmt.Println("Ingresa el tamaño del arreglo y la division, separados por coma: ")
  fmt.Scanln(&datos)
  arreglo := strings.Split(datos,",")
  tam,_ := strconv.Atoi(arreglo[0])
  div,_ := strconv.Atoi(arreglo[1])
  for i:=0 ; i<tam ; i++{
    cadena = append(cadena,(random.Intn(100)*1))
  }
  fmt.Println("Inicial: ",cadena)
  res := divide(div,cadena)
  fmt.Println("resultado: ",res)
}
