package repeat

import (
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

func NextDate(now time.Time, date string, repeat string) (string, error) {
	if repeat == "" {
		return "", errors.WithMessage(nil, "repeat is empty")
	}

	d, err := time.Parse(`20060102`, date)
	if err != nil {
		return "", errors.WithMessage(err, "parse date")
	}

	checkSlice := strings.Split(repeat, " ")

	switch checkSlice[0] {
	case "y":
		if len(checkSlice) != 1 {
			return "", errors.New("invalid repeat")
		}

		for {
			d = d.AddDate(1, 0, 0)
			if d.After(now) {
				break
			}
		}
	case "m":
		return "", errors.New("invalid repeat")
	case "w":
		return "", errors.New("invalid repeat")
	case "d":
		if len(checkSlice) == 2 {
			days, err := strconv.Atoi(checkSlice[1])
			if days > 400 || err != nil {
				return "", errors.New("invalid repeat")
			}
			for {
				d = d.AddDate(0, 0, days)
				if d.After(now) {
					break
				}

			}
		} else {
			return "", errors.New("invalid repeat")
		}

	default:
		return "", errors.New("invalid repeat")
	}

	return d.Format(`20060102`), nil
}
