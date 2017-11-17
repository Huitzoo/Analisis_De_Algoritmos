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
fmt.Println("ordena: ",arreglo)
  return arreglo
}
func divide(div float64,arreglo []int) []int{
  fmt.Println("arreglo: ",arreglo)
  if len(arreglo) < int(div) { //caso base
    return ordenar(arreglo)
  }
  var (
    res []int
    i,j,k,r int = 0,1,0,0
  )
  var lista = make([][]int,int(div))
  tam := float64(len(arreglo))
  r = int(math.Ceil(tam/div))
  fmt.Println(r)
  for i,k,j=0,0,1 ;(r*j)-1 < len(arreglo);j++{
    lista[k] = divide(div,arreglo[i:(r*j)])
    k++
    i=i+r
    fmt.Println("lista: ",lista)
  }
  if (i<len(arreglo)){
    lista[k] = divide(div,arreglo[i:r*j-1])
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
  fmt.Println("inici: ",cadena)
  res := divide(float64(div),cadena)
  fmt.Println(len(res))
  for i := 0; i < len(res); i++ {
    if res[i] != 0{
      fmt.Print(" ",res[i])
    }
  }
}
