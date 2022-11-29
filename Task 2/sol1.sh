cat ip_list.txt | sort | uniq -c | sort -nr | head -n 1 | awk '{$1="Most frequent IP is : "}1'

# Explaination of the code
# 1. cat ip_list.txt -> Reads the IP Address File
# 2. sort -> Sorts the file contents based on the IP
# 3. uniq -c -> Finds all the unique IPs
# 4. sort -nr -> Sorts the list in reverse order keeping the count
# 5. head -n 1 -> Prints the most recurring IP