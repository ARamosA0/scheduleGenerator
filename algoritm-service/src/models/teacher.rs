use genevo::{operator::prelude::*, prelude::*, random::Rng, types::fmt::Display};
use std::fmt;

// ==============================
// Estructura para cursos
// ==============================
#[derive(Clone, Debug, PartialEq)]
pub struct Teacher {
    pub id: usize,
    pub name: String,
    pub last_name: String,
    pub email: String,
    pub phone: String,
    pub speciality: String,
    pub available_days: Vec<String>,
}
