package customdb

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sbapi"
	"sort"
)

type CustomDB struct{}

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

func ReadDir(dirname string) ([]os.FileInfo, error) {
	f, err := os.Open(dirname)

	if err != nil {
		return nil, err
	}

	list, err := f.Readdir(-1)

	f.Close()

	if err != nil {
		return nil, err
	}

	sort.Slice(list, func(i, j int) bool { return list[i].ModTime().String() > list[j].ModTime().String() })

	if len(list) > 100 {
		return list[0:100], nil
	}

	return list, nil
}

func (cdb *CustomDB) GetList(time string) []*sbapi.Event {
	list, err := ReadDir("./events")
	if err != nil {
		log.Printf("couldn't get the list of events")
	}

	includedEvents := make([]*sbapi.Event, 0)
	for _, file := range list {
		if file.ModTime().String() > time {
			eventContent, err := ioutil.ReadFile("./events/" + file.Name())
			if err != nil {
				log.Printf("couldn't read event file")
				return nil
			}

			evt := sbapi.Event{}
			err = json.Unmarshal(eventContent, &evt)
			if err != nil {
				log.Printf("couldn't unmarshal data")
				return nil
			}

			includedEvents = append(includedEvents, &evt)
		}
	}

	return includedEvents
}

func (cdb *CustomDB) Save(event *sbapi.Event) (bool, string) {
	return false, ""
}
