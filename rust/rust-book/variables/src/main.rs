fn main() {

    // mutating vars
    let mut x = 5;
    println!("THe value of x is {}", x);
    x = 6;
    println!("The value of x is {}", x);

    // shadowing

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
