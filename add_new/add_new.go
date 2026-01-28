package addnew

import "errors"

type Relationship string

const (
	Father      = Relationship("father")
	Mother      = Relationship("mother")
	Child       = Relationship("child")
	GrandMother = Relationship("grandMother")
	GrandFather = Relationship("grandFather")
)

type Family struct {
	Members map[Relationship]Person
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

var (
	ErrRelationshipAlreadyExists = errors.New("relationship already exists")
)

func (f *Family) AddNew(r Relationship, p Person) error {
	if f.Members == nil {
		f.Members = map[Relationship]Person{}
	}
	if _, ok := f.Members[r]; ok {
		return ErrRelationshipAlreadyExists
	}
	f.Members[r] = p
	return nil
}
