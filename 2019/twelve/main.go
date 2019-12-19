package main

import (
	"fmt"
)

type moon struct {
	name string
	x    int
	y    int
	z    int
	vX   int
	vY   int
	vZ   int
}

func main() {
	initMoons := [4]moon{
		moon{
			name: "A",
			x:    -1, y: 7, z: 3,
			vX: 0, vY: 0, vZ: 0},
		moon{
			name: "B",
			x:    12, y: 2, z: -13,
			vX: 0, vY: 0, vZ: 0},
		moon{
			name: "C",
			x:    14, y: 18, z: -8,
			vX: 0, vY: 0, vZ: 0},
		moon{
			name: "D",
			x:    17, y: 4, z: -4,
			vX: 0, vY: 0, vZ: 0}}
	// initMoons := [4]moon{
	// 	moon{
	// 		name: "A",
	// 		x:    -1, y: 0, z: 2,
	// 		vX: 0, vY: 0, vZ: 0},
	// 	moon{
	// 		name: "B",
	// 		x:    2, y: -10, z: -7,
	// 		vX: 0, vY: 0, vZ: 0},
	// 	moon{
	// 		name: "C",
	// 		x:    4, y: -8, z: 8,
	// 		vX: 0, vY: 0, vZ: 0},
	// 	moon{
	// 		name: "D",
	// 		x:    3, y: 5, z: -1,
	// 		vX: 0, vY: 0, vZ: 0}}

	moons := initMoons
	stepX, stepY, stepZ := 0, 0, 0
	for true {
		moons = applyForces(moons)
		stepX++

		if moons[0].x == initMoons[0].x && moons[1].x == initMoons[1].x && moons[2].x == initMoons[2].x && moons[3].x == initMoons[3].x && moons[0].vX == 0 && moons[1].vX == 0 && moons[2].vX == 0 && moons[3].vX == 0 {
			break
		}
	}
	moons = initMoons
	for true {
		moons = applyForces(moons)
		stepY++

		if moons[0].y == initMoons[0].y && moons[1].y == initMoons[1].y && moons[2].y == initMoons[2].y && moons[3].y == initMoons[3].y && moons[0].vY == 0 && moons[1].vY == 0 && moons[2].vY == 0 && moons[3].vY == 0 {
			break
		}
	}
	moons = initMoons
	for true {
		moons = applyForces(moons)
		stepZ++

		if moons[0].z == initMoons[0].z && moons[1].z == initMoons[1].z && moons[2].z == initMoons[2].z && moons[3].z == initMoons[3].z && moons[0].vZ == 0 && moons[1].vZ == 0 && moons[2].vZ == 0 && moons[3].vZ == 0 {
			break
		}
	}
	fmt.Println(stepX, stepY, stepZ)
	fmt.Println(LCM(stepX, stepY, stepZ))
}

func applyForces(moons [4]moon) [4]moon {
	newMoons := moons
	for i := range moons {
		for j := i + 1; j < 4; j++ {
			if moons[i].x < moons[j].x {
				newMoons[i].vX++
			} else if moons[i].x > moons[j].x {
				newMoons[i].vX--
			}

			if moons[j].x < moons[i].x {
				newMoons[j].vX++
			} else if moons[j].x > moons[i].x {
				newMoons[j].vX--
			}

			if moons[i].y < moons[j].y {
				newMoons[i].vY++
			} else if moons[i].y > moons[j].y {
				newMoons[i].vY--
			}

			if moons[j].y < moons[i].y {
				newMoons[j].vY++
			} else if moons[j].y > moons[i].y {
				newMoons[j].vY--
			}
			if moons[i].z < moons[j].z {
				newMoons[i].vZ++
			} else if moons[i].z > moons[j].z {
				newMoons[i].vZ--
			}

			if moons[j].z < moons[i].z {
				newMoons[j].vZ++
			} else if moons[j].z > moons[i].z {
				newMoons[j].vZ--
			}
		}
	}
	for i := range newMoons {
		newMoons[i].x += newMoons[i].vX
		newMoons[i].z += newMoons[i].vZ
		newMoons[i].y += newMoons[i].vY
	}
	return newMoons
}

func calcEnergy(moons [4]moon) int {
	total := 0
	for _, m := range moons {
		pot := Abs(m.x) + Abs(m.y) + Abs(m.z)
		kin := Abs(m.vX) + Abs(m.vY) + Abs(m.vZ)
		total += pot * kin
	}
	return total
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
