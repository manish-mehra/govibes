package ui

import (
	"github.com/manish-mehra/go-vibes/lib"
	"log"
)

var preference = lib.PreferenceManager{
	Preferences: lib.UserPreferences{
		InputDevice:   "",
		KeyboardSound: "",
	},
	Path: "preference.json",
}

type LoadedPreference struct {
	lastKeyboardSound   string
	lastKeyboardDev     string
	lastKeyboardDevPath string
}

func loadPreferences() (LoadedPreference, error) {

	lp := LoadedPreference{}
	err := preference.InitPreferences()
	if err != nil {
		log.Fatal("error initializing preferences ", err)
		return LoadedPreference{}, err
	}

	// lp.lastKeyboardSound = preference.Preferences.KeyboardSound
	lp.lastKeyboardSound = preference.Preferences.KeyboardSound

	inputDevLs, err := lib.GetDeviceInfoFromProcBusInputDevices()
	if err != nil {
		log.Fatal(err)
		return LoadedPreference{}, err
	}
	// find the list device name based on the path
	for path, devName := range inputDevLs {
		if preference.Preferences.InputDevice == devName {
			lp.lastKeyboardDevPath = path
			lp.lastKeyboardDev = devName
		}
	}
	return lp, nil
}
