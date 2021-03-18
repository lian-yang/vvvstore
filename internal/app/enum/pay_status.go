package enum

type PayStatus int

const (
	Unpay PayStatus = 0
	Payed PayStatus = 1
)

func (p PayStatus) String() string  {
	switch p {
		case Unpay: return "未支付"
		case Payed: return "已支付"
	default:
		return "unknown"
	}
}