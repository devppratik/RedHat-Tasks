# Approach 1 : 
# Time Complexity : 0(n)
awk '{ sum += $1 } END { print sum }' random_numbers

# Approach 2
# Time Complexity : 0(n)
s=0;
while read l; # Read operation is slower and thus takes more time even if TC is same
do ((s+=l));
done<random_numbers;
echo $s;
