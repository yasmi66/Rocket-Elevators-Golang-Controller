package main

//FloorRequestButton is a button on the pannel at the lobby to request any floor
type FloorRequestButton struct {
	ID, floor int
	status, direction string
}

func NewFloorRequestButton(_id int, _floor int, _status string, _direction string) *FloorRequestButton {

	fb := new(FloorRequestButton)

	fb.ID = _id
	fb.status = _status
	fb.floor = _floor
	fb.direction = _direction 

	return fb
}
