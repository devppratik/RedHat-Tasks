# Converts the file contest into a list
lines = [line.strip() for line in open("ip_list.txt", 'r')]

# Makes a dictionary of the unique IPs from list and stores the value of the occurance
cnt = {x:lines.count(x) for x in set(lines)}

# Finds the maxm value from the dictionary
n = max(cnt.values())

# Finds the IP that is most occuring
res = [k for k in cnt if cnt[k] == n]

# Prints the solution
print("Most recurring IP is : " + res[0]) 