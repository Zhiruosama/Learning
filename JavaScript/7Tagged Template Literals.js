// 标签函数
// 简单来说，标签函数是一个在模板字符串前面直接写上的函数。 
// 当你这么做时，JavaScript 会自动调用这个函数，并将模板字符串的内容以及其中嵌入的表达式
// （就是 ${...} 包起来的部分）作为参数传递给它。

// 它的主要目的是让你能够处理模板字符串的原始文本和嵌入的变量，从而实现更复杂的字符串处理或 DSL（领域特定语言）的构建。
// 常见的就是处理sql语句

// 我们再来看这段代码：

const email = "test@example.com";
const password = 'hello123';

function sql(strings, ...exps) {
    console.log("--- sql 函数内部发生了什么 ---");
    console.log("strings (静态文本部分):", strings);
    console.log("exps (表达式值部分):", exps);

    // 假设这是一个真正的数据库工具会做的：
    // 1. 构建一个带占位符的 SQL 模板
    let queryTemplate = strings[0];
    for (let i = 0; i < exps.length; i++) {
        queryTemplate += `?` + strings[i + 1]; // 每个表达式位置放一个问号占位符
    }
    console.log("构建的 SQL 模板 (带占位符):", queryTemplate);

    // 2. 将参数值独立出来
    const parameters = exps;
    console.log("分离出的参数:", parameters);

    console.log("--- sql 函数内部处理结束 ---");

    // 实际的数据库操作会在这里发生，它会接收 queryTemplate 和 parameters
    // 为了演示，我们只是返回一个模拟的结果
    return {
        executedQuery: queryTemplate,
        executedParameters: parameters,
        mockResult: { name: '小明', age: 20 } // 假设这是数据库返回的数据
    };
}

// 这是我们调用标签函数的方式：
const result = sql`SELECT * FROM users WHERE email=${email} AND password=${password}`;

console.log("\n--- sql 标签函数调用后的结果 ---");
console.log("最终返回的对象:", result);
// 如果是真实数据库，你会拿到数据库查询到的数据，而不是一个模拟对象
console.log("模拟查询结果数据:", JSON.stringify(result.mockResult));

// 当 JavaScript 引擎看到 sql 后面紧跟着一个反引号包裹的模板字符串时：

// const result = sql`SELECT * FROM users WHERE email=${email} AND password=${password}`;
// 它就会自动触发 sql 这个标签函数的调用，并以一种特殊的方式给它传递参数：
// 第一个参数：strings 数组 (静态文本)
// JavaScript 会把模板字符串中所有固定不变的文本部分（不包括 ${} 里面的变量）提取出来，按顺序放到一个数组里，作为 sql 函数的第一个参数 strings。
// 在我们的例子中：
// SELECT * FROM users WHERE email= 是第一个固定文本。
//  AND password= 是第二个固定文本。
// 模板字符串的末尾（在最后一个 ${password} 之后）虽然没有明文字符，但也被视为一个空字符串 ''。
// 所以，strings 数组会是：
// strings: ["SELECT * FROM users WHERE email=", " AND password=", ""]
// 剩余参数：...exps 数组 (动态表达式的值)
// JavaScript 会把模板字符串中所有 ${} 表达式的值提取出来，按顺序放到一个数组里，作为 sql 函数的剩余参数 exps。
// 在我们的例子中：
// 第一个表达式是 ${email}，它的值是 "test@example.com"。
// 第二个表达式是 ${password}，它的值是 "hello123"。
// 所以，exps 数组会是：
// exps: ["test@example.com", "hello123"]
// sql 函数内部如何处理这些参数？
// 一旦 sql 函数接收到 strings 和 exps 这两个数组，它就可以开始进行我们想要的处理了：
// 构建占位符 SQL 模板：
// 函数内部通过遍历 strings 和 exps，将 exps 中的每一个动态值替换成一个占位符（通常是 ? 或 $1, $2 等）。
// queryTemplate = strings[0];  // "SELECT * FROM users WHERE email="
// queryTemplate += '?' + strings[1]; // "SELECT * FROM users WHERE email=? AND password="
// queryTemplate += '?' + strings[2]; // "SELECT * FROM users WHERE email=? AND password=?"
// 这样我们就得到了一个纯粹的 SQL 模板："SELECT * FROM users WHERE email=? AND password=?"。
// 分离参数：
// exps 数组本身就已经包含了所有动态参数的值：["test@example.com", "hello123"]。
// 通过这种方式，我们成功地将 SQL 语句的 结构（模板） 和 数据（参数） 完全分离开了。
// 为什么这很重要？（关键的应用价值）
// 防止 SQL 注入！
// 这是最主要的原因。如果用户输入的 email 是 "test@example.com'; DROP TABLE users; --"：
// 标签函数会把它作为一个完整的值放入 exps 数组。
// 数据库系统在执行 queryTemplate 和 parameters 时，会把 'test@example.com'; DROP TABLE users; -- 作为一个字符串数据来处理，而不是 SQL 代码。
// 数据库会试图在 email 字段中查找这个长字符串，而不是执行 DROP TABLE users。这就从根本上避免了恶意代码的执行。
// 提高 SQL 可读性：
// 你可以像写普通的 SQL 语句一样写在反引号里，包括换行和缩进，这比传统的字符串拼接方式（"SELECT * FROM users WHERE email='" + email + "' AND password='" + password + "'"）清晰多了。
// 标准化数据库操作：
// 在大型项目中，通常会有一个统一的数据库操作层。标签函数提供了一种优雅的方式来封装这些操作，确保所有 SQL 查询都通过相同的安全机制处理。