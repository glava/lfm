package helper

import (
	".././lastfm"
	"fmt"
	"github.com/fatih/color"
	"io"
	"os"
	"strings"
)

func ReadStdin() []string {

	text := make([]string, 50)
	buf := make([]byte, 1024)
	var n int
	var err error
	for err != io.EOF {
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) != 0 {
			break
		}
		n, err = os.Stdin.Read(buf)

		if n > 0 {
			text = append(text, strings.Split(string(buf[0:n]), "\n")...)
		}

	}

	return text
}

func Output(response lastfm.Response, optionalTitle string) {
	if optionalTitle != "" {
		color.Green(optionalTitle)
	}
	//TODO fix this by passing coloring flag
	fmt.Println(response.ToString())
}
