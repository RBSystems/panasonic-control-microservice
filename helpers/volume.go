package helpers

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/bobziuchkovski/digest"
	"github.com/byuoitav/common/log"
)

//PanasonicVolumeResponse is a struct for parsing the XML responses
type PanasonicVolumeResponse struct {
	Result  xml.Name `xml:"RESULT"`
	AVolume string   `xml:"AVOLUME"`
	AMute   string   `xml:"AMUTE"`
}

//SetVolume sets the volume of the projector
func SetVolume(address string, volumeLevel uint8) error {

	command := fmt.Sprintf("http://%s/cgi-bin/controlCmd.cgi?param=AVOLUME&value=%v", address, volumeLevel)
	t := digest.NewTransport("byuav", "test")
	req, err := http.NewRequest("GET", command, nil)
	if err != nil {
		log.L.Errorf("Nope Didn't work! - %v", err.Error())
		return err
	}
	_, err = t.RoundTrip(req)
	if err != nil {
		log.L.Errorf("Nope still didn't work! - %v", err.Error())
		return err
	}
	return nil
}

//SetAudioMute sets the projector on Mute, or UnMute
func SetAudioMute(address string, muteValue string) error {

	command := fmt.Sprintf("http://%s/cgi-bin/controlCmd.cgi?param=AMUTE&value=%s", address, muteValue)
	t := digest.NewTransport("byuav", "test")
	req, err := http.NewRequest("GET", command, nil)
	if err != nil {
		log.L.Errorf("Nope Didn't work! - %v", err.Error())
		return err
	}
	_, err = t.RoundTrip(req)
	if err != nil {
		log.L.Errorf("Nope still didn't work! - %v", err.Error())
		return err
	}
	return nil
}

//GetVolume returns the status of the projector, returning if it is on or on standby
func GetVolume(address string) (string, error) {
	command := fmt.Sprintf("http://%s/cgi-bin/queryCmd.cgi?param=AVOLUME", address)

	t := digest.NewTransport("byuav", "test")
	req, err := http.NewRequest("GET", command, nil)
	if err != nil {
		log.L.Errorf("Nope Didn't work! - %v", err.Error())
		return "", err
	}
	resp, err := t.RoundTrip(req)
	if err != nil {
		log.L.Errorf("Nope still didn't work! - %v", err.Error())
		return "", err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.L.Errorf("Error retreiving Body. Error:", err)
		return "", err
	}
	log.L.Debugf("%s", b)
	var status PanasonicVolumeResponse
	err = xml.Unmarshal(b, &status)
	if err != nil {
		log.L.Errorf("Error:", err)
		return "", err
	}

	log.L.Infof("Volume is at: %s", status.AVolume)
	return status.AVolume, nil

}

//GetMute returns the status of the projector, returning if it is on or on standby
func GetMute(address string) (string, error) {
	command := fmt.Sprintf("http://%s/cgi-bin/queryCmd.cgi?param=AMUTE", address)

	t := digest.NewTransport("byuav", "test")
	req, err := http.NewRequest("GET", command, nil)
	if err != nil {
		log.L.Errorf("Nope Didn't work! - %v", err.Error())
		return "", err
	}
	resp, err := t.RoundTrip(req)
	if err != nil {
		log.L.Errorf("Nope still didn't work! - %v", err.Error())
		return "", err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.L.Errorf("Error retreiving Body. Error:", err)
		return "", err
	}
	// log.L.Infof("%s", b)
	var status PanasonicVolumeResponse
	err = xml.Unmarshal(b, &status)
	if err != nil {
		log.L.Errorf("Error:", err)
		return "", err
	}

	log.L.Infof("Mute Status: %s", status.AMute)
	return status.AMute, nil

}
