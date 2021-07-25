// ---------------------------QC-------------------------------//
package main
import (
 "fmt"
 "net"
 "sort"
 "time"
 "flag"
)
func worker(ports, results chan int, host string) {
   for p := range ports {
      address := fmt.Sprintf("%s:%d", host,p)
      conn, err := net.Dial("tcp", address)
      if err != nil {
        results <- 0
        continue
      }
      conn.Close()
      results <- p
      // fmt.Printf("%d open\n", p)
   }
}

func main() {
    t0 := time.Now()

    ip := flag.String("ip", "127.0.0.1", "ip or host")
    flag.Parse()
    if *ip == "127.0.0.1"{
      fmt.Println("Please enter ip or host")
      fmt.Println("example: scan.exe -ip=10.0.2.15")
    } else{
      ports := make(chan int, 500)
      results := make(chan int)
      var openports []int

      for i := 0; i < cap(ports); i++ {
        go worker(ports, results, *ip)
      }

      go func() {
        for i := 1; i <= 1024; i++ {
          ports <- i
        }
      }()

      for i := 0; i < 1024; i++ {
        port := <-results
        if port != 0 {
          openports = append(openports, port)
        }
      }

      close(ports)
      close(results)
      sort.Ints(openports)
      for _, port := range openports {
          fmt.Printf("%d open\n", port)
      }
      t1 := time.Now()
      fmt.Printf("Run time %v \n", t1.Sub(t0))
    }
}
