package domain

import (
	"gopkg.in/mgo.v2/bson"
	"time"
	"springo/core"
	"fmt"
	"reflect"
)

type GenericInterface interface {
	ChangeId()
	ChangeUI(ui string)
	ChangeGI(gi string)
	ChangeRights(rights int32)
	GetRights() int32
	Value() GenericDomain
	ChangeCreated()
}

type GenericDomain struct {
	GI      string         `json:"gi,omitempty" bson:",omitempty"`
	UI      string         `json:"ui,omitempty" bson:",omitempty"`
	Rights  int32          `json:"rights" bson:",omitempty"`
	Extra   string         `json:"extra" bson:",omitempty"`
	ID      *bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Created time.Time      `json:"created"`
}

func (g *GenericDomain) ChangeId() {
	var id bson.ObjectId = bson.NewObjectId()
	fmt.Println(reflect.TypeOf(g))
	g.ID = &id
}

func (g GenericDomain) ChangeUI(ui string) {
	g.UI = ui
}

func (g GenericDomain) ChangeGI(gi string) {
	g.GI = gi
}

func (g GenericDomain) ChangeRights(rights int32) {
	g.Rights = rights
}

func (g GenericDomain) WithDefaultRights() *GenericDomain {
	g.Rights = core.DEFAULT_RIGHTS.Value
	return &g
}

func (g GenericDomain) GetRights() int32 {
	return g.Rights
}

func (g *GenericDomain) ChangeCreated() {
	g.Created = time.Now()
}

func (g GenericDomain) Value() GenericDomain {
	return g
}