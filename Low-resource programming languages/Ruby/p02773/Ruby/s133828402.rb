n = gets.to_i
strs = n.times.map{ gets.chomp.to_s }.sort.append('1')

max_str_ary = []
max_num = 1
count_hash = Hash.new(0)
pre_elem = ''
pre_num = 0

strs.each_with_index do |elem, i|
  count_hash[elem] += 1
  if pre_elem != elem
    pre_elem = elem
    cnt = i+1 - pre_num
    pre_num = i+1
    if cnt > max_num
      max_num = cnt
    end
  end
end

puts count_hash.find_all { |_, v| v == max_num }.to_h.keys
