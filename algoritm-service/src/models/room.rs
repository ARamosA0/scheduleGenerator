use genevo::{operator::prelude::*, prelude::*, random::Rng, types::fmt::Display};
use std::fmt;

// ==============================
// Estructura para cursos
// ==============================
#[derive(Clone, Debug, PartialEq)]
pub struct Room {
    pub id: usize,
    pub code: String,
    pub name: String,
    pub capacity: u32,
    pub room_type: usize,
    pub floor: u32,
    pub building: String,
    pub observations: String,
}
