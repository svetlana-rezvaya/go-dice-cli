package dice

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// ParseDiceNotation ...
// Example of input data: "2d20", "d20" (it means "1d20").
func ParseDiceNotation(text string) (throwCount int, faceCount int, err error) {
	parts := strings.Split(text, "d")
	if len(parts) != 2 {
		return 0, 0, errors.New("invalid count of parts")
	}

	if parts[0] == "" {
		throwCount = 1
	} else {
		throwCount, err = strconv.Atoi(parts[0])
		if err != nil {
			return 0, 0, fmt.Errorf("unable to parse the throw count: %s", err)
		}
	}

	faceCount, err = strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, fmt.Errorf("unable to parse the face count: %s", err)
	}

	return throwCount, faceCount, nil
}
