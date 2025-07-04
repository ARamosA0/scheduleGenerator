use genevo::{operator::prelude::*, prelude::*, random::Rng, types::fmt::Display};
use std::fmt;

// ==============================
// Estructura para cursos
// ==============================
#[derive(Clone, Debug, PartialEq)]
pub struct Subject {
    pub id: usize,
    pub nombre: String,
    pub credits: usize,
    pub requiredRoomType: usize,
}
