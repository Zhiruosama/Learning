// Promise（承诺）是一个 JavaScript 对象，它代表了一个异步操作的最终完成（或失败），以及该异步操作的结果值

// 一个 Promise 对象在它生命周期中，只会处于以下三种状态之一：
// Pending (进行中)：初始状态，既不是成功也不是失败。异步操作正在执行中。
// Fulfilled (已成功)：异步操作成功完成，并返回了一个结果值。
// Rejected (已失败)：异步操作失败，并返回了一个错误原因。

// 一个 Promise 一旦从 Pending 状态变为 Fulfilled 或 Rejected，它的状态就不可逆转，也不能再改变。
// 一旦 Promise 状态确定（Fulfilled 或 Rejected），它就是 Settled (已敲定) 状态。

// Promise 对象主要通过它的两个方法来处理异步结果：
// .then() 方法：
// 用于处理 Promise **成功（Fulfilled）**时的结果，或者处理 成功和失败 两种情况。
// .then(onFulfilled)：只处理成功的情况。onFulfilled 是一个回调函数，它会接收到 Promise 成功时的结果值。
// .then(onFulfilled, onRejected)：同时处理成功和失败。onFulfilled 接收结果值，onRejected 接收错误原因。
// .catch() 方法：
// 用于专门处理 Promise **失败（Rejected）**时的错误。它是 .then(null, onRejected) 的语法糖，更推荐用来处理错误，因为它能捕获链中任何一步的错误。

// Promise 最强大的特性之一就是它的链式调用能力。then() 和 catch() 方法都会返回一个新的 Promise，这允许你将多个异步操作串联起来，形成一个清晰的序列。
// 假设 fetch 函数返回一个 Promise
fetch('/api/users') // 第一个异步操作：获取用户数据
  .then(response => {
    // 检查HTTP响应是否成功
    if (!response.ok) {
      throw new Error('网络请求失败'); // 如果不成功，抛出错误，进入下一个 .catch
    }
    return response.json(); // 将响应体解析成 JSON，这个操作本身也是异步的，所以这里返回一个新的 Promise
  })
  .then(users => {
    // 第二个 then：处理上一个 then 返回的 JSON 数据（用户的数组）
    console.log('获取到的用户数据:', users);
    // 假设我们现在想根据第一个用户的数据再发起一个请求
    return fetch(`/api/user/${users[0].id}`); // 返回一个新的 Promise
  })
  .then(response => response.json()) // 第三个 then：解析第二个请求的响应
  .then(singleUser => {
    // 第四个 then：处理获取到的单个用户数据
    console.log('获取到的第一个用户详情:', singleUser);
    // 这里可以返回一个普通值，或者另一个 Promise
    return '所有数据都已处理完毕！';
  })
  .catch(error => {
    // 统一的错误处理：捕获链中任何一个 then 抛出的错误
    console.error('操作链中发生错误:', error.message);
  })
  .finally(() => {
    // 无论成功还是失败，都会执行（比如清理资源、隐藏加载指示器）
    console.log('所有异步操作已完成（或终止）。');
  });

// Promise 的创建
// 可以使用 new Promise() 构造函数来创建自己的 Promise。它接受一个函数作为参数，这个函数被称为 执行器（Executor），执行器函数会接收两个参数：
// resolve：一个函数，当异步操作成功时调用它，并将结果值作为参数传递。调用 resolve() 会将 Promise 的状态从 pending 变为 fulfilled。
// reject：一个函数，当异步操作失败时调用它，并将错误原因作为参数传递。调用 reject() 会将 Promise 的状态从 pending 变为 rejected。

function loadDataAsync(url) {
  return new Promise((resolve, reject) => {
    const xhr = new XMLHttpRequest();
    xhr.open('GET', url);
    xhr.onload = function() {
      if (xhr.status === 200) {
        resolve(JSON.parse(xhr.responseText)); // 成功时调用 resolve
      } else {
        reject(new Error(`加载失败: ${xhr.statusText}`)); // 失败时调用 reject
      }
    };
    xhr.onerror = function() {
      reject(new Error('网络请求错误')); // 网络错误时调用 reject
    };
    xhr.send();
  });
}

// 使用自定义的 Promise
loadDataAsync('/api/products')
  .then(data => {
    console.log('产品数据:', data);
  })
  .catch(error => {
    console.error('加载产品数据出错:', error);
  });

  // 当Promise成功时  resolve(JSON.parse(xhr.responseText)); 
  // JSON.parse(xhr.responseText会被传到resolve处 也就是后面的data
  // 接下来就能对data进行一系列操作了