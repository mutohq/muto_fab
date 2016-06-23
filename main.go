package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
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

//MakeDirlib creates new directory in the destination directory
func MakeDirlib(fp string) {
	str := destFolder
	cp := fp
	storesub := strings.Split(cp, os.PathSeparator)
	fmt.Println(storesub)

	os.Mkdir(str, 0711)
}

// Copyfilelib copies the file to the destination
func Copyfilelib(fp string) {
	str := destFolder
	cp := fp

	lines, err := readLines(cp)
	if err != nil {
		log.Fatalf("readLines h : %s", err)
	}
	if err := writeLines(lines, str); err != nil {
		log.Fatalf("writeLines h : %s", err)
	}
	fmt.Println("Crap ", str)
}

// WriteLibTodocument writes the content of lib into document/lib
func WriteLibTodocument(fp string, fi os.FileInfo, err error) error {
	if fi.IsDir() {
		MakeDirlib(fp)
		fmt.Println("Src Dir : ", fp)
		// fmt.Println("Dest Dir : ", str)
	} else {
		// Copyfilelib(fp)
		fmt.Println("Src File : ", fp)
		// fmt.Println("Dest File : ", str)
	}
	return nil
}

//MakeDirassests creates new directory in the destination directory
func MakeDirassests(fp string) {
	str := destFolder
	cp := fp
	// storesub := strings.Split(cp, os.PathSeparator)
	// fmt.Println(storesub)

	os.Mkdir(str, 0711)
}

// Copyfileassests copies the file to the destination
func Copyfileassests(fp string) {
	str := destFolder
	cp := fp

	lines, err := readLines(cp)
	if err != nil {
		log.Fatalf("readLines h : %s", err)
	}
	if err := writeLines(lines, str); err != nil {
		log.Fatalf("writeLines h : %s", err)
	}
	fmt.Println("Crap ", str)
}

//WriteAssestsTodocument function
func WriteAssestsTodocument(fp string, fi os.FileInfo, err error) error {
	if fi.IsDir() {
		// MakeDirassests(fp)
		fmt.Println("Src Assests Dir : ", fp)
		// fmt.Println("Dest Dir : ", str)
	} else {
		// Copyfileassests(fp)
		fmt.Println("Src Lib File : ", fp)
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
			srcFolder := checkassests
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

			}
		}
	}
}