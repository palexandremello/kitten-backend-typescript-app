package users

// User is the entity about users in application
type User struct {
	ID          string
	Name        *Name
	Email       *Email
	Password    *Password
	Age         *Age
	MonthlyWage *Wage
	UserPicture *UserPicture
}
