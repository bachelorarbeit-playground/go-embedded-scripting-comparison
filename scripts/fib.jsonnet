local fibonacci(n) =
    if n < 2 then
        n
    else
        fibonacci(n - 2) + fibonacci(n - 1);

fibonacci(stdin.n)