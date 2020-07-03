#!/usr/bin/python
#coding=utf8

import random
import matplotlib.pyplot as plt

def Subset(backends, client_id, subset_size):
  subset_count = int(len(backends) / subset_size)

  # Group clients into rounds; each round uses the same shuffled list:
  round = client_id / subset_count
  random.seed(round)
  random.shuffle(backends)

  # The subset id corresponding to the current client:
  subset_id = client_id % subset_count

  start = subset_id * subset_size
  return backends[start:start + subset_size]

def printbar(subset_dict):
  x = []
  y = []
  for key in subset_dict:
    x.append(key)
    y.append(subset_dict[key])

  plt.bar(x, y)
  plt.show()


def main():
  backends = list(range(6))
  client_id = list(range(360000))
  subset_size = 3
  subset_list = []

  for id in client_id:
    subset = Subset(backends, id, subset_size)
    subset_list.append(subset)

  print(subset_list)
  subset_dict = {}
  for k in subset_list:
    for v in k:
      if v in subset_dict:
        subset_dict[v] = subset_dict[v]+1
      else:
        subset_dict[v] = 1

  print(subset_dict)
  printbar(subset_dict)

if __name__ == "__main__":
    main()
