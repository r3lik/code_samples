fn main() {
    let mut count = 0;
    'counting_up: loop {
        println!("count = {}", count);
        let mut remaining = 10;

        loop {
            println!("remaining = {}", remaining);
            if remaining == 9 {
                break;
            }
            if count == 2 {
                break 'counting_up;
            }
            remaining -= 1;
        }

        count += 1;
    }

    println!("End count = {}", count);

    // while loops
    let mut number = 3;

    while number != 0 {
        println!("{}!", number);

        number -= 1;
    }

    println!("LIFTOFF!!!");

    // as for loop

    for n in (1..4).rev() {
        println!("{}!", n);
    }
    println!("LIFTOFF!!!");

    // looping through array
    let a = [10, 20, 30, 40, 50, 60, 70];
    //let a = [10, 20, 30, 40, 50]; // error prone: this causes a panic
    let mut index = 0;

    println!("\nwhile loop");
    while index <= 5 { // loop until this is no longer true
        println!("the value is: {}", &a[index]);

        index += 1;
    }
    // for loops
    // safer and better 
    println!("\nfor loop");
    for element in &a {
        println!("the value is: {}", element);
    }
}
