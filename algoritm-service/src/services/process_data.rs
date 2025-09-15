use crate::models::{
    algoritm_config::{AlgorithmConfig, DaySchedule, RawData},
    group::Group,
    room::Room,
    subject::Subject,
    teacher::Teacher,
};
use serde_json::Value;

pub fn format_json(data: RawData) -> AlgorithmConfig {
    let formated_days_range = format_days_range(&data.template);
    let formated_techers = format_teachers(&data.teachers);
    let formated_subjects = format_subjects(&data.subjects);
    let formated_rooms = format_rooms(&data.rooms);
    let formated_groups = format_groups(&data.groups);


    let num_subjects = arrays_data_counter(&data.subjects);
    let num_rooms = arrays_data_counter(&data.rooms);
    let num_teachers = arrays_data_counter(&data.teachers);
    let num_groups = arrays_data_counter(&data.groups);
    let num_days = count_days(&formated_days_range);
    let num_periods = periods_definition(&formated_days_range);

    println!("FORMATED_DAYS_RANGE:{:?}", formated_days_range);

    AlgorithmConfig {
        num_subjects,
        num_rooms,
        num_teachers,
        num_groups,
        num_days,
        num_periods,

        population: data.population as usize,
        generations: data.generations as u64,
        mutation: data.mutation,
        cross_over: data.cross_over,
        reinsertion: data.reinsertion,
        selection: data.selection,

        template: formated_days_range,
        subjects: formated_subjects,
        teachers: formated_techers,
        rooms: formated_rooms,
        groups: formated_groups,
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

pub fn format_subjects(data: &Value) -> Vec<Subject> {
    match serde_json::from_value::<Vec<Subject>>(data.clone()) {
        Ok(subjects) => subjects,
        Err(e) => {
            eprintln!("Error al deserializar subjects: {}", e);
            vec![]
        }
    }
}
pub fn format_teachers(data: &Value) -> Vec<Teacher> {
    match serde_json::from_value::<Vec<Teacher>>(data.clone()) {
        Ok(teacher) => teacher,
        Err(e) => {
            eprintln!("Error al deserializar teachers: {}", e);
            vec![]
        }
    }
}
pub fn format_rooms(data: &Value) -> Vec<Room> {
    match serde_json::from_value::<Vec<Room>>(data.clone()) {
        Ok(room) => room,
        Err(e) => {
            eprintln!("Error al deserializar rooms: {}", e);
            vec![]
        }
    }
}

pub fn format_groups(data: &Value) -> Vec<Group> {
    print!("GROUPS RAW VALUE: {}", data);
    match serde_json::from_value::<Vec<Group>>(data.clone()) {
        Ok(group) => group,
        Err(e) => {
            eprintln!("Error al deserializar rooms: {}", e);
            vec![]
        }
    }
}

pub fn arrays_data_counter(data: &Value) -> usize {
    match data {
        Value::Array(arr) => arr.len(),
        _ => 0,
    }
}
