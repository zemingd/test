N = gets.to_i
array = gets.strip.split.map(&:to_i)
array_a = array.sort
cnt = 0
for i in 0..(N-1) do
  if array[i] !=  array_a[i]
    cnt += 1
  end
end
puts cnt < 3 ? 'YES':'NO'