package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "regexp"
    "strings"

    "github.com/urfave/cli"
)

const (
    defaultMessage = "continue?"
    acceptPatternNo = "^[yY]$"
    acceptPatternYes = "^[yY]*$"
)

func curriedDetermineYesNo(acceptPattern string) func (string) int {
    return func (text string) int {
        pattern := regexp.MustCompile(acceptPattern)
        if (pattern.MatchString(text)) {
            return 0
        }
        return 1
    }
}

func createDetermineYesNo(isDefaultYes bool) func (text string) int {
    // create accept regexp
    acceptPattern := acceptPatternNo
    if isDefaultYes {
        acceptPattern = acceptPatternYes
    }
    return curriedDetermineYesNo(acceptPattern)
}

func createShowMessage(message string, isDefaultYes bool) string {
    // select show message
    if isDefaultYes {
        return fmt.Sprintf("%s [Y/n]: ", message)
    }
    return fmt.Sprintf("%s [y/N]: ", message)
}

func doYesNo(showMessage string, determineYesNo func (string) int) (int, error) {
    fmt.Print(showMessage)

    // read yes/no string
    reader := bufio.NewReader(os.Stdin)
    text, err := reader.ReadString('\n')
    if (err != nil) {
        return 1, err
    }
    text = strings.TrimRight(text, "\n")

    return determineYesNo(text), nil
}

func yesNoAction(message string, isDefaultYes bool) (int, error) {
    // create doYesNo parameters
    showMessage := createShowMessage(message, isDefaultYes)
    determineYesNoFunc := createDetermineYesNo(isDefaultYes)

    return doYesNo(showMessage, determineYesNoFunc)
}

func main() {
    app := cli.NewApp()
    app.Name = "yesno"
    app.Usage = "yes/no prompt"
    app.UsageText = "yesno [options]"
    app.Version = "0.1.0"

    // add flags
    app.Flags = []cli.Flag {
        cli.StringFlag{
            Name: "message, m",
            Value: defaultMessage,
            Usage: "specify message",
        },
        cli.BoolFlag{
            Name: "yes, y",
            Usage: "set default to yes",
        },
    }

    // set action
    var exitCode int
    app.Action = func(ctx *cli.Context) error {
        isDefaultYes := ctx.Bool("y")

        var err error
        exitCode, err = yesNoAction(ctx.String("message"), isDefaultYes);
        return err
    }

    // execute
    err := app.Run(os.Args)
    if err != nil {
        log.Fatal(err)
    }
    os.Exit(exitCode)
}
