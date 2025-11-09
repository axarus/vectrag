package domain

type Model struct {
	ID          string
	Name        string
	Slug        string
	Description string
	Fields      []Field
	// Relations   []Relation
	Status        Status
	SchemaVersion int // TODO: This could be a string, what could be the best way to handle this?
}
