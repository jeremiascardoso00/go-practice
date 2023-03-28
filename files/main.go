package main

import (
	"io/ioutil"
	"log"
)

func main() {
	// newFile, err := os.Create("aa.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// newFile.Close()

	//

	// newFile, err := os.Create("aa.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// newFile.Close()

	// // checking if file exists
	// _, err = os.Stat("aa.txt")
	// if err != nil {
	// 	if os.IsNotExist(err) {
	// 		log.Fatal("The file does not exist!")
	// 	}
	// }

	// err := os.Rename("data.txt", "info.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = os.Remove("data.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// file, err := os.Open("info.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()

	// // data, err := ioutil.ReadAll(file)
	// // if err != nil {
	// // 	log.Fatal(err)
	// // }

	// // fmt.Printf("Data as string: %s\n", data)

	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }

	// if err := scanner.Err(); err != nil {
	// 	log.Fatal(err)
	// }

	// ioutil.WriteFile() handles creating, opening, writing a slice of bytes and closing the file.
	// if the file doesn't exist WriteFile() creates it
	// and if it already exists the function will truncate it before writing to file.

	bs := []byte("The Go gopher is an iconic mascot!")
	err := ioutil.WriteFile("info.txt", bs, 0644)
	// error handling
	if err != nil {
		log.Fatal(err)
	}
}
