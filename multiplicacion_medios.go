package main

import (
  "fmt"
  "math"
  "strconv"
  "math/big"
)
func min(tam1 int, tam2 int) int{
  if tam1 < tam2{
    return tam1
  }else{
    return tam2
  }
}

func max(tam1 int, tam2 int) int{
  if tam1 < tam2{
    return tam2
  }else{
    return tam1
  }
}
var sa int = 0
func multiplica(numero1 []int64,numero2 []int64) (*big.Int){
  num1 := ""
  num2 := ""
  for i := 0; i < len(numero1); i++ {
      num1 = num1+strconv.Itoa(int(numero1[i]))
  }
  for i := 0; i < len(numero2); i++ {
      num2 = num2+strconv.Itoa(int(numero2[i]))
  }
  n1,_ := strconv.ParseInt(num1,10,64)
  n2,_ := strconv.ParseInt(num2,10,64)
  mul1 := big.NewInt(n1)
  mul2 := big.NewInt(n2)
  res := big.NewInt(0)
  res.Mul(mul1,mul2)
  fmt.Println("factores: ",mul1," ",mul2)
  fmt.Println("Mul: ",res)
  return res

}

func conquer(numero1 []int64, numero2 []int64) (*big.Int){
  n := min(len(numero1),len(numero2))
  sa++
  fmt.Println("Entra: ","numero1: ",numero1,"numero2: ", numero2)
  if n < 2 {
    return multiplica(numero1,numero2)
  }else{
    pare1 := int(math.Ceil(float64(float64(len(numero1))/float64(2))))
    pare2 := int(math.Ceil(float64(float64(len(numero2))/float64(2))))
    par1 := int(math.Floor(float64(float64(len(numero1))/float64(2))))
    par2 := int(math.Floor(float64(float64(len(numero2))/float64(2))))
    s := pare1+pare2
    exp := big.NewInt(int64(math.Pow10(s)))
    exp1 := big.NewInt(int64(math.Pow10(pare1)))
    exp2 := big.NewInt(int64(math.Pow10(pare2)))
    res1 := big.NewInt(0)
    res2 := big.NewInt(0)
    res3 := big.NewInt(0)
    res := big.NewInt(0)
    //fmt.Println("par2: ",par2)
    //fmt.Println("s: ",s)
    fmt.Println(numero1[par1:len(numero1)])
    fmt.Println(numero2[par2:len(numero2)])
    //(fmt.Println("1: ","exp: ",exp,"n1: ",numero1[0:par1],"n2: ",numero2[0:par2],"\n")
    res1.Mul(conquer(numero1[0:par1],numero2[0:par2]),exp)
    //fmt.Println("res1: ",res1)
    //fmt.Println("2: ","exp: ",exp1,"n1: ",numero1[0:par1],"n2: ",numero2[par2:len(numero2)],"\n")
    res2.Mul(conquer(numero1[0:par1],numero2[par2:len(numero2)]),exp1)
    //fmt.Println("res2: ",res2)
    //fmt.Println("3: ","exp: ",exp2,"n1: ",numero1[par1:len(numero1)],"n2: ",numero2[0:par2],"\n")
    res3.Mul(conquer(numero1[par1:len(numero1)],numero2[0:par2]),exp2)
    //fmt.Println("res3: ",res3)
    //fmt.Println("3: ","n1: ",numero1[par1:len(numero1)],"n2: ",numero2[par2:len(numero2)],"\n")
    res4 := conquer(numero1[par1:len(numero1)],numero2[par2:len(numero2)])
    //fmt.Println("res4: ",res4)
    res.Add(res,res1)
    res.Add(res,res2)
    res.Add(res,res3)
    res.Add(res,res4)
    return res
  }
}

func main()  {
    var (
      numero1 string
      numero2 string
      num1 []int64
      num2 []int64
    )
    fmt.Println("Primer numero: ")
    fmt.Scanln(&numero1)
    fmt.Println("Segundo numero: ")
    fmt.Scanln(&numero2)
    for i := 0; i < len(numero1); i++ {
      n1,_ := strconv.ParseInt(string(numero1[i]),10,64)
      num1 = append(num1,n1)
    }
    for i := 0; i < len(numero2); i++ {
      n2,_ := strconv.ParseInt(string(numero2[i]),10,64)
      num2 = append(num2,n2)
    }
    res := conquer(num1,num2)
    fmt.Println("final:",res)

}
