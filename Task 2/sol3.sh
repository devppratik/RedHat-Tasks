start=`date +%s%N`

awk '{
    a[$0]++;   # Keeps a hasmap count of IP address
    if(m<a[$0]){ 
        m=a[$0]; # Identifies the IP with most occurance
        s[m]=$0
    }} 
    END{print s[m]}' book.txt # Prints the result.

end=`date +%s%N`
echo Time - `expr $end - $start` ns
