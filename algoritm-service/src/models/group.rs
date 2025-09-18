use genevo::{operator::prelude::*, prelude::*, random::Rng, types::fmt::Display};
use serde::{Deserialize, Serialize};
use std::fmt;

// ==============================
// Estructura para cursos
// ==============================
#[derive(Clone, Debug, PartialEq, Deserialize, Serialize)]
pub struct Group {
    #[serde(rename = "id")]
    pub id: usize,
    pub name: String,
    pub size: u32,
    pub subjects: Vec<usize>,
}
