use crate::models::algoritm_config::{AlgorithmConfig, DaySchedule, RawData};
use serde_json::Value;

pub fn format_json(data: RawData) -> AlgorithmConfig {
    let formated_days_range = format_days_range(&data.template);
    println!("FORMATED DAYS: {:?}", formated_days_range);
    let num_subjects = arrays_data_counter(&data.subjects);
    let num_rooms = arrays_data_counter(&data.rooms);
    let num_teachers = arrays_data_counter(&data.teachers);
    let num_days = count_days(&formated_days_range);
    let num_periods = periods_definition(&formated_days_range);

    println!("NUM_DAYS:\n{:#?}", num_days);

    AlgorithmConfig {
        num_subjects,
        num_rooms,
        num_teachers,
        num_days,
        num_periods,

        population: data.population as usize,
        generations: data.generations as u64,
        mutation: data.mutation,
        cross_over: data.cross_over,
        reinsertion: data.reinsertion,
        elitism: data.selection,

        template: formated_days_range,
        subjects: data.subjects,
        teachers: data.teachers,
        rooms: data.rooms,
    }
}

pub fn count_days(data: &Vec<DaySchedule>) -> usize {
    return data.len();
}

pub fn periods_definition(data: &Vec<DaySchedule>) -> usize {
    let mut all_periods = Vec::new();

    for day in data {
        all_periods.push(day.periods.len());
    }

    let mut average = 0;
    for item in &all_periods {
        average += item;
    }
    return average / all_periods.len();
}

pub fn format_days_range(data: &Value) -> Vec<DaySchedule> {
    if let Some(Value::String(days_range_str)) = data.get("daysRange") {
        let result: Result<Vec<DaySchedule>, _> = serde_json::from_str(days_range_str);
        match result {
            Ok(days) => days,
            Err(err) => {
                eprintln!("Error al parsear daysRange: {}", err);
                Vec::new()
            }
        }
    } else {
        Vec::new()
    }
}

pub fn arrays_data_counter(data: &Value) -> usize {
    match data {
        Value::Array(arr) => arr.len(),
        _ => 0,
    }
}
