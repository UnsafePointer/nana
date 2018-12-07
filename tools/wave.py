#!/usr/bin/env python2
from scipy.io.wavfile import write
from numpy import array
from numpy import int16

data = [int16(line) for line in open('wave.txt')]
ndarray = array(data)
write("wave.wav", 44100, ndarray)
