//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

type Player struct {
	name              string
	health, maxHealth uint
	energy, maxEnergy uint
}

func (p *Player) addHealth(value uint) {
	p.health += value
	if p.health > p.maxHealth {
		p.health = p.maxHealth
	}
	fmt.Println(p.name, "Adding", value, "Points of HP: ", value, "New HP: ", p.health)
}

func (p *Player) attack(value uint) {
	if p.health-value > p.health {
		p.health = 0
	} else {
		p.health -= value
	}
	fmt.Println(p.name, "Damage", value, " -> ", p.health)
}

func (player *Player) addEnergy(amount uint) {
	player.energy += amount
	if player.energy > player.maxEnergy {
		player.energy = player.maxEnergy
	}
	fmt.Println(player.name, "Add", amount, "energy -> ", player.energy)
}

func (player *Player) consumeEnergy(amount uint) {
	// overflow check
	if player.energy-amount > player.energy {
		player.energy = 0
	} else {
		player.energy -= amount
	}
	fmt.Println(player.name, "Consume", amount, "energy -> ", player.energy)
}

func main() {
	fmt.Println("========== WORLD OF WARCRAFT - BETA TEST CHARACTERS ==========")
	player := Player{
		name:      "Paladin",
		health:    100,
		maxHealth: 100,
		energy:    500,
		maxEnergy: 500,
	}

	player.attack(99)
	player.addHealth(10)
	player.consumeEnergy(20)
	player.addEnergy(10)

	player.consumeEnergy(2222)
}
