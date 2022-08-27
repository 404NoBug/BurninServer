package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	//reader := bufio.NewReader(os.Stdin)
	//ch := make(chan []string)
	//go func() {
	//	for {
	//		select {
	//		case input := <-ch:
	//			fmt.Println("===============", input)
	//		}
	//	}
	//}()
	//for {
	//	readString, err := reader.ReadString('\n')
	//	if err != nil {
	//		fmt.Println("input err ,check your input and  try again !!!")
	//		continue
	//	}
	//	strings.TrimSpace(readString)
	//	readString = strings.Replace(readString, "\n", "", -1)
	//	readString = strings.Replace(readString, "\r", "", -1)
	//	split := strings.Split(readString, " ")
	//	if len(split) == 0 {
	//		fmt.Println("input err, check your input and  try again !!! ")
	//		continue
	//	}
	//	ch <- split
	//}
	i := uint64(123456789)

	fmt.Println(i)

	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, i)

	fmt.Println(b[:])

	i = uint64(binary.BigEndian.Uint64(b))
}
