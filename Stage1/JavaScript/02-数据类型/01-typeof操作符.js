// * typeof 操作符

const a = undefined;
const b = true;
const c = "true";
const d = 12;
const e = null;
const f = () => {};
const g = Symbol("foo");

console.log(typeof a); // undefined 表示值未定义
console.log(typeof b); // boolean 表示值为布尔值
console.log(typeof c); // string 表示值为字符串
console.log(typeof d); // number 表示值为数值
console.log(typeof e); // object 表示值为对象（而不是函数）或 null
console.log(typeof f); // function 表示值为函数
console.log(typeof g); // symbol 表示值为符号
