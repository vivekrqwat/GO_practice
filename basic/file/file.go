package main

import (
	"os"
)

func main() {
	// f, er := os.Open("example.txt")

	// if er != nil {
	// 	panic(er)
	// }
	// fileinfo, er := f.Stat()
	// if er != nil {
	// 	panic(er)
	// }
	// fmt.Println("file name is", fileinfo.Name())
	// fmt.Println("file size is", fileinfo.Size())
	// defer f.Close()
	// buf := make([]byte, 13)
	// d, er := f.Read(buf)
	// if er != nil {
	// 	panic(er)
	// }
	// fmt.Println("data read is", string(buf[:d]), buf)
	// f1, er := os.ReadFile("example.txt")
	// if er != nil {
	// 	panic(er)
	// }
	// fmt.Println("data read is", string(f1), f1)
	//reading directory

	// dir, er := os.Open(("../"))
	// if er != nil {
	// 	panic(er)
	// }
	// f, e := dir.ReadDir(0)
	// if e != nil {
	// 	panic(er)
	// }
	// fmt.Println("files in the directory are", f)

	//writing of file
	// f, er := os.Create("golang.txt")
	// if er != nil {
	// 	panic(er)
	// }
	// defer f.Close()
	// f.WriteString("hi go hi ")
	// bytes := []byte("hi go how r u")
	// f.Write(bytes)

	// srcFile, er := os.Open("example.txt")
	// if er != nil {
	// 	panic(er)
	// }
	// defer srcFile.Close()
	// buf := make([]byte, 100) // buffer to read data
	// read, er := srcFile.Read(buf)

	// if er != nil {
	// 	panic(er)
	// }

	// writefile, er := os.OpenFile("golang.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	// if er != nil {
	// 	panic(er)
	// }
	// defer writefile.Close()

	// ans1 := string(buf[:read])
	// fmt.Println(ans1, read)
	// _, er = writefile.WriteString(ans1)
	// if er != nil {
	// 	panic(er)
	// }
	os.Remove("example.txt") // remove the source file if needed

}
