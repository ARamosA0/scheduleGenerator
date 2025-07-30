use chrono::NaiveDate;
use genevo::{operator::prelude::*, prelude::*, random::Rng, types::fmt::Display};
use serde::{Deserialize, Serialize};
use std::fmt;

// ==============================
// Estructura para schedule
// ==============================
#[derive(Clone, Debug, PartialEq, Deserialize, Serialize)]
pub struct ScheduleResponse {
    pub id: usize,
    pub startedDate: NaiveDate,
    pub endDate: NaiveDate,
    pub title: String,
    pub tooltip: String,
}
