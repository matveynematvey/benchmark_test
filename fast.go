package main

import (
	"bufio"
	"fmt"
	"github.com/mailru/easyjson"
	usertool "hw3_bench/user"
	"io"
	"os"
	"strings"
)

func FastSearch(out io.Writer) {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	seenBrowsers := make(map[string]bool)
	uniqueBrowsers := 0
	foundUsers := strings.Builder{}
	i := 0
	isAndroid := false
	isMSIE := false
	line := make([]byte, 0)
	user := usertool.User{}

	for scanner.Scan() {

		line = scanner.Bytes()
		err := easyjson.Unmarshal(line, &user)
		if err != nil {
			panic(err)
		}

		isAndroid = false
		isMSIE = false

		for _, browser := range user.Browsers {
			if strings.Contains(browser, "MSIE") {
				isMSIE = true
				if _, ok := seenBrowsers[browser]; !ok {
					seenBrowsers[browser] = true
					uniqueBrowsers++
				}
			} else if strings.Contains(browser, "Android") {
				isAndroid = true
				if _, ok := seenBrowsers[browser]; !ok {
					seenBrowsers[browser] = true
					uniqueBrowsers++
				}
			}
		}

		if isAndroid && isMSIE {
			foundUsers.WriteString(fmt.Sprintf("[%d] %s <%s>\n", i, user.Name, strings.ReplaceAll(user.Email, "@", " [at] ")))
		}
		i++
	}

	fmt.Fprintln(out, "found users:\n"+foundUsers.String())
	fmt.Fprintln(out, "Total unique browsers", uniqueBrowsers)
}
