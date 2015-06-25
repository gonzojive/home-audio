#!/usr/bin/python
# Example output: 01:45:89:ab:cd:ef
import random

def random_char():
  return random.choice('0123456789abcdef')

print ':'.join(random_char() + random_char() for n in xrange(6))
