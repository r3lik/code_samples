#[derive(Debug)]
struct Rectangle {
    width: u32,
    height: u32,
}
// This is starting to look better with Structs
fn area(rectangle: &Rectangle) -> u32 {
    rectangle.width * rectangle.height
}

fn main() {
    let scale = 2;
    let rect1 = Rectangle {
        //width: 420,
        // can debug the expression to see what it evaluates to
        width: dbg!(420 * scale),
        height: 69,
    };
    println!(
        "The area of the rectangle is {} square pixels", area(&rect1)
    );

    println!("rect1 is {:?}", rect1);

    dbg!(&rect1);
}
