void main () {
  List<int> list = [10, 5, 7, 3, 2, 8, 9, 6, 1, 4];
  int min = findMin(list);
  print(min);
}

int findMin(List<int> list) {
  int listLength = list.length;
  int currentMin = list[0];
  for (int x = 0; x < listLength; x += 1) {
    if (list[x] < currentMin) {
      currentMin = list[x];
    }
  }
  return currentMin;
}