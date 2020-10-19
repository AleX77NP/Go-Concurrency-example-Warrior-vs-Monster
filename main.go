package main

import (
	"fmt"
	"time"
	"strconv"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(1)

	warrior1 := Warrior{"Arnold Hagss", 35,120,510}
	monster1 := Monster{"KrakenusX", "Water", 1,12,230,440}

	go func() {
		monsterAttack(&warrior1, &monster1)
		wg.Done()
	}()

	go func() {
		warriorAttack(&warrior1, &monster1)
		wg.Done()
	}()


	wg.Wait()

}

func warriorAttack(warrior *Warrior, monster *Monster)  {
  time.Sleep(time.Millisecond * 1000)
  for monster.defense > 0 && warrior.defense > 0{
  fmt.Println("Warrior " + warrior.name + " attacked " + monster.name + " !")
  if warrior.attack > monster.defense {
	monster.defense = 0;
	monster.lifes--;
  } else {
  monster.defense -= warrior.attack
  }
  fmt.Println("Monster defense: " + strconv.Itoa(monster.defense))
  if checkOver(warrior, monster) == true {
	  break
  }
 }
}

func monsterAttack(warrior *Warrior, monster *Monster)  {
	time.Sleep(time.Second*1)
	for warrior.defense > 0 && monster.defense > 0{	
	fmt.Println("Monster " + monster.name + " attacked " + warrior.name + " !")
	if monster.attack >= warrior.defense {
		warrior.defense = 0;
	} else {
	warrior.defense -= monster.attack
	}
	fmt.Println("Warrior defense: " + strconv.Itoa(warrior.defense))
	if checkOver(warrior, monster) == true {
	break
	}
  }
}

func checkOver(warrior *Warrior, monster *Monster) bool {
   if warrior.defense == 0 || monster.defense == 0 {
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