package main

import (
	"fmt"
)

func main() {
	useCountingSort()
	countingAlphabet()
	countingHeight()
}

// 쉬우면서 효율적인 알고리즘 , 하지만 실제로 쓸일이 많지는않음..
//countSort 최적의 조건일때 가장 빠른 정렬이라고 해도 무난함 (한정된 범위)
// O(N + K) 값의 범위가 작을때 가장 빠름(k가 작을때)

func useCountingSort() {

	// 1 ~ 10 범위를 가지는 배열
	arr := []int{4, 3, 2, 1, 5, 6, 7, 8, 9, 5, 5, 6, 7, 4, 5, 3, 1, 10, 4}

	// 범위 +1  크기만큼의 배열을 생성
	var count [11]int

	for i := 0; i < len(arr); i++ {
		count[arr[i]]++
	}
	sorted := make([]int, 0, len(arr))
	for i := 0; i < 11; i++ {
		for j := 0; j < count[i]; j++ {
			sorted = append(sorted, i)
		}
	}
	fmt.Println(sorted)
}

// 소문자 알파뱃 최대  갯수를 세시오
func countingAlphabet() {
	str := "adsfkjfjshbbbbbbiufwegnbmcvzbmnvhweyotyerwyfgbnmzvnbcvxzjhgz"
	//str := "asssssscxvzjhweryio"
	// a~z = 26
	var count [26]int
	for i := 0; i < len(str); i++ {
		count[str[i]-'a']++
	}

	maxCnt := 0
	var maxCh byte
	for i := 0; i < 26; i++ {
		if count[i] > maxCnt {
			maxCnt = count[i]
			maxCh = byte('a' + i)
		}
	}
	fmt.Printf("%c 가 %d 개로 가장 많습니다", maxCh, maxCnt)
}

// 학생 키 범위 소수 첫번째 자리까지 유효할경우
func countingHeight() {
	students := []struct {
		Name   string
		Height float64
	}{
		{Name: "joy", Height: 173.4},
		{Name: "world", Height: 151.2},
		{Name: "warm", Height: 164.6},
		{Name: "oil", Height: 180.2},
		{Name: "kelly", Height: 151.7},
		{Name: "ryu", Height: 175.9},
		{Name: "jax", Height: 154.8},
		{Name: "ashe", Height: 165.6},
	}

	// 키범위가  0.0 ~ 300.0 까지라고 과정
	var heightMap [3000][]string

	for i := 0; i < len(students); i++ {
		// 소수점 첫번째 자리를  없애기 위해 10을 곱해줌
		idx := int(students[i].Height) * 10
		heightMap[idx] = append(heightMap[idx], students[i].Name)
	}

	// 키 140~170 까지 구할경우
	for i := 1400; i < 1700; i++ {
		for _, name := range heightMap[i] {
			fmt.Println("name", name, "height", float64(i)/10)
		}
	}

}
