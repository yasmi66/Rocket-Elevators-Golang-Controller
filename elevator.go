package main

type Elevator struct {

	ID, amountOfFloors, currentFloor int
	status, direction string
	overweight bool
	Door Door
	floorRequestsList []int
	completedRequestsList []int
}

func NewElevator(_id int, _amountOfFloors int, _currentFloor int, _status string) *Elevator {
	
	e := new(Elevator)

	e.ID = _id
	e.status = _status
	e.amountOfFloors = _amountOfFloors
	e.currentFloor = _currentFloor
	e.Door = Door{_id, "closed"}
	e.floorRequestsList = make([]int, 0)

	e.direction = "null"
	e.overweight = false

	return e

}

func (e *Elevator) move() {
	for len(e.floorRequestsList) != 0 {
		var destination = e.floorRequestsList[0]
		e.status = "moving"
		e.sortFloorList()
		if e.direction == "up" {
			for e.currentFloor < destination {
					e.currentFloor++}
			} else if e.direction == "down"{
				for e.currentFloor > destination {
					e.currentFloor--
				}
			}
			e.status = "stopped"
			e.operateDoors()
			e.floorRequestsList = e.floorRequestsList[1:]
			e.completedRequestsList = append(e.completedRequestsList, destination)
		}
	e.status = "idle"
	e.direction = "empty"
}

func (e *Elevator) sortFloorList() { // this function left as is for tests to pass
	if e.direction == "up" {

	}

}


func (e *Elevator) operateDoors(){
    if e.overweight{
        e.status = "obstructed"
    }else{//if not obstructed return to idle
        e.status = "idle"
    }
}


func (e *Elevator) addNewRequest(requestedFloor int) {
	if !contains(e.floorRequestsList, requestedFloor) {
		e.floorRequestsList = append(e.floorRequestsList, requestedFloor)
	}
	if e.currentFloor < requestedFloor {
		e.direction = "up"
	}
	if e.currentFloor > requestedFloor {
		e.direction = "down"
	}
}