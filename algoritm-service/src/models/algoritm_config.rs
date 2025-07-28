use serde::{Deserialize, Serialize};
use serde_json::Value;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct AlgorithmConfig {
    pub num_subjects: usize,
    pub num_rooms: usize,
    pub num_teachers: usize,
    pub num_days: usize,
    pub num_periods: usize,

    pub population: usize,
    pub generations: u64,
    pub mutation: f64,
    pub cross_over: f64,
    pub reinsertion: f64,
    pub elitism: f64,

    pub template: Vec<DaySchedule>,
    pub subjects: Value,
    pub teachers: Value,
    pub rooms: Value,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct RawData {
    pub id: Option<u64>,
    pub created_at: Option<String>,
    pub updated_at: Option<String>,
    pub deleted_at: Option<String>,

    pub template: Value,
    pub subjects: Value,
    pub teachers: Value,
    pub rooms: Value,

    #[serde(rename = "processName")]
    pub process_name: String,

    pub population: i32,
    pub generations: i32,
    pub mutation: f64,

    #[serde(rename = "cross_over")]
    pub cross_over: f64,
    #[serde(rename = "selction")]
    pub selection: f64,
    pub reinsertion: f64,
}

#[derive(Debug, Deserialize, Serialize, Clone)]
pub struct Period {
    pub name: String,
    pub startHour: String,
    pub endHour: String,
}

#[derive(Debug, Deserialize, Serialize, Clone)]
pub struct DaySchedule {
    pub day: String,
    pub periods: Vec<Period>,
    pub status: bool,
    pub startHour: String,
    pub endHour: String,
}
