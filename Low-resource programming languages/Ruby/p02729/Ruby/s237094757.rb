s = gets.chomp
n = s.length
num1 = ((n-1)/2)
num2 = ((n+3)/2)-1
 
str1 = s[0,num1]
str2 = s[num2,n]
 
print s == s.reverse && str1 == str1.reverse && str2 == str2.reverse ? 'Yes' : 'No'