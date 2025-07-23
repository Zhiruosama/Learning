// DOM操作
// 1.首先 我们得获取到元素
// 最常见就是通过ID来获取
document.getElementById('你的ID');
// 接着就是通过 CSS 选择器获取 (更灵活)：
document.querySelector('你的CSS选择器');
document.querySelectorAll('你的CSS选择器');

// 2.有时 需要创建一些DOM元素
document.createElement('标签名');

// 3.有了元素后 我们开始做更新修改
// 修改文本内容：
// 元素.innerText = '新文本内容'; (只修改可见文本，忽略 HTML 标签)
// 元素.innerHTML = '新的HTML内容'; (可以修改文本，也可以插入新的 HTML 结构)
// 修改属性：
// 元素.setAttribute('属性名', '属性值');

// 4.插入 DOM 元素 
// 追加子元素 (最常用)：
// 父元素.appendChild(子元素);
// 插入到指定子元素之前：
// 父元素.insertBefore(新元素, 参考元素);

// 5.删除 DOM 元素 
// 要从页面上移除一个元素，你需要通过它的父元素来执行删除操作。
// 删除子元素：
// 父元素.removeChild(要删除的子元素);

// 元素自删除 (知道父元素)：
// 如果你只有子元素本身，没有直接引用父元素，但又想删除这个子元素，你可以：
// 子元素.parentNode.removeChild(子元素);

// 删除子元素的时候很容易遇到索引变化的问题
// 对于list.children 这样的“活的集合”，就很可能会遇到索引问题
// 所以 要么从后往前删
// 要么将集合转换成数组再去做操作

