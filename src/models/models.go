package models

type Todos struct {
	bun.BaseModel `bun:"table:todos,alias:u"`

	ID   int64 `bun:",pk,autoincrement"`
	Item string
}
