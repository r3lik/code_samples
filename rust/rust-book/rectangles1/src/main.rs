fn main() {
    let width1 = 30;
    let height1 = 50;

    println!(
        "The area of the rectangle is {} square pixels.", area(width1, height1)
    );
}

// The area function is supposed to calculate the area of one rectangle,
// but the function we wrote has two parameters, and it’s not clear 
// anywhere in our program that the parameters are related. 
// It would be more readable and more manageable to group width and height together. 
fn area(width: u32, height: u32) -> u32 {
    width * height
}
