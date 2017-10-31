package main

import (
    "fmt"
)

func main() {
  var low int = 2;
  var high int = 100;
  var flag int = 0;
  
  for low < high {
    flag = 0
    
    for i := 2; i <= low/2; i++ {
      if low % i == 0 {
        flag = 1
	break
      }
    }
    
    if flag == 0 {
      fmt.Printf("%d\n", low);
    }
  
    //fmt.Printf("low:%d, hight: %d, flag:%d\n", low, high, flag)
    
    low++
  }
}
