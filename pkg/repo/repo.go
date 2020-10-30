package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"gopkg.in/src-d/go-git.v4"
)

// Node is the repository type
type Node struct {
	Repository struct {
		Name           string
		URL            string
		CreatedAt      string
		UpdatedAt      string
		StargazerCount int
		ForkCount      int
		Owner          struct {
			Login string
		}
		PrimaryLanguage struct {
			Name string
		}
		Watchers struct {
			TotalCount int
		}
		Releases struct {
			TotalCount int
		}
	}
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
	// AppendAllCsv()
}

func run() (err error) {
	r := make(chan [][]string, 2)
	// r <- ReadCSV("java")
	r <- ReadCSV("python")
	pydata := <-r
	close(r)

	// java := make(chan string, 100)
	// quitJv := make(chan string)
	// go ForEachLanguage(java, quitJv, jvdata, "java")

	python := make(chan string, 100)
	quitPy := make(chan string)
	go ForEachLanguage(python, quitPy, pydata, "python")

	// var j, p int
	for {
		select {
		// case <-java:
		// fmt.Println(<-java)
		case <-python:
			fmt.Println(<-python)
		default:
			qp := <-quitPy
			if qp == "quit" {
				return
			}
		}
	}
}

// ReadCSV ...
func ReadCSV(f string) (d [][]string) {
	df, err := os.Open("./data/csv/" + f + ".csv")
	if err != nil {
		panic(err)
	}
	defer df.Close()

	r := csv.NewReader(df)
	d, err = r.ReadAll()
	if err != nil {
		panic(err)
	}
	return
}

// ForEachLanguage ...
func ForEachLanguage(ch, quit chan string, lang [][]string, name string) {
	for _, r := range lang {
		repo, err := CloneRepository(r, name)
		if err != nil {
			fmt.Printf("Error while clone: %e", err)
		}
		ch <- repo
	}
	quit <- "quit"
	close(ch)
}

// CloneRepository ...
func CloneRepository(r []string, l string) (repo string, err error) {
	repo = fmt.Sprintf("%s-%s", r[0], r[1])
	url := fmt.Sprintf("%s.git", r[2])
	// Tempdir to clone the repository
	dir, err := ioutil.TempDir("./repositories", repo)
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir) // clean up

	// Clones the repository into the given dir, just as a normal git clone does
	_, err = git.PlainClone(dir, false, &git.CloneOptions{
		URL: url,
	})
	if err != nil {
		log.Fatal(err)
	}
	if err = WriteData(dir, repo, l); err != nil {
		fmt.Printf("Error while writing data for repository: %s | %s", repo, url)
	}
	return repo, err
}

// WriteData ...
func WriteData(dir string, repo string, l string) (err error) {
	ch := make(chan error, 3)
	ch <- ExecCommand("csv", dir, repo, l)
	// ch <- ExecCommand("tabular", dir, repo)
	// ch <- ExecCommand("html", dir, repo)e
	if err = <-ch; err != nil {
		fmt.Println(err)
	}
	close(ch)

	return
}

// ExecCommand ...
func ExecCommand(ext, dir, repo, l string) (err error) {
	cmd := exec.Command("scc", "-f", ext, "-o", "./../../"+ext+"/"+l+"/"+repo+"."+ext)
	cmd.Dir = dir
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput

	if err = cmd.Run(); err != nil {
		fmt.Printf("Error while clone: %e", err)
	}
	fmt.Print(string(cmdOutput.Bytes()))

	return
}

// // AppendAllCsv write a single file with all csv data collected.
// func AppendAllCsv() {
// 	dir, _ := ioutil.ReadDir("./csv")
// 	var mr io.Reader

// 	f, err := os.OpenFile("./data.csv", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	rw := bufio.NewReadWriter(bufio.NewReader(mr), bufio.NewWriter(f))

// 	for _, file := range dir {
// 		b, err := ioutil.ReadFile("./csv/" + file.Name())
// 		if err != nil {
// 			log.Fatalln(err)
// 		}
// 		rw.Write(b)
// 		fmt.Printf("- %q\n", file)
// 	}
// 	rw.Flush()
// }

// func writeTXT() {
// 	cmd := exec.Command("scc", "-f", "html", "-o", "./../output.html")
// 	cmd.Dir = dir
// 	if err != nil {
// 		os.Stderr.WriteString(err.Error())
// 	}

// 	cmdOutput := &bytes.Buffer{}
// 	cmd.Stdout = cmdOutput
// 	err = cmd.Run()
// 	if err != nil {
// 		os.Stderr.WriteString(err.Error())
// 	}
// 	fmt.Print(string(cmdOutput.Bytes()))
// }

// func writeHTML() {
// 	cmd := exec.Command("scc", "-f", "html", "-o", "./../output.html")
// 	cmd.Dir = dir
// 	if err != nil {
// 		os.Stderr.WriteString(err.Error())
// 	}

// 	cmdOutput := &bytes.Buffer{}
// 	cmd.Stdout = cmdOutput
// 	err = cmd.Run()
// 	if err != nil {
// 		os.Stderr.WriteString(err.Error())
// 	}
// 	fmt.Print(string(cmdOutput.Bytes()))
// }

// func writeCSV() {
// 	cmd := exec.Command("scc", "-f", "html", "-o", "./../output.html")
// 	cmd.Dir = dir
// 	if err != nil {
// 		os.Stderr.WriteString(err.Error())
// 	}

// 	cmdOutput := &bytes.Buffer{}
// 	cmd.Stdout = cmdOutput
// 	err = cmd.Run()
// 	if err != nil {
// 		os.Stderr.WriteString(err.Error())
// 	}
// 	fmt.Print(string(cmdOutput.Bytes()))
// }
