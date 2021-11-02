#### 一、题目链接
[有效数字](https://leetcode-cn.com/problems/valid-number/solution/you-xiao-shu-zi-by-leetcode-solution-298l/)

#### 二、题目大意
给定一个字符串，判断其是否符合我们指定的规则，能够被转化成一个整数。

#### 三、解题思路
参考官方解题报告，如何利用 FSM（有限状态自动机）来解决判定问题。主要两个核心点：
- 如何定义出所有的状态合集：初始状态、接受状态、其它状态（拒绝状态）
- 如何定义状态之间的转换
上述两件事情确定以后，接下来我们只需要遍历字符串进行状态转移，如果最终字符串遍历完成时状态停留在可接受状态则认为字符串是合法数字

#### 四、复杂度分析

#### 五、代码
```go
// 构造状态机利用状态机状态转移判定是否合法，这是一种很巧妙的解题思路
type (
	State    = int
	CharType = int
)

const (
	StateInitial             State = iota // 0: 初始状态
	StateSign                             // 1: 符号位
	StateInteger                          // 2: 整数部分
	StatePointWithInteger                 // 3: 左侧有整数的小数点
	StatePointWithoutInteger              // 4: 左侧无整数的小数点
	StateFraction                         // 5: 小数部分
	StateExp                              // 6: 指数 e
	StateExpSign                          // 7: 指数符号位 e+1, e-1
	StateExpInteger                       // 8: 指数整数部分
	StateEnd                              // 9: 结束状态
)

const (
	CharInteger CharType = iota // 0: 整数
	CharExp                     // 1: 指数 e 或 E
	CharPoint                   // 2: 小数点
	CharSign                    // 3: 符号位 -/+
	CharIllegal                 // 4: 非法字符
)

// 定义 FSM 状态转换关系即每种状态下遇到某类字符时的状态转移
// 其实本质上状态转换已经对整数的表示形式进行了相对应的限制了
var transation = map[State]map[CharType]State{
	StateInitial: {
		CharInteger: StateInteger,
		CharPoint:   StatePointWithoutInteger,
		CharSign:    StateSign,
	},
	StateSign: {
		CharInteger: StateInteger,
		CharPoint:   StatePointWithoutInteger,
	},
	StateInteger: {
		CharInteger: StateInteger,
		CharExp:     StateExp,
		CharPoint:   StatePointWithInteger,
	},
	StatePointWithInteger: {
		CharInteger: StateFraction, // 进入小数部分了
		CharExp:     StateExp,
	},
	StatePointWithoutInteger: {
		CharInteger: StateFraction,
	},
	StateFraction: {
		CharInteger: StateFraction, // 小数部分
		CharExp:     StateExp,
	},
	StateExp: {
		CharInteger: StateExpInteger,
		CharSign:    StateExpSign,
	},
	StateExpSign: {
		CharInteger: StateExpInteger,
	},
	StateExpInteger: {
		CharInteger: StateExpInteger,
	},
}

func toCharType(ch byte) CharType {
	switch ch {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return CharInteger
	case 'e', 'E':
		return CharExp
	case '.':
		return CharPoint
	case '+', '-':
		return CharSign
	default:
		return CharIllegal
	}
}

func isLegalEndState(state State) bool {
	return state == StateInteger || state == StatePointWithInteger || state == StateFraction || state == StateExpInteger
}

func isNumber(s string) bool {
	state := StateInitial
	for i := 0; i < len(s); i++ {
		typ := toCharType(s[i])
		if typ == CharIllegal {
			return false
		}
		if _, ok := transation[state][typ]; !ok {
			return false
		}
		state = transation[state][typ]
	}
	return isLegalEndState(state)
}
```