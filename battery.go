package main

import (
    "math"
   )



type Battery struct {
	ID, columnID, elevatorID, floorRequestButtonID, callButtonID int 
	status string
	columnsList [] Column
	floorRequestsButtonsList []FloorRequestButton
	servedFloors []int

	
}

func NewBattery(_id int, _amountOfColumns int, _amountOfFloors int, _amountOfBasements int, _amountOfElevatorPerColumn int) *Battery {

	b := new(Battery)

	b.ID = _id
	b.status = "online"


	b.columnID = 1
	b.elevatorID = 1
	b.floorRequestButtonID = 1
	b.callButtonID = 1


	if _amountOfBasements > 0 {
		b.createBasementFloorRequestButtons(_amountOfBasements)
		b.createBasementColumn(_amountOfBasements, _amountOfElevatorPerColumn)
		_amountOfColumns--
	}

	b.createFloorRequestButtons(_amountOfFloors)
	b.createColumns(_amountOfColumns, _amountOfFloors, _amountOfElevatorPerColumn)

	return b

}

func (b *Battery) createBasementColumn(_amountOfBasements int, _amountOfElevatorPerColumn int ){
	b.servedFloors = make([]int, 0)
	floor := -1 //floors served per column

	for i := 0; i < _amountOfBasements ; i++ {
		b.servedFloors = append(b.servedFloors, floor)
		floor--
	}
	column := NewColumn(b.columnID, "online", _amountOfBasements, _amountOfElevatorPerColumn, b.servedFloors, true)
	b.columnsList = append(b.columnsList, *column)
	b.columnID++
}

func (b *Battery) createColumns(_amountOfColumns int, _amountOfFloors int, _amountOfElevatorPerColumn int){
	amountOfFloorsPerColumn := int(math.Ceil((float64(_amountOfFloors)/(float64(_amountOfColumns)))))
	floor := 1
	for i :=0; i < _amountOfColumns; i++ {
		b.servedFloors = make ([]int, 0)
	
		for x := 0; x< amountOfFloorsPerColumn; x++{
			if floor <= _amountOfFloors{
			b.servedFloors = append(b.servedFloors, floor)
			floor++
			}
		}
		column := NewColumn(b.columnID, "online", _amountOfFloors, _amountOfElevatorPerColumn, b.servedFloors, false)
		b.columnsList = append(b.columnsList, *column)
		b.columnID++
	}
}


func (b *Battery) createFloorRequestButtons(_amountOfFloors int) {
	var buttonFloor = 1
	for y := 0; y < _amountOfFloors; y++{
		floorRequestButton := NewFloorRequestButton(b.floorRequestButtonID, buttonFloor, "OFF", "Down")
		b.floorRequestsButtonsList = append(b.floorRequestsButtonsList, *floorRequestButton)
		buttonFloor--
		b.floorRequestButtonID++
	}
}

func (b *Battery) createBasementFloorRequestButtons(_amountOfBasements int){
	var buttonFloor = -1
	for i := 0; i < _amountOfBasements; i++ {
		basementFloorRequestButton := NewFloorRequestButton(b.floorRequestButtonID, buttonFloor,"OFF", "Down")
		b.floorRequestsButtonsList = append(b.floorRequestsButtonsList, *basementFloorRequestButton)
		buttonFloor--
		b.floorRequestButtonID++
	}

}

func (b *Battery) findBestColumn(_requestedFloor int) *Column {
	for _, column := range b.columnsList {
		if contains(column.servedFloorsList, _requestedFloor){
		return &column 
		}
	}
	return nil
}

//Simulate when a user press a button at the lobby
func (b *Battery) assignElevator(_requestedFloor int, _direction string) (*Column, *Elevator) {
	column := b.findBestColumn(_requestedFloor)
	elevator := column.findElevator(1, _direction)
	elevator.addNewRequest(1)
	elevator.move()

	elevator.addNewRequest(_requestedFloor)
	elevator.move()
	return column, elevator
}


