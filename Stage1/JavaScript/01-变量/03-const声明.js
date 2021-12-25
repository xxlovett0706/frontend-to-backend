// * const 声明

// ? 声明变量时必须同时初始化变量
{
    const name = "MasK";
    // const age;  // SyntaxError: Missing initializer in const declaration
}

// ? 重新赋值会报错
{
    const school = "shz";
    school = "bd"; // TypeError: Assignment to constant variable.
}
