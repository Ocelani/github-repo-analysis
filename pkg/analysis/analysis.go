package analysis

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"
)

// Repository ...
type Repository struct {
	Name            string
	Owner           string
	URL             string
	CreatedAt       string
	UpdatedAt       string
	StargazerCount  int
	ForkCount       int
	PrimaryLanguage string
	Watchers        int
	Releases        int
}

// Lang ...
type Lang struct {
	StargazerCount     int
	StargazerCountMean int
	ForkCount          int
	ForkCountMean      int
	Watchers           int
	WatchersMean       int
	Releases           int
	ReleasesMean       int
	ReleasesFreq       int
	Age                int
	AgeMean            int
	Lines              int
	LinesMean          int
	Code               int
	CodeMean           int
	Comments           int
	CommentsMean       int
	Blanks             int
	BlanksMean         int
	Complexity         int
	ComplexityMean     int
}

// ReadLanguagesCsv ...
func ReadLanguagesCsv() (*Lang, *Lang) {
	jvch := make(chan [][]string)
	pych := make(chan [][]string)

	java := &Lang{}
	python := &Lang{}

	go func() {
		time.Sleep(200 * time.Millisecond)
		jvch <- ReadCSV("./data/csv/java.csv")
		pych <- ReadCSV("./data/csv/python.csv")
	}()

	i := 0
	for {
		select {
		case jvs := <-jvch:
			for _, jv := range jvs {
				j := NewRepo(jv)
				java.CountInfo(j)
			}
			java.CountAllSlocsCsv("./data/csv/slocs/java/")
			java.MeanInfo()
			java.FreqReleases()
			i++

		case pys := <-pych:
			for _, py := range pys {
				p := NewRepo(py)
				python.CountInfo(p)
			}
			python.CountAllSlocsCsv("./data/csv/slocs/python/")
			python.MeanInfo()
			python.FreqReleases()
			i++

		default:
			time.Sleep(200 * time.Millisecond)
			if i > 1 {
				return java, python
			}
		}
	}
}

// CountInfo ...
func (lang *Lang) CountInfo(repo *Repository) *Lang {
	lang.StargazerCount += repo.StargazerCount
	lang.Watchers += repo.Watchers
	lang.ForkCount += repo.ForkCount
	lang.Releases += repo.Releases
	lang.Age += repo.GetAge()
	return lang
}

// MeanInfo ...
func (lang *Lang) MeanInfo() *Lang {
	lang.StargazerCountMean = lang.StargazerCount / 100
	lang.WatchersMean = lang.Watchers / 100
	lang.ForkCountMean = lang.ForkCount / 100
	lang.ReleasesMean = lang.Releases / 100
	lang.AgeMean = lang.Age / 100
	lang.LinesMean = lang.Lines / 100
	lang.CodeMean = lang.Code / 100
	lang.CommentsMean = lang.Comments / 100
	lang.BlanksMean = lang.Blanks / 100
	lang.ComplexityMean = lang.Complexity / 100
	return lang
}

// FreqReleases ...
func (lang *Lang) FreqReleases() *Lang {
	lang.ReleasesFreq = lang.Releases / lang.Age
	return lang
}

// NewRepo ...
func NewRepo(row []string) *Repository {
	s, _ := strconv.Atoi(row[5])
	f, _ := strconv.Atoi(row[6])
	w, _ := strconv.Atoi(row[8])
	r, _ := strconv.Atoi(row[9])
	return &Repository{
		Name:            row[0],
		Owner:           row[1],
		URL:             row[2],
		CreatedAt:       row[3],
		UpdatedAt:       row[4],
		StargazerCount:  s, // 5
		ForkCount:       f, // 6
		PrimaryLanguage: row[7],
		Watchers:        w, // 8
		Releases:        r, // 9
	}
}

// GetAge ...
func (repo *Repository) GetAge() int {
	born, _ := time.Parse(time.RFC3339, repo.CreatedAt)
	dur := time.Since(born)
	h := dur.Hours()
	return int(h / 8040)
}

// CountSlocs ...
func (lang *Lang) CountSlocs(xx [][]string) *Lang {
	for _, row := range xx {
		lines, _ := strconv.Atoi(row[3])
		code, _ := strconv.Atoi(row[4])
		comments, _ := strconv.Atoi(row[5])
		blanks, _ := strconv.Atoi(row[6])
		complexity, _ := strconv.Atoi(row[7])

		lang.Lines += lines
		lang.Code += code
		lang.Comments += comments
		lang.Blanks += blanks
		lang.Complexity += complexity
	}
	return lang
}

// CountAllSlocsCsv ...
func (lang *Lang) CountAllSlocsCsv(path string) *Lang {
	ch := make(chan [][]string)
	defer close(ch)
	quit := make(chan bool)
	defer close(quit)

	dir, _ := os.Open(path)
	files, _ := dir.Readdirnames(0)
	go func() {
		for _, f := range files {
			ch <- ReadCSV(path + f)
		}
		quit <- true
	}()

	for {
		select {
		case xx := <-ch:
			lang.CountSlocs(xx)
		case <-quit:
			return lang
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}

// ReadCSV ...
func ReadCSV(f string) (d [][]string) {
	df, err := os.Open(f)
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
