// * Undifined 类型

// ! 注意事项
// * 1. Undefined 类型只有一个值，就是特殊值 undefined。
// * 2. 声明变量的时候最好初始化，避免自动赋予 undefined。
// * 3. 对未声明的变量，只能执行一个有用的操作，就是对它调用 typeof。

// ? 声明未初始化
{
    let message;
    console.log(typeof message === undefined);
}
// ? undefined 和未定义变量区别
{
    let name;
    // let age;

    console.log(name); // undefined
    // console.log(age); // Uncaught ReferenceError: age is not defined
}

// ? undefined 是假值
{
    let flag;
    if (!flag) {
        console.log("flag is false");
    }
}
