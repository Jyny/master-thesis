import numpy as np
from scipy import signal

def compressorList(data, max):
	ret = np.zeros(shape=(data.shape[0],data.shape[1]))
	for i, v in enumerate(data):
		ret[i] = compressor(v, max)

	return ret


def compressor(data, max):
	if abs(data[0]) > max:
		if data[0] > 0:
			data[0] = max
		else:
			data[0] = -max
	if abs(data[1]) > max:
		if data[1] > 0:
			data[1] = max
		else:
			data[1] = -max

	return data

def lowpass(data, rate, freq):
	sos = signal.butter(8, freq, fs=rate, output='sos')
	data[:, 0] = signal.sosfilt(sos, data[:, 0])
	data[:, 1] = signal.sosfilt(sos, data[:, 1])

	return data