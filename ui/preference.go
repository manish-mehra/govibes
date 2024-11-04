package ui

import (
	"github.com/manish-mehra/govibes/lib"
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
	LastKeyboardSound   string
	LastKeyboardDev     string
	LastKeyboardDevPath string
}

func LoadPreferences() (LoadedPreference, error) {

	lp := LoadedPreference{}
	err := preference.InitPreferences()
	if err != nil {
		log.Fatal("error initializing preferences ", err)
		return LoadedPreference{}, err
	}

	// lp.lastKeyboardSound = preference.Preferences.KeyboardSound
	lp.LastKeyboardSound = preference.Preferences.KeyboardSound

	inputDevLs, err := lib.GetDeviceInfoFromProcBusInputDevices()
	if err != nil {
		log.Fatal(err)
		return LoadedPreference{}, err
	}
	// find the list device name based on the path
	for path, devName := range inputDevLs {
		if preference.Preferences.InputDevice == devName {
			lp.LastKeyboardDevPath = path
			lp.LastKeyboardDev = devName
		}
	}
	return lp, nil
}
