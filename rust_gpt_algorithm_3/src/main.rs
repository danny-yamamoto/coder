use std::vec;

fn find_common_elements(input1: &str, input2: &str) -> String {
    let elments = "".to_string();
    let mut common_numbers = Vec::new();
    let mut index = 0;
    let numbers1: Vec<i32> = input1.split(",").map(|s| s.trim()).filter_map(|s| s.parse().ok()).collect();
    let numbers2: Vec<i32> = input2.split(",").map(|s| s.trim()).filter_map(|s| s.parse().ok()).collect();
    for num1 in numbers1 {
        for num2 in numbers2 {
            if num1 == num2 {
                common_numbers.insert(index, num1.to_string());
                index += 1;
                continue;
            }
        }
    }
    return elments;
}

fn main() {
    // array1
    // array2
    // find_common_elements
    // print
    println!("hello")
}
