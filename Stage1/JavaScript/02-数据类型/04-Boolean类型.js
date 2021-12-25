// * Boolean 类型

// ! 注意事项
// * 1. Boolean 类型同样有两个值，true 和 false。
// * 2. 这两个布尔值不同于数值，因此 true 不等于 1，false 不等于 0。
// * 3. 等价形式用于流程控制语句的自动转换

// ?等价形式
{
    // String 类型：非空字符串和空字符串
    console.log("非空字符串", Boolean("s1"));
    console.log("空字符串", Boolean(""));

    // Number 类型：非零数值（包括无穷大）、0、NaN
    console.log("非零数值（包括无穷大）", Boolean(10));
    console.log("0", Boolean(0));
    console.log("NaN", Boolean(NaN));

    // Object 类型：任意对象和 null
    console.log("任意对象", Boolean({}));
    console.log("null", Boolean(null));

    // Undefined 类型： undefined
    console.log("undefined", Boolean(undefined));
}
