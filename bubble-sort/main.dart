void main () {
  List<int> list = [10, 5, 7, 3, 2, 8, 9, 6, 1, 4];
  List<int> sortedList = bubbleSort(list);
  print(sortedList);
}
List<int> bubbleSort(List<int> list) {
  int listLength = list.length;
  for (int x = 0; x < listLength; x += 1) {
    for (int y = 0; y < listLength - 1; y+= 1) {
      if (list[y] > list[y + 1]) {
        int temp = list[y];
        list[y] = list[y + 1];
        list[y + 1] = temp;
      }
    }
  }
  return list;
}