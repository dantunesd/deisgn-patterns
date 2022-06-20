package gear

type StateReverse struct {
	gear *Gear
}

func NewStateReverse(gear *Gear) *StateReverse {
	return &StateReverse{
		gear: gear,
	}
}

func (s *StateReverse) Accelerate() {
	println("Accelerating backwards")
}

func (s *StateReverse) Brake() {
	println("Braking the car")
}

func (s *StateReverse) SwitchToNeutral() {
	s.gear.setNewState(NewStateNeutral(s.gear))
}

func (s *StateReverse) SwitchToDrive() {
	s.gear.setNewState(NewStateDrive(s.gear))
}

func (s *StateReverse) SwitchToReverse() {
	println("The gear is already Reverse")
}
