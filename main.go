package main

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/manish-mehra/go-vibes/lib"
	"github.com/manish-mehra/go-vibes/ui"
)

const baseAudioFilesPath string = "./audio"

// Color codes
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Cyan   = "\033[36m"
	White  = "\033[37m"
)

func main() {

	paths := lib.GetAudioFilesPath(baseAudioFilesPath)

	var wg sync.WaitGroup
	// get args
	args := os.Args[1:]
	var ctx context.Context
	// var cancel context.CancelFunc
	// if no args, start application in interactive mode
	if len(args) == 0 {
		ui.Ui_Main()
	} else {
		arg := args[0]
		switch arg {
		// govibes list
		case getKeyAsString(paths, arg):
			configPaths, err := lib.GetConfigPaths(paths[arg])
			if err != nil {
				panic(err)
			}

			ctx, _ = context.WithCancel(context.Background())
			wg.Add(1)
			go lib.ListenKeyboardInput(ctx, configPaths.ConfigJson, configPaths.SoundFilePath)
			fmt.Println(Cyan + "Playing " + Yellow + arg + "🎧" + Reset)
		default:
			fmt.Println("unknown args")
		}
	}

	wg.Wait()
	// cancel()
}

func getKeyAsString[T any](mymap map[string]T, arg string) string {
	_, ok := mymap[arg]
	if !ok {
		return ""
	}
	return arg
}
