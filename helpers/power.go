package helpers

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/bobziuchkovski/digest"
	"github.com/byuoitav/av-api/statusevaluators"
	"github.com/byuoitav/common/log"
)

//PanasonicPowerResponse is a struct to help with XML parsing
type PanasonicPowerResponse struct {
	Result xml.Name `xml:"RESULT"`
	Power  string   `xml:"POWER"`
}

//SetPower sets the power of the projector, turning it either on or off.
func SetPower(address string, powerValue string) error {

	command := fmt.Sprintf("http://%s/cgi-bin/controlCmd.cgi?param=POWER&value=%s", address, powerValue)
	t := digest.NewTransport("byuav", "test")
	req, err := http.NewRequest("GET", command, nil)
	if err != nil {
		log.L.Info("Nope Didn't work! - %v", err.Error())
		return err
	}
	_, err = t.RoundTrip(req)
	if err != nil {
		log.L.Infof("Nope still didn't work! - %v", err.Error())
		return err
	}
	return nil
}

//GetPower gets the status of the projector, returning if it is on or on standby
func GetPower(address string) (statusevaluators.PowerStatus, error) {
	command := fmt.Sprintf("http://%s/cgi-bin/queryCmd.cgi?param=POWER", address)

	t := digest.NewTransport("byuav", "test")
	req, err := http.NewRequest("GET", command, nil)
	if err != nil {
		log.L.Infof("Nope Didn't work! - %v", err.Error())
		return statusevaluators.PowerStatus{}, err
	}
	resp, err := t.RoundTrip(req)
	if err != nil {
		log.L.Info("Nope still didn't work! - %v", err.Error())
		return statusevaluators.PowerStatus{}, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.L.Infof("Error retreiving Body. Error:", err)
		return statusevaluators.PowerStatus{}, err
	}
	log.L.Debugf("%s", b)
	var response PanasonicPowerResponse
	err = xml.Unmarshal(b, &response)
	if err != nil {
		log.L.Info("Error:", err)
		return statusevaluators.PowerStatus{}, err
	}
	var status statusevaluators.PowerStatus

	if response.Power == "OFF" {
		status.Power = "standby"
	} else if response.Power == "ON" {
		status.Power = "on"
	}

	log.L.Infof("Power status: %s", status.Power)

	return status, nil

}
