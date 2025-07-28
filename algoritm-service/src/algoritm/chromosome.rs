use crate::models::algoritm_models::*;
use genevo::{operator::prelude::*, prelude::*, random::Rng, types::fmt::Display};

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
                // salon_id: rng.gen_range(0..self.config.num_salones),
                salon_id: rng.gen_range(0..10),
                // dia: rng.gen_range(0..self.config.num_dias),
                dia: rng.gen_range(0..5),
                // bloque: rng.gen_range(0..self.config.num_bloques),
                bloque: rng.gen_range(0..4),
            })
            .collect()
    }
}
