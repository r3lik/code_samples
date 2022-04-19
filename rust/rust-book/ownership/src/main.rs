fn main() {
    let mut s = String::from("hello"); // allocated on heap and can be mutable
    s.push_str(", world!");
    println!("{}", s);

    let rs = "hello"; // string literal gets allocated on stack
    println!("String literal on stack: {}", rs);

    // MOVING

    // doesn't compile because ownership is moved and s1 is invalidated 
    //let s1 = String::from("hello");
    //let s2 = s1; // double free error in most languages
    // To ensure memory safety, after the line let s2 = s1, Rust considers s1 as no longer valid.
    // Therefore, Rust doesn’t need to free anything when s1 goes out of scope

    //println!("{}, world!", s1);

    // can clone heap data (expensive)
    let s1 = String::from("hello");
    let s2 = s1.clone();

    println!("s1 = {}, s2 = {}", s1, s2);

    let x = 5;
    let y = x; // on stack, so we can Copy

    println!("x = {}, y = {}", x, y);

    // BORROWING

    let r1 = String::from("hello");
    let len = calc_length(&r1);

    println!("The length of '{}' is {}", r1, len);
    
    fn calc_length(s: &String) -> usize {
        s.len()
    }

    // mutable references
    // references are immutable by default

    let mut r = String::from("hello");
    change(&mut r);

    fn change(some_string: &mut String) {
        some_string.push_str(", world");
    }

}
