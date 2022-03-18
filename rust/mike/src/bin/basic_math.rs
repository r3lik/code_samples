fn sum(a: i32, b: i32) -> i32 {
    a + b
}

fn display_result(result: i32) {
    println!("{:?}", result);
}

fn main() {
    
    let result_sum = sum(6, 999999);
    display_result(result_sum);
}
