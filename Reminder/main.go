package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
)

const (
	Mark_name = "CLI_REMINDER"
	Markvalue = "1"
)

func main() {
	fmt.Println("hello")
	if len(os.Args) < 3 {
		fmt.Println("set timer")
		os.Exit(1)
	}
	now := time.Now()
	w := when.New(nil)
	w.Add(en.All...)
	w.Add(common.All...)
	value, er := w.Parse(os.Args[1], now)
	if er != nil {
		fmt.Println("not parsed")
		os.Exit(2)
	}
	if value == nil {
		fmt.Println("not parsed")
		os.Exit(2)
	}
	if now.After(value.Time) {
		fmt.Println("set reminder for future")
		os.Exit(2)
	}

	diff := value.Time.Sub(now) //we get the difference
	if os.Getenv(Mark_name) == Markvalue {
		time.Sleep(diff)
		err := beeep.Alert("REminder", strings.Join(os.Args[2:], ""), "assets/information.png")
		if err != nil {
			panic(err)
		}
	} else {
		cmd := exec.Command(os.Args[0], os.Args[1:]...)
		cmd.Env = append(cmd.Environ(), fmt.Sprintf("%s=%s", Mark_name, Markvalue))
		err := cmd.Start()
		if err != nil {
			panic(err)
			os.Exit(4)
		}
		fmt.Println("Reminder will be display after", diff.Round(time.Second))
		os.Exit(0)
	}

}
