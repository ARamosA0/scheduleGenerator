use genevo::{operator::prelude::*, prelude::*, random::Rng, types::fmt::Display};
use std::fmt;
use std::collections::HashMap;

use crate::models::{
    algoritm_config::AlgorithmConfig, algoritm_models::ClaseProgramada, subject::Subject,
};


type HorarioGenome = Vec<ClaseProgramada>;

fn calcular_colisiones(
    genoma: &HorarioGenome,
    subjects: &[Subject],
    params: &AlgorithmConfig,
) -> usize {
    let mut colisiones = 0;
    let mut periodo_ocupado = vec![vec![false; params.num_periods]; params.num_days];
    
    let mut cursos_por_dia: HashMap<(usize, usize), Vec<usize>> = HashMap::new(); // (dia, curso_id) -> [periodos]
    
    for clase in genoma.iter() {
        if clase.dia >= params.num_days || clase.bloque >= params.num_periods {
            colisiones += 100; 
            continue;
        }
        
        if periodo_ocupado[clase.dia][clase.bloque] {
            colisiones += 50; 
        } else {
            periodo_ocupado[clase.dia][clase.bloque] = true;
        }
        
        cursos_por_dia
            .entry((clase.dia, clase.curso_id))
            .or_insert_with(Vec::new)
            .push(clase.bloque);
    }
    
    for ((dia, curso_id), mut periodos) in cursos_por_dia {
        periodos.sort(); 
        
        if periodos.len() > 3 {
            colisiones += (periodos.len() - 3) * 30; 
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
                colisiones += 40; 
            }
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
        if genome.is_empty() {
            return 0;
        }

        let mut penalty = 0;
        
        // Cambio crítico: rastrear qué profesor está en cada slot
        let mut teacher_schedule: HashMap<(usize, usize), usize> = HashMap::new(); // (dia, bloque) -> teacher_id
        let mut room_schedule: HashMap<(usize, usize), usize> = HashMap::new(); // (dia, bloque) -> room_id
        let mut cursos_por_dia: HashMap<(usize, usize), Vec<usize>> = HashMap::new();

        for clase in genome.iter() {
            // Validar rangos
            if clase.dia >= self.param.num_days || clase.bloque >= self.param.num_periods {
                penalty += 1000;
                continue;
            }

            let slot = (clase.dia, clase.bloque);

            // Verificar colisión de profesores - CORRECCIÓN PRINCIPAL
            if let Some(existing_teacher) = teacher_schedule.get(&slot) {
                if *existing_teacher != clase.teacher_id {
                    // Solo penalizar si hay diferentes profesores en el mismo slot
                    penalty += 300;
                }
                // Si es el mismo profesor, también es problemático (doble clase)
                penalty += 200;
            } else {
                teacher_schedule.insert(slot, clase.teacher_id);
            }

            // Verificar colisión de aulas
            if let Some(existing_room) = room_schedule.get(&slot) {
                if *existing_room != clase.salon_id {
                    penalty += 250;
                }
                penalty += 200; // Doble uso del aula
            } else {
                room_schedule.insert(slot, clase.salon_id);
            }

            // Agrupar cursos por día
            cursos_por_dia
                .entry((clase.dia, clase.curso_id))
                .or_insert_with(Vec::new)
                .push(clase.bloque);
        }

        // Penalizar cursos no consecutivos y exceso de períodos por día
        for ((_, _), mut periodos) in cursos_por_dia {
            periodos.sort();
            periodos.dedup(); // Eliminar duplicados para contar correctamente
            
            if periodos.len() > 3 {
                penalty += (periodos.len() - 3) * 200;
            }

            // Verificar consecutividad solo si hay más de un período
            if periodos.len() >= 2 {
                for i in 1..periodos.len() {
                    if periodos[i] != periodos[i-1] + 1 {
                        penalty += 150;
                    }
                }
            }
        }

        // Convertir penalty a fitness
        let max_penalty = genome.len() * 1000;
        
        // Debug: uncomment para ver los valores
        if penalty == 0 {
            println!("¡Genoma perfecto encontrado! Penalty = 0");
        }
        
        if penalty >= max_penalty {
            0
        } else {
            let fitness = ((max_penalty - penalty) * 100) / max_penalty;
            // Asegurar que no todos tengan fitness 100
            if fitness >= 99 && penalty > 0 {
                99 - (penalty % 20) // Introducir más variabilidad
            } else {
                fitness
            }
        }
    }

    fn average(&self, values: &[usize]) -> usize {
        if values.is_empty() { 0 } else { values.iter().sum::<usize>() / values.len() }
    }

    fn highest_possible_fitness(&self) -> usize { 100 }
    fn lowest_possible_fitness(&self) -> usize { 0 }
}

// impl<'a> FitnessFunction<HorarioGenome, usize> for FitnessCalc<'a> {
//     fn fitness_of(&self, genome: &HorarioGenome) -> usize {
//         let colisiones = calcular_colisiones(genome, &self.subject, &self.param);
        
//         if genome.len() == 0 {
//             return 0;
//         }
        
//         let max_possible_collisions = genome.len() * 100; 
//         let normalized_collisions = colisiones.min(max_possible_collisions);
        
//         let score = if max_possible_collisions > 0 {
//             ((max_possible_collisions - normalized_collisions) * 100) / max_possible_collisions
//         } else {
//             100
//         };
        
//         score
//     }
    
//     fn average(&self, values: &[usize]) -> usize {
//         if values.is_empty() {
//             0
//         } else {
//             values.iter().sum::<usize>() / values.len()
//         }
//     }
    
//     fn highest_possible_fitness(&self) -> usize {
//         100
//     }
    
//     fn lowest_possible_fitness(&self) -> usize {
//         0
//     }
// }

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
