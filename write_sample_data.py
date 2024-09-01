import numpy as np

NUM_ITEMS = 10

data = np.arange(NUM_ITEMS, dtype="int64").tobytes()
with open('foo', 'wb') as f:
    f.write(data)
