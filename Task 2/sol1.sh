start=`date +%s%N` # Keeping count of time in nanoseconds

cat book.txt | sort | uniq -c | sort -nr | head -n 1

end=`date +%s%N` # Getting the time after exceution
echo Time - `expr $end - $start` ns # Caluclation of time taken