cat ip_list.txt | sort | uniq -c | sort -nr | head -n 1 | awk '{$1="Most frequent IP is : "}1'