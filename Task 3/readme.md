## Task-3

### _Problem Statement_
```
1. Write a script to separate day, month, year, hour, minute and second values of the current system date and time.
2. Write a script to remove duplicated lines from a file.
3. Write a script to find the sum of all numbers in a file in Linux
```

### Script 1 - (script1.sh)
- How to run : `sh script1.sh`
- Time Complexity Analysis - `O(1)`
- Space Complexity Analysis - `O(1)`
- Code Explaination
    
    - `$(date  +'%Y %m %d %H %M %S')` -> Formats current date into Year, Month, Day, Hour, Minutes, Seconds

### Script 2 - (script2.sh)
- How to run : `sh script2.sh`
- ### Approach 1
    - Time Complexity Analysis - `O(n)`
    - Space Complexity Analysis - `O(n)`
    - Code Explaination

        - `awk '{gsub(/\? |! |\. /,ORS)} 1'` -> Splits every paragraph into separate lines by subsiuting the delimiter characters (?, !, .) with new line.

        - `awk '!seen[$0]++'` -> Keeps a hashmap of visited lines. If not visited then prints the line. The input for this is the output provided by previous awk command

### Script 3 - (script3.sh)
- How to run : `sh script3.sh`
- ### Approach 1 - Assuming only the sum of digits
    - Time Complexity Analysis - `O(n)` where n : length of string
    - Code Explaination

        - `awk '{gsub(/([^[:digit:]]+|)/,"",$0);printf $0}'` -> Subsitute every group of non-digits with "" i.e remove them and concatenates all digits into one string

        - `awk '{for(i=1;i<=NF;i++)s+=$i} END{print s}' FS=""` -> Runs through each digit of the string, add to sum and finally prints the sum. The input for this is the output provided by previous awk command


- ### Approach 2 - Assuming group of digits as  integers/decimals
    - Time Complexity Analysis - `O(n)` where n : length of string
    - Code Explaination

        - `awk '{ gsub(/([^[[:digit:].-]+|[^[:alnum:].-]+])/,"\n",$0); printf $0 }'` - Subsitute every group of alphabets with newline, thus keeping numbers in individual lines

        - `awk '{ sum += $1 } END { print sum }'` - Sums the numbers in the individual lines and finally prints the sum
       