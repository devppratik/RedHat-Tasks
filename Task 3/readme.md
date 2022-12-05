## Task-2

### _Problem Statement_
```
1. Write a script to separate day, month, year, hour, minute and second values of the current system date and time.
2. Write a script to remove duplicated lines from a file.
3. Write a script to find the sum of all numbers in a file in Linux
```

### Script 1 - (script1.sh)
- How to run : `sh script1.sh`
- Time Complexity Analysis - `O(1)`
- Code Explaination
    
    - `$(date  +'%Y %m %d %H %M %S')` -> Formats current date into Year, Month, Day, Hour, Minutes, Seconds
    - `read` -> Reads the input 

### Script 2 - (script2.sh)
- How to run : `sh script2.sh`
- ### Approach 1 & 2
    - Time Complexity Analysis - `O(nlogn)`
    - Code Explaination

        - `sort` file.txt -> Reads the file and sorts according to alpabetical order.
        - `sort -u file.txt` -> Sorts and keeps only unique values
        - `uniq` -> Keeps the unique in sorted file

- ### Approach 3
    - Time Complexity Analysis - `O(n)`
    - Code Explaination

        - `awk '!seen[$0]++'` -> Keeps a hashmap of visited lines. If not visited then prints the line


### Script 3 - (script3.sh)
- How to run : `sh script3.sh`
- To generate file with random numbers, run
```
for ((i=0; i<1000000; i++)) ; do echo $RANDOM; done > random_numbers
```
- ### Approach 1
    - Time Complexity Analysis - `O(n)`
    - Code Explaination

        - `awk '{ sum += $1 } END { print sum }'` -> Sums the contents of the files and finally prints the sum


- ### Approach 2
    - Time Complexity Analysis - `O(n)`
    - Code Explaination

        - `while read l`; -> Read from file and start loop
        - `do ((s+=l));` - Add the current value to sum
        - `done<random_numbers;` - Break when file ends


## Runtime Analysis

1. Script - 2
    - Approach 1 & 2 uses sort and thus in case of larger files will perform slower
    - Approach 3 is faster as it directly loops over the file

2. Script - 3
    - Approach 1 is faster than Approach 2 because of `read` operation
    - It is because read needs to perform system call before even file is read and thus is slower

        