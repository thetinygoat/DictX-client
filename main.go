//Copyright 2020 Sachin Saini

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at

//http://www.apache.org/licenses/LICENSE-2.0

//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

func encode(msg string) string {
	tokens := strings.Split(msg, " ")
	var encodedTokens []string
	for _, token := range tokens {
		prefix := "$"
		size := len(token)
		sizeStr := strconv.Itoa(size)
		crlf := "\r\n"
		encodedTokens = append(encodedTokens, prefix+sizeStr+crlf+token+crlf)
	}
	arrPrefix := "#"
	size := strconv.Itoa(len(encodedTokens))
	query := strings.Join(encodedTokens, "")
	return arrPrefix + size + "\r\n" + query
}

func main() {
	port := flag.String("port", "8080", "specify dictX port")
	flag.Parse()
	fmt.Println(*port)
	conn, err := net.Dial("tcp", ":"+*port)
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		buf, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		buf = buf[:len(buf)-1]
		query := encode(buf)
		conn.Write([]byte(query))
	}
}
