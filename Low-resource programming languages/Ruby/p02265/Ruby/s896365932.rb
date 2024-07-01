def array_converter(i)
  num_array=[]
  for num in 1..i
    line = readline.split(/[[:space:]]/)
    if judge_key_and_calc(num_array,line[0],line[1])
      num_array.push line[1]
    end
  end
  num_array
end

def judge_key_and_calc(array,key,val)
  if key == "delete"
    array.delete_at(array.index(val))
    false
  elsif key == "deleteFirst"
    array.delete(array.last)
    #array.shift
    false
  elsif key == "deleteLast"
    array.shift
    #array.delete(array.last)
    false
  else
    true
  end
end

def queue_calc
  num = readline.to_i
  array_converter(num).reverse.join(" ")
end

puts queue_calc