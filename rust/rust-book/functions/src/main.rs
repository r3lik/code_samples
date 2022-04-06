fn main() {

// Statements are instructions that perform some action and do not return a value 
// Expressions evaluate to a resulting value

    let y = {
        let x = 3;
        x + 1 // expression
    };

    println!("y is {}", y);
    
    print_labeled_measurement(5, 'h');
}

fn print_labeled_measurement(value: i32, unit_label: char) {
    println!("The measurement is: {}{}", value, unit_label);

}
