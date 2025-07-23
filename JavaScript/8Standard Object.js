// 标准对象

// 1.Date 用于表示日期和时间
let now = new Date();
now; // Wed Jun 24 2015 19:49:22 GMT+0800 (CST)
now.getFullYear(); // 2015, 年份
now.getMonth(); // 5, 月份，注意月份范围是0~11，5表示六月
now.getDate(); // 24, 表示24号
now.getDay(); // 3, 表示星期三
now.getHours(); // 19, 24小时制
now.getMinutes(); // 49, 分钟
now.getSeconds(); // 22, 秒
now.getMilliseconds(); // 875, 毫秒数
now.getTime(); // 1435146562875, 以number形式表示的时间戳

// 2.RegExp 正则表达式
// 正则的两种写法
let re1 = /ABC\-001/;
let re2 = new RegExp('ABC\\-001');

re1; // /ABC\-001/
re2; // /ABC\-001/
// 正则的匹配
let re3 = /^\d{3}\-\d{3,8}$/;
re3.test('010-12345'); // true
re3.test('010-1234x'); // false
re3.test('010 12345'); // false
// 用()进行分组提取
let re = /^(\d{3})-(\d{3,8})$/;
re.exec('010-12345'); // ['010-12345', '010', '12345']
re.exec('010 12345'); // null
// 实话 看多少遍都还是记不得 还是交给厉害的ai来写正则表达式吧！
