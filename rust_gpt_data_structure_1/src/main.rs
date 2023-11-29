fn sorting(output: &str) -> Vec<i32> {
    let mut ret = Vec::new();
    if output.is_empty() {
        return  ret;
    }
    if !output.contains(",") {
        return ret;
    }
    let mut numbers: Vec<i32> = Vec::new();
    let num_ary:Vec<i32> = output
        .split(",")
        .map(|s| s.trim())
        .filter_map(|s| s.parse().ok())
        .collect();
    for row in num_ary {
        numbers.push(row);
    }
    numbers.sort();
    let min_value = numbers[0];
    let mut max_value = numbers[0];
    if numbers.len() > 1 {
        max_value = numbers[numbers.len()-1];
    }
    ret.push(min_value);
    ret.push(max_value);
    return  ret;
}

fn main() {
    let ary = "3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5".to_string();
    let sort_ary = sorting(&ary);
    for row in sort_ary {
        let str_value = row.to_string();
        println!("{}", str_value);
    }
}
