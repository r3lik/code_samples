fn main() {

    infinite_loop();
    while_loop()
}

fn infinite_loop() {
    
    let mut a = 0;
    loop {
        if a == 5 {
            println!("a is now 5, breaking infinite loop");
            break;
        }
        println!("{:?}", a);
        a += 1;
    }

}

fn while_loop() {

    let mut a = 0;
    while a != 5 {
        println!("{:?}", a);
        a += 1;
    }
}
