package phonebook

/*
 * Syntax Go Homework 4
 * Stepanov Anton, 06.05.2019
 */

type phones []int

type Contact struct {
	FirstName string
	LastName  string
}

type Phones struct {
	Mobile string
	Work   string
	Home   string
	Other  string
}

type AddressBookEntity struct {
	Contact Contact
	Phones  Phones
	Email   string
	Address string
	Notes   string
}

type AddressBook []AddressBookEntity

func (ab AddressBook) Len() int {
	return len(ab)
}

func (ab AddressBook) Less(i, j int) bool {
	return ab[i].Contact.LastName < ab[j].Contact.LastName
}

func (ab AddressBook) Swap(i, j int) {
	ab[i], ab[j] = ab[j], ab[i]
}
