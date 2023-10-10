package testhelp

import "fmt"

type ResultsMap map[string]string

func ValidateResults(res, wanted ResultsMap) error {
	for k, v := range wanted {
		if rv, ok := res[k]; !ok {
			return fmt.Errorf("result missing key %s", k)
		} else if rv != v {
			return fmt.Errorf("wanted '%s' for result key %s, got '%s'", v, k, res[k])
		}
	}
	return nil
}
