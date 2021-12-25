// * Null 类型

// ! 注意事项
// * 1. Null 类型同样只有一个值，即特殊值 null。
// * 2. null 值表示一个空对象指针，这也是给 typeof 传一个 null 会返回"object"的原因。
// * 3. 在定义将来要保存对象值的变量时，建议使用 null 来初始化，不要使用其他值。

// ? typeof null 为 object
{
    let car = null;
    console.log(typeof car); // object
}
// ? null 是假值
{
    let car = null;
    if (!car) {
        console.log("null is false");
    }
}
