/*
Улица, на которой хочет жить Тимофей, имеет длину n, то есть состоит из n одинаковых идущих подряд участков. На каждом участке либо уже построен дом, либо участок пустой. Тимофей ищет место для строительства своего дома. Он очень общителен и не хочет жить далеко от других людей, живущих на этой улице. Чтобы оптимально выбрать место для строительства, Тимофей хочет для каждого участка знать расстояние до ближайшего пустого участка. (Для пустого участка эта величина будет равна нулю –— расстояние до самого себя). Ваша задача –— помочь Тимофею посчитать искомые расстояния. Для этого у вас есть карта улицы. Дома в городе Тимофея нумеровались в том порядке, в котором строились, поэтому их номера на карте никак не упорядочены. Пустые участки обозначены нулями.

Формат ввода В первой строке дана длина улицы —– n (1 ≤ n ≤ 106). В следующей строке записаны n целых неотрицательных чисел — номера домов и обозначения пустых участков на карте (нули). Гарантируется, что в последовательности есть хотя бы один ноль. Номера домов (положительные числа) уникальны и не превосходят 109.

Формат вывода Для каждого из участков выведите расстояние до ближайшего нуля. Числа выводите в одну строку, разделяя их пробелами.

Пример 1:
Ввод:
5
0 1 4 9 0

Вывод:
0 1 2 1 0

Пример 2:
Ввод:
6
0 7 9 4 8 20

Вывод:
0 1 2 3 4 5
*/
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main()  {
	var n int

	reader := bufio.NewReader(os.Stdin)

	line, _ := reader.ReadString('\n')
	tmp, _ := strconv.ParseInt(strings.TrimSpace(line), 10, 32)

	n = int(tmp)
	fmt.Println(n)

	line, _ = reader.ReadString('\n')
	splitted := strings.Split(strings.TrimSpace(line), " ")

	arr := make([]int, len(splitted))

	for i := range splitted {
		v, _ := strconv.ParseInt(splitted[i], 10, 32)
		arr[i] = int(v)
	}

	fmt.Println(arr)

	ans := solution1(arr)
	fmt.Println(ans)

	ans = solution2(arr)
	fmt.Println(ans)
}

// O(n^2)
func solution1(arr []int) []int {
	rng := make([]int, len(arr))
	for i := range rng {
		rng[i] = int(math.Inf(1)) - 1
	}
	// fmt.Println(rng)

	for i, v := range arr {
		if v == 0 {
			for j := range rng {
				if int(math.Abs(float64(i-j))) < rng[j] {
					rng[j] = int(math.Abs(float64(i-j)))
				} 
			}
		}
	}

	return rng
}

//O(n)
func solution2(arr []int) []int {
	rng := make([]int, len(arr))
	nearest_zero := -1

	for i, v := range arr {
		if v == 0 {
			nearest_zero = i
			rng[i] = 0
			continue
		}
		rng[i] = i - nearest_zero
		
	}

	for i := len(arr) - 1; i > -1; i-- {
		if arr[i] == 0 {
			nearest_zero = i
			continue
		}
		if nearest_zero - i < rng[i] {
			rng[i] = nearest_zero - i
		}
	}

	return rng
}