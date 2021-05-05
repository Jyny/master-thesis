import numpy as np

def dejammer_with_shift(data, jammer_data, samplerate, shift):
	dejammed_data = np.zeros(shape=(data.shape[0],data.shape[1]))
	for i, v in enumerate(data):
		idx = i - shift
		if idx < 0 or idx >= jammer_data.shape[0]:
			channel0 = v[0]
			channel1 = v[1]
		else:
			channel0 = v[0] + -jammer_data[idx, 0]
			channel1 = v[1] + -jammer_data[idx, 1]
		dejammed_data[i] = np.array([channel0, channel1])

	return samplerate, dejammed_data


def dejammer(data, jammer_data, samplerate):
	return dejammer_with_shift(data, jammer_data, samplerate, 0)