package  main

import(
	"bufio"
	"fmt"
	"os"
	"net"
	"log"
)

func main() {
	conn ,err := net.Dial("tcp",":3031")
	if err != nil {
		log.Panic(err)
	}
	scanner := bufio.NewScanner(os.Stdin)
   for scanner.Scan() {
	   ln := scanner.Text()
	   if ln == "end" {
		   break
	   }
	   fmt.Fprintln(conn,ln)
	   msg , _ := bufio.NewReader(conn).ReadString('\n')
	   fmt.Printf(msg)
   }
   conn.Close()
}