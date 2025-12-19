package chooser

import "math/rand/v2"

func Choose(choices []string) string {
	if len(choices) == 0 {
		return ""
	}
	index := rand.IntN(len(choices))
	return choices[index]
}
