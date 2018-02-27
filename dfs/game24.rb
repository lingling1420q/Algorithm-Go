
class Game

  class << self

    def dfs24(item, index)
      $r << item
      $book[index] = true

      if $r.size == 7
        rc = $r.join('')
        rs = eval rc
        if rs == 24
          puts rc+"=24"
        end
      end

      $a.each_with_index do |ai, ix|
        if $book[ix] != true
          $mc.each do |mi|
            $r << mi
            dfs(ai, ix)
            $r.pop
            $r.pop
            $book[ix] = false
          end
        end
      end

    end

    def count24
      $a = [1,2,3,4]
      $mc = ['+', '-', '*', '/']

      $a.each_with_index do |item, index|
        $r = []
        $book = [false, false, false, false]
        dfs24(item, index)
      end
    end
  end
end

Game.count24

# 想必大家都玩过一个游戏，叫做“24点”：给出4个整数，要求用加减乘除4个运算使其运算结果变成24，4个数字要不重复的用到计算中。

# 例如给出4个数：1、2、3、4。我可以用以下运算得到结果24：

# 1*2*3*4 = 24；2*3*4/1 = 24