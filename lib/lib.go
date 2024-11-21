package lib

import (
	"bufio"
	"context"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func PrintHelp() string {
	return `Govibes - Mechanical Keyboard Sound Simulator
An unnecessary rewrite of mechvibes.com disguised as a CLI tool

Usage:
  govibes [command]

Commands:
  sounds           List available keyboard sound profiles
                   Example output:
                   > cherrymx-brown-pbt
                   > cherrymx-red-abs
                   > cherrymx-red-pbt
                   > eg-oreo

  <profile>        Play a specific keyboard sound profile
                   Example: govibes eg-oreo
                   If no sound is selected, interactive mode will be launched

  default          Play the last used sound profile and input device

  (no command)     Run interactive mode
                   In interactive mode, you can:
                   - Select sound profiles
                   - Change input audio channel

  help             Display this help information

Examples:
  govibes sounds        # List available sound profiles
  govibes eg-oreo       # Play 'eg-oreo' sound profile
  govibes default       # Play last used sound and input device
  govibes               # Enter interactive mode

Project Repository: https://github.com/manish-mehra/govibes`
}

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

const BaseAudioFilesPath string = "../audio"

type ConfigPaths struct {
	ConfigJson    string
	SoundFilePath string
}

// return json and soundfile path
func GetConfigPaths(selectedAudioFile []string) (ConfigPaths, error) {
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
		ConfigJson:    configJsonPath,
		SoundFilePath: soundFilePath,
	}, nil
}

/*
* execute sound based on keyevent and config.json mapping
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
func GetAudioFilesPath(baseAudioFilesPath string) map[string][]string {
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

/*
* listen keyboard input from linux file system
 */
func ListenKeyboardInput(ctx context.Context, configJsonPath string, soundFilePath string, keyboardFilePath string) {

	file, err := os.Open(keyboardFilePath) // Correct event device for your keyboard
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

// getDeviceInfoFromProcBusInputDevices parses the /proc/bus/input/devices file
// to map each event device to its human-readable name.
func GetDeviceInfoFromProcBusInputDevices() (map[string]string, error) {
	deviceInfo := make(map[string]string)

	file, err := os.Open("/proc/bus/input/devices")
	if err != nil {
		return nil, fmt.Errorf("error reading /proc/bus/input/devices: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var currentName string
	var currentEvent string

	// Regular expressions to match the relevant lines
	nameRegex := regexp.MustCompile(`^N: Name="(.+)"`)
	eventRegex := regexp.MustCompile(`H: Handlers=.*(event[0-9]+)`)

	for scanner.Scan() {
		line := scanner.Text()

		// Check for device name
		if nameMatches := nameRegex.FindStringSubmatch(line); len(nameMatches) > 1 {
			currentName = nameMatches[1]
		}

		// Check for event handler (eventX)
		if eventMatches := eventRegex.FindStringSubmatch(line); len(eventMatches) > 1 {
			currentEvent = eventMatches[1]
			deviceInfo["/dev/input/"+currentEvent] = currentName
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error scanning /proc/bus/input/devices: %w", err)
	}

	return deviceInfo, nil
}

type UserPreferences struct {
	InputDevice   string `json:"input_device"`
	KeyboardSound string `json:"keyboard_sound"`
}

type PreferenceManager struct {
	Preferences UserPreferences
	Path        string
}

func (s *PreferenceManager) InitPreferences() error {
	file, err := os.Open(s.Path)
	if err != nil {
		return fmt.Errorf("error opening settings.json: %w", err)
	}
	defer file.Close()

	// Read the file content
	preferenceJson, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error reading json file: %w", err)
	}

	// Unmarshal JSON data into s.config
	if err = json.Unmarshal(preferenceJson, &s.Preferences); err != nil {
		return fmt.Errorf("error unmarshalling json file: %w", err)
	}
	return nil
}

func (s *PreferenceManager) UpdatePreferences(newPreference UserPreferences) error {
	// Open the file in read-write mode
	file, err := os.OpenFile(s.Path, os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("error opening json file: %w", err)
	}
	defer file.Close()

	// Read the current contents of the file
	preferenceJson, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error reading json file: %w", err)
	}

	// Unmarshal the file content into the Settings struct
	if len(preferenceJson) > 0 {
		if err := json.Unmarshal(preferenceJson, &s.Preferences); err != nil {
			return fmt.Errorf("error unmarshalling json file: %w", err)
		}
	}

	// Update only specific fields
	if newPreference.InputDevice != "" {
		s.Preferences.InputDevice = newPreference.InputDevice
	}
	if newPreference.KeyboardSound != "" {
		s.Preferences.KeyboardSound = newPreference.KeyboardSound
	}

	// Seek to the beginning of the file and truncate it
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		return fmt.Errorf("error seeking json file: %w", err)
	}
	if err := file.Truncate(0); err != nil {
		return fmt.Errorf("error truncating json file: %w", err)
	}

	// Marshal the updated config and write it back to the file
	updatedPreferenceJSON, err := json.MarshalIndent(s.Preferences, "", "\t")
	if err != nil {
		return fmt.Errorf("error marshalling updated json file: %w", err)
	}

	// Write the updated JSON to the file
	if _, err := file.Write(updatedPreferenceJSON); err != nil {
		return fmt.Errorf("error writing to json file: %w", err)
	}

	return nil
}
