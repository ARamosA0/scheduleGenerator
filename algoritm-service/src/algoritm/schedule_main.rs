use crate::models::{algoritm_config::AlgorithmConfig, subject::Subject};
// use create::{chromosome, crossover, fitnes, mutation};
// use crate::algoritm::chromosome::HorarioBuilder;
use crate::algoritm::fitnes::FitnessCalc;

use crate::models::algoritm_models::HorarioBuilder;
use genevo::{operator::prelude::*, prelude::*, random::Rng, types::fmt::Display};
use rocket::config;
use std::fmt;

// ==============================
// Configuración del problema
// ==============================
const NUM_PROFESORES: usize = 4;
const NUM_CURSOS: usize = 8;
const NUM_SALONES: usize = 5;
const NUM_DIAS: usize = 5; // Lunes a Viernes
const NUM_BLOQUES: usize = 4; // Bloques horarios por día
const POPULATION_SIZE: usize = 200;
const GENERATION_LIMIT: u64 = 1000;
const SELECTION_RATIO: f64 = 0.7;
const MUTATION_RATE: f64 = 0.05;
const REINSERTION_RATIO: f64 = 0.7;

// ==============================
// Estructuras de datos
// ==============================
#[derive(Clone, Debug, PartialEq)]
struct Profesor {
    id: usize,
    nombre: String,
}

#[derive(Clone, Debug, PartialEq)]
struct Curso {
    id: usize,
    nombre: String,
    profesor_id: usize,
}

#[derive(Clone, Debug, PartialEq)]
struct Salon {
    id: usize,
    nombre: String,
}

// Representación de una clase programada
#[derive(Clone, Debug, PartialEq, PartialOrd)]
struct ClaseProgramada {
    curso_id: usize,
    salon_id: usize,
    dia: usize,    // 0-4 (Lunes-Viernes)
    bloque: usize, // 0-3 (Bloques horarios)
}

// Genotipo: Vector de clases programadas
type HorarioGenome = Vec<ClaseProgramada>;

// ==============================
// Datos de ejemplo
// ==============================
fn crear_profesores() -> Vec<Profesor> {
    (0..NUM_PROFESORES)
        .map(|i| Profesor {
            id: i,
            nombre: format!("Profesor_{}", (b'A' + i as u8) as char),
        })
        .collect()
}

fn crear_cursos(profesores: &[Profesor]) -> Vec<Subject> {
    (0..NUM_CURSOS)
        .map(|i| Subject {
            id: i,
            nombre: format!("Curso_{}", i + 1),
            profesor_id: profesores[i % profesores.len()].id,
        })
        .collect()
}

fn crear_salones() -> Vec<Salon> {
    (0..NUM_SALONES)
        .map(|i| Salon {
            id: i,
            nombre: format!("Salon_{}", (b'A' + i as u8) as char),
        })
        .collect()
}

// ==============================
// Función de Fitness
// ==============================
// fn calcular_colisiones(horario: &HorarioGenome, cursos: &[Curso]) -> usize {
//     let mut colisiones = 0;
//     let mut salon_ocupado = vec![vec![vec![false; NUM_BLOQUES]; NUM_DIAS]; NUM_SALONES];
//     let mut profesor_ocupado = vec![vec![vec![false; NUM_BLOQUES]; NUM_DIAS]; NUM_PROFESORES];
//     println!("HORARIO: {:?}", horario);
//     println!("CURSO: {:?}", cursos);
//     for clase in horario {
//         // Verificar colisiones de salón
//         if salon_ocupado[clase.salon_id][clase.dia][clase.bloque] {
//             colisiones += 1;
//         } else {
//             salon_ocupado[clase.salon_id][clase.dia][clase.bloque] = true;
//         }

//         // Verificar colisiones de profesor
//         if let Some(curso) = cursos.iter().find(|c| c.id == clase.curso_id) {
//             if profesor_ocupado[curso.profesor_id][clase.dia][clase.bloque] {
//                 colisiones += 1;
//             } else {
//                 profesor_ocupado[curso.profesor_id][clase.dia][clase.bloque] = true;
//             }
//         }
//     }
//     colisiones
// }

// #[derive(Clone, Debug)]
// struct FitnessCalc {
//     cursos: Vec<Curso>,
// }

// impl FitnessFunction<HorarioGenome, usize> for FitnessCalc {
//     fn fitness_of(&self, genome: &HorarioGenome) -> usize {
//         let colisiones = calcular_colisiones(genome, &self.cursos);
//         let max_colisiones = genome.len() * 2; // Máximo posible de colisiones

//         // Fitness más alto = menos colisiones
//         if max_colisiones == 0 {
//             100
//         } else {
//             let score = (max_colisiones - colisiones) * 100 / max_colisiones;
//             score as usize
//         }
//     }

//     fn average(&self, values: &[usize]) -> usize {
//         values.iter().sum::<usize>() / values.len()
//     }

//     fn highest_possible_fitness(&self) -> usize {
//         100
//     }

//     fn lowest_possible_fitness(&self) -> usize {
//         0
//     }
// }

// ==============================
// Operadores Genéticos
// ==============================
// impl RandomValueMutation for ClaseProgramada {
//     fn random_mutated<R>(value: Self, _min: &Self, _max: &Self, rng: &mut R) -> Self
//     where
//         R: Rng + Sized,
//     {
//         ClaseProgramada {
//             dia: rng.gen_range(0..NUM_DIAS),
//             bloque: rng.gen_range(0..NUM_BLOQUES),
//             salon_id: rng.gen_range(0..NUM_SALONES),
//             ..value
//         }
//     }
// }

// struct HorarioBuilder {
//     cursos: Vec<Curso>,
// }

// impl GenomeBuilder<HorarioGenome> for HorarioBuilder {
//     fn build_genome<R>(&self, _: usize, rng: &mut R) -> HorarioGenome
//     where
//         R: Rng + Sized,
//     {
//         self.cursos
//             .iter()
//             .map(|curso| ClaseProgramada {
//                 curso_id: curso.id,
//                 salon_id: rng.gen_range(0..NUM_SALONES),
//                 dia: rng.gen_range(0..NUM_DIAS),
//                 bloque: rng.gen_range(0..NUM_BLOQUES),
//             })
//             .collect()
//     }
// }

// ==============================
// Visualización
// ==============================
// impl Display for HorarioGenome {
//     fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
//         let mut output = String::new();
//         let dias = ["Lunes", "Martes", "Miércoles", "Jueves", "Viernes"];
//         let bloques = ["8-10", "10-12", "12-14", "14-16"];

//         for clase in self {
//             output += &format!(
//                 "Curso {}: {} - Salon {}, {} - {}\n",
//                 clase.curso_id,
//                 dias[clase.dia],
//                 (b'A' + clase.salon_id as u8) as char,
//                 bloques[clase.bloque],
//                 dias[clase.dia]
//             );
//         }
//         write!(f, "{}", output)
//     }
// }

// ==============================
// Función principal
// ==============================

pub fn execute_process(config: &AlgorithmConfig) -> Result<(), String> {
    let profesores = crear_profesores();
    let cursos = crear_cursos(&profesores);
    let _salones = crear_salones();

    let fitness_calc = FitnessCalc {
        cursos: cursos.clone(),
    };

    let horario_builder = HorarioBuilder {
        subject: cursos.clone(),
        config: config.clone(),
    };

    let initial_population: Population<HorarioGenome> = build_population()
        .with_genome_builder(horario_builder)
        .of_size(config.population)
        .uniform_at_random();

    let mut simulacion = simulate(
        genetic_algorithm()
            .with_evaluation(fitness_calc.clone())
            .with_selection(RouletteWheelSelector::new(
                config.elitism,
                config.population,
            ))
            .with_crossover(UniformCrossBreeder::new())
            .with_mutation(RandomValueMutator::new(
                config.mutation,
                ClaseProgramada {
                    curso_id: 0,
                    salon_id: 0,
                    dia: 0,
                    bloque: 0,
                },
                ClaseProgramada {
                    // Todos estos datos tienen que llegar de config
                    // numero de cursos
                    curso_id: 10,
                    // numero de salones
                    salon_id: 10,
                    // numero de dias
                    dia: 4,
                    // numero de bloques
                    bloque: 30,
                },
            ))
            .with_reinsertion(ElitistReinserter::new(
                fitness_calc,
                false,
                config.reinsertion,
            ))
            .with_initial_population(initial_population)
            .build(),
    )
    .until(or(
        FitnessLimit::new(100),
        GenerationLimit::new(config.generations),
    ))
    .build();

    loop {
        match simulacion.step() {
            Ok(SimResult::Intermediate(step)) => {
                let best_fitness = step.result.best_solution.solution.fitness;
                println!(
                    "Generación {}: Fitness promedio = {}, Mejor fitness = {}",
                    step.iteration,
                    step.result.evaluated_population.average_fitness(),
                    best_fitness
                );
            }
            Ok(SimResult::Final(step, _time, _duration, stop_reason)) => {
                let best = step.result.best_solution;
                println!("\n--- RESULTADO FINAL ---");
                println!("{}", stop_reason);
                println!("Generación: {}", step.iteration);
                println!("Mejor fitness: {}", best.solution.fitness);
                // println!("\nMEJOR HORARIO ENCONTRADO:\n{}", best.solution.genome);
                break;
            }
            Err(error) => {
                println!("Error: {}", error);
                break;
            }
        }
    }

    Ok(())
}

pub fn ejecutar_algoritmo_horario() {}
