fn sum_of_numbers_in_string(input: &str) -> i32 {
    let mut total = 0;
    let numbers: Vec<i32> = input.split(",").map(|s| s.trim()).filter_map(|s| s.parse().ok()).collect();
    for num in numbers {
        total += num;
    }
    total
}

fn main() {
    let input = "1, 2, 3, 4, 5".to_string();
    let total = sum_of_numbers_in_string(&input);
    println!("{}", total);
}
