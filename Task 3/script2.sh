# Time Complexity : O(n)
# Space Complexity : O(n)
awk '{gsub(/\? |! |\. /,ORS)} 1' input.txt  > lines.txt # Split line by delimiter
awk '!seen[$0]++' lines.txt # Get the unique lines from the total lines