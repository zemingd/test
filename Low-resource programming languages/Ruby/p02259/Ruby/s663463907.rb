# ??????????????????
def bubble_sort(ary, len)
  swp = 0
  (0...len-1).each {|n|
    (0...len-1-n).each{|j|
      if ary[j+1] < ary[j]
        ary[j], ary[j+1] = ary[j+1], ary[j] # swap????????????
        swp += 1
      end
    }
  }
  return ary, swp
end

# main
len = gets.to_i
ary = gets.split(' ').map(&:to_i) # ????§????????????????????????§???????????°?????????
sorted_ary, swp = bubble_sort(ary, len)
puts sorted_ary.join(' ') # ????§????????????????????????§??????
puts swp