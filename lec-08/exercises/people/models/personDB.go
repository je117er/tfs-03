package models

type PersonDBModel struct {
	ID           uint
	TrainerName  string
	TrainerEmail string
	TraineeName  string
	TraineeEmail string
}

func (PersonDBModel) TableName() string {
	return "persons"
}
