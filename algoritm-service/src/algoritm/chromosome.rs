use genevo::{operator::prelude::*, prelude::*, random::Rng, types::fmt::Display};
use rocket::serde::json::Json;
use serde::Serialize;
use std::fmt;

#[derive(Clone, Debug, PartialEq)]
pub struct ClassAssignment {
    pub group_id: u32,
    pub subject_id: u32,
    pub professor_id: u32,
    pub room_id: u32,
    pub day: u8,
    pub slot: u8,
}
