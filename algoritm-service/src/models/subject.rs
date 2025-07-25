use genevo::{operator::prelude::*, prelude::*, random::Rng, types::fmt::Display};
use std::fmt;

// ==============================
// Estructura para cursos
// ==============================
#[derive(Clone, Debug, PartialEq)]
pub struct Subject {
    pub id: usize,
    pub name: String,
    pub credits: usize,
    pub hours: u32,
    pub semester: u32,
    pub career: String,
    pub requirementes: String,
    pub description: String,
    pub requiredRoomType: usize,
}
