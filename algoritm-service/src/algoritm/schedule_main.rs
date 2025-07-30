use crate::models::algoritm_models::ClaseProgramada;
use crate::models::{
    algoritm_config::AlgorithmConfig, room::Room, subject::Subject, teacher::Teacher,
};
// use create::{chromosome, crossover, fitnes, mutation};
// use crate::algoritm::chromosome::HorarioBuilder;
use crate::algoritm::fitnes::FitnessCalc;

use crate::models::algoritm_models::HorarioBuilder;
use crate::models::schedule_model::ScheduleResponse;
use chrono::Utc;
use genevo::{operator::prelude::*, prelude::*, random::Rng, types::fmt::Display};
use rocket::config;
use std::fmt;
use std::vec::Vec;

// Genotipo: Vector de clases programadas
type HorarioGenome = Vec<ClaseProgramada>;
pub struct HorarioGenomeDisplay(pub Vec<ClaseProgramada>);
// ==============================
// Visualización
// ==============================
impl Display for HorarioGenomeDisplay {
    fn fmt(&self) -> String {
        let dias = [
            "Lunes",
            "Martes",
            "Miércoles",
            "Jueves",
            "Viernes",
            "Sábado",
        ];
        let bloques = ["8-10", "10-12", "12-14", "14-16"];

        let mut output = String::new();

        for clase in &self.0 {
            output += &format!(
                "Curso {}: {} - Salón {}, Bloque {}\n",
                clase.curso_id,
                dias[clase.dia],
                (b'A' + clase.salon_id as u8) as char,
                bloques[clase.bloque]
            );
        }

        output
    }
}

pub fn execute_process(config: &AlgorithmConfig) -> Vec<ScheduleResponse> {
    let subjects = &config.subjects;

    let fitness_calc = FitnessCalc {
        subject: subjects.clone(),
        param: &config,
    };

    let horario_builder = HorarioBuilder {
        subject: subjects.clone(),
        config: &config,
    };

    let initial_population: Population<HorarioGenome> = build_population()
        .with_genome_builder(horario_builder)
        .of_size(config.population)
        .uniform_at_random();

    let mut simulacion = simulate(
        genetic_algorithm()
            .with_evaluation(fitness_calc.clone())
            .with_selection(RouletteWheelSelector::new(
                config.selection,
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
                    curso_id: config.num_subjects,
                    salon_id: config.num_rooms,
                    dia: config.num_days,
                    bloque: config.num_periods,
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
                // println!("\nMEJOR HORARIO ENCONTRADO:\n{:?}", best.solution.genome);
                let best_genome = best.solution.genome;
                return format_schedule_response(best_genome, &config);
                break;
            }
            Err(error) => {
                println!("Error: {}", error);
                break;
            }
        }
    }
    vec![]
}

pub fn format_schedule_response(
    best_genome: Vec<ClaseProgramada>,
    config: &AlgorithmConfig,
) -> Vec<ScheduleResponse> {
    println!("BEST GENOME:{:?}", best_genome);
    println!("CONFIG:{:?}", config);
    let mut response = vec![];

    for chromosome in best_genome {
        let subject_name = config
            .subjects
            .iter()
            .find(|s| s.id == chromosome.curso_id)
            .map(|s| s.name.clone())
            .unwrap_or_else(|| "Desconocido".to_string());

        let room_name = config
            .rooms
            .iter()
            .find(|r| r.id == chromosome.salon_id)
            .map(|s| s.name.clone())
            .unwrap_or_else(|| "Desconocido".to_string());

        // let day = config

        response.push(ScheduleResponse {
            id: chromosome.curso_id,
            startedDate: Utc::now().date_naive(),
            endDate: Utc::now().date_naive(),
            title: subject_name,
            tooltip: room_name,
        });
    }
    response
}

pub fn ejecutar_algoritmo_horario() {}
