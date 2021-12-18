# Results

## Version 1

This approach is simply iterating over the arguments provided,
starting at index 1.

The execution time in average is **0.099s**.

## Version 2

This approach also iterates over the arguments, just that it
iterates over a slice that starts at index 1.

The execution time in average is **0.182s**.

## Version 3

This approach uses `strings.Join` to join its command-line arguments
and then prints them to the standard output.

The execution time in average is **0.180s**.

## Conclusion

They are not so different, and using range was actually slower.
