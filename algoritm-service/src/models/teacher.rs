use genevo::{operator::prelude::*, prelude::*, random::Rng, types::fmt::Display};
use serde::{Deserialize, Serialize};
use std::fmt;

// ==============================
// Estructura para cursos
// ==============================
#[derive(Clone, Debug, PartialEq, Deserialize, Serialize)]
pub struct Teacher {
    #[serde(rename = "ID")]
    pub id: usize,
    pub name: String,
    #[serde(rename = "lastName")]
    pub last_name: String,
    pub email: String,
    pub phone: String,
    #[serde(rename = "specialty")]
    pub speciality: usize,
    pub available_days: Vec<String>,
}
