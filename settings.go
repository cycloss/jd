package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

const (
	qFlagVal  = false
	qFlagExpl = "Silences the output of the program (the filename of the page created)."
	tFlagExpl = "Adds the given `string` as a tag to the file which can be easily targeted by the grep utility. `grep -Er Tags:.*(<tag1>|<tag2>) .` will list the files in the current directory where <tag1> and <tag2> are given. This flag can be specified multiple times."
	fFlagExpl = "Changes the filename from DD-MM-YYYY-jd.md to the string given as the flag's value."
)

type settings struct {
	quiet    bool
	filename string
	filepath string
	tags     strSliceFlag
	title    string
	date     string
}

func (s *settings) toTemplateArgs() *templateArgs {
	return &templateArgs{
		Title: s.title,
		Date:  s.date,
		Tags:  s.tags,
	}
}

func (s *settings) filenameFullPath() string {
	return fmt.Sprintf("%s%s", s.filepath, s.filename)
}

func getSettings() *settings {
	var tags strSliceFlag
	// backticks used in a description are put as the value type
	// in the flags usage description.
	// flag.Var adds a random `value` word for some reason, so give empty backticks to remove
	flag.Var(&tags, "t", "``")
	flag.Var(&tags, "tag", tFlagExpl)

	// bool flags do not show their value type in the usage description
	var quiet bool
	flag.BoolVar(&quiet, "q", qFlagVal, "")
	flag.BoolVar(&quiet, "quiet", qFlagVal, qFlagExpl)

	var filename string
	// empty back ticks remove flag value type
	flag.StringVar(&filename, "f", "", "``")
	flag.StringVar(&filename, "filename", "", fFlagExpl)

	// allows changing the usage description. Click through on `Usage` to see default
	// must be set before `flag.Parse``
	flag.Usage = func() {
		fmt.Fprint(flag.CommandLine.Output(), "Usage: jd [OPTIONS] [TITLE]\n\n")
		fmt.Fprint(flag.CommandLine.Output(), "A tool for assisting in journaling by creating structured, partially filled markdown format journal templates\n\n")
		fmt.Fprint(flag.CommandLine.Output(), "Options:\n")
		flag.PrintDefaults()
		fmt.Fprint(flag.CommandLine.Output(), "\n\nOther information:\n")
		fmt.Fprint(flag.CommandLine.Output(), "- The JD_PATH environment variable can be set to an absolute path, which is where new files will always be outputted to.\n- The template is located at `/usr/local/etc/jd/journalTemplate.txt`, and modifications can be made to its structure; however, care should be taken when modifying template components.")
	}

	flag.Parse()

	if filename == "" {
		filename = generateFilename()
	}

	pt := generatePageTitle()
	shortDate := getShortDate()
	filepath := getFilePath()

	return &settings{quiet: quiet, tags: tags, filename: filename, filepath: filepath, title: pt, date: shortDate}
}

type strSliceFlag []string

func (i *strSliceFlag) String() string {
	return fmt.Sprint(*i)
}

func (i *strSliceFlag) Set(value string) error {
	// must dereference pointer, as we are passing in a pointer
	// to a slice, which is itself a pointer
	*i = append(*i, value)
	return nil
}

func generatePageTitle() string {
	args := flag.Args() // gets non flag args
	argCount := len(args)
	// return arg as title if found
	if argCount == 1 {
		return args[0]
	}
	if argCount > 1 {
		exitFatal("jd accepts only 0-1 arguments")
	}
	now := time.Now()
	weekday := now.Weekday()
	day := now.Day()

	suffix := "th"
	switch day {
	case 1, 21, 31:
		suffix = "st"
	case 2, 22:
		suffix = "nd"
	case 3, 23:
		suffix = "rd"
	}

	month := now.Month()
	year := now.Year()

	return fmt.Sprintf("%s %d%s of %s %d", weekday, day, suffix, month, year)
}

func getFilePath() string {
	path := os.Getenv(JD_PATH_KEY)
	if path == "" {
		return "./"
	}
	// last character should be a slash
	// if it's not, add it
	last := path[len(path)-1:]
	if last != "/" {
		path = path + "/"
	}
	return path
}

func generateFilename() string {
	shortDate := getShortDate()
	return fmt.Sprintf("%s-jd.md", shortDate)
}
