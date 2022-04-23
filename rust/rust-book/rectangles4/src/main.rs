// The main reason for using methods instead of functions, 
// in addition to providing method syntax and not having to 
// repeat the type of self in every method’s signature, 
// is for organization. We’ve put all the things we can 
// do with an instance of a type in one impl block rather than 
// making future users of our code search for capabilities 
// of Rectangle in various places in the library we provide.

#[derive(Debug)]
struct Rectangle {
    width: u32,
    height: u32,
}

impl Rectangle {
    fn area(&self) -> u32 {
        self.width * self.height
    }

    fn can_hold(&self, other: &Rectangle) -> bool {
        self.width > other.width && self.height > other.height 
    }

    // we can also give a method the same name as one of the struct fields
    fn width(&self) -> bool {
        self.width > 0
    }

    // Associated function (associated with type but aren't methods)
    fn square(size: u32) -> Rectangle {
        Rectangle {
            width: size,
            height: size,
        }
    }
}

fn main() {
    let rect1 = Rectangle {
        width: 30,
        height: 50,
    };
    let rect2 = Rectangle {
        width: 10,
        height: 40,
    };
    let rect3 = Rectangle {
        width: 60,
        height: 45,
    };

    // To call this associated function, we use the :: syntax with the struct name
    let square = Rectangle::square(10);

    println!(
        "The area of the rectangle is {} square pixels.", rect1.area()
    );

    println!("Can rect1 hold react2? {}", rect1.can_hold(&rect2));
    println!("Can rect1 hold react3? {}", rect1.can_hold(&rect3));

    if rect1.width() {
        println!("The rectangle has a nonzero width; it is {}", rect1.width)
    };

    println!("Square: {:?}", square);
}
