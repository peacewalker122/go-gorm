package enum

type UserStatus string

const (
	Active   UserStatus = "active"
	Deactive UserStatus = "deactive"
)

func (u UserStatus) String() string {
	return string(u)
}
