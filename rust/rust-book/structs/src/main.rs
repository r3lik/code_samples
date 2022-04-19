
#[derive(Debug)]
struct User {
    active: bool,
    username: String,
    email: String,
    sign_in_count: u64,
}

fn build_user(email: String, username: String) -> User {
    User {
        email, // Field Init Shorthand
        username,
        active: true,
        sign_in_count: 1,
    }
}

// tuple structs
struct Color(i32, i32, i32);
struct Point(i32, i32, i32);

fn main() {
    let mut user1 = User {
        email: String::from("test@example.com"),
        username: String::from("r3lik"),
        active: true,
        sign_in_count: 1,
    };

    // access struct values with dot notation
    user1.email = String::from("r3lik@g.com");

    // we can create a User with pre-populated values
    let user2 = build_user(
        String::from("test@yahoo.com"),
        String::from("r3lik2"),
    );

    // we can create an instance from another instance with some updated values
    let user3 = User {
        email: String::from("test2@yahoo.com"),
        ..user1
    };

    println!("{:?}", user3);
    //println!("{:?}", user1); // String in the username field was moved into user2

    let black = Color(1, 2, 3);
    let origin = Point(1, 2, 3);

    println!("accessing tuple values {} {} {}", black.0, black.1, black.2)
}
