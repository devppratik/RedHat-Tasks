# Splits the date into various components
read YYYY month day hh mm ss <<<$(date  +'%Y %m %d %H %M %S') 

# Print the result on the screen
echo "Today is Day: $day Month: $month Year: $YYYY"
echo "Time is Hour: $hh Minutes: $mm Seconds: $ss"