# My Solution for the IE Gophers 2018-02-09 Challenge

The [challenge](https://github.com/IEGophers/Challenge-02-09-2018) was to implement `wc` in Go in about an hour.

## My Design

I wanted my solution to be able to read from either stdin or a file, so using an io.Reader made sense.

I designed my solution to use interface as the abstraction for each command as I realized each command:

* Had the same input (an io.Reader)
* Returned a single value (a count) plus an error

My design was built keeping in mind that the input could be larger than the available memory, so it operates on a data stream rather than reading an entire file into memory. I arbitrarily set the size of the buffer to 5 bytes as that way I was sure to test edge cases like words getting truncated in the middle of a buffer. If this was a real application, I'd set the buffer much higher and allow some option for changing it either at compile time (via LD flags) or runtime (via an environment variable, a flag, or both).

## Testing

I built a simple test harness that is set up to use [Table Driven Tests](https://github.com/golang/go/wiki/TableDrivenTests) although in practice I only set up a single file for testing. I used subtests to tests the various individual commands.
