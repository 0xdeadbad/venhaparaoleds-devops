package api

// type IModels interface {
// 	Applicant | Concourse | Vacancy | Profession
// }

// Generic interface so the CRUD can be can technology agnostic
type ApiCrud[T any] interface {
	Create(...any) (T, error)
	Read(...any) (T, error)
	Update(...any) (T, error)
	Delete(...any) (T, error)
}
