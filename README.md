Finite State Machine (FSM)

Standard Exercise

Example 1
For input string ”110” the machine will operate as follows:
1. Initial state = S0, Input = 1, result state = S1
2. Current state = S1, Input = 1, result state = S0
3. Current state = S0, Input = 0, result state = S0
4. No more input
5. Print output value (output for state S0 = 0) <---- This is the answer
Example 2
For input string ”1010” the machine will operate as follows:
1. Initial state = S0, Input = 1, result state = S1
2. Current state = S1, Input = 0, result state = S2
3. Current state = S2, Input = 1, result state = S2
4. Current state = S2, Input = 0, result state = S1
5. No more input
6. Print output value (output for state S1 = 1) <---- This is the answer


RUN
go run main.go
OR
go test -v 
