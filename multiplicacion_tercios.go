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
  //fmt.Println("factores: ",mul1," ",mul2)
  //fmt.Println("res: ",res)
  return res

}

func conquer(numero1 []int64, numero2 []int64) (*big.Int){
  n := min(len(numero1),len(numero2))
  fmt.Println("Entra: ","numero1: ",numero1,"numero2: ", numero2)
  if n < 3 {
    return multiplica(numero1,numero2)
  }else{
    par1 := int(math.Floor(float64(len(numero1))/float64(3)))
    par2 := int(math.Floor(float64(len(numero2))/float64(3)))
    //fmt.Println("fiv: ",int(math.Ceil(2.0*float64(len(numero1))/float64(3))))
    //fmt.Println("f: ",int(math.Ceil(2.0*float64(len(numero2))/float64(3))))
    pare1 := int(math.Ceil(2.0*float64(len(numero1))/float64(3)))
    pare2 := int(math.Ceil(2.0*float64(len(numero2))/float64(3)))
    pares1 := int(math.Ceil(float64(len(numero1))/float64(3)))
    pares2 := int(math.Ceil(float64(len(numero2))/float64(3)))
    pare := pare1+pare2
    s := big.NewInt(int64(math.Pow10(pare)))
    s1 := big.NewInt(int64(math.Pow10(pare1+pares2)))
    s3 := big.NewInt(int64(math.Pow10(pare2+pares1)))
    s4 := big.NewInt(int64(math.Pow10(pare1)))
    s5 := big.NewInt(int64(math.Pow10(pare2)))
    s6 := big.NewInt(int64(math.Pow10(pares1+pares2)))
    s7 := big.NewInt(int64(math.Pow10(pares1)))
    s8 := big.NewInt(int64(math.Pow10(pares2)))
    res := big.NewInt(0)
    res1 := big.NewInt(0)
    res2 := big.NewInt(0)
    res3 := big.NewInt(0)
    res4 := big.NewInt(0)
    res5 := big.NewInt(0)
    res6 := big.NewInt(0)
    res7 := big.NewInt(0)
    res8 := big.NewInt(0)
    //(fmt.Println("1","s: ",s)
    res1.Mul(conquer(numero1[0:par1],numero2[0:par2]),s)
    //fmt.Println("res1: ",res1)
    //fmt.Println("2","s: ",s1)
    res2.Mul(conquer(numero1[0:par1],numero2[par2:2*par2]),s1)
    //fmt.Println("res2: ",res2)
    //fmt.Println("3","s: ",s4)
    res3.Mul(conquer(numero1[0:par1],numero2[2*par2:len(numero2)]),s4)
    //fmt.Println("res3: ",res3)
    //fmt.Println("4","s: ",s3)
    res4.Mul(conquer(numero1[par1:2*par1],numero2[0:par2]),s3)
    //fmt.Println("res4: ",res4)
    //fmt.Println("5","s: ",s6)
    res5.Mul(conquer(numero1[par1:2*par1],numero2[par2:2*par2]),s6)
    //fmt.Println("res5: ",res5)
    //fmt.Println("6","s: ",s7)
    res6.Mul(conquer(numero1[par1:2*par1],numero2[2*par2:len(numero2)]),s7)
    //fmt.Println("res6: ",res6)
    //fmt.Println("7","s: ",s5)
    res7.Mul(conquer(numero1[2*par1:len(numero1)],numero2[0:par2]),s5)
    //fmt.Println("res7: ",res7)
    //fmt.Println("8","s: ",s8)
    res8.Mul(conquer(numero1[2*par1:len(numero1)],numero2[par2:2*par2]),s8)
    //fmt.Println("res8: ",res8)
    //fmt.Println("9")
    res9 := conquer(numero1[2*par1:len(numero1)],numero2[2*par2:len(numero2)])
    //fmt.Println("res9: ",res9)
    res.Add(res,res1)
    res.Add(res,res2)
    res.Add(res,res3)
    res.Add(res,res4)
    res.Add(res,res5)
    res.Add(res,res6)
    res.Add(res,res7)
    res.Add(res,res8)
    res.Add(res,res9)
    //fmt.Println("res: ",res)
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
