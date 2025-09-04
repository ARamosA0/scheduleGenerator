use crate::models::{algoritm_config::AlgorithmConfig, algoritm_models::*, subject::Subject, group::Group};
use rocket::{config, serde::json::Json};

use crate::algoritm::demo::{run_genetic_algorithm, BestSolutionResult};
use genevo::operator::MutationOp;
use genevo::{operator::prelude::*, prelude::*, random::Rng, types::fmt::Display};
use serde::Serialize;
use rand::seq::SliceRandom;
use rand::thread_rng;

impl RandomValueMutation for ClaseProgramada {
    fn random_mutated<R>(value: Self, _min: &Self, _max: &Self, rng: &mut R) -> Self
    where
        R: Rng + Sized,
    {
        ClaseProgramada {
            dia: rng.gen_range(_min.dia.._max.dia),
            bloque: rng.gen_range(_min.bloque.._max.bloque),
            salon_id: rng.gen_range(_min.salon_id.._max.salon_id),
            curso_id: rng.gen_range(_min.curso_id.._max.curso_id),
            teacher_id: rng.gen_range(_min.teacher_id.._max.teacher_id),
            ..value
        }
    }
}

impl HorarioMutator {
    pub fn new(config: AlgorithmConfig) -> Self {
        Self { config }
    }
}

fn create_subject_by_group(group: &Group) -> Option<usize> {
    let mut rng = thread_rng();
    group.subjects.choose(&mut rng).copied()
}

fn create_room_by_subject(subjects: &Vec<Subject>, subject_id: usize) -> Option<usize> {
    let mut rng= thread_rng();
    if let Some(subject) = subjects.iter().find(|s| s.id == subject_id) {
        Some(subject.required_room_type)
    } else {
        None
    }

}

fn create_teacher_by_subject(subject: &Vec<Subject>) -> Option<usize> {
    
    Some(0)
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
