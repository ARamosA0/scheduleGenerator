use genevo::{operator::prelude::*, prelude::*, random::Rng, types::fmt::Display};
use std::fmt;
use std::collections::HashMap;

use crate::models::{
    algoritm_config::AlgorithmConfig, algoritm_models::ClaseProgramada, subject::{self, Subject},
};


type HorarioGenome = Vec<ClaseProgramada>;

#[derive(Clone, Debug)]
pub struct FitnessCalc<'a> {
    pub subject: Vec<Subject>,
    pub param: &'a AlgorithmConfig,
}

impl<'a> FitnessFunction<HorarioGenome, usize> for FitnessCalc<'a> {
    fn fitness_of(&self, genome: &HorarioGenome) -> usize {
        if genome.is_empty() {
            return 0;
        }

        let mut penalty = 0;
        
        // Cambio crítico: rastrear qué profesor está en cada slot
        let mut subject_in_schedule: HashMap<usize, usize> = HashMap::new();
        let mut teacher_schedule: HashMap<(usize, usize), usize> = HashMap::new(); // (dia, bloque) -> teacher_id
        let mut room_schedule: HashMap<(usize, usize), usize> = HashMap::new(); // (dia, bloque) -> room_id
        let mut clases_por_dia: HashMap<(usize, usize), Vec<usize>> = HashMap::new();

        for clase in genome.iter() {

            // Guardar la repeticion de cursos en el cromosoma
            if let Some(subject) = self.subject.iter().find(|s| s.id == clase.curso_id) {
                subject_in_schedule
                    .entry(subject.id) 
                    .and_modify(|count| *count += 1) 
                    .or_insert(1);           
            }


        }

        for clase in genome.iter() {
            if clase.dia >= self.param.num_days || clase.bloque >= self.param.num_periods {
                penalty += 1000;
                continue;
            }

            let slot = (clase.dia, clase.bloque);
            
            // Validacion de cantidad de cursos en el cromosoma
            if let Some(subject) = &self.subject.iter().find(|s| s.id == clase.curso_id) {
                if let Some(&cantidad) = subject_in_schedule.get(&subject.id){
                    if cantidad > subject.hours as usize{
                        penalty += 1000;
                    } 
                }
            } 

            // Validacion de repeticion de profesores
            if let Some(existing_teacher) = teacher_schedule.get(&slot) {
                if *existing_teacher != clase.teacher_id {
                    penalty += 500;
                }
                penalty += 300;
            } else {
                teacher_schedule.insert(slot, clase.teacher_id);
            }

            // Validacion de salon no solapado en la misma clase
            if let Some(existing_room) = room_schedule.get(&slot) {
                if *existing_room != clase.salon_id {
                    penalty += 550;
                }
                penalty += 300; 
            } else {
                room_schedule.insert(slot, clase.salon_id);
            }

        }

        for clase in genome.iter() {
            clases_por_dia
                .entry((clase.dia, clase.curso_id))
                .or_default()
                .push(clase.bloque);
        }

        // Revisar cada grupo
        for ((_dia, _curso), mut bloques) in clases_por_dia {
            bloques.sort();
            for w in bloques.windows(2) {
                let diff = w[1] as i32 - w[0] as i32;
                if diff > 1 {
                    // hay hueco → no son adyacentes
                    penalty += 700;
                }
            }
        }

        // Calculo del fitness
        let max_penalty = genome.len() * 1000;
        
        if penalty == 0 {
            // println!("¡Genoma perfecto encontrado! Penalty = 0");
        }
        
        // println!("------ INICIO FITNES ------");
        // println!("PENALTY: {:?}", penalty);
        
        if penalty >= max_penalty {
        // println!("FITNESS 0 ");
            0
        } else {
            let fitness = ((max_penalty - penalty) * 100) / max_penalty;
            // println!("FITNESS: {:?}", fitness);
            // println!("------ FIN FITNES ------");
            if fitness >= 99 && penalty > 0 {
                99 - (penalty % 20) 
            } else {
                fitness
            }
        }
    }

    fn average(&self, values: &[usize]) -> usize {
        println!("VALUES: {:?}", values);
        if values.is_empty() { 0 } else { values.iter().sum::<usize>() / values.len() }
    }

    fn highest_possible_fitness(&self) -> usize { 100 }
    fn lowest_possible_fitness(&self) -> usize { 0 }
}


fn imprimir_diagnostico_colisiones(genoma: &HorarioGenome, params: &AlgorithmConfig) {
    let mut cursos_por_dia: HashMap<(usize, usize), Vec<usize>> = HashMap::new();
    
    for clase in genoma.iter() {
        cursos_por_dia
            .entry((clase.dia, clase.curso_id))
            .or_insert_with(Vec::new)
            .push(clase.bloque);
    }
    
    println!("=== DIAGNÓSTICO DE COLISIONES ===");
    for ((dia, curso_id), mut periodos) in cursos_por_dia {
        periodos.sort();
        println!("Día {}, Curso {}: períodos {:?}", dia, curso_id, periodos);
        
        if periodos.len() > 3 {
            println!(" VIOLACIÓN: Más de 3 períodos ({} períodos)", periodos.len());
        }
        
        if periodos.len() >= 2 {
            let mut es_consecutivo = true;
            for i in 1..periodos.len() {
                if periodos[i] != periodos[i-1] + 1 {
                    es_consecutivo = false;
                    break;
                }
            }
            
            if !es_consecutivo {
                println!("VIOLACIÓN: Períodos no consecutivos");
            }
        }
    }
}
