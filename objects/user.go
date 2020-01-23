package objects

type MenuId int

const (
	Menu_Init        MenuId = 0
	Menu_AskLocation MenuId = 100
	Menu_Feed        MenuId = 200
	Menu_Post        MenuId = 300
)

type User struct {
	UserId int64
	MenuId MenuId
}
