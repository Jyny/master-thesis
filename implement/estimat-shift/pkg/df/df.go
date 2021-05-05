package df

import "math"

func UpdateDF(n int, sOld, meanOld, input float64) (s, mean float64) {
	if n < 2 {
		return 0, float64(input)
	}

	nfloat := float64(n)
	mean = (meanOld*(nfloat-1) + input) / nfloat

	s = math.Sqrt(
		(nfloat-2)/(nfloat-1)*math.Pow(sOld, 2) + (math.Pow((input-meanOld), 2))/nfloat,
	)

	return s, mean
}
