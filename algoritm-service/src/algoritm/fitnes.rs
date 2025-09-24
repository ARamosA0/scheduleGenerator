use genevo::{operator::prelude::*, prelude::*, random::Rng, types::fmt::Display};
use std::collections::HashMap;
use std::fmt;

use crate::models::{
    algoritm_config::AlgorithmConfig, 
    algoritm_models::ClaseProgramada, 
    subject::{self, Subject},
};

type HorarioGenome = Vec<ClaseProgramada>;
type Slot = (usize, usize); // (dia, bloque)
type CourseDay = (usize, usize); // (dia, curso_id)

// Constantes para penalties - más fácil de mantener y configurar
const PENALTY_INVALID_SLOT: usize = 1000;
const PENALTY_EXCESS_HOURS: usize = 1000;
const PENALTY_TEACHER_CONFLICT: usize = 500;
const PENALTY_TEACHER_OVERLAP: usize = 300;
const PENALTY_ROOM_CONFLICT: usize = 550;
const PENALTY_ROOM_OVERLAP: usize = 300;
const PENALTY_NON_ADJACENT: usize = 700;
const PENALTY_TOO_MANY_CONSECUTIVE: usize = 400;
const PENALTY_TOO_MANY_FREE: usize = 400;

const MAX_CONSECUTIVE_PERIODS: usize = 3;
const MAX_FREE_PERIODS: usize = 2;

#[derive(Clone, Debug)]
pub struct FitnessCalc<'a> {
    pub subjects: &'a [Subject], // Más idiomático usar slice
    pub config: &'a AlgorithmConfig, // Mejor nombre que 'param'
}

impl<'a> FitnessCalc<'a> {
    pub fn new(subjects: &'a [Subject], config: &'a AlgorithmConfig) -> Self {
        Self { subjects, config }
    }

    // Constructor alternativo desde Vec (para compatibilidad)
    pub fn from_vec(subjects: &'a Vec<Subject>, config: &'a AlgorithmConfig) -> Self {
        Self { 
            subjects: subjects.as_slice(), 
            config 
        }
    }

    /// Validaciones de slots válidos
    fn validate_slot_bounds(&self, genome: &HorarioGenome) -> usize {
        genome
            .iter()
            .filter(|clase| {
                clase.dia >= self.config.num_days || clase.bloque >= self.config.num_periods
            })
            .count() * PENALTY_INVALID_SLOT
    }

    /// Validación de horas excedidas por materia
    fn validate_subject_hours(&self, genome: &HorarioGenome) -> usize {
        let mut subject_counts = HashMap::new();
        
        // Contar occurrencias de cada materia
        for clase in genome {
            *subject_counts.entry(clase.curso_id).or_insert(0) += 1;
        }

        let mut penalty = 0;
        for (subject_id, count) in subject_counts {
            if let Some(subject) = self.subjects.iter().find(|s| s.id == subject_id) {
                if count > subject.hours as usize {
                    penalty += PENALTY_EXCESS_HOURS * (count - subject.hours as usize);
                }
            }
        }
        
        penalty
    }

    /// Validación de conflictos de profesores y salones
    fn validate_resource_conflicts(&self, genome: &HorarioGenome) -> usize {
        let mut teacher_schedule = HashMap::new();
        let mut room_schedule = HashMap::new();
        let mut penalty = 0;

        for clase in genome {
            if clase.dia >= self.config.num_days || clase.bloque >= self.config.num_periods {
                continue; // Ya penalizado en validate_slot_bounds
            }

            let slot = (clase.dia, clase.bloque);
            
            // Validación de profesores
            match teacher_schedule.get(&slot) {
                Some(&existing_teacher) if existing_teacher != clase.teacher_id => {
                    penalty += PENALTY_TEACHER_CONFLICT;
                }
                Some(_) => {
                    penalty += PENALTY_TEACHER_OVERLAP;
                }
                None => {
                    teacher_schedule.insert(slot, clase.teacher_id);
                }
            }

            // Validación de salones
            match room_schedule.get(&slot) {
                Some(&existing_room) if existing_room != clase.salon_id => {
                    penalty += PENALTY_ROOM_CONFLICT;
                }
                Some(_) => {
                    penalty += PENALTY_ROOM_OVERLAP;
                }
                None => {
                    room_schedule.insert(slot, clase.salon_id);
                }
            }
        }

        penalty
    }

    /// Validación de clases adyacentes por día y curso
    fn validate_adjacent_classes(&self, genome: &HorarioGenome) -> usize {
        let mut classes_by_day: HashMap<CourseDay, Vec<usize>> = HashMap::new();
        
        // Agrupar por día y curso
        for clase in genome {
            classes_by_day
                .entry((clase.dia, clase.curso_id))
                .or_default()
                .push(clase.bloque);
        }

        let mut penalty = 0;
        for (_, mut periods) in classes_by_day {
            periods.sort_unstable(); // más rápido que sort()
            
            // Verificar que los períodos sean consecutivos
            for window in periods.windows(2) {
                if window[1] - window[0] > 1 {
                    penalty += PENALTY_NON_ADJACENT;
                }
            }
        }

        penalty
    }

    /// Validación de períodos consecutivos ocupados/libres
    fn validate_consecutive_periods(&self, genome: &HorarioGenome) -> usize {
        let mut schedule_grid = vec![vec![false; self.config.num_periods]; self.config.num_days];
        
        // Marcar períodos ocupados
        for clase in genome {
            if clase.dia < self.config.num_days && clase.bloque < self.config.num_periods {
                schedule_grid[clase.dia][clase.bloque] = true;
            }
        }

        let mut penalty = 0;
        
        // Validar cada día
        for day_schedule in &schedule_grid {
            let mut consecutive_occupied = 0;
            let mut consecutive_free = 0;

            for &is_occupied in day_schedule {
                if is_occupied {
                    consecutive_occupied += 1;
                    consecutive_free = 0;

                    if consecutive_occupied > MAX_CONSECUTIVE_PERIODS {
                        penalty += PENALTY_TOO_MANY_CONSECUTIVE;
                    }
                } else {
                    consecutive_free += 1;
                    consecutive_occupied = 0;

                    if consecutive_free > MAX_FREE_PERIODS {
                        penalty += PENALTY_TOO_MANY_FREE;
                    }
                }
            }
        }

        penalty
    }

    /// Calcula el fitness final basado en el penalty total
    fn calculate_final_fitness(&self, total_penalty: usize, genome_size: usize) -> usize {
        let max_penalty = genome_size * PENALTY_INVALID_SLOT;
        
        if total_penalty == 0 {
            return 100; // Fitness perfecto
        }

        if total_penalty >= max_penalty {
            return 0;
        }

        let base_fitness = ((max_penalty - total_penalty) * 100) / max_penalty;
        
        // Ajuste fino para fitness altos
        if base_fitness >= 99 && total_penalty > 0 {
            99_usize.saturating_sub(total_penalty % 20)
        } else {
            base_fitness
        }
    }
}

impl<'a> FitnessFunction<HorarioGenome, usize> for FitnessCalc<'a> {
    fn fitness_of(&self, genome: &HorarioGenome) -> usize {
        if genome.is_empty() {
            return 0;
        }

        let mut total_penalty = 0;

        // Ejecutar todas las validaciones
        total_penalty += self.validate_slot_bounds(genome);
        total_penalty += self.validate_subject_hours(genome);
        total_penalty += self.validate_resource_conflicts(genome);
        total_penalty += self.validate_adjacent_classes(genome);
        total_penalty += self.validate_consecutive_periods(genome);

        self.calculate_final_fitness(total_penalty, genome.len())
    }

    fn average(&self, values: &[usize]) -> usize {
        if values.is_empty() { 
            0 
        } else { 
            values.iter().sum::<usize>() / values.len() 
        }
    }

    fn highest_possible_fitness(&self) -> usize { 100 }
    fn lowest_possible_fitness(&self) -> usize { 0 }
}

#[derive(Debug)]
pub struct ScheduleDiagnostics {
    pub violations: Vec<String>,
    pub fitness_score: usize,
    pub total_penalty: usize,
}

impl<'a> FitnessCalc<'a> {
    /// Función mejorada de diagnóstico que retorna información estructurada
    pub fn diagnose(&self, genome: &HorarioGenome) -> ScheduleDiagnostics {
        let mut violations = Vec::new();
        let mut classes_by_course_day: HashMap<CourseDay, Vec<usize>> = HashMap::new();
        
        // Agrupar clases por día y curso
        for clase in genome {
            classes_by_course_day
                .entry((clase.dia, clase.curso_id))
                .or_default()
                .push(clase.bloque);
        }
        
        // Verificar violaciones
        for ((day, course_id), mut periods) in classes_by_course_day {
            periods.sort_unstable();
            
            if periods.len() > MAX_CONSECUTIVE_PERIODS {
                violations.push(format!(
                    "Día {}, Curso {}: {} períodos exceden el máximo de {}", 
                    day, course_id, periods.len(), MAX_CONSECUTIVE_PERIODS
                ));
            }
            
            if periods.len() >= 2 {
                let is_consecutive = periods.windows(2)
                    .all(|w| w[1] == w[0] + 1);
                
                if !is_consecutive {
                    violations.push(format!(
                        "Día {}, Curso {}: períodos no consecutivos {:?}", 
                        day, course_id, periods
                    ));
                }
            }
        }

        let fitness_score = self.fitness_of(genome);
        let total_penalty = self.calculate_total_penalty(genome);
        
        ScheduleDiagnostics {
            violations,
            fitness_score,
            total_penalty,
        }
    }

    fn calculate_total_penalty(&self, genome: &HorarioGenome) -> usize {
        let mut total = 0;
        total += self.validate_slot_bounds(genome);
        total += self.validate_subject_hours(genome);
        total += self.validate_resource_conflicts(genome);
        total += self.validate_adjacent_classes(genome);
        total += self.validate_consecutive_periods(genome);
        total
    }
}

impl fmt::Display for ScheduleDiagnostics {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        writeln!(f, "=== DIAGNÓSTICO DE HORARIO ===")?;
        writeln!(f, "Fitness Score: {}", self.fitness_score)?;
        writeln!(f, "Total Penalty: {}", self.total_penalty)?;
        writeln!(f, "Violaciones encontradas: {}", self.violations.len())?;
        
        for violation in &self.violations {
            writeln!(f, "  - {}", violation)?;
        }
        
        Ok(())
    }
}
