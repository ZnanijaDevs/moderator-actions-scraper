package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var MAX_USER_IDS_LENGTH_IN_QUERY = 500

func ExtractUserIdsFromQuery(query string) ([]string, error) {
	chunks := strings.Split(query, ",")

	if (len(chunks) > MAX_USER_IDS_LENGTH_IN_QUERY) {
		return nil, fmt.Errorf("too many users")
	}

	var keys = make(map[string]bool)
	var slice []string

	for _, chunk := range chunks {
		// Check for duplicates
		if _, value := keys[chunk]; value {
			continue
		} else {
			keys[chunk] = true
		}

		i, err := strconv.Atoi(chunk)
		if err != nil {
			return nil, err
		}

		if i < 0 || i > math.MaxInt32 {
			return nil, fmt.Errorf("invalid number: %d", i)
		}

		
		slice = append(slice, chunk)
	}

	return slice, nil
}