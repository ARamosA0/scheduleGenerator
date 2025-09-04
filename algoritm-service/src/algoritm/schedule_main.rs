use crate::models::algoritm_models::ClaseProgramada;
use crate::models::group::{self, Group};
use crate::models::{
    algoritm_config::AlgorithmConfig, room::Room, subject::Subject, teacher::Teacher, algoritm_config::DaySchedule
};
// use create::{chromosome, crossover, fitnes, mutation};
// use crate::algoritm::chromosome::HorarioBuilder;
use crate::algoritm::fitnes::FitnessCalc;

use crate::models::algoritm_models::HorarioBuilder;
use crate::models::schedule_model::ScheduleResponse;
use genevo::{operator::prelude::*, prelude::*, random::Rng, types::fmt::Display};
use rocket::config;
use std::fmt::{self, format};
use std::vec::Vec;
use chrono::{Local, Datelike, Duration, NaiveDate, Utc};
use std::collections::HashMap;

// Genotipo: Vector de clases programadas
pub type HorarioGenome = Vec<ClaseProgramada>;
pub struct HorarioGenomeDisplay(pub Vec<ClaseProgramada>);
#[derive(Clone, Debug)]
pub struct HorarioWrapper(pub HorarioGenome);


// ==============================
// Visualización
// ==============================
impl Display for HorarioGenomeDisplay {
    fn fmt(&self) -> String {
        let dias = [1, 2, 3, 4, 5, 6];
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
    let subjects: &Vec<Subject> = &config.subjects;
    let group: &Vec<Group> = &config.groups;

    let fitness_calc = FitnessCalc {
        subject: subjects.clone(),
        param: &config,
    };

    let horario_builder = HorarioBuilder {
        subject: subjects.clone(),
        group: group.clone(),
        config: &config,
    };

    let initial_population: Population<HorarioGenome> = build_population()
        .with_genome_builder(horario_builder)
        .of_size(config.population)
        .uniform_at_random();

    let mut simulacion = simulate(
        genetic_algorithm()
            .with_evaluation(fitness_calc.clone())
            .with_selection(MaximizeSelector::new(
                config.selection,
                config.population,
            ))
            .with_crossover(UniformCrossBreeder::new())
            .with_mutation(RandomValueMutator::new(
                config.mutation,
                ClaseProgramada {
                    group_id: 0,
                    curso_id: 0,
                    salon_id: 0,
                    teacher_id: 0,
                    dia: 0,
                    bloque: 0,
                },
                ClaseProgramada {
                    group_id: config.num_groups,
                    curso_id: config.num_subjects,
                    salon_id: config.num_rooms,
                    teacher_id: config.num_teachers,
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
    let mut response = vec![];

    let base_date = get_current_week_dates();

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
        let format_day = format_day_period(&chromosome, &config.template);
        let (start_date, end_date) = get_period_datetime(&chromosome, &config.template, &base_date);
        println!("Format DAY:{:?}", format_day);
        // println!("Format PERIOD:{:?}", format_period);
        response.push(ScheduleResponse {
            id: chromosome.curso_id,
            startDate: start_date,
            endDate: end_date,
            title: subject_name,
            tooltip: room_name,
        });
    }
    response
}

pub fn get_current_week_dates() -> HashMap<u32, NaiveDate> {
    let today = Local::now().date_naive();
    let days_since_monday = today.weekday().num_days_from_monday() as i64;
    let monday = today - Duration::days(days_since_monday);
    
    let mut week_dates = HashMap::new();
    
    for day in 0..7 {
        let current_date = monday + Duration::days(day);
        let day_number = (day as u32) + 1;
        week_dates.insert(day_number, current_date);
    }
    
    week_dates
}

pub fn format_day_period(chromosome: &ClaseProgramada, day_schedule: &Vec<DaySchedule>) -> String {
    // Encontrar el día correspondiente en el schedule
    let day_info = day_schedule.iter()
        .find(|d| d.day == chromosome.dia as usize + 1)  // Convertir de 0-indexed a 1-indexed
        .expect("Día no encontrado en el schedule");
    
    // Obtener el período correspondiente
    let period_info = &day_info.periods[chromosome.bloque as usize];
    
    // Formatear la salida
    format!("Día {}, {} ({} a {})", 
            day_info.day,
            period_info.name,
            period_info.startHour,
            period_info.endHour)
}

// Si necesitas una fecha real en lugar de solo formatear el string:
pub fn get_period_datetime(chromosome: &ClaseProgramada, day_schedule: &Vec<DaySchedule>, base_date: &HashMap<u32, NaiveDate>) -> (chrono::NaiveDateTime, chrono::NaiveDateTime) {
    let day_info = day_schedule.iter()
        .find(|d| d.day == chromosome.dia as usize + 1)
        .expect("Día no encontrado en el schedule");
    
    let period_info = &day_info.periods[chromosome.bloque as usize];
    
    // Parsear las horas del período
    let start_time = chrono::NaiveTime::parse_from_str(&period_info.startHour[11..16], "%H:%M")
        .expect("Formato de hora inválido");
    let end_time = chrono::NaiveTime::parse_from_str(&period_info.endHour[11..16], "%H:%M")
        .expect("Formato de hora inválido");
    
    let day_number = day_info.day as u32;
    let target_date = *base_date.get(&day_number).unwrap_or(&NaiveDate::from_ymd(2025, 1, 1));
    
    let start_datetime = chrono::NaiveDateTime::new(target_date, start_time);
    let end_datetime = chrono::NaiveDateTime::new(target_date, end_time);
    
    (start_datetime, end_datetime)
}

pub fn ejecutar_algoritmo_horario() {}
