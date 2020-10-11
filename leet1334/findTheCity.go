package leet1334

import "fmt"

type cityNode struct {
	city     int
	toCity   []*cityNode
	distance []int
}

// todo:行吧这题朴素解法我只能想到：先把数据具像化为图关系，再转化为树结构，再对树就行遍历求值

func FindTheCity(n int, edges [][]int, distanceThreshold int) int {

	cityMap := make([]*cityNode, n)
	for _, item := range edges {
		if cityMap[item[0]] == nil {
			c := cityNode{}
			c.city = item[0]
			cityMap[item[0]] = &c
		}
		if cityMap[item[1]] == nil {
			c := cityNode{}
			c.city = item[1]
			cityMap[item[1]] = &c
		}
		cityMap[item[0]].toCity = append(cityMap[item[0]].toCity, cityMap[item[1]])
		cityMap[item[1]].toCity = append(cityMap[item[1]].toCity, cityMap[item[0]])
		cityMap[item[0]].distance = append(cityMap[item[0]].distance, item[2])
		cityMap[item[1]].distance = append(cityMap[item[1]].distance, item[2])

	}

	for _, c := range cityMap {
		fmt.Println(c)

	}

	return 1
}

func getResult(m map[int][]int) int {
	mLen := 0
	res := -1
	for k, v := range m {
		if mLen < len(v) {
			mLen = len(v)
			res = k
		}
	}
	return res
}

func getTree(city *cityNode, cityMap *map[int]*cityNode, treed []*cityNode) {
	for _, cN := range *cityMap {
		saved := false
		for _, v := range treed {
			if cN == v {
				saved = true
				break
			}
		}
		if saved {
			continue
		}


	}

}
