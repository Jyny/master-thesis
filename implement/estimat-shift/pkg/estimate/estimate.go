package estimate

import (
	"estimat-shift/pkg/df"
	"estimat-shift/pkg/wave"
	"fmt"
	"math"
	"os"
	"time"
)

var (
	DEBUG     = false
	threshold = 7
)

func Estimate() int {
	startTime := time.Now()
	if DEBUG {
		fmt.Println("Start: ", startTime)
	}

	if len(os.Args) < 3 {
		panic("need 2 wav file")
	}

	jammed, err := wave.ReadWave(os.Args[1])
	if err != nil {
		panic(err)
	}

	subjammer, err := wave.ReadWave(os.Args[2])
	if err != nil {
		panic(err)
	}

	if DEBUG {
		fmt.Println("file: ", os.Args[1], "samples: ", len(jammed.Data)/jammed.Format.NumChannels)
		fmt.Println("file: ", os.Args[2], "samples: ", len(subjammer.Data)/subjammer.Format.NumChannels)
	}

	min_power := float32(10.0)
	var min_power_idx, df_idx int
	var df_s, df_mean float64
	var power float32

	for i := 0; i <= len(subjammer.Data)/2; i += 1 {
		// positive side
		df_idx++
		power = PowersumDejammerWithShift(jammed.Data, subjammer.Data, i)
		df_s, df_mean = df.UpdateDF(df_idx, df_s, df_mean, float64(power))
		if power < min_power {
			// update res
			min_power = power
			min_power_idx = i
			if DEBUG {
				fmt.Printf("duration: %s, index: %d, s: %f, power: %f, mean: %f \n", time.Since(startTime), min_power_idx, df_s, min_power, df_mean)
			}

			// reset estimate
			df_s = 0
			df_mean = float64(power)
			df_idx = 1
		}

		// negative side
		df_idx++
		power = PowersumDejammerWithShift(jammed.Data, subjammer.Data, -i)
		df_s, df_mean = df.UpdateDF(df_idx, df_s, df_mean, float64(power))
		if power < min_power {
			// update res
			min_power = power
			min_power_idx = -i
			if DEBUG {
				fmt.Printf("duration: %s, index: %d, s: %f, power: %f, mean: %f \n", time.Since(startTime), min_power_idx, df_s, min_power, df_mean)
			}

			// reset estimate
			df_s = 0
			df_mean = float64(power)
			df_idx = 1
		}

		if DEBUG {
			fmt.Printf("duration: %s, index: %d, s: %f, power: %f, mean: %f \r", time.Since(startTime), i, df_s, power, df_mean)
		}

		if math.Abs(df_mean-float64(min_power)) > df_s*float64(threshold) {
			break
		}
	}

	return min_power_idx
}

func PowersumDejammerWithShift(jammedData []float32, jammerData []float32, shift int) (ret float32) {
	for i, v := range jammedData {
		var ii int
		if i%2 != 0 {
			continue
		} else {
			ii = i / 2
		}

		var data float32
		idx := ii - shift
		if idx < 0 || idx >= len(jammerData)/2 {
			data = v
		} else {
			data = v + -jammerData[idx*2]
		}

		ret += float32(math.Abs(float64(data)))
	}

	return ret / float32(len(jammedData)/2)
}
