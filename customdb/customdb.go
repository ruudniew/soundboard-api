package customdb

import (
  "sbapi"
  "io/ioutil"
  "log"
  "encoding/json"
)

type CustomDB struct {}

func (cdb *CustomDB) Get(id string) *sbapi.Event {
  // insecure if id isn't checked to not contain a special character
  // example: id = "../api" would read "../api.json"
  dat, err := ioutil.ReadFile("./events/" + id + ".json")
  
  if err != nil {
    log.Printf("couldn't read the file by id.json: %+v", err)
    return nil
  }

  evt := sbapi.Event{}
  err = json.Unmarshal(dat, &evt)

  if err != nil {
    log.Printf("couldn't unmarshal the evt while getting event by id: %+v", err)
    return nil
  }

  return &evt
}

func (cdb *CustomDB) GetList() []*sbapi.Event {
    return nil
}

func (cdb *CustomDB) Save(event *sbapi.Event) (bool, string) {
  return false, ""
}