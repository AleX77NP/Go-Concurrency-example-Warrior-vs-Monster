package main

import (
	"fmt"
	"time"
	"strconv"
)

func main() {

	warrior1 := Warrior{"Arnold Hagss", 35,120,510}
	monster1 := Monster{"KrakenusX", "Water", 1,12,230,440}

	go warriorAttack(&warrior1, &monster1)

	go monsterAttack(&warrior1, &monster1)

	fmt.Scanln()

}

func warriorAttack(warrior *Warrior, monster *Monster)  {
  for monster.defense > 0{
  time.Sleep(time.Millisecond * 500)
  fmt.Println("Warrior " + warrior.name + " attacked " + monster.name + " !")
  if warrior.attack > monster.defense {
	monster.defense = 0;
  } else {
  monster.defense -= warrior.attack
  }
  fmt.Println("Monster defense: " + strconv.Itoa(monster.defense))
  if checkOver(warrior, monster) == true {
	  return
  }
 }
}

func monsterAttack(warrior *Warrior, monster *Monster)  {
	for warrior.defense > 0 {	
	time.Sleep(time.Second*2)
	fmt.Println("Monster " + monster.name + " attacked " + warrior.name + " !")
	if monster.attack >= warrior.defense {
		warrior.defense = 0;
	} else {
	warrior.defense -= monster.attack
	}
	fmt.Println("Warrior defense: " + strconv.Itoa(warrior.defense))
	if checkOver(warrior, monster) == true {
	return
	}
  }
}

func checkOver(warrior *Warrior, monster *Monster) bool {
   if warrior.defense < 1 || monster.defense <1 {
	   return true
   }
   return false
}

//Warrior ...
type Warrior struct {
	name string
	age int
	attack int
	defense int
}
//Monster ...
type Monster struct {
	name string
    attribute string
	lifes int
	level int
	attack int
	defense int
}