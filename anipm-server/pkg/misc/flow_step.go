package misc

type FlowStep int

const (
	KA FlowStep = iota
	AD
	IBA
	AC
	IB
)

func (s FlowStep) String() string {
	switch s {
	case KA:
		return "一原"
	case AD:
		return "作监"
	case IBA:
		return "二原"
	case AC:
		return "动检"
	case IB:
		return "中割"
	default:
		return "Unknown"
	}
}
