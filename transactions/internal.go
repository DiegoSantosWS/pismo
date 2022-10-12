package transactions

func validateValue(limit, value float64) (ok bool) {
	if limit >= value {
		ok = true
	} else if limit < value {
		ok = false
	}

	return
}
