package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func scanDirectory(path string) {
	fmt.Println(path)

	files, err := ioutil.ReadDir(path)
	if err != nil {
		// 호출에서 발생한 에러의 디버깅 정보 출력
		// fmt.Printf("Returning error from scanDirectory(\"%s\") call\n", path)
		panic(err)
	}

	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			// err := scanDirectory(filePath)
			scanDirectory(filePath)
			//if err != nil {
			//	 fmt.Printf("Returning error from scanDirectory(\"%s\") call\n", path)
			//}
		} else {
			fmt.Println(filePath)
		}
	}
	// return nil
}

func main() {

	if len(os.Args) <= 1 {
		log.Println("Need directory name as 1 argument")
		return
	}

	defer fmt.Println("Directory Scanning is finished")
	directoryName := os.Args[1]

	//err := scanDirectory(directoryName)
	//if err != nil {
	//	log.Fatalln(err)
	//}

	// defer reportPanic()
	scanDirectory(directoryName)
}
