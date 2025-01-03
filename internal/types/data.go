package types

type Data struct {
	Id     string `db:"id" json:"id"`
	UserId string `db:"user_id" json:"-"`
}
