package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type recipe struct {
	ingredients, allergens []string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var recipes []recipe
	for scanner.Scan() {
		var r recipe
		txt := scanner.Text()
		segs := strings.Split(txt, " (contains ")
		if len(segs) != 2 {
			continue
		}
		r.ingredients = strings.Split(segs[0], " ")
		r.allergens = strings.Split(strings.TrimSuffix(segs[1], ")"), ", ")
		recipes = append(recipes, r)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	p1, p2 := findNoAllergens(recipes)
	fmt.Println(p1, p2)
}

func findNoAllergens(recipes []recipe) (int, string) {
	allergenIngredients := make(map[string]map[string]int, len(recipes))
	allIngs := make(map[string]int, len(recipes))
	for _, r := range recipes {
		for _, i := range r.ingredients {
			allIngs[i]++
		}
	}
	for _, r := range recipes {
		for _, a := range r.allergens {
			var aIngs map[string]int
			var ok bool

			if aIngs, ok = allergenIngredients[a]; !ok {
				aIngs = map[string]int{a: 0}
				allergenIngredients[a] = aIngs
			}

			for _, i := range r.ingredients {
				aIngs[i]++
			}
			aIngs[a]++
		}
	}

	possibleIngs := make(map[string]bool, len(allIngs))

	for a, aIngs := range allergenIngredients {
		for i, c := range aIngs {
			if i == a {
				continue
			}
			if c == aIngs[a] {
				possibleIngs[i] = true
			} else {
				delete(aIngs, i)
			}
		}
		delete(aIngs, a)
	}

	var c int

	for i, v := range allIngs {
		if !possibleIngs[i] {
			c += v
		}
	}

	found := make(map[string]string, len(allergenIngredients))

	for len(found) != len(allergenIngredients) {
		// random iteration order, could be improved by length order
	aLoop:
		for a, aIngs := range allergenIngredients {
			if aIngs == nil {
				continue
			}
			var possible string

			for i := range aIngs {
				if found[i] == "" {
					if possible == "" {
						possible = i
					} else {
						continue aLoop
					}
				}
			}
			allergenIngredients[a] = nil
			found[possible] = a
		}
	}

	mapSlice := make([][2]string, 0, len(found))

	for i, a := range found {
		mapSlice = append(mapSlice, [2]string{i, a})
	}

	sort.Slice(mapSlice, func(i, j int) bool {
		return mapSlice[i][1] < mapSlice[j][1]
	})

	var s strings.Builder

	for i, p := range mapSlice {
		if i != 0 {
			s.WriteRune(',')
		}
		s.WriteString(p[0])
	}

	return c, s.String()
}
