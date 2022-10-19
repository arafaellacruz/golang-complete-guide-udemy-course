//--Summary:
//  Copy your rcv-func solution to this directory and write unit tests.
//
//--Requirements:
//* Write unit tests that ensure:
//  - Health & energy can not go below 0
//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
//
//--Notes:
//* Use `go test -v ./exercise/testing` to run these specific tests

package main

import (
	"testing"
)

func newPlayer() Player {
	return Player{
		name:      "testPlayer",
		health:    100,
		maxHealth: 100,
		energy:    350,
		maxEnergy: 500,
	}
}

//  - Health & energy can not go above their maximums
//  - Health & energy can not go below 0
func TestAddHealth(t *testing.T) {
	p := newPlayer()
	p.addHealth(777)

	if p.health > p.maxHealth {
		t.Fatalf("health is above their limits: %v, want <= : %v", p.health, p.maxHealth)
	}
	p.attack(p.maxHealth + 1)
	if p.health < 0 {
		t.Fatalf("health: %v, want: >= 100", p.health)
	}
}

//  - Health & energy can not go above their maximums
//  - Health & energy can not go below 0
func TestAddEnergy(t *testing.T) {
	p := newPlayer()
	p.addEnergy(888)

	if p.energy > p.maxEnergy {
		t.Fatalf("energy is above their limits: %v, want %v", p.energy, p.maxEnergy)
	}
	p.consumeEnergy(p.maxEnergy + 1)
	if p.energy < 0 {
		t.Fatalf("health is below their limits: %v, minimum: 0", p.energy)
	}
}
