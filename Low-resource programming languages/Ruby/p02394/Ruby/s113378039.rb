while str = STDIN.gets
  
  n = str.split(" ")
  w = n[0].to_i
  h = n[1].to_i
  x = n[2].to_i
  y = n[3].to_i
  r = n[4].to_i

  if ((r <= x) && (x <= w - r)) && (( r <= y) && (y <= h - r)) #????§???¢???????????????????????´???
    puts "Yes"
  else
    puts "No"
end
end