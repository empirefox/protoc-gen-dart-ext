package zero

func (m *Affect) BuildOn(plan string) bool {
	if m == nil {
		return true
	}
	if m.GetAll() {
		return true
	}
	if m.GetNone() {
		return false
	}
	return plan == "" || plan == m.GetPlan()
}
