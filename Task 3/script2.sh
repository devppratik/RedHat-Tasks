# Approach 1
# Time Complexity : 0(nlog(n))
sort -u lines.txt # Output is sorted in alphabetical order

# Approach 2
# Time Complexity : 0(nlog(n))
sort lines.txt | uniq # Output is sorted in alphabetical order

# Approach 3
# Time Complexity : 0(n)
awk '!seen[$0]++' lines.txt # Output is in order of appearance in the file
