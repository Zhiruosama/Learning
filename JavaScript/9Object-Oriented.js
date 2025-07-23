// 面向对象
// 除了一般的{}创建对象
// 还可以通过构造函数来创建对象
function Student(name) {
    this.name = name;
    this.hello = function () {
        alert('Hello, ' + this.name + '!');
    }
}

// let xiaoming = new Student('小明');
// xiaoming.name; // '小明'
// xiaoming.hello(); // Hello, 小明!
// 绑定的this指向新创建的对象，并默认返回this

// 原型继承
function PrimaryStudent(props) {
    // 调用Student构造函数，绑定this变量:
    Student.call(this, props);
    this.grade = props.grade || 1;
}
// 这时候的PrimaryStudent的原型链是
// new PrimaryStudent() --> PrimaryStudent.prototype --> Object.prototype --> null
// 正确应该是
// new PrimaryStudent() --> PrimaryStudent.prototype --> Student.prototype --> Object.prototype --> null

// 所以我们需要用一个空函数F 来进行原型继承
// PrimaryStudent构造函数:
function PrimaryStudent(props) {
    Student.call(this, props);
    this.grade = props.grade || 1;
}

// 空函数F:
function F() {
}

// 把F的原型指向Student.prototype:
F.prototype = Student.prototype;

// 把PrimaryStudent的原型指向一个新的F对象，F对象的原型正好指向Student.prototype:
PrimaryStudent.prototype = new F();

// 把PrimaryStudent原型的构造函数修复为PrimaryStudent:
PrimaryStudent.prototype.constructor = PrimaryStudent;

// 继续在PrimaryStudent原型（就是new F()对象）上定义方法：
PrimaryStudent.prototype.getGrade = function () {
    return this.grade;
};

// 创建xiaoming:
let xiaoming = new PrimaryStudent({
    name: '小明',
    grade: 2
});
xiaoming.name; // '小明'
xiaoming.grade; // 2

// 验证原型:
xiaoming.__proto__ === PrimaryStudent.prototype; // true
xiaoming.__proto__.__proto__ === Student.prototype; // true

// 验证继承关系:
xiaoming instanceof PrimaryStudent; // true
xiaoming instanceof Student; // true

function inherits(Child, Parent) {
    let F = function () {};
    F.prototype = Parent.prototype;
    Child.prototype = new F();
    Child.prototype.constructor = Child;
}
// 为了方便 可以将这一过程给封装起来


// 随着技术的迭代 后来ES6引进了class
// 学过别的面向对象编程例如c++ 对于这个就会比较熟悉了
class Student {
    constructor(name) {
        this.name = name;
    }

    hello() {
        alert('Hello, ' + this.name + '!');
    }
}
// 创建一个对象跟前面一样 还是需要new关键词
let xiaomin = new Student('小明');
xiaomin.hello();

// 但是对于派生类 继承等问题就无需考虑了
class PrimaryStudent extends Student {
    constructor(name, grade) {
        super(name); // 记得用super调用父类的构造方法!
        this.grade = grade;
    }

    myGrade() {
        alert('I am at grade ' + this.grade);
    }
}
// 一个extends直接解决问题