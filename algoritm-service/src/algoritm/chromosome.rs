use crate::models::algoritm_models::*;
use genevo::{operator::prelude::*, prelude::*, random::Rng, types::fmt::Display};

type HorarioGenome = Vec<ClaseProgramada>;

impl<'a> GenomeBuilder<HorarioGenome> for HorarioBuilder<'a> {
    fn build_genome<R>(&self, _: usize, rng: &mut R) -> HorarioGenome
    where
        R: Rng + Sized,
    {
        // println!("CONFIG:{:?}", self.config);
        self.subject
            .iter()
            .map(|curso| ClaseProgramada {
                curso_id: curso.id,
                salon_id: rng.gen_range(0..self.config.num_rooms),
                dia: rng.gen_range(0..self.config.num_days),
                bloque: rng.gen_range(0..self.config.num_periods),
            })
            .collect()
    }
}
