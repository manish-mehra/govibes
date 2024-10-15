/*
======================
 Where We Are:
 - bare minimum features are done: list sounds, pass sound flavour, and play keyboard sound

 Next Steps:
 - pass sound flavour as arg from terminal prompt #
 - be able to set default from terminal
 - be able to change the sound flavour without terminating the program

 ======================
*/

package main

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"

	"github.com/manish-mehra/go-vibes/ui"

	"github.com/cqroot/prompt"
	// "github.com/cqroot/prompt/choose"
)

type inputEvent struct {
	Time  [16]byte
	Type  uint16
	Code  uint16
	Value uint32
}

type SoundPack struct {
	ID             string           `json:"id"`
	Name           string           `json:"name"`
	KeyDefineType  string           `json:"key_define_type"`
	IncludesNumpad bool             `json:"includes_numpad"`
	Sound          string           `json:"sound"`
	Defines        map[string][]int `json:"defines"`
}

const baseAudioFilesPath string = "./audio"

func CheckErr(err error) {
	if err != nil {
		if errors.Is(err, prompt.ErrUserQuit) {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		} else {
			panic(err)
		}
	}
}

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

	paths := getAudioFilesPath(baseAudioFilesPath)

	var wg sync.WaitGroup

	// get args
	args := os.Args[1:]

	var cancel context.CancelFunc // holds the cancel function of the previous sound
	var ctx context.Context

	// if no args, start application in interactive mode
	if len(args) == 0 {
		ui.Ui_Main()

		// --------------
		configPaths, err := getConfigPaths(paths["cherrymx-black-abs"])
		if err != nil {
			panic(err)
		}
		// Cancel previous sound if it's playing
		if cancel != nil {
			cancel()
		}

		ctx, cancel = context.WithCancel(context.Background())
		wg.Add(1)
		go listenKeyboardInput(ctx, configPaths.configJson, configPaths.soundFilePath)

	} else {
		arg := args[0]
		switch arg {
		// govibes list
		case "default":
			configPaths, err := getConfigPaths(paths["cherrymx-black-abs"])
			if err != nil {
				panic(err)
			}

			ctx, _ = context.WithCancel(context.Background())
			wg.Add(1)
			go listenKeyboardInput(ctx, configPaths.configJson, configPaths.soundFilePath)
			// govibes nk-cream
		case getKeyAsString(paths, arg):
			configPaths, err := getConfigPaths(paths[arg])
			if err != nil {
				panic(err)
			}

			ctx, _ = context.WithCancel(context.Background())
			wg.Add(1)
			go listenKeyboardInput(ctx, configPaths.configJson, configPaths.soundFilePath)
			fmt.Println(Cyan + "Playing " + Yellow + arg + "🎧" + Reset)
		default:
			fmt.Println("unknown args")
		}
	}

	wg.Wait()
}

func getKeyAsString[T any](mymap map[string]T, arg string) string {
	_, ok := mymap[arg]
	if !ok {
		return ""
	}
	return arg
}

/*
* listent keyboard input from linux file system
 */
func listenKeyboardInput(ctx context.Context, configJsonPath string, soundFilePath string) {

	// TODO: find the right input channel
	file, err := os.Open("/dev/input/event2") // Correct event device for your keyboard
	if err != nil {
		fmt.Println("Error opening input device:", err)
		return
	}
	defer file.Close()

	// open config.json and marshal to json
	jsonFile, err := os.Open(configJsonPath)
	if err != nil {
		panic(err)
	}

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}

	var soundData SoundPack
	if err := json.Unmarshal(jsonData, &soundData); err != nil {
		panic(err)
	}

	var event inputEvent

	for {
		select {
		case <-ctx.Done():
			return
		default:
			// reading for key press from linux file
			err := binary.Read(file, binary.LittleEndian, &event)
			if err != nil {
				fmt.Println("Error reading input event:", err)
				break
			}
			// Check if the event type is EV_KEY (1)
			if event.Type == 1 {
				if event.Value == 1 { // Key press event
					go playAudio(event, soundData, soundFilePath)
				} else if event.Value == 0 { // Key release event
					// let see what can be done here!!!
				}
			}
		}
	}
}

type ConfigPaths struct {
	configJson    string
	soundFilePath string
}

// return json and soundfile path
func getConfigPaths(selectedAudioFile []string) (ConfigPaths, error) {
	var configJsonPath string
	var soundFilePath string
	// set json & ogg sound path
	for _, file := range selectedAudioFile {
		if has := strings.HasSuffix(file, ".json"); has {
			configJsonPath = file
		}
		if has := strings.HasSuffix(file, ".ogg"); has {
			soundFilePath = file
		}
	}

	if configJsonPath == "" || soundFilePath == "" {
		return ConfigPaths{}, errors.New("no confiig json path or sound path found")
	}

	return ConfigPaths{
		configJson:    configJsonPath,
		soundFilePath: soundFilePath,
	}, nil
}

/*
* unmarshal config.json
* and execute sound based on keyevent and config.json mapping
 */
func playAudio(input inputEvent, soundData SoundPack, audioFilePath string) {
	keyCodeStr := strconv.Itoa(int(input.Code))

	if values, ok := soundData.Defines[keyCodeStr]; ok {
		if len(values) >= 2 {

			t1 := fmt.Sprintf("%.3f", float64(values[0])/1000)
			t2 := fmt.Sprintf("%.3f", float64(values[1])/1000)

			cmd := exec.Command("play", audioFilePath, "trim", t1, t2)

			// Use exec.Command to run the aplay command
			err := cmd.Run()

			if err != nil {
				log.Fatalf("Error playing audio: %v", err)
			}

		} else {
			fmt.Println("not enough values")
		}
	} else {
	}

}

/*
* open audio dir and return the map of sound flavours with config and json
 */
func getAudioFilesPath(baseAudioFilesPath string) map[string][]string {
	audioDir, err := os.ReadDir(baseAudioFilesPath)
	if err != nil {
		panic(err)
	}

	paths := make(map[string][]string)
	// audio
	for _, dir := range audioDir {
		// audio/cherrymx-black-abs
		if dir.IsDir() {
			subDir, err := os.ReadDir(baseAudioFilesPath + "/" + dir.Name())
			if err != nil {
				panic(err)
			}

			for _, subFiles := range subDir {
				path := baseAudioFilesPath + "/" + dir.Name() + "/" + subFiles.Name()
				paths[dir.Name()] = append(paths[dir.Name()], path)
			}
		}
	}

	return paths
}
