package mongo

import (
	"github.com/yaronsumel/simpleuser/server/user"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const db = "simpledb"
const simpleuserns = "simpleuserns"

type mongo struct {
	session *mgo.Session
}

func NewHandler() *mongo {
	m := &mongo{}
	//
	mgoSession, err := mgo.DialWithInfo(&mgo.DialInfo{
		PoolLimit: 4096,
		Timeout:   time.Second,
		FailFast:  true,
		Addrs:     []string{"localhost:27017"},
	})
	if err != nil {
		panic(err)
	}
	// Switch the session to a monotonic behavior.
	mgoSession.SetMode(mgo.Monotonic, true)
	m.session = mgoSession
	return m
}

func (m *mongo) Insert(u *user.Object) error {
	s := m.session.Clone()
	defer s.Close()
	selector := bson.D{bson.DocElem{Name: "email", Value: u.Email}}
	update := bson.M{
		"$set": bson.M{"email": u.Email, "name": u.Name},
		"$inc": bson.M{"timesreceived": 1},
	}
	_, err := m.session.DB(db).C(simpleuserns).Upsert(selector, update)
	return err
}
