# dual-number-fourier-series-go

This program uses dual numbers for forward-mode auto-differentiation to express a target function as a sum of sine waves.

The target function is itself a sum of sine waves in the same parameterized family, though this is easily changed.

The learning process is visualized live. 

The target function, the 'pursuit' function, and the individual component waves of the pursuit function are all graphed as learning proceeds. 

The target function is approximated with something like a Fourier series.

The frequencies are real numbers, however, and not just natural numbers (this is easily changed.)
