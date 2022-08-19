export const RandInt = (n: number) => {
    let minNum = 0;
    let maxNum = n;
    let value_a = Math.random() * (maxNum - minNum + 1) + minNum + "";
    let value_b = parseInt(value_a);
    return value_b
}