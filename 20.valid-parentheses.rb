#
# @lc app=leetcode id=20 lang=ruby
#
# [20] Valid Parentheses
#

# @lc code=start
# @param {String} s
# @return {Boolean}
def is_valid(s)
  stack = []
  s.each_char do |c|
    case c
    when '(', '[', '{'
      stack.push(c)
    when ')'
      return false if stack.pop() != '('
    when ']'
      return false if stack.pop() != '['
    when '}'
      return false if stack.pop() != '{'
    end
  end
  stack.empty?
end
# @lc code=end

