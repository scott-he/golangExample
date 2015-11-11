type TZ int // HL

const (
	UTC TZ = 0*60*60 // HL
	EST TZ = -5*60*60
)

// iota枚举:
const (
	bit0, mask0 uint32 = 1<<iota, 1<<iota - 1 // HL
	bit1, mask1 uint32 = 1<<iota, 1<<iota - 1
	bit2, mask2 // 缺省时, 和上一行相同 // HL
)

// 高精度:
const Ln2= 0.693147180559945309417232121458176568075500134360255254120680009
const Log2E= 1/Ln2 // 高精度 // HL


//（ 简单地讲，每遇到一次 const 关键字，iota 就重置为 0 ）。
//与c 的enum一样（未定义值，则默认从0开始）

//当然，常量之所以为常量就是恒定不变的量，因此我们无法在程序运行过程中修改它的值；如果你在代码中试图修改常量的值则会引发编译错误。

//引用 time 包中的一段代码作为示例：一周中每天的名称。

const (
    Sunday = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)
//你也可以使用某个类型作为枚举常量的类型：

type Color int
const (
    RED Color = iota // 0
    ORANGE // 1
    YELLOW // 2
    GREEN // ..
    BLUE
    INDIGO
    VIOLET // 6
)



