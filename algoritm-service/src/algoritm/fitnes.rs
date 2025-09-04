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
    let mut colisiones_periodo = 0;
    let mut salon_ocupado =
        vec![vec![vec![false; params.num_periods]; params.num_days]; params.num_rooms];

    let mut periodo_ocupado = vec![vec![false; params.num_periods]; params.num_days]; 

    let mut profesor_ocupado =
        vec![vec![vec![false; params.num_periods]; params.num_days]; params.num_teachers];
    // println!("PERIODO OCUPADO:{:?}", salon_ocupado);
    // println!("salones:{:?}", params.num_rooms);
    // for (i, clase) in genoma.iter().enumerate() {
    //     if salon_ocupado[clase.salon_id][clase.dia][clase.bloque] {
    //         colisiones += 1;
    //     } else {
    //         salon_ocupado[clase.salon_id][clase.dia][clase.bloque] = true;
    //     }
    // }

    for clase in genoma.iter() {
        // Verificar que los índices estén dentro de los límites
        if clase.dia >= params.num_days || clase.bloque >= params.num_periods {
            colisiones += 1; // Fuera de límites cuenta como colisión
            continue;
        }
        
        if periodo_ocupado[clase.dia][clase.bloque] {
            colisiones += 1;
        } else {
            periodo_ocupado[clase.dia][clase.bloque] = true;
        }
    }
    colisiones
}

#[derive(Clone, Debug)]
pub struct FitnessCalc<'a> {
    pub subject: Vec<Subject>,
    pub param: &'a AlgorithmConfig,
}


impl<'a> FitnessFunction<HorarioGenome, usize> for FitnessCalc<'a> {
    fn fitness_of(&self, genome: &HorarioGenome) -> usize {
        let colisiones = calcular_colisiones(genome, &self.subject, &self.param);

        if genome.len() == 0 {
            return 0;
        }

        // porcentaje de clases SIN colisiones
        let score = ((genome.len().saturating_sub(colisiones)) * 100) / genome.len();

        println!(
            "COLISIONES:{} LEN:{} FITNESS:{}",
            colisiones,
            genome.len(),
            score
        );

        score
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
