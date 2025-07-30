package util

func Map[S any, D any](source []S, fn func(S) D) []D {
	res := make([]D, 0, len(source))

	for _, item := range source {
		res = append(res, fn(item))
	}

	return res
}
