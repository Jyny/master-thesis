import sys
from scipy.io import wavfile
import numpy as np

sys.path.append('./')
from explib.draw import wavegram, voicegram
from explib.dejam import dejammer_with_shift
from explib.anc import dejammer_ANC_with_shift
from explib.filter import compressorList
from explib.filter import lowpass
from explib.align import powersum


if __name__ == '__main__':
    if len(sys.argv) < 4:
        sys.exit()

    # output dejammed
    dejammed_wavefile = sys.argv[3]

    # input
    jammed_wavefile = sys.argv[1]
    jammer_wavefile = sys.argv[2]

    # align with shift
    shift = int(sys.argv[4])

    jammed_samplerate, jammed_data = wavfile.read(jammed_wavefile)
    jammer_samplerate, jammer_data = wavfile.read(jammer_wavefile)

    # compressor
    compressorMax = 0.1
    jammed_data = compressorList(jammed_data, compressorMax)
    jammer_data = compressorList(jammer_data, compressorMax)

    # band filter
    jammed_data = lowpass(jammed_data, jammed_samplerate, 2000)
    jammer_data = lowpass(jammer_data, jammer_samplerate, 2000)

    # ANC
    dejammed_samplerate, dejammed_data = dejammer_ANC_with_shift(jammed_data, jammer_data, jammed_samplerate, shift, 20, 0.8)
    dejammed_data = compressorList(dejammed_data, compressorMax)

    # output
    wavfile.write(dejammed_wavefile, dejammed_samplerate, dejammed_data.astype(np.float32))