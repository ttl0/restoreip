# Simple leetcode medium challenge turned into a CLI demo using Cobra CLI

Demo usage:
$ ./restoreip 2222222\
[2.2.22.222 2.2.222.22 2.22.2.222 2.22.22.22 2.22.222.2 2.222.2.22 2.222.22.2 22.2.2.222 22.2.22.22 22.2.222.2 22.22.2.22 22.22.22.2 22.222.2.2 222.2.2.22 222.2.22.2 222.22.2.2]\
% ./restoreip 0000

[0.0.0.0]

$ ./restoreip 16925444444

[]

$ ./restoreip 1692544444

[169.254.44.44]

$ ./restoreip 169254444

[169.25.44.44 169.254.4.44 169.254.44.4]

$ ./restoreip 2561

[2.5.6.1]

$ ./restoreip 257

[]

$ ./restoreip 2571

[2.5.7.1]
