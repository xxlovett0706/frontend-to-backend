// * var 声明

// ? 作用域为函数作用域
{
    function test1() {
        var name = "MasK";
    }
    test1();
    // console.log(name); // ReferenceError: name is not defined
}

// ? 声明提升
{
    function test2() {
        console.log(age);
        var age = 18;
    }
    test2(); // undefined
}
