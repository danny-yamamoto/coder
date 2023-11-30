use std::collections::HashMap;

fn parse_log(output: &str) -> HashMap<String, String> {
    let mut ret = HashMap::new();
    if output.is_empty() {
        return ret;
    }
    if !output.contains(":") {
        return  ret;
    }
    for row in output.lines() {
        if row.is_empty() {
            continue;
        }
        if !row.contains(":") {
            continue;
        }
        if let Some((log_date, log_value)) = row.split_once(":") {
            ret.insert(log_date.trim().to_string(), log_value.trim().to_string());
        }
    }
    ret
}

fn main() {
    let log_data = "2023-01-01: New Year's Day\n2023-02-14: Valentine's Day\n2023-03-17: St. Patrick's Day\n2023-03-20\n".to_string();
    let log_json = parse_log(&log_data);
    println!("{:#?}", log_json);
}
