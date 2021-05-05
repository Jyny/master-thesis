import numpy as np
import padasip as pa

def dejammer_ANC_with_shift(data, jammer_data, samplerate, shift, filt_len, mu):
	dejammed_data = np.zeros(shape=(data.shape[0],data.shape[1]))
	filt = pa.filters.FilterLMS(filt_len, mu=mu, w="zeros")
	for i, v in enumerate(data):
		idx = i - shift
		if idx < 0 or idx >= jammer_data.shape[0] or i < filt_len or idx < filt_len:
			continue
		else:
			x = jammer_data[idx-filt_len:idx, 0]
			d = data[i, 0]
			filt.adapt(d, x)

	for i, v in enumerate(data):
		idx = i - shift
		if idx < 0 or idx >= jammer_data.shape[0] or i < filt_len or idx < filt_len:
			channel0 = v[0]
			channel1 = v[1]
		else:
			x = jammer_data[idx-filt_len:idx, 0]
			y = filt.predict(x)

			d = data[i, 0]
			filt.adapt(d, x)

			channel0 = v[0] - y
			channel1 = v[0] - y

		dejammed_data[i] = np.array([channel0, channel1])

	return samplerate, dejammed_data