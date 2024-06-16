package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net"

	"github.com/hhow09/protohackers/server"
)

// problem: https://protohackers.com/problem/1

func main() {
	s := server.New("tcp", "localhost:8080")
	err := s.Handle(handlePrimeTime)
	if err != nil {
		log.Fatal(err)
	}
}

type request struct {
	Method string   `json:"method"`
	Number *float64 `json:"number"`
}

type response struct {
	Method string `json:"method"`
	Prime  bool   `json:"prime"`
}

func handlePrimeTime(conn net.Conn) {
	defer conn.Close()
	bfio := bufio.NewReader(conn)
	for {
		buf, err := bfio.ReadBytes('\n')
		if err != nil {
			// EOF
			fmt.Println(fmt.Errorf("could not read data: %w", err))
			break
		}
		buf = buf[:len(buf)-1]
		req := &request{}
		if err := json.Unmarshal(buf, req); err != nil {
			_, err = conn.Write([]byte{0x00})
			if err != nil {
				log.Fatal(fmt.Errorf("could not unmarshal data: %w", err))
			}
			return
		}
		if req.Method != "isPrime" {
			_, err = conn.Write([]byte{0x00})
			if err != nil {
				log.Fatal(fmt.Errorf("invalid method: %s", req.Method))
			}
			return
		}
		if req.Number == nil {
			_, err = conn.Write([]byte{0x00})
			if err != nil {
				log.Fatal(fmt.Errorf("number is nil"))
			}
			return
		}
		ip := isPrime(*req.Number)

		resp := &response{
			Method: "isPrime",
			Prime:  ip,
		}
		respBytes, err := json.Marshal(resp)
		if err != nil {
			log.Fatal(err)
			return
		}
		_, err = conn.Write(append(respBytes, '\n'))
		if err != nil {
			log.Fatal(err)
		}
	}

}

func isPrime(n float64) bool {
	if math.Trunc(n) != n {
		return false
	}
	if n < 2 {
		return false
	}
	num := int64(n)
	var i int64
	for i = 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
