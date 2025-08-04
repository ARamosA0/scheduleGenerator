use genevo::{operator::prelude::*, prelude::*, random::Rng, types::fmt::Display};
use std::fmt;

use crate::models::{
    algoritm_config::AlgorithmConfig, algoritm_models::ClaseProgramada, subject::Subject,
};

// ==============================
// Función de Fitness
// ==============================
type HorarioGenome = Vec<ClaseProgramada>;
fn calcular_colisiones(
    genoma: &HorarioGenome,
    subjects: &[Subject],
    params: &AlgorithmConfig,
) -> usize {
    let mut colisiones = 0;
    let mut salon_ocupado =
        vec![vec![vec![false; params.num_periods]; params.num_days]; params.num_rooms];

    let mut profesor_ocupado =
        vec![vec![vec![false; params.num_periods]; params.num_days]; params.num_teachers];
    for (i, clase) in genoma.iter().enumerate() {
        // println!(
        //     "Clase #{i}: curso_id={}, salon_id={}, dia={}, bloque={}",
        //     clase.curso_id, clase.salon_id, clase.dia, clase.bloque
        // );
        // Verificar colisiones de salón
        if salon_ocupado[clase.salon_id][clase.dia][clase.bloque] {
            colisiones += 1;
        } else {
            salon_ocupado[clase.salon_id][clase.dia][clase.bloque] = true;
        }

        // Verificar colisiones de profesor
        // if let Some(subject) = subjects.iter().find(|c| c.id == clase.curso_id) {
        //     if profesor_ocupado[subject.profesor_id][clase.dia][clase.bloque] {
        //         colisiones += 1;
        //     } else {
        //         profesor_ocupado[subject.profesor_id][clase.dia][clase.bloque] = true;
        //     }
        // }
    }
    colisiones
}

#[derive(Clone, Debug)]
pub struct FitnessCalc<'a> {
    pub subject: Vec<Subject>,
    pub param: &'a AlgorithmConfig,
}
// pub struct FitnessCalc {
//     pub subject: Vec<Subject>,
//     pub param: AlgorithmConfig,
// }

impl<'a> FitnessFunction<HorarioGenome, usize> for FitnessCalc<'a> {
    fn fitness_of(&self, genome: &HorarioGenome) -> usize {
        let colisiones = calcular_colisiones(genome, &self.subject, &self.param);
        // println!(
        //     "GENOME:{:?} \nSUBJECT:{:?} \nPARAM:{:?}",
        //     genome, &self.subject, &self.param
        // );
        let max_colisiones = genome.len() * colisiones; // Máximo posible de colisiones

        // Fitness más alto = menos colisiones
        if max_colisiones == 0 {
            100
        } else {
            let score = (max_colisiones - colisiones) * 100 / max_colisiones;
            score as usize
        }
    }

    fn average(&self, values: &[usize]) -> usize {
        values.iter().sum::<usize>() / values.len()
    }

    fn highest_possible_fitness(&self) -> usize {
        100
    }

    fn lowest_possible_fitness(&self) -> usize {
        0
    }
}
