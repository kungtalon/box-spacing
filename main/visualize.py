import numpy as np
import matplotlib.pyplot as plt
import matplotlib.patches as patches
import csv

fig = plt.figure()


def plot_boxes_csv(file, subplot_idx):
    ax = fig.add_subplot(subplot_idx)
    ax.set_xlim(xmin=-2, xmax=6)
    ax.set_ylim(ymin=-2, ymax=6)
    with open(file) as csvfile:
        reader = csv.reader(csvfile)
        for row in reader:
            values = list(map(float, row))
            ax.add_patch(patches.Rectangle(
                [values[0], values[1]],
                values[2] - values[0], values[3] - values[1],
                color=list(np.random.choice(range(256), size=3) / 256)
            ))


plot_boxes_csv('./boxes.csv', 121)
plot_boxes_csv('./result.csv', 122)
plt.show()
