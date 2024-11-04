// TODO: style fixes
package main

import (
	"context"
	"fmt"
	"log"
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
	var cancel context.CancelFunc // holds the cancel function of the previous sound

	// if no args, start application in interactive mode
	if len(args) == 0 {
		ui.Ui_Main()
	} else {
		arg := args[0]
		switch arg {
		// govibes list
		case "list":
			fmt.Printf("%s \n\n", ui.TitleStyle("Available Sounds"))

			audio := lib.GetAudioFilesPath("./audio")
			for key := range audio {
				fmt.Printf("> %s \n", key)
			}
		case "default":
			fmt.Printf("%s \n\n", ui.AsciiTitle)

			loadedPreferences, err := ui.LoadPreferences()
			if err != nil {
				log.Fatal(err)
			}
			if loadedPreferences.LastKeyboardDev != "" && loadedPreferences.LastKeyboardDevPath != "" && loadedPreferences.LastKeyboardSound != "" {
				// get config json & sound file path based on selected sound
				configPaths, err := lib.GetConfigPaths(paths[loadedPreferences.LastKeyboardSound])
				if err != nil {
					panic(err)
				}
				// Cancel previous sound if it's playing
				if cancel != nil {
					cancel()
				}
				ctx, cancel = context.WithCancel(context.Background())
				wg.Add(1)
				go lib.ListenKeyboardInput(ctx, configPaths.ConfigJson, configPaths.SoundFilePath, loadedPreferences.LastKeyboardDevPath)

				fmt.Printf("%s %s \n\n", ui.InputDeviceStyle(loadedPreferences.LastKeyboardDev), ui.SoundStyle(loadedPreferences.LastKeyboardSound))
			}

		case getKeyAsString(paths, arg):
			/**
			configPaths, err := lib.GetConfigPaths(paths[arg])
			  if err != nil {
				panic(err)
			}

			ctx, _ = context.WithCancel(context.Background())
			wg.Add(1)
			go lib.ListenKeyboardInput(ctx, configPaths.ConfigJson, configPaths.SoundFilePath, "/dev/input4")
			**/
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
