package event

func (ae ArkEvent) GetColor() int {
	switch ae.Kind {
	case KillEvent:
		return 0xb36f6f
	case TameEvent:
		return 0x4db329
	case AdminCmdEvent:
		return 0x828282
	case JoinEvent:
		return 0x3496fe
	case LeaveEvent:
		return 0x8dacce
	case DefaultEvent:
		fallthrough
	default:
		return 0xdedede
	}
}

func (ae ArkEvent) GetEventTitle() string {
	switch ae.Kind {
	case KillEvent:
		return "Killed"
	case TameEvent:
		return "Tamed"
	case AdminCmdEvent:
		return "AdminCmd"
	case JoinEvent:
		return "User Joined"
	case LeaveEvent:
		return "User Left"
	case DefaultEvent:
		fallthrough
	default:
		return "New Event"
	}
}
