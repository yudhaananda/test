package main

import "fmt"

// Case : You're looking for an apartment and you received a list of blocks with their own service/facility.

// To get your ideal apartment, you want to find the block with the closest distance from all of the facilities. Please create an algorithm for that solution.

// Blocks = [
// { "gym":false,
//  "school":true,
//  "store":false,},
// { "gym":true,
//  "school":false,
//  "store":false,},
// {"gym":true,
// "school":true,
//  "store":false,},
// {"gym":false,
//  "school":true,
//  "store":false,},
// {"gym":false,
//  "school":true,
//  "store":true,}
// ]

// For the example above, if [gym, school, store] is the requirement, index 3 is the best location for your apartment.
// Example:
// Reqs: ["gym", school, store]
// Block 0: [1, 0, 4] = 4
// Block 1: [0, 1, 3] = 3
// Block 2: [0, 0, 2] = 2
// Block 3: [1, 0, 1] = 1
// Block 4: [2,0,0] = 2
// Block with the most minimum distance to all facilities: Block 3

// Func: Input: Blocks, Reqs
// Output: Block with the most minimum distance to all facilities

func main() {
	result := idealApartment([]map[string]bool{
		{"gym": false,
			"school": true,
			"store":  false},
		{"gym": true,
			"school": false,
			"store":  false},
		{"gym": true,
			"school": true,
			"store":  false},
		{"gym": false,
			"school": true,
			"store":  false},
		{"gym": false,
			"school": true,
			"store":  true},
	}, []string{"gym", "school", "store"})
	fmt.Println(result)
}

func idealApartment(blocks []map[string]bool, reqs []string) string {
	distance := []map[string]int{}
	longestDistance := 0
	blockString := ""
	for blockIndex, block := range blocks {
		temp := make(map[string]int)
		for _, req := range reqs {
			if block[req] {
				temp[req] = 0
			} else {
				count := 0
				for i := blockIndex; i < len(blocks); i++ {
					if blocks[i][req] {
						break
					}
					count++
					if i == len(blocks)-1 && !blocks[i][req] {
						count = -1
					}
				}
				if (blockIndex != 0 && distance[blockIndex-1][req]+1 < count) || count == -1 {
					temp[req] = distance[blockIndex-1][req] + 1
				} else {
					temp[req] = count
				}
			}
		}
		tempLongestDistance := 0
		for _, v := range temp {
			if v > tempLongestDistance {
				tempLongestDistance = v
			}
		}
		if blockIndex == 0 || longestDistance > tempLongestDistance {
			longestDistance = tempLongestDistance
			blockString = "Block " + fmt.Sprint(blockIndex)
		}
		distance = append(distance, temp)
	}
	return blockString
}
