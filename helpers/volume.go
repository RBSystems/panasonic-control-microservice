package helpers

import (
	"fmt"
	"net/http"

	"github.com/bobziuchkovski/digest"
	"github.com/byuoitav/common/log"
)

func SetVolume(address string, volumeLevel uint8) error {

	command := fmt.Sprintf("http://%s/cgi-bin/controlCmd.cgi?param=AVOLUME&value=%v", address, volumeLevel)
	t := digest.NewTransport("byuav", "test")
	req, err := http.NewRequest("GET", command, nil)
	if err != nil {
		log.L.Debugf("Nope Didn't work! - %v", err.Error())
		return err
	}
	_, err = t.RoundTrip(req)
	if err != nil {
		log.L.Debugf("Nope still didn't work! - %v", err.Error())
		return err
	}
	return nil
}

//Function sets the projector on Mute, or UnMute
func SetAudioMute(address string, muteValue string) error {

	command := fmt.Sprintf("http://%s/cgi-bin/controlCmd.cgi?param=AMUTE&value=%s", address, muteValue)
	t := digest.NewTransport("byuav", "test")
	req, err := http.NewRequest("GET", command, nil)
	if err != nil {
		log.L.Debugf("Nope Didn't work! - %v", err.Error())
		return err
	}
	_, err = t.RoundTrip(req)
	if err != nil {
		log.L.Debugf("Nope still didn't work! - %v", err.Error())
		return err
	}
	return nil
}
