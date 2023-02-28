```
go run main.go
roll
got: 6!
6
[6]
[]
roll
got: 6!
12
[6 6]
[]
roll
got: 1!
13
[6 6 1]
[]
roll
got: 3!
16
[6 6 1 3]
[]
roll
got: 3!
19
[6 6 1 3 3]
[]
roll
got: 2!
21
[6 6 1 3 3 2]
[]
undo
19
[6 6 1 3 3]
[2]
undo
16
[6 6 1 3]
[2 3]
undo
13
[6 6 1]
[2 3 3]
undo
12
[6 6]
[2 3 3 1]
redo
13
[6 6 1]
[2 3 3]
redo
16
[6 6 1 3]
[2 3]
roll
got: 2!
18
[6 6 1 3 2]
[]
roll
got: 5!
23
[6 6 1 3 2 5]
[]
undo
18
[6 6 1 3 2]
[5]
redo
23
[6 6 1 3 2 5]
[]
exit
23
[6 6 1 3 2 5]
[]
```