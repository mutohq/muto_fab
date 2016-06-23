package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	// "strconv"
	// "os/exec"
	"path/filepath"
	"strings"
)

var destFolder string

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

//MakeDirLib creates new directory in the destination directory
func MakeDirLib(fp string) {
	str := destFolder
	cp := fp
	storesub := strings.Split(cp, string(os.PathSeparator))
	// fmt.Println(storesub)
	var newpath []string
	i := 0
	for {
		if strings.Compare(storesub[i], "muto-fab") == 0 {
			i = i + 1
			break
		} else {
			i = i + 1
		}
	}

	var j int
	for j = i; j < len(storesub); j = j + 1 {
		newpath = append(newpath, storesub[j])
	}

	for _, x := range newpath {
		str = filepath.Join(str, x)
	}
	// fmt.Println("---------------------------------------")
	// fmt.Println("Inside :  ",str)
	os.Mkdir(str, 0711)
}

// CopyfileLib copies the file to the destination
func CopyfileLib(fp string) {

	str := destFolder
	cp := fp
	dir, file := filepath.Split(cp)
	// fmt.Println(dir)
	// fmt.Println(file)
	storesub := strings.Split(dir, string(os.PathSeparator))
	var newpath []string
	i := 0
	for {
		if strings.Compare(storesub[i], "muto-fab") == 0 {
			i = i + 1
			break
		} else {
			i = i + 1
		}
	}

	var j int
	for j = i; j < len(storesub); j = j + 1 {
		newpath = append(newpath, storesub[j])
	}

	for _, x := range newpath {
		str = filepath.Join(str, x)
	}

	str = filepath.Join(str, file)
	// fmt.Println(str)
	lines, err := readLines(fp)
	if err != nil {
		log.Fatalf("readLines h : %s", err)
	}
	if err := writeLines(lines, str); err != nil {
		log.Fatalf("writeLines h : %s", err)
	}

	// fmt.Println("Crap ", str)
}

// WriteLibTodocument writes the content of lib into document/lib
func WriteLibTodocument(fp string, fi os.FileInfo, err error) error {
	if fi.IsDir() {
		MakeDirLib(fp)
		// fmt.Println("Src Dir : ", fp)
		// fmt.Println("Dest Dir : ", str)
	} else {
		CopyfileLib(fp)
		// fmt.Println("Src Lib File : ", fp)
		// fmt.Println("Dest File : ", str)
	}
	return nil
}

//MakeDirassests creates new directory in the destination directory
func MakeDirassests(fp string) {
	str := destFolder
	cp := fp
	storesub := strings.Split(cp, string(os.PathSeparator))
	// fmt.Println(storesub)
	var newpath []string
	i := 0
	for {
		if strings.Compare(storesub[i], "doc-templates") == 0 {
			i = i + 1
			break
		} else {
			i = i + 1
		}
	}

	var j int
	for j = i; j < len(storesub); j = j + 1 {
		newpath = append(newpath, storesub[j])
	}

	for _, x := range newpath {
		str = filepath.Join(str, x)
	}
	// fmt.Println("---------------------------------------")
	// fmt.Println("Inside Assests :  ",str)
	os.Mkdir(str, 0711)
}

// Copyfileassests copies the file to the destination
func Copyfileassests(fp string) {
	str := destFolder
	cp := fp
	dir, file := filepath.Split(cp)
	// fmt.Println(dir)
	// fmt.Println(file)
	storesub := strings.Split(dir, string(os.PathSeparator))
	var newpath []string
	i := 0
	for {
		if strings.Compare(storesub[i], "doc-templates") == 0 {
			i = i + 1
			break
		} else {
			i = i + 1
		}
	}

	var j int
	for j = i; j < len(storesub); j = j + 1 {
		newpath = append(newpath, storesub[j])
	}

	for _, x := range newpath {
		str = filepath.Join(str, x)
	}

	str = filepath.Join(str, file)
	// fmt.Println(str)
	lines, err := readLines(fp)
	if err != nil {
		log.Fatalf("readLines h : %s", err)
	}
	if err := writeLines(lines, str); err != nil {
		log.Fatalf("writeLines h : %s", err)
	}
	// fmt.Println("Crap ", str)
}

//WriteAssestsTodocument function
func WriteAssestsTodocument(fp string, fi os.FileInfo, err error) error {
	if fi.IsDir() {
		MakeDirassests(fp)
		// fmt.Println("Src Assests Dir : ", fp)
		// fmt.Println("Dest Dir : ", str)
	} else {
		Copyfileassests(fp)
		// fmt.Println("Src Assests File : ", fp)
		// fmt.Println("Dest File : ", str)
	}
	return nil
}

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// fmt.Println(pwd)
	// constpath := "/home/ashu/Desktop/intern/git/muto-fab/"
	// var checkdocument, checklib, checkdoc string
	// constpath := pwd + "/"

	checkdocument := filepath.Join(pwd, "document")
	// Initially check fo the file if it doesn't exist , creates it
	if _, err := os.Stat(checkdocument); os.IsNotExist(err) {
		fmt.Println("Creating document directory.....")
		os.Mkdir(checkdocument, 0711)
	}

	checklib := filepath.Join(pwd, "lib")

	if _, err := os.Stat(checklib); err == nil {
		srcFolder := checklib
		destFolder = checkdocument
		// filepath.Walk(srcFolder)

		filepath.Walk(srcFolder, WriteLibTodocument)
		// cpCmd := exec.Command("cp", "-rf", srcFolder, destFolder)
		// err := cpCmd.Run()
		// if err != nil {
		// 	fmt.Println("Couldn't copy lib")
		// 	log.Fatalln("%s", err)
		// }
		fmt.Println("copied lib to document")
	}

	checkdoc := filepath.Join(pwd, "doc-templates")
	checkdocuments := checkdocument
	header := filepath.Join(checkdoc, "header.tmpl")
	footer := filepath.Join(checkdoc, "footer.tmpl")
	checkassests := filepath.Join(checkdoc, "assests")
	// fmt.Println(header, footer, checkassests)
	files, _ := ioutil.ReadDir(checkdoc)

	
	for _, file := range files {

		str := filepath.Join(checkdoc, file.Name())
		if strings.Compare(str, checkassests) == 0 {

			srcFolder := str
			destFolder = checkdocuments

			filepath.Walk(srcFolder, WriteAssestsTodocument)
			// cpCmd := exec.Command("cp", "-rf", srcFolder, destFolder)
			// err := cpCmd.Run()
			// if err != nil {
			// 	fmt.Println("Couldn't copy assests")
			// 	log.Fatalln("%s", err)
			// }
			fmt.Println("copied assests to document")
		} else if strings.Compare(header, str) != 0 && strings.Compare(footer, str) != 0 {

			_, sepr := filepath.Split(str)
			middlefile := sepr
			// fmt.Printf("%q\n", sepr)
			filename := strings.Split(sepr, ".")
			// fmt.Println(filename[1])
			if strings.Compare(filename[1], "tmpl") == 0 {

				destfilename := filename[0] + ".html"
				// fmt.Println(destfilename)
				destpath := filepath.Join(checkdocuments, destfilename)

				lines, err := readLines(header)
				if err != nil {
					log.Fatalf("readLines h : %s", err)
				}
				// for i, line := range lines {
				// 	fmt.Println(i, line)
				// }
				if err := writeLines(lines, destpath); err != nil {
					log.Fatalf("writeLines h : %s", err)
				}

				middle := filepath.Join(checkdoc, middlefile)
				linesmiddle, errm := readLines(middle)
				if errm != nil {
					log.Fatalf("readLines m : %s", err)
				}

				f, err1 := os.OpenFile(destpath, os.O_APPEND|os.O_WRONLY, 0600)
				if err1 != nil {
					panic(err1)
				}
				defer f.Close()

				for _, linem := range linesmiddle {
					if _, err1 = f.WriteString(linem + "\n"); err1 != nil {
						panic(err1)
					}
				}

				linesfooter, errf := readLines(footer)
				if errf != nil {
					log.Fatalf("readLines f : %s", err)
				}
				for _, linefooter := range linesfooter {
					if _, errf = f.WriteString(linefooter + "\n"); errf != nil {
						panic(errf)
					}
				}

				fmt.Println("Created file : ", destfilename)

			} else {
				// fmt.Println(str)
				_, fileExtra := filepath.Split(str)
				destination := filepath.Join(checkdocuments, fileExtra)
				lines, err := readLines(str)
				if err != nil {
					log.Fatalf("readLines h : %s", err)
				}
				if err := writeLines(lines, destination); err != nil {
					log.Fatalf("writeLines h : %s", err)
				}
			}
		}
	}
}

/*
	if strings.Compare(str, checkassests) == 0 {

		srcFolder := str
		destFolder = checkdocuments

		filepath.Walk(srcFolder, WriteAssestsTodocument)
		// cpCmd := exec.Command("cp", "-rf", srcFolder, destFolder)
		// err := cpCmd.Run()
		// if err != nil {
		// 	fmt.Println("Couldn't copy assests")
		// 	log.Fatalln("%s", err)
		// }
		fmt.Println("copied assests to document")
	} else
*/
