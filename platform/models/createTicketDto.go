package models

type CreateTicketDto struct {
	Vatin     string `json:"vatin"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func NewCreateTicketDto(vatin, firstName, lastName string) *CreateTicketDto {
	return &CreateTicketDto{
		Vatin:     vatin,
		FirstName: firstName,
		LastName:  lastName,
	}
}
