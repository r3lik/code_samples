fn main() {
    let mut s = String::from("hello world");

    //let word = first_word(&s);
    let word = first_word_slice(&s);
    s.clear();
    println!("the first word is {}", word);
}
// without slices
fn first_word(s: &String) -> usize {
    let bytes = s.as_bytes(); //array of bytes

    for (i, &item) in bytes.iter().enumerate() {
        if item == b' ' {
            println!("index is {}", i);
            return i;
        }
    }
    // If we find a space, we return the position.
    // Otherwise, return the length of the string by using s.len():
    println!("length of word is {}", s.len());
    s.len()

}

// String Slices
fn first_word_slice(s: &str) -> &str {
    let bytes = s.as_bytes();

    for (i, &item) in bytes.iter().enumerate() {
        if item == b' ' {
            return &s[0..i];
        }
    }

    &s[..]
}
