package picker

import "math/rand/v2"

func Pick(choices []string) string {
	if len(choices) == 0 {
		return ""
	}
	index := rand.IntN(len(choices))
	return choices[index]
}
