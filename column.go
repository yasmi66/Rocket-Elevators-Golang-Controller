package main

import (
    "math"
   )

   var callButtonID int = 1
   var elevatorID int = 1

type Column struct {

    ID int
    status string
    servedFloorsList []int
    isBasement bool
    elevatorsList []*Elevator
    callButtonsList []CallButton
}

func NewColumn(_id int, _status string, _amountOfFloors int,_amountOfElevators int, _servedFloors []int, _isBasement bool) *Column {
	
    c := new(Column)

    c.ID = _id
    c.status = _status 
    c.servedFloorsList = _servedFloors
   

    c.createElevators(_amountOfFloors, _amountOfElevators)
    c.createCallButtons(_amountOfFloors, _isBasement)

    return c
}

func (c *Column) createCallButtons(_amountOfFloors int, _isBasement bool){
    if _isBasement {
        var buttonFloor int = -1
        for i:= 0; i < _amountOfFloors; i++ {
            callButton := NewCallButton(callButtonID, "OFF", buttonFloor, "up")
            c.callButtonsList = append(c.callButtonsList, *callButton)
            buttonFloor--
            callButtonID++
        } 
    }else{
        var buttonFloor = 1
        for i := 0; i < _amountOfFloors; i++ {
            callButton := NewCallButton(callButtonID, "OFF", buttonFloor, "down")
            c.callButtonsList = append(c.callButtonsList, *callButton)
            buttonFloor++
            callButtonID++
        }
    }
}

func (c *Column) createElevators(_amountOfFloors int, _amountOfElevators int ){
    for x := 0; x < _amountOfElevators; x++ {
        elevator := NewElevator(elevatorID, _amountOfFloors, 1, "idle")
        c.elevatorsList = append(c.elevatorsList, elevator)
        elevatorID++
    }
}

//Simulate when a user press a button on a floor to go back to the first floor
func (c *Column) requestElevator(_requestedFloor int, _direction string) *Elevator {
	elevator := c.findElevator(_requestedFloor, _direction)
    elevator.addNewRequest(_requestedFloor)
    elevator.move()

    elevator.addNewRequest(1)
    elevator.move()

    return elevator
}

func (c *Column) findElevator(requestedFloor int, requestedDirection string) *Elevator {

    var bestElevator *Elevator
    var bestScore int = 6
    var referenceGap int = 10000000
        
    
    if requestedFloor == 1 {
        for _, elevator := range c.elevatorsList{
            if 1 == elevator.currentFloor && elevator.status == "stopped" {
                bestScore, referenceGap, bestElevator = c.checkIfElevatorIsBetter(1, elevator, bestScore, referenceGap, bestElevator, requestedFloor)
            }else if 1 == elevator.currentFloor && elevator.status == "idle" {
                bestScore, referenceGap, bestElevator = c.checkIfElevatorIsBetter(2, elevator, bestScore, referenceGap, bestElevator, requestedFloor) 
            }else if 1 > elevator.currentFloor && elevator.direction == "up" {
                bestScore, referenceGap, bestElevator = c.checkIfElevatorIsBetter(3, elevator, bestScore, referenceGap, bestElevator, requestedFloor) 
            }else if 1 < elevator.currentFloor && elevator.direction == "down" {
                bestScore, referenceGap, bestElevator = c.checkIfElevatorIsBetter(3, elevator, bestScore, referenceGap, bestElevator, requestedFloor) 
            }else if elevator.status == "idle" {
                bestScore, referenceGap, bestElevator = c.checkIfElevatorIsBetter(4, elevator, bestScore, referenceGap, bestElevator, requestedFloor) 
            }else {
                bestScore, referenceGap, bestElevator = c.checkIfElevatorIsBetter(5, elevator, bestScore, referenceGap, bestElevator, requestedFloor) 
            }
        }
    }else {
        for _, elevator := range c.elevatorsList{
            if requestedFloor == elevator.currentFloor && elevator.status == "stopped" && requestedDirection == elevator.direction {
                bestScore, referenceGap, bestElevator = c.checkIfElevatorIsBetter(1, elevator, bestScore, referenceGap, bestElevator, requestedFloor) 
            }else if requestedFloor > elevator.currentFloor && elevator.direction == "up" && requestedDirection == "up" {
                bestScore, referenceGap, bestElevator = c.checkIfElevatorIsBetter(2, elevator, bestScore, referenceGap, bestElevator, requestedFloor) 
            }else if requestedFloor < elevator.currentFloor && elevator.direction == "down" && requestedDirection == "down" {
                bestScore, referenceGap, bestElevator = c.checkIfElevatorIsBetter(2, elevator, bestScore, referenceGap, bestElevator, requestedFloor) 
            }else if elevator.status == "idle" {
                bestScore, referenceGap, bestElevator = c.checkIfElevatorIsBetter(4, elevator, bestScore, referenceGap, bestElevator, requestedFloor) 
            }else {
                bestScore, referenceGap, bestElevator = c.checkIfElevatorIsBetter(5, elevator, bestScore, referenceGap, bestElevator, requestedFloor)
            }
        }
    }
    return bestElevator
}

func (c *Column) checkIfElevatorIsBetter(scoreToCheck int, newElevator *Elevator, bestScore int, referenceGap int, bestElevator *Elevator, floor int) (int, int, *Elevator) {
    if scoreToCheck < bestScore {
        bestScore = scoreToCheck
        bestElevator = newElevator
        referenceGap = int(math.Abs(float64(newElevator.currentFloor) - float64(floor)))
    }else if bestScore == scoreToCheck{
        var gap int = int(math.Abs(float64(newElevator.currentFloor) - float64(floor)))
        if referenceGap > gap {
            bestElevator = newElevator
            referenceGap = gap
        }
    }
    return bestScore, referenceGap, bestElevator 
        
}