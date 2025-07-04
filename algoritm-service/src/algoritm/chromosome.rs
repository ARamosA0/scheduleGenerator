use crate::models::{algoritm_config::AlgorithmConfig, algoritm_models::*, subject::Subject};
use genevo::{operator::prelude::*, prelude::*, random::Rng, types::fmt::Display};
use rocket::serde::json::Json;
use serde::Serialize;
use std::fmt;

type HorarioGenome = Vec<ClaseProgramada>;

impl GenomeBuilder<HorarioGenome> for HorarioBuilder {
    fn build_genome<R>(&self, _: usize, rng: &mut R) -> HorarioGenome
    where
        R: Rng + Sized,
    {
        self.subject
            .iter()
            .map(|curso| ClaseProgramada {
                curso_id: curso.id,
                salon_id: rng.gen_range(0..self.config.num_salones),
                dia: rng.gen_range(0..self.config.num_dias),
                bloque: rng.gen_range(0..self.config.num_bloques),
            })
            .collect()
    }
}
