require 'prime'
require 'set'
require 'tsort'
include Math
ALP = ('a'..'z').to_a
INF = 0xffffffffffffffff
def max(a,b);              a > b ? a : b                              end
def min(a,b);              a < b ?  a : b                             end
def swap(a,b);             a, b = b, a                                end
def gif;                   gets.to_i                                  end
def gff;                   gets.to_f                                  end
def gsf;                   gets.chomp                                 end
def gi;                    gets.split.map(&:to_i)                     end
def gf;                    gets.split.map(&:to_f)                     end
def gs;                    gets.chomp.split.map(&:to_s)               end
def gc;                    gets.chomp.split('')                       end
def pr(num);               num.prime_division                         end
def pr?(num);              Prime.prime?(num)                          end
def digit(num);            num.to_s.length                            end
def array(s,ini=nil);      Array.new(s){ini}                          end
def darray(s1,s2,ini=nil); Array.new(s1){Array.new(s2){ini}}          end
def rep(num);              num.times{|i|yield(i)}                     end
def repl(st,en,n=1);       st.step(en,n){|i|yield(i)}                 end

s = gc
x,y = gi
tx,ty = 0,0
dir = 0

s.each do |ss|
  case ss
  when 'F'
    case dir
    when 0
      tx += 1
    when 1
      ty += 1
    when 2
      tx -= 1
    when 3
      ty -= 1
    end
  when 'T'
    case
    when dir == 0 && y - ty > 0
      dir = 1
    when dir == 0 && y - ty < 0
      dir = 3
    when dir == 0 && y - ty == 0
      dir = 1
    when dir == 1 && x - tx > 0
      dir = 0
    when dir == 1 && x - tx < 0
      dir = 2
    when dir == 1 && x - tx == 0
      dir = 0
    when dir == 2 && y - ty > 0
      dir = 1
    when dir == 2 && y - ty < 0
      dir = 3
    when dir == 2 && y = ty == 0
      dir = 1
    when dir == 3 && x - tx > 0
      dir = 0
    when dir == 3 && x - tx < 0
      dir = 2
    when dir == 3 && x - tx == 0
      dir = 0
    end
  end
end

puts tx == x && ty == y ? 'Yes' : 'No'
