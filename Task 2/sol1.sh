start=`date +%s%N`

cat book.txt | sort | uniq -c | sort -nr | head -n 1

end=`date +%s%N`
echo Time - `expr $end - $start` ns

# Explaination of the code
# 1. cat ip_list.txt -> Reads the IP Address File and pipes the output
# 2. sort -> Sorts the file contents based on the IP
# 3. uniq -c -> Finds all the unique IPs and keeps the count
# 4. sort -nr -> Sorts the list in reverse order
# 5. head -n 1 -> Prints the most recurring IP