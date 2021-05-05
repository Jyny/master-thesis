import matplotlib.pyplot as plt
import numpy as np

def wavegram(samplerate, data, size):
	plt.plot(
		np.linspace(0., data.shape[0]/samplerate, data.shape[0]),
		data[:, 0],
		label="channel 0",
		figure=plt.figure(figsize=size)
	)
	plt.xlabel("Time [s]")
	plt.ylabel("Amplitude")
	plt.show()

def voicegram(samplerate, data, size):
	plt.specgram(data[:, 0], Fs=samplerate, figure=plt.figure(figsize=size))
	plt.xlabel('Time')
	plt.ylabel('Frequency')
	plt.show()
