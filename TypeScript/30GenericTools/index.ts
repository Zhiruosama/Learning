interface User {
    address: string;
    name?: string;
    age: number;
}

type PartialUser = Partial<User>;//所有属性都可选
// type PartialUser = {
//     address?: string | undefined;
//     name?: string | undefined;
//     age?: number | undefined;
// }

type RequiredUser = Required<User>//所有属性必选

type test = Pick<User, 'age' | 'name'>//提取部分属性

type test1 = Exclude<'a' | 'b' | 'c', 'c'>//排除部分属性

type test2 = Omit<User, 'age' | 'name'>//排除部分属性 并返回新的类型