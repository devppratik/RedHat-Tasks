# Approach 1 - Assuming only the sum of digits
# Time Complexity - O(n) where n : length of string
# Subsitute every group of non-digits with ""
# In the single string obtained loop for the sum
# FS -> Field Separator
awk '{for(i=1;i<=NF;i++)s+=$i} END{print s}' FS="" <<< $(awk '{gsub(/([^[:digit:]]+|)/,"",$0);printf $0}' nums.txt)

# Approach 2- Assuming group of digits as  integers/decimals
# Time Complexity - O(n) where n : length of string
# Subsitute every group of alphabets with newline
# In the list of numbers obtained, add for sum
awk '{ sum += $1 } END { print sum }' <<< $(awk '{ gsub(/([^[[:digit:].-]+|[^[:alnum:].-]+])/,"\n",$0); printf $0 }' nums.txt)
