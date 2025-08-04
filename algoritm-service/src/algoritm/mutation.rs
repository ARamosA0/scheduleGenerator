use crate::models::{algoritm_config::AlgorithmConfig, algoritm_models::*, subject::Subject};
use rocket::{config, serde::json::Json};

use crate::algoritm::demo::{run_genetic_algorithm, BestSolutionResult};
use genevo::operator::MutationOp;
use genevo::{operator::prelude::*, prelude::*, random::Rng, types::fmt::Display};
use serde::Serialize;

impl RandomValueMutation for ClaseProgramada {
    fn random_mutated<R>(value: Self, _min: &Self, _max: &Self, rng: &mut R) -> Self
    where
        R: Rng + Sized,
    {
        // println!("value:{:?} \nmin:{:?} \nmax:{:?} \n", value, _min, _max);
        ClaseProgramada {
            dia: rng.gen_range(_min.dia.._max.dia),
            bloque: rng.gen_range(_min.bloque.._max.bloque),
            salon_id: rng.gen_range(_min.salon_id.._max.salon_id),
            ..value
        }
    }
}

impl HorarioMutator {
    pub fn new(config: AlgorithmConfig) -> Self {
        Self { config }
    }
}

// impl MutationOp<HorarioGenome> for HorarioMutator {
//     fn mutate<HorarioGenome>(&self, genome: &mut HorarioGenome, rng: &mut R) -> usize
//     where
//         R: Rng + Sized,
//     {
//         let mut mutations_applied = 0;

//         for clase in genome.iter_mut() {
//             if rng.gen_bool(self.mutation_rate) {
//                 clase.dia = rng.gen_range(0..self.config.num_dias);
//                 clase.bloque = rng.gen_range(0..self.config.num_bloques);
//                 clase.salon_id = rng.gen_range(0..self.config.num_salones);
//                 mutations_applied += 1;
//             }
//         }

//         mutations_applied
//     }
// }
