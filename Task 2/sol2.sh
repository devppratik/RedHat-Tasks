start=`date +%s%N` # Keeping count of time in nanoseconds

sort --parallel=8 book.txt | uniq -c | sort | tail -n1

end=`date +%s%N` # Getting the time after exceution
echo Time - `expr $end - $start` ns # Caluclation of time taken