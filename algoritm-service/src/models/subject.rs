use genevo::{operator::prelude::*, prelude::*, random::Rng, types::fmt::Display};
use serde::{Deserialize, Serialize};
use std::fmt;

// ==============================
// Estructura para cursos
// ==============================
#[derive(Clone, Debug, PartialEq, Deserialize, Serialize)]
pub struct Subject {
    #[serde(rename = "ID")]
    pub id: usize,
    pub name: String,
    pub credits: usize,
    pub hours: u32,
    pub semester: u32,
    pub career: String,
    pub requirements: String,
    pub description: String,
    pub required_room_type: usize,
    #[serde(rename = "specialty")]
    pub speciality: usize,
}
