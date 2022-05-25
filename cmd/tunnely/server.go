package main

// func main() {
// 	l, err := net.Listen("tcp", ":6969")
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer l.Close()

// 	for {
// 		c, err := l.Accept()
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}

// 		go func() {
// 			// for {
// 			reader := bufio.NewReader(c)
// 			request, err := http.ReadRequest(reader)
// 			if err != nil {
// 				fmt.Println(err)
// 				return
// 			}

// 			fmt.Println(request)
// 			response := http.Response{
// 				StatusCode: 200,
// 			}

// 			response.Write(c)
// 			// c.Close()
// 			// }
// 		}()
// 	}
// }
