def bubble(a)
  n = a.length
  t = 0           
  /t is exchange_times/
  for i in 0..n-1
    j = 0
    while a[i-j] < a[i-j-1] and j <= i-1
      k = a[i-j]
      a[i-j] = a[i-j-1]
      a[i-j-1] = k
      j += 1
      t += 1
    end
  end
  print (a)
  print ("\n")
  print (t)
  print("\stimes_exhange\n")
  retun 0
end

