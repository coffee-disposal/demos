void main() {
  int n = 10;
  int fib_n = fib(n);
  String formattedString = "The ${n}th number of the fibonacci sequence is ${fib_n}.";
  print(formattedString);
}

int fib (int n) {
  if (n <= 1)
    return n;
  return fib (n - 1) + fib (n - 2);
}