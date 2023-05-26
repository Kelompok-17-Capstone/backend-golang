package payload

type GetMember struct {
	ID         uint   `json:"id" form:"id"`
	Name       string `json:"name" form:"name"`
	Image      string `json:"image" form:"image"`
	MemberCode string `json:"member_code" form:"member_code"`
}
