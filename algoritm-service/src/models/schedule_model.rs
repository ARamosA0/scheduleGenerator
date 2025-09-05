use chrono::{NaiveDate, NaiveDateTime};
use genevo::{operator::prelude::*, prelude::*, random::Rng, types::fmt::Display};
use serde::{Deserialize, Serialize};
use std::fmt;

// ==============================
// Estructura para schedule
// ==============================
#[derive(Clone, Debug, PartialEq, Deserialize, Serialize)]
pub struct ScheduleResponse {
    pub bestGeneration: Vec<BestGenome>,  
    pub bestFitness: usize,
    pub iteration: u64,
}

#[derive(Clone, Debug, PartialEq, Deserialize, Serialize)]
pub struct BestGenome {
    pub id: usize,
    pub startDate: NaiveDateTime,
    pub endDate: NaiveDateTime,
    pub title: String,
    pub tooltip: String,
}

pub struct AlgoritmInformation {
    pub duration: usize,
    pub state: String,
    pub fitness: i64,
    pub generations: usize,
}
