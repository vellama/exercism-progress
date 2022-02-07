// Package census simulates a system used to collect census data.
package census

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	resident := Resident{
		Name:    name,
		Age:     age,
		Address: address,
	}

	return &resident
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
	if r.Name == "" {
		return false
	}

	if r.Address["street"] == "" {
		return false
	}

	return true
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
	r.Name = ""
	r.Age = 0
	r.Address = nil
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
	var rCount int = 0

	for i := 0; i < len(residents); i++ {
		if residents[i].HasRequiredInfo() {
			rCount += 1
		}
	}

	return rCount
}
