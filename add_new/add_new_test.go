package addnew

import "testing"

func TestAddNew(t *testing.T) {
	type NewPerson struct {
		relationshop Relationship
		person       Person
	}

	tests := []struct {
		name           string
		existedMembers map[Relationship]Person
		newMember      NewPerson
		wantError      bool
	} {
		{
			name: "Test #1: adding child",
			existedMembers: map[Relationship]Person{
				Child: {
					FirstName: "Zaza",
					LastName: "Zerno",
					Age: 5,
				},
			},
			newMember: NewPerson{
				relationshop: Father,
				person: Person{
					FirstName: "Zakabuddin",
					LastName: "Zernovich",
					Age: 34,
				},
			},
			wantError: false,
		},

		{
			name: "Test #1: adding mother and get error",
			existedMembers: map[Relationship]Person{
				Child: {
					FirstName: "Zaza",
					LastName: "Zerno",
					Age: 5,
				},
			},
			newMember: NewPerson{
				relationshop: Mother,
				person: Person{
					FirstName: "Galya",
					LastName: "Zernovna",
					Age: 30,
				},
			},
			wantError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			family := &Family{
				Members: test.existedMembers,
			}

			if err := family.AddNew(test.newMember.relationshop, test.newMember.person); (err != nil) != test.wantError {
				t.Errorf("AddNew() error = %v, wantErr %v", err, test.wantError)
			}
		})
	}
}
