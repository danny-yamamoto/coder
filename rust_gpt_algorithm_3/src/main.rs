use std::collections::HashSet;

fn find_common_elements(input1: &str, input2: &str) -> Vec<String> {
    let set_1: HashSet<i32> = input1.split(",").map(|s| s.trim()).filter_map(|s| s.parse().ok()).collect();
    let set_2: HashSet<i32> = input2.split(",").map(|s| s.trim()).filter_map(|s| s.parse().ok()).collect();
    let mut common_numbers: Vec<String> = set_1.intersection(&set_2).into_iter().map(|n| n.to_string()).collect();
    common_numbers.sort();
    common_numbers
}

fn main() {
    let input1 = "1, 3, 4, 6, 7, 9".to_string();
    let input2 = "1, 2, 4, 5, 9, 10".to_string();
    let match_numbers = find_common_elements(&input1, &input2);
    let result = match_numbers.join(",");
    println!("{}", result);
}
