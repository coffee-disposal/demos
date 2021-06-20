void main () {
  List<int> list = [10, 5, 7, 3, 2, 8, 9, 6, 1, 4];
  int max = findMax(list);
  print(max);
}

int findMax(List<int> list) {
  int listLength = list.length;
  int currentMax = list[0];
  for (int x = 0; x < listLength; x += 1) {
    if (list[x] > currentMax) {
      currentMax = list[x];
    }
  }
  return currentMax;
}