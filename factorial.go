/*Program that get the factorial for the method of divide and conquer*/
package main
import (
  "fmt"
  "math"
  "math/big"
  "os"
)
func multiplica(fac []*big.Int) (*big.Int){
  mul := big.NewInt(1)
  for i := 0;  i< len(fac); i++ {
    mul.Mul(mul,fac[i])
  }
  return mul
}

func divide(fac []int64, div int64) (*big.Int){
  var (
    j,i int64
    res []*big.Int
  )
  fmt.Println(fac)
  if int64(len(fac)) < div {
    if len(fac) == 1 {
      return big.NewInt(fac[0])
    }else{
      var rm []*big.Int
      for p:=0;p<len(fac);p++{
        rm = append(rm,big.NewInt(fac[p]))
      }
      return multiplica(rm)
    }
  }else{
    tam := float64(len(fac))
    divi := float64(div)
    mod := int64(len(fac))%div
    r := int64(math.Floor(tam/divi))
    res = append(res,divide(fac[0:r+mod],div))
    for j,i=0,r+mod; j < div-1; j++{
        fmt.Println(fac[i:i+r])
        res = append(res,divide(fac[i:i+r],div))
        i = i+r
    }
  }
  return multiplica(res)
}
func main(){
  var (
    num int64
    div int64
    factorial []int64
    i int64 = 0
  )
  fmt.Println("Factorial: ")
  fmt.Scanln(&num)
  fmt.Println("Divide by: ")
  fmt.Scanln(&div)
  if div == 1 || div == 0 {
    fmt.Println("Not valid, div have to > 2")
    os.Exit(0)
  }
  for i=1; i<=num; i++ {
    factorial = append(factorial,i)
  }
  fmt.Println("Resultado: " ,divide(factorial,div))
}
