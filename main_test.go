package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func MainCreateHypothesis(t *testing.T) {
	var jsonMap map[string]interface{}
	testJSON :=
		`{
		"Key" : "button_color",
		"Options" : {
			"#FF0000" : 33.3,
			"#00FF00" : 33.3,
			"#0000FF" : 33.3
		}
	}`
	json.Unmarshal([]byte(testJSON), &jsonMap)
	h := createHypothesis(jsonMap)
	if reflect.TypeOf(h).String() != "*Hypothesis" && reflect.TypeOf(h).String() != "*main.Hypothesis" {
		t.Fatal(fmt.Errorf("Type Error expected *Hypothesis, have %T", h))
	}
	if h.Key != "button_color" {
		t.Fatal(fmt.Errorf("Error in creating the structure. Incorrect key expected button_color, have %v", h.Key))
	}
	if h.Options[1].Percent != 33.3 {
		t.Fatal(fmt.Errorf("Error in creating the structure. Incorrect Percent value expected 33.3, have %v", h.Options[1].Percent))
	}
}

func TestGetHypothesis(t *testing.T) {
	MainCreateHypothesis(t)
	key := "button_color"
	user := "333f1"
	option, err := getHypothesis(key, user)
	if err != nil {
		t.Fatal(err)
	}
	o, err := getHypothesis(key, user)
	if err != nil {
		t.Fatal(err)
	}
	if o != option {
		t.Fatal(fmt.Errorf("The error of obtaining a hypothesis. Incorrect option received"))
	}
}

func TestDeleter(t *testing.T) {
	MainCreateHypothesis(t)
	err := deleter("button_color")
	if err != nil {
		t.Fatal(err)
	}
	MainCreateHypothesis(t)
	err = deleter("")
	if err != nil {
		t.Fatal(err)
	}

}
