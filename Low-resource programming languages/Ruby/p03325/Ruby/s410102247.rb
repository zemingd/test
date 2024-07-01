def progression_operation(num, arr)
  count = 0
  while true do
    index_dev2 = -1
    arr.each_with_index do | n, i |
      if (n%2 == 0) then
        index_dev2 = i
        arr[i] /= 2
      else
        arr[i] *= 3
      end
    end
    count += 1
    if index_dev2 == -1 then
      break
    end
  end
  p count
end

input = gets()
number = input.to_i()
input = gets()
progression = input.split(" ")
progression.map!(&:to_i)
p progression

progression_operation(number, progression)