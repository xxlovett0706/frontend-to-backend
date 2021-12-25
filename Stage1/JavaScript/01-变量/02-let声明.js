// * let声明

// ! 注意事项
// * 1. 在定义将来要保存对象值的变量时，建议使用 null 来初始化，不要使用其他值。
// * 2. let 关键字不依赖条件声明。

// ? 作用域为块级作用域
{
    if (true) {
        let name = "MasK";
        console.log(name); // MasK
    }
    // console.log(name); // ReferenceError: name is not defined
}

// ? 暂时性死区
{
    function test1() {
        // console.log(age); // ReferenceError: Cannot access 'age' before initialization
        let age = 18;
    }
    test1();
}

// ? 全局声明
{
    var name1 = "John";
    let name2 = "Juice";
    console.log(window.name1); // John
    console.log(window.name2); // undefined
}

// ?条件声明
{
    // let school = "shz";
    // let grade = 1;

    if (typeof school === "undefined") {
        let school;
    }
    school = "bd"; // 如果 36 行存在，则修改的是 36 行 school；反之，则声明了一个全局变量 school
    console.log(school);

    try {
        console.log(grade);
    } catch (error) {
        let grade;
    }
    grade = 2; // 如果 37 行存在，则对 37 行的 grade 重新赋值；反之，则声明了一个全局变量 grade
    console.log(grade);
}

// ? for 循环中的 let 声明
{
    for (let i = 0; i < 5; i++) {
        setTimeout(() => {
            console.log("let中的i：", i); // 0, 1, 2, 3, 4
        }, 100);
    }

    for (var i = 0; i < 5; i++) {
        setTimeout(() => {
            console.log("var中的i：", i); // 5, 5, 5, 5, 5
        }, 100);
    }
}
