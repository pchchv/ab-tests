package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"

	"github.com/joho/godotenv"
)

type Hypothesis struct {
	Id      int
	Key     string
	Options []Option
}

type Option struct {
	Name    string
	Percent float64
	UsersId []string
}

var repository = NewInMemoryRepository()

func init() {
	// Load values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Panic("No .env file found")
	}
}

func getEnvValue(v string) string {
	// Getting a value. Outputs a panic if the value is missing
	value, exist := os.LookupEnv(v)
	if !exist {
		log.Panicf("Value %v does not exist", v)
	}
	return value
}

func createHypothesis(jsonMap map[string]interface{}) *Hypothesis {
	h := new(Hypothesis)
	var o Option
	h.Key = fmt.Sprint(jsonMap["Key"])
	iter := reflect.ValueOf(jsonMap["Options"]).MapRange()
	for iter.Next() {
		o.Name = iter.Key().String()
		value, err := strconv.ParseFloat(fmt.Sprint(iter.Value().Interface()), 64)
		if err != nil {
			log.Panic(err)
		}
		o.Percent = value
		h.Options = append(h.Options, o)
	}
	repository.Create(h)
	return h
}

func getHypothesis(hypothesis string, userId string) (string, error) {
	var option string
	h, err := repository.GetByTitle(hypothesis)
	if err != nil {
		return option, errors.New("Hypothesis not found")
	}
	options := h.Options
	// If one of the options is not already in use, use it
	for _, o := range options {
		if len(o.UsersId) == 0 {
			o.UsersId = append(o.UsersId, userId)
			repository.Update(h)
			return o.Name, nil
		}
	}
	// TODO: Implement an algorithm for allocating users to options
	return option, nil
}

func main() {
	server()
}
