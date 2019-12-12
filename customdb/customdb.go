package customdb

import sbapi "soundboard-api"

type CustomDB struct {}

func (cdb *CustomDB) Get(id string) sbapi.Event {

}

func (cdb *CustomDB) GetList() []*sbapi.Event {

}

func (cdb *CustomDB) Save(event *sbapi.Event) (bool, string) {

}