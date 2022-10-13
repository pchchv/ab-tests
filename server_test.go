package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
)

func Test(t *testing.T) {
	testURL = "http://" + getEnvValue("HOST") + ":" + getEnvValue("PORT")
	testClient = http.Client{}
}

func TestServerPing(t *testing.T) {
	res, err := http.Get(testURL + "/ping")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("status not OK")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.Error(err)
		}
	}(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	b := string(body)
	if !strings.Contains(b, "AB-tests") {
		t.Fatal()
	}
}

func TestLoadPing(t *testing.T) {
	rate := vegeta.Rate{Freq: 1000, Per: time.Second}
	duration := 5 * time.Second
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    testURL + "/ping",
	})
	attacker := vegeta.NewAttacker()
	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
		metrics.Add(res)
	}
	metrics.Close()
	log.Printf("99th percentile: %s\n", metrics.Latencies.P99)
}

func TestAllHypothesis(t *testing.T) {
	res, err := http.Get(testURL + "/all")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("status not OK")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.Error(err)
		}
	}(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
}

func TestLoadAllHypothesis(t *testing.T) {
	rate := vegeta.Rate{Freq: 1000, Per: time.Second}
	duration := 5 * time.Second
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    testURL + "/all",
	})
	attacker := vegeta.NewAttacker()
	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
		metrics.Add(res)
	}
	metrics.Close()
	log.Printf("99th percentile: %s\n", metrics.Latencies.P99)
}

func TestCreateHypothesis(t *testing.T) {
	body := []byte(`{
		"Key" : "button_color",
		"Options" : {
			"#FF0000" : 33.3,
			"#00FF00" : 33.3,
			"#0000FF" : 33.3
		}
	}`)
	res, err := http.Post(testURL+"/create", "application/json", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("status not OK")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.Error(err)
		}
	}(res.Body)
	if err != nil {
		t.Fatal(err)
	}
}

func TestLoadCreate(t *testing.T) {
	rate := vegeta.Rate{Freq: 1000, Per: time.Second}
	duration := 5 * time.Second
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "POST",
		URL:    testURL + "/create",
		Body: []byte(`{
			"Key" : "price",
			"Options" : {
				"10" : 70.0,
				"20" : 15.0,
				"22" : 10.0,
				"25" : 5.0
			}
		}`),
	})
	attacker := vegeta.NewAttacker()
	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
		metrics.Add(res)
	}
	metrics.Close()
	log.Printf("99th percentile: %s\n", metrics.Latencies.P99)
}

// TODO: Correct the test. EOF error occurs.
/*
func TestGetUserHypothesis(t *testing.T) {
	body := []byte(`[{"hypothesis":  "button_color"}, {"user": "354f1"}]`)
	req, err := http.NewRequest(http.MethodPatch, testURL+"/forUser?", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	res, err := testClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("status not OK")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.Error(err)
		}
	}(res.Body)
	if err != nil {
		t.Fatal(err)
	}
}*/

func TestLoadGetUserHypothesis(t *testing.T) {
	user := "354f1"
	rate := vegeta.Rate{Freq: 1000, Per: time.Second}
	duration := 5 * time.Second
	user += "0"
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "PATCH",
		URL:    testURL + "/forUser?hypothesis=button_color,user=" + user,
	})
	attacker := vegeta.NewAttacker()
	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
		metrics.Add(res)
	}
	metrics.Close()
	log.Printf("99th percentile: %s\n", metrics.Latencies.P99)
}

func TestDeleteOne(t *testing.T) {
	body := []byte(`[{"hypothesis":  "button_color"}]`)
	req, err := http.NewRequest(http.MethodDelete, testURL+"/one?", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	res, err := testClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("status not OK")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.Error(err)
		}
	}(res.Body)
	if err != nil {
		t.Fatal(err)
	}
}

func TestLoadDeleteOne(t *testing.T) {
	rate := vegeta.Rate{Freq: 1000, Per: time.Second}
	duration := 5 * time.Second
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "DELETE",
		URL:    testURL + "/one?hypothesis=button_color",
	})
	attacker := vegeta.NewAttacker()
	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
		metrics.Add(res)
	}
	metrics.Close()
	log.Printf("99th percentile: %s\n", metrics.Latencies.P99)
}

func TestDelete(t *testing.T) {
	req, err := http.NewRequest(http.MethodDelete, testURL+"/", nil)
	if err != nil {
		t.Fatal(err)
	}
	res, err := testClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("status not OK")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.Error(err)
		}
	}(res.Body)
	if err != nil {
		t.Fatal(err)
	}
}
