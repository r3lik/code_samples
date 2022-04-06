fn main() {

    // mutating vars
    let mut x = 5;
    println!("THe value of x is {}", x);
    x = 6;
    println!("The value of x is {}", x);

    // shadowing
    // declare a new variable with the same name as a previous variable

    let x = 5;

    let x = x +1;

    {
        let x = x * 2;
        println!("The value of x in the inner scope is: {}", x);
    }

    println!("THe value of x is: {}", x);

    // changing type when shadowing

    let spaces = "         ";
    let spaces = spaces.len();
    println!("The length of spaces is {}", spaces);

}
