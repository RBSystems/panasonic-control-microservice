package helpers

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/bobziuchkovski/digest"
	se "github.com/byuoitav/av-api/statusevaluators"
	"github.com/byuoitav/common/log"
	"github.com/byuoitav/common/nerr"
)

//PanasonicVolumeResponse is a struct for parsing the XML responses
type PanasonicVolumeResponse struct {
	Result  xml.Name `xml:"RESULT"`
	AVolume string   `xml:"AVOLUME"`
	AMute   string   `xml:"AMUTE"`
}

//SetVolume sets the volume of the projector
func SetVolume(address string, volumeLevel float64) error {
	log.L.Infof("Setting Volume to %v on %s", volumeLevel, address)

	if volumeLevel > 100 || volumeLevel < 0 {
		return nerr.Create(fmt.Sprintf("Invalid volume level %v: must be in range 0-100", volumeLevel), "params")
	}

	//Do some quick mafs for optimizing the volume since the panasonic can only take volume from 0-62

	level := (volumeLevel / 100) * 62

	command := fmt.Sprintf("http://%s/cgi-bin/controlCmd.cgi?param=AVOLUME&value=0%v", address, level)
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
func GetVolume(address string) (se.Volume, error) {
	command := fmt.Sprintf("http://%s/cgi-bin/queryCmd.cgi?param=AVOLUME", address)

	t := digest.NewTransport("byuav", "test")
	req, err := http.NewRequest("GET", command, nil)
	if err != nil {
		log.L.Errorf("Nope Didn't work! - %v", err.Error())
		return se.Volume{}, err
	}
	resp, err := t.RoundTrip(req)
	if err != nil {
		log.L.Errorf("Nope still didn't work! - %v", err.Error())
		return se.Volume{}, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.L.Errorf("Error retreiving Body. Error:", err)
		return se.Volume{}, err
	}
	log.L.Debugf("%s", b)
	var status PanasonicVolumeResponse
	err = xml.Unmarshal(b, &status)
	if err != nil {
		log.L.Errorf("Error:", err)
		return se.Volume{}, err
	}

	log.L.Infof("Volume is at: %s", status.AVolume)

	level, err := strconv.Atoi(status.AVolume)
	if err != nil {
		return se.Volume{}, err
	}

	return se.Volume{Volume: level}, nil

}

//GetMute returns the status of the projector, returning if it is on or on standby
func GetMute(address string) (se.MuteStatus, error) {
	command := fmt.Sprintf("http://%s/cgi-bin/queryCmd.cgi?param=AMUTE", address)
	log.L.Infof("Getting mute status of %s...", address) //Print that the device is powering on

	t := digest.NewTransport("byuav", "test")
	req, err := http.NewRequest("GET", command, nil)
	if err != nil {
		log.L.Errorf("Nope Didn't work! - %v", err.Error())
		return se.MuteStatus{}, err
	}
	resp, err := t.RoundTrip(req)
	if err != nil {
		log.L.Errorf("Nope still didn't work! - %v", err.Error())
		return se.MuteStatus{}, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.L.Errorf("Error retreiving Body. Error:", err)
		return se.MuteStatus{}, err
	}
	// log.L.Infof("%s", b)
	var response PanasonicVolumeResponse
	err = xml.Unmarshal(b, &response)
	if err != nil {
		log.L.Errorf("Error:", err)
		return se.MuteStatus{}, err
	}

	log.L.Infof("Mute Status: %s", response.AMute)

	if response.AMute == "ON" {
		return se.MuteStatus{Muted: true}, nil
	} else {
		return se.MuteStatus{Muted: false}, nil
	}

}
