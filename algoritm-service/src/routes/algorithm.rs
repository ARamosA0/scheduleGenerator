use crate::models::algoritm_config::AlgorithmConfig;
use rocket::serde::json::Json;

use crate::algoritm::demo::{run_genetic_algorithm, BestSolutionResult};
use serde::Serialize;

use crate::algoritm::schedule_main::{ejecutar_algoritmo_horario, execute_process};

#[get("/run-algorithm")]
pub async fn run_algorithm() -> Json<BestSolutionResult> {
    Json(run_genetic_algorithm())
}

#[get("/run-schedule")]
pub async fn run_schedule() {
    println!("\n--- INICIO EL ALGORITMO ---");
    ejecutar_algoritmo_horario()
}

#[post("/generar", format = "json", data = "<config>")]
pub async fn generar_horario(config: Json<AlgorithmConfig>) -> String {
    let mut config = config.into_inner();

    config.generation_limit = 100;
    config.mutation_rate = 0.01;
    config.num_bloques = 1;
    config.num_dias = 6;
    config.num_profesores = 10;
    config.num_salones = 3;
    config.population_size = 1000;
    config.reinsertion_ratio = 10.4;
    config.selection_ratio = 3.0;

    execute_process(&config);
    // match execute_process(&config) {
    //     Ok(_) => "Algoritmo ejecutado con éxito".into(),
    //     Err(e) => {
    //         eprintln!("Error al ejecutar el algoritmo: {}", e);
    //         format!("Error: {}", e)
    //     }
    // }
    "Algoritmo ejecutado con éxito".into()
}
