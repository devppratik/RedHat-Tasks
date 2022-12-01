start=`date +%s%N`

sort --parallel=8 book.txt | uniq -c | sort | tail -n1

end=`date +%s%N`
echo Time - `expr $end - $start` ns

# Explaination of the code
# 1. sort --parallel=8 book.txt -> Sorts the file contents parallely on n threads
# 3. uniq -c -> Finds all the unique IPs and keeps the count of occurances
# 4. sort > Sorts the list  based on the count of uniq characters
# 5. tail -n 1 -> Prints the most recurring IP